package main

import (
	"log"

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

func as(cond bool) {
	if !cond {
		log.Fatal("assertion failed")
	}
}

func main() {
	addr := "localhost:3000"
	ready := make(chan bool)
	st := store.NewStorage()
	var b bool
	ne(st.Set(&trib.KeyValue{"hello", "hi"}, &b))
	as(b == true)

	go func() {
		e := entries.ServeBackSingle(addr, st, ready)
		ne(e)
	}()

	r := <-ready
	if !r {
		log.Fatal("not ready")
	}

	s := triblab.NewClient(addr)

	var str string
	ne(s.Get("hello", &str))

	if str != "hi" {
		log.Fatal("not using the storage provided")
	}
}
