package db_struct

import "template/controller/rt_struct"

type Sample struct {
	Name string `bson:"name"`
}

func (sample *Sample) ToController() *rt_struct.Sample {
	return &rt_struct.Sample{
		Name: sample.Name,
	}
}

func InitSample(sample *rt_struct.Sample) *Sample {
	return &Sample{
		Name: sample.Name,
	}
}
