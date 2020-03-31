package main

import (
	"log"
	"time"

	"trib/entries"
	"trib/store"
)

func main() {
	addr := "^_^"
	ready := make(chan bool, 0)

	go func() {
		e := entries.ServeBackSingle(addr, store.NewStorage(), ready)
		if e == nil {
			log.Fatal("address was error")
		}
	}()

	t := time.After(time.Second * 2)
	select {
	case <-t:
		log.Fatal("should return error first")
	case r := <-ready:
		if r != false {
			log.Fatal("should send false on Ready")
		}
	}
}
