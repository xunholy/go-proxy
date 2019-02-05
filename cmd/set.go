package cmd

import (
	"fmt"

	"github.com/urfave/cli"
)

func SetCommand() cli.Command {
	return cli.Command{
		Name:        "set",
		Aliases:     []string{""},
		Usage:       "proxy set",
		Description: "Set CNTLM Proxy Config",
		Subcommands: []cli.Command{
			{
				Name:        "npm",
				Usage:       "set npm proxy config",
				Description: "additional description?",
				Flags: []cli.Flag{
					cli.IntFlag{
						Name:        "port, p",
						Value:       3128,
						Usage:       "set custom CNTLM `PORT`",
						Destination: &port,
					},
				},
				Action: func(c *cli.Context) {
					p := fmt.Sprintf("http://localhost:%v", port)
					a := []string{"config", "set", "proxy", p}
					e := execCommand{cmd: "npm", args: a}
					executeCommand(e)
					fmt.Println("Set npm config successfully")
				},
			},
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "password, p",
				Usage:       "set CNTLM `PASSWORD` config",
				Value:       "",
				Destination: &password,
			},
		},
		Action: func(c *cli.Context) {
			if password != "" {
				// process password update here
			}
		},
	}
}
