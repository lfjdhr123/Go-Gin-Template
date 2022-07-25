package request

import "template/controller/rt_struct"

type Sample struct {
	Name string `json:"name"`
}

func (sample *Sample) ToController() *rt_struct.Sample {
	return &rt_struct.Sample{
		Name: sample.Name,
	}
}
