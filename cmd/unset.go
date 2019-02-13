package cmd

import (
	"fmt"

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
				Action: func(c *cli.Context) {
					cmds := []execute.NewCommand{}
					cmds = append(cmds, execute.NewCommand{Cmd: "npm", Args: []string{"config", "delete", "proxy"}})
					execute.Commands(cmds)
					fmt.Println("Unset npm config successfully")
				},
			},
			{
				Name:        "git",
				Usage:       "unset git proxy config",
				Description: "additional description?",
				Action: func(c *cli.Context) {
					cmds := []execute.NewCommand{}
					cmds = append(cmds, execute.NewCommand{Cmd: "git", Args: []string{"config", "--global", "--unset", "http.proxy"}})
					cmds = append(cmds, execute.NewCommand{Cmd: "git", Args: []string{"config", "--global", "--unset", "https.proxy"}})
					execute.Commands(cmds)
					fmt.Println("Unset git config successfully")
				},
			},
		},
	}
}
