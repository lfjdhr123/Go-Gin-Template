package database

import (
	"fmt"
	"template/conf"
	"template/controller/rt_struct"
	"template/database/mongodb"
)

type DB struct {
	mongo *mongodb.MongoDB
}

// Init the databases' connection
func (db *DB) Init(mongoConfig conf.MongoDB) error {
	db.mongo = &mongodb.MongoDB{
		Host:         mongoConfig.Host,
		Port:         mongoConfig.Port,
		DatabaseName: mongoConfig.DBName,
	}
	err := db.mongo.Init()
	if err != nil {
		return fmt.Errorf("cannot connect to mongodb due to: %v", err)
	}
	return nil
}

func (db *DB) SaveSample(sample *rt_struct.Sample) error {
	return db.mongo.SaveSample(sample)
}

func (db *DB) GetSampleByName(name string) (*rt_struct.Sample, error) {
	return db.mongo.GetSampleByName(name)
}


func (db *DB) GetSampleList() ([]*rt_struct.Sample, error) {
	return db.mongo.GetSampleList()
}
