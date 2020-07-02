package main

import (
	"activities/gateway/router"
	_ "activities/library/clients/nats"
	"activities/library/config"
	_ "activities/service"
	"fmt"
)

func main() {
	conf := config.GetServerConf()
	router.NewRouter().Run(fmt.Sprintf("%v:%v", conf.Addr, conf.Post))
}
