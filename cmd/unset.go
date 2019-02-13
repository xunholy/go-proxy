package cmd

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/xUnholy/go-proxy/execute"
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
					cmds := []execCommand{}
					cmds = append(cmds, execCommand{cmd: "npm", args: []string{"config", "delete", "proxy"}})
					execute.Commands(cmds)
					fmt.Println("Unset npm config successfully")
				},
			},
			{
				Name:        "git",
				Usage:       "unset git proxy config",
				Description: "additional description?",
				Action: func(c *cli.Context) {
					cmds := []execCommand{}
					cmds = append(cmds, execCommand{cmd: "git", args: []string{"config", "--global", "--unset", "http.proxy"}})
					cmds = append(cmds, execCommand{cmd: "git", args: []string{"config", "--global", "--unset", "https.proxy"}})
					execute.Commands(cmds)
					fmt.Println("Unset git config successfully")
				},
			},
		},
	}
}
