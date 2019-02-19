package cmd

import (
	"fmt"
	"log"

	"github.com/urfave/cli"
	"github.com/xUnholy/go-proxy/pkg/execute"
)

func UnsetCommand() cli.Command {
	return cli.Command{
		Name:        "unset",
		Aliases:     []string{""},
		Usage:       "proxy unset",
		Description: "Unset CNTLM Proxy Config",
		Subcommands: []cli.Command{
			{
				Name:        "npm",
				Usage:       "unset npm proxy config",
				Description: "additional description?",
				Action: func(_ *cli.Context) {
					cmds := []execute.NewCommand{}
					cmds = append(cmds, execute.NewCommand{Cmd: "npm", Args: []string{"config", "delete", "proxy"}})
					output, err := execute.Commands(cmds)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println(output)
					fmt.Println("Unset npm config successfully")
				},
			},
			{
				Name:        "git",
				Usage:       "unset git proxy config",
				Description: "additional description?",
				Action: func(_ *cli.Context) {
					cmds := []execute.NewCommand{}
					http := execute.NewCommand{Cmd: "git", Args: []string{"config", "--global", "--unset", "http.proxy"}}
					https := execute.NewCommand{Cmd: "git", Args: []string{"config", "--global", "--unset", "https.proxy"}}
					cmds = append(cmds, http, https)
					output, err := execute.Commands(cmds)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println(output)
					fmt.Println("Unset git config successfully")
				},
			},
		},
	}
}
