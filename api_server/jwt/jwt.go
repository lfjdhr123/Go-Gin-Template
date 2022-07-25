package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

const (
	secretKey     = "my-jwt-secret-key"
	tokenDuration = 365 * 24 * time.Hour
)

// Manager is a JSON web token manager
type Manager struct {
	secretKey     string
	tokenDuration time.Duration
}

// UserClaims is a custom JWT claims that contains some user's information
type UserClaims struct {
	jwt.StandardClaims
	AccountID string `json:"accountID"`
}

// NewJWTManager returns a new JWT manager
func NewJWTManager() *Manager {
	return &Manager{secretKey, tokenDuration}
}

// Generate generates and signs a new token for a user
func (manager *Manager) Generate(accountID string) (string, error) {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.tokenDuration).Unix(),
		},
		AccountID: accountID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.secretKey))
}

// Verify verifies the access token string and return a user claim if the token is valid
func (manager *Manager) Verify(accessToken string) (*UserClaims, error) {
	if accessToken == "" {
		return nil, fmt.Errorf("invalid token, cannot be empty")
	}
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return []byte(manager.secretKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
