package controller

import (
	"template/conf"
	"template/controller/rt_struct"
	"template/database"
)

type Controller struct {
	db *database.DB
}

func Init(mongoConfig conf.MongoDB) (*Controller, error) {
	db := &database.DB{}
	if err := db.Init(mongoConfig); err != nil {
		return nil, err
	}
	return &Controller{db: db}, nil
}

func (controller *Controller) CreateSample(sample *rt_struct.Sample) error {
	return controller.db.SaveSample(sample)
}

func (controller *Controller) GetSampleByName(name string) (*rt_struct.Sample, error) {
	return controller.db.GetSampleByName(name)
}

func (controller *Controller) GetSampleList() ([]*rt_struct.Sample, error) {
	return controller.db.GetSampleList()
}
