package main

import (
	"fmt"
	"github.com/caibei1/mydocker/container"
	"os"
)

func Run(tty bool, command string) {
	parent := container.NewParentProcess(tty, command)
	if err := parent.Start(); err != nil {
		fmt.Println(err)
	}
	parent.Wait()
	os.Exit(-1)
}
