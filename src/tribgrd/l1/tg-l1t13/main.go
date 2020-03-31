package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
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
	okay := make(chan bool, 1)
	ready := make(chan bool)
	go func() {
		e := entries.ServeBackSingle(addr, store.NewStorage(), ready)
		okay <- (e != nil)
	}()

	<-ready
	pid := os.Getpid()
	out, e := exec.Command("ls", fmt.Sprintf("/proc/%d/fd", pid)).Output()
	ne(e)
	fs := strings.Fields(string(out))
	for _, f := range fs {
		stat, e := os.Stat(fmt.Sprintf("/proc/%d/fd/%s", pid, f))
		if e == nil && (stat.Mode()&os.ModeSocket) != 0 {
			fd, e := strconv.Atoi(f)
			ne(e)

			e = syscall.Shutdown(fd, syscall.SHUT_RD)
			ne(e)
		}
	}

	time.Sleep(time.Second * 1)
	if len(okay) == 0 {
		log.Fatal("not returning anything")
	}
	if <-okay == false {
		log.Fatal("not returning an error")
	}
}
