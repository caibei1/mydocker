package main

import (
	"fmt"
	"github.com/caibei1/mydocker/container"
	"github.com/urfave/cli/v2"
)

var runCommand = &cli.Command{
	Name: "run",
	Usage: `Create a container with namespace and cgroups limit
			mydocker run -ti [command]`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "ti",
			Usage: "enable tty",
		},
	},
	Action: func(context *cli.Context) error {
		if context.Args().Len() < 1 {
			return fmt.Errorf("%s \n", "Missing container command")
		}
		cmd := context.Args().Get(0)
		tty := context.Bool("ti")
		Run(tty, cmd)
		return nil
	},
}

var initCommand = &cli.Command{
	Name:  "init",
	Usage: "Init container process run user's process in container. Do not call it outside",
	Action: func(context *cli.Context) error {
		fmt.Println("init come on")
		cmd := context.Args().Get(0)
		fmt.Printf("command %s \n", cmd)
		err := container.RunContainerInitProcess(cmd, nil)
		return err
	},
}
