package main

import "log"

func main() {
	log.Println("successful")
	signalChan := make(chan int)
	<-signalChan
}
