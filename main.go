package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

const usage = `mydocker is a simple container runtime implementation.
			   The purpose of this project is to learn how docker works and how to write a docker by ourselves
			   Enjoy it, just for fun.`

func main() {
	app := cli.NewApp()
	app.Name = "mydocker"
	app.Usage = usage
	app.Commands = []*cli.Command{
		initCommand,
		runCommand,
	}

	app.Before = func(context *cli.Context) error {
		fmt.Println("start...")
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
