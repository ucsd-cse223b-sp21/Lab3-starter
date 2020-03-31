package main

import (
	"log"
	"os"
	"os/exec"
	"time"

	"trib"
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
	c := exec.Command("kv-server", "-addr=localhost:3000")
	e := c.Start()
	ne(e)

	time.Sleep(time.Second)

	s := triblab.NewClient("localhost:3000")

	var b bool
	var str string
	e1 := s.Set(&trib.KeyValue{"hello", "hi"}, &b)

	c.Process.Signal(os.Interrupt)
	time.Sleep(time.Second)
	c.Process.Kill()

	time.Sleep(time.Second)

	e2 := s.Get("hello", &str)

	ne(e1)
	as(b == true)
	if e2 == nil {
		log.Fatal("not returning error when server is down")
	}
}
