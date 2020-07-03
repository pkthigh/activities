package main

import (
	"activities/gateway/router"
	_ "activities/library/clients/nats"
	"activities/library/config"
	"activities/library/logger"
	_ "activities/service"
)

func main() {
	if err := router.NewRouter().Run(config.GetServerConf().Address()); err != nil {
		logger.ErrorF("Gateway Router Run Fail: %v", err)
	}
}
