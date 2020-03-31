package main

import (
	"log"
	"time"

	"trib/entries"
	"trib/store"
)

func ne(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	addr := "localhost:3000"
	ready := make(chan bool, 1)

	done := make(chan bool, 1)

	go func() {
		e := entries.ServeBackSingle(addr, store.NewStorage(), ready)
		ne(e)
		done <- true
	}()

	t := time.After(time.Second * 2)
	select {
	case <-t:
	case <-done:
		if len(t) == 0 {
			log.Fatal("server is not a blocking call")
		}
	}
}
