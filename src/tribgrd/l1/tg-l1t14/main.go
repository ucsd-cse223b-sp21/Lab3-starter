package main

import (
	"log"
	"runtime"

	"trib"
	"trib/entries"
	"trib/store"
	"triblab"
)

func ne(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	addr := "localhost:3000"
	ready := make(chan bool)

	go func() {
		e := entries.ServeBackSingle(addr, store.NewStorage(), ready)
		ne(e)
	}()

	<-ready

	done := make(chan bool, 2)
	f := func() {
		c := triblab.NewClient(addr)
		var b bool
		for i := 0; i < 10; i++ {
			b = false
			ne(c.ListAppend(&trib.KeyValue{"lst", "item"}, &b))
			if !b {
				log.Fatal("append failed")
			}
			runtime.Gosched()
		}
		done <- true
	}

	go f()
	go f()

	<-done
	<-done

	var lst trib.List
	c := triblab.NewClient(addr)
	c.ListGet("lst", &lst)
	if len(lst.L) != 20 {
		log.Fatal("list count incorrect")
	}
}
