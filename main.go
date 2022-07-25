package main

import (
	"log"
	"template/api_server"
	"template/conf"
)

func main() {
	config, err := conf.GetConfig()
	if err != nil {
		log.Fatalln("fail to get the config", err)
	}
	if err = api_server.Start(config.Server, config.MongoDB, config.Log); err != nil {
		log.Fatalln("fail to start http server", err)
	}
}
