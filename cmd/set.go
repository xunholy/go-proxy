package cmd

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/xUnholy/go-proxy/pkg/execute"
)

func SetCommand() cli.Command {
	return cli.Command{
		Name:        "set",
		Aliases:     []string{""},
		Usage:       "proxy set",
		Description: "Set CNTLM Proxy Config",
		Subcommands: []cli.Command{
			{
				Name:        "npm",
				Usage:       "set npm proxy config",
				Description: "additional description?",
				Flags: []cli.Flag{
					cli.IntFlag{
						Name:        "port, p",
						Value:       3128,
						Usage:       "set custom CNTLM `PORT`",
						Destination: &port,
					},
				},
				Action: func() {
					p := SetProxyPort(port)
					cmds := []execute.NewCommand{}
					cmds = append(cmds, execute.NewCommand{Cmd: "npm", Args: []string{"config", "set", "proxy", p}})
					execute.Commands(cmds)
					fmt.Println("Set npm config successfully")
				},
			},
			{
				Name:        "git",
				Usage:       "set git proxy config",
				Description: "additional description?",
				Flags: []cli.Flag{
					cli.IntFlag{
						Name:        "port, p",
						Value:       3128,
						Usage:       "set custom CNTLM `PORT`",
						Destination: &port,
					},
				},
				Action: func() {
					p := SetProxyPort(port)
					cmds := []execute.NewCommand{}
					http := execute.NewCommand{Cmd: "git", Args: []string{"config", "--global", "http.proxy", p}}
					https := execute.NewCommand{Cmd: "git", Args: []string{"config", "--global", "https.proxy", p}}
					cmds = append(cmds, http, https)
					execute.Commands(cmds)
					fmt.Println("Set npm config successfully")
				},
			},
			{
				Name:        "password",
				Usage:       "proxy set password",
				Description: "additional description?",
				Action: func() {
					fmt.Printf("Enter Password: ")
					e := execute.NewCommand{Cmd: "cntlm", Args: []string{"-H"}}
					o := execute.Command(e)
					UpdatePassword(o)
					fmt.Println("Set cntlm config successfully")
				},
			},
		},
		Action: func() {
			fmt.Println("Set command invoked")
		},
	}
}

func SetProxyPort(port int) (p string) {
	p = fmt.Sprintf("http://localhost:%v", port)
	return
}
