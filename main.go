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
			configFile := c.Args().First()
			kubetmuxp.CreateKubeConfig(configFile)
			return nil
		},
		Version: fmt.Sprintf("0.1.0-%s", version),
	}

	app.Run(os.Args)
}
