package main

import (
	"fmt"
	"os"

	kubetmuxp "github.com/karuppiah7890/kube-tmuxp-config-gen/kubetmuxp"
	cli "gopkg.in/urfave/cli.v2"
)

var version = "dev"

func main() {
	app := &cli.App{
		Name:  "kube-tmuxp-config-gen",
		Usage: "Generate configurations for tmuxp to work with multiple GKE kubernetes clusters",
		Action: func(c *cli.Context) error {

			numberOfArguments := c.NArg()

			if numberOfArguments != 1 {
				return cli.Exit("Please pass exactly one configuration file", 1)
			}

			configFile := c.Args().First()
			kubetmuxp.CreateKubeTmuxpConfig(configFile)
			return nil
		},
		Version: fmt.Sprintf("0.1.0-%s", version),
	}

	app.Run(os.Args)
}
