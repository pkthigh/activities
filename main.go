package main

import (
	_ "activities/library/clients/nats"
	"log"
)

func main() {
	log.Println("successful")
	signalChan := make(chan int)
	<-signalChan
}
