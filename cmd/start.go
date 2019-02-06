package cmd

import (
	"fmt"

	"github.com/urfave/cli"
)

func StartCommand() cli.Command {
	return cli.Command{
		Name:        "start",
		Aliases:     []string{""},
		Usage:       "proxy start",
		Description: "Start CNTLM Proxy",
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:        "port, p",
				Value:       3128,
				Usage:       "set custom CNTLM `PORT`",
				Destination: &port,
			},
			cli.BoolFlag{
				Name:        "all, a",
				Usage:       "set all CNTLM config",
				Destination: &setAll,
			},
		},
		Action: func() {
			UpdatePort(port)
			proxy := fmt.Sprintf("http://localhost:%v", port)
			fmt.Println("CNTLM Proxy Started On", proxy)
		},
	}
}
