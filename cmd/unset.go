package cmd

import (
	"fmt"

	"github.com/urfave/cli"
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
				Action: func() {
					a := []string{"config", "delete", "proxy"}
					e := execCommand{cmd: "npm", args: a}
					executeCommand(e)
					fmt.Println("Unset npm config successfully")
				},
			},
		},
	}
}
