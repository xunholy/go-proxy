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
					cmds := []execute.Command{}
					cmds = append(cmds, execute.Command{Cmd: "npm", Args: []string{"config", "delete", "proxy"}})
					output, err := execute.RunCommands(cmds)
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
					cmds := []execute.Command{}
					http := execute.Command{Cmd: "git", Args: []string{"config", "--global", "--unset", "http.proxy"}}
					https := execute.Command{Cmd: "git", Args: []string{"config", "--global", "--unset", "https.proxy"}}
					cmds = append(cmds, http, https)
					output, err := execute.RunCommands(cmds)
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
