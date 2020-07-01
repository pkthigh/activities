package main

import (
	"activities/common"
	"activities/library/storage"
	"fmt"
)

func main() {
	itemStore, _ := storage.GetRdsDB(common.ItemRecordStore)
	// handStore, _ := storage.GetRdsDB(common.HandOverRecordStore)
	// insuranceStore, _ := storage.GetRdsDB(common.InsuranceRecordStore)
	cmd := itemStore.Keys("*")
	fmt.Printf("ids: %+v", cmd.Val())
	// log.Println("successful")
	// signalChan := make(chan int)
	// <-signalChan
}
