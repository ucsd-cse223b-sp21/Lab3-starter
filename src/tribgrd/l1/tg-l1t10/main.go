package main

import (
	"log"
	"os/exec"

	"trib/entries"
	"trib/store"
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

	go func() {
		e := entries.ServeBackSingle(addr, st, ready)
		ne(e)
	}()

	r := <-ready
	as(r)

	c := exec.Command("kv-client", "localhost:3000", "clock", "2999")
	e := c.Run()
	ne(e)

	var clk uint64
	ne(st.Clock(0, &clk))
	as(clk == 3000)
}
