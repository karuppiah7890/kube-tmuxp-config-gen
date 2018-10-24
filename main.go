package main

import (
	"fmt"
	"os"

	cli "gopkg.in/urfave/cli.v2"
)

var version = "dev"

func main() {
	app := &cli.App{
		Name:  "kube-tmuxp-config-gen",
		Usage: "Generate configurations for tmuxp to work with multiple GKE kubernetes clusters",
		Action: func(c *cli.Context) error {
			fmt.Println("This feature has not yet been implemented")
			return nil
		},
		Version: fmt.Sprintf("0.1.0-%s", version),
	}

	app.Run(os.Args)
}
