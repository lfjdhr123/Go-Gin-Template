package response

import "template/controller/rt_struct"

type Sample struct {
	Name string `json:"name"`
}

func InitSample(sample *rt_struct.Sample) *Sample {
	return &Sample{
		Name: sample.Name,
	}
}
