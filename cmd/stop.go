package cmd

import (
	"fmt"

	"github.com/urfave/cli"
)

func StopCommand() cli.Command {
	return cli.Command{

		Name:        "stop",
		Aliases:     []string{""},
		Usage:       "proxy stop",
		Description: "Stop CNTLM Proxy",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:        "all, a",
				Usage:       "unset all CNTLM config",
				Destination: &setConfig,
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("All Command Executed: ", c.Args().First())
			return nil
		},
	}

}
