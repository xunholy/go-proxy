package cmd

import (
	"github.com/urfave/cli"
	"github.com/xUnholy/go-proxy/file"
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
				Destination: &setAll,
			},
		},
		Action: func(c *cli.Context) {
			if file.Contains(c.FlagNames(), "all") {
				println("true")
			}
		},
	}

}
