package cmd

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/xUnholy/go-proxy/file"
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
		Action: func(c *cli.Context) {
			if file.Contains(c.FlagNames(), "all") {
				println("true")
			}

			UpdatePort(port)
			fmt.Printf("CNTLM Proxy Started On http://localhost:%v\n", port)
		},
	}
}
