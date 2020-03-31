package main

import (
	"log"
	"runtime/debug"

	"trib"
	"trib/entries"
	"trib/store"
	"triblab"
)

func main() {
	addr := "localhost:3000"
	ready := make(chan bool)

	go func() {
		e := entries.ServeBackSingle(addr, store.NewStorage(), ready)
		if e != nil {
			log.Fatal(e)
		}
	}()

	r := <-ready
	if !r {
		log.Fatal("not ready")
	}

	ne := func(e error) {
		if e != nil {
			debug.PrintStack()
			log.Fatal(e)
		}
	}

	s := triblab.NewClient(addr)

	var l = new(trib.List)

	l.L = []string{"foo"}
	ne(s.Keys(new(trib.Pattern), l))
	if len(l.L) != 0 {
		log.Printf("first item: %q", l.L[0])
		log.Fatal("not reseting the list for empty list")
	}
}
