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
				Description: "This command will unset the NPM proxy values. Both https-proxy and proxy will be unset",
				Action: func(_ *cli.Context) {
					cmds := []execute.Command{}
					cmds = append(cmds, execute.Command{Cmd: "npm", Args: []string{"config", "delete", "proxy"}})
					cmds = append(cmds, execute.Command{Cmd: "npm", Args: []string{"config", "delete", "https-proxy"}})
					_, err := execute.RunCommands(cmds)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println("Unset npm config successfully")
				},
			},
			{
				Name:        "git",
				Usage:       "unset git proxy config",
				Description: "This command will unset the GIT global proxy values. Both http.proxy and https.proxy will be unset",
				Action: func(_ *cli.Context) {
					cmds := []execute.Command{}
					http := execute.Command{Cmd: "git", Args: []string{"config", "--global", "--unset", "http.proxy"}}
					https := execute.Command{Cmd: "git", Args: []string{"config", "--global", "--unset", "https.proxy"}}
					cmds = append(cmds, http, https)
					_, err := execute.RunCommands(cmds)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println("Unset git config successfully")
				},
			},
		},
	}
}
