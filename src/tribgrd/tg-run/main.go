package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const TEST_TIMEOUT_MIN = 2

func ne(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		panic("bug")
	}

	p := args[0]
	ne(os.MkdirAll(p, 0700))

	done := make(chan error)
	cmd := exec.Command(args[1])

	go func() {
		output, e := cmd.CombinedOutput()
		ne(ioutil.WriteFile(filepath.Join(p, "out"), output, 0600))
		done <- e
	}()

	res := func(s string) {
		ne(ioutil.WriteFile(filepath.Join(p, "res"), []byte(s), 0600))
	}

	timeout := time.After(TEST_TIMEOUT_MIN * time.Minute)

	exitcode := 0
	select {
	case e := <-done:
		if e == nil {
			res("pass")
			fmt.Println("ok")
			exitcode = 0
		} else {
			res("error: " + e.Error())
			fmt.Println("FAILED")
			exitcode = 1
		}
	case <-timeout:
		res("timeout")
		fmt.Println("TIMEOUT")
		exitcode = 1

		// clean up
		cmd.Process.Kill()
		<-done
	}

	exec.Command("pkill", "kv-server").Run()
	exec.Command("pkill", "kv-client").Run()
	exec.Command("pkill", "bins-back").Run()
	exec.Command("pkill", "bins-keeper").Run()

	os.Exit(exitcode)
}
