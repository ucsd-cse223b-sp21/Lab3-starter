package main

import (
	"log"
	"time"

	"trib/entries"
	"trib/store"
)

func main() {
	addr := "^_^"
	done := make(chan bool)

	go func() {
		e := entries.ServeBackSingle(addr, store.NewStorage(), nil)
		if e == nil {
			log.Fatal("address was error")
		}
		done <- true
	}()

	t := time.After(time.Second * 2)
	select {
	case <-t:
		log.Fatal("should return error first, might be blocking on nil channel")
	case <-done:
	}
}
