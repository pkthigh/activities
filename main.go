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
	fmt.Printf("%+v", conf)
	router.NewRouter().Run(fmt.Sprintf("%s:%s", conf.Addr, conf.Port))
}
