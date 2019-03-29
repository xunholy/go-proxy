package cmd

import (
	"fmt"
	"log"

	"github.com/urfave/cli"
	"github.com/xUnholy/go-proxy/pkg/execute"

	"github.com/xUnholy/go-proxy/internal/cntlm"
	"github.com/xUnholy/go-proxy/internal/profile"
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
		},
		Action: func(_ *cli.Context) {
			proxyURL := makeProxyURL(port)
			profile.UpdateGlobalEnvironmentVariables(proxyURL)
			p := fmt.Sprintf("Listen\t%v", port)
			cntlm.UpdateFile(cntlmFile, p)
			cmds := execute.Command{Cmd: "cntlm", Args: []string{"-g"}}
			_, err := execute.RunCommand(cmds)
			if err != nil {
				fmt.Println("CNTLM Proxy couldn't be started. Is it already running?")
				log.Fatal(err)
			}
			fmt.Println("CNTLM Proxy Started On", proxyURL)
		},
	}
}
