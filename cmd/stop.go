package cmd

import (
	"fmt"
	"log"

	"github.com/urfave/cli"
	"github.com/xUnholy/go-proxy/pkg/execute"
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
			cmds := execute.NewCommand{Cmd: "pkill", Args: []string{"cntlm"}}
			_, err := execute.Command(cmds)
			if err != nil {
				fmt.Println("Couldn't kill CNTLM process")
				log.Fatal(err)
			}
			fmt.Println("CNTLM proxy stopped")
		},
	}

}
