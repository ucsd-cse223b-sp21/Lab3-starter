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
	addr1 := "localhost:3000"
	addr2 := "localhost:3001"
	ready := make(chan bool, 2)

	go func() {
		e := entries.ServeBackSingle(addr1, store.NewStorage(), ready)
		ne(e)
	}()

	go func() {
		e := entries.ServeBackSingle(addr2, store.NewStorage(), ready)
		ne(e)
	}()

	go func() {
		time.Sleep(time.Second * 2)
		log.Fatal("time out")
	}()

	r := <-ready
	r = r && <-ready

	if r != true {
		log.Fatal("not ready")
	}
}
