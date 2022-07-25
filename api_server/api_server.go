package api_server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"template/api_server/jwt"
	"template/conf"
	"template/controller"
)

const (
	ModeProduction  = "production"
	ModeDevelopment = "dev"
)

type RouterHandler struct {
	controller *controller.Controller
	jwtManager *jwt.Manager
}

func InitRouterHandler(mongoConfig conf.MongoDB) (*RouterHandler, error) {
	c, err := controller.Init(mongoConfig)
	if err != nil {
		return nil, err
	}
	return &RouterHandler{controller: c, jwtManager: jwt.NewJWTManager()}, nil
}
func Start(httpConfig conf.Server, mongoConfig conf.MongoDB, logConfig conf.Log) error {
	//response.InitErrorMap()
	routerHandler, err := InitRouterHandler(mongoConfig)
	if err != nil {
		log.Fatal(err)
	}
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	if os.Getenv("EVENTOR_MODE") == ModeProduction {
		f, _ := os.Create(logConfig.Path)
		gin.DefaultWriter = io.MultiWriter(f)
	}

	router := gin.Default()
	cwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	router.Static("/statics", fmt.Sprintf("%s/assets", cwd))
	router.Use(corsMiddleware)
	loadRouterHandler(router, routerHandler)

	if err = router.Run(fmt.Sprintf(":%s", httpConfig.Port)); err != nil {
		return err
	}
	return nil
}
func corsMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, x-auth-token")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
	c.Next()
}
