package conf

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	MongoDB
	Server
	Log
}

type MongoDB struct {
	Host   string `yaml:"host"`
	Port   uint64 `yaml:"port"`
	DBName string `yaml:"db_name"`
}

type Server struct {
	Port string `yaml:"port"`
}

type Log struct {
	Path string `yaml:"path"`
}

func GetConfig() (Config, error) {
	result := Config{}
	if config, err := ioutil.ReadFile("./conf/config.yml"); err == nil {
		if err := yaml.Unmarshal(config, &result); err != nil {
			return result, err
		}
	} else {
		return result, err
	}
	return result, nil
}
