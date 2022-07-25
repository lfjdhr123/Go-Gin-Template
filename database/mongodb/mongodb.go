package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"template/controller/rt_struct"
	"template/database/mongodb/db_struct"
	"time"
)

const (
	CollectionSample = "sample"
)
type MongoDB struct {
	Host         string
	Port         uint64
	DatabaseName string
	Timeout      uint8
	DB           *mongo.Database
	Client       *mongo.Client
}

func (mdb *MongoDB) Init() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	uri := fmt.Sprintf("mongodb://%s:%d", mdb.Host, mdb.Port)
	log.Printf("Connecting to %s, DB name: %s \n", uri, mdb.DatabaseName)
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}
	mdb.Client = client
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return fmt.Errorf("ping mongodb failed: %v", err)
	}
	mdb.DB = client.Database(mdb.DatabaseName)
	log.Printf("mongodb connected to host: [%s], port [%d]", mdb.Host, mdb.Port)
	return nil
}


func (mdb *MongoDB) SaveSample(sample *rt_struct.Sample) error {
	collection := mdb.DB.Collection(CollectionSample)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	dbSample := db_struct.InitSample(sample)
	if _, err := collection.InsertOne(ctx, dbSample); err != nil {
		return err
	}
	return nil
}

func (mdb *MongoDB) GetSampleByName(name string) (*rt_struct.Sample, error) {
	collection := mdb.DB.Collection(CollectionSample)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var result db_struct.Sample
	if err := collection.FindOne(ctx, bson.M{"name": name}).Decode(&result); err != nil {
		return nil, err
	}
	return result.ToController(), nil
}

func (mdb *MongoDB) GetSampleList() ([]*rt_struct.Sample, error) {
	collection := mdb.DB.Collection(CollectionSample)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var samples []*rt_struct.Sample

	for cur.Next(ctx) {
		var result db_struct.Sample
		err = cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		samples = append(samples, result.ToController())
	}
	if err = cur.Err(); err != nil {
		return nil, err
	}
	return samples, nil
}