package cmd

import (
	"fmt"
	"log"

	"github.com/urfave/cli"
	"github.com/xUnholy/go-proxy/pkg/execute"
	"github.com/xUnholy/go-proxy/pkg/prompt"

	"github.com/xUnholy/go-proxy/internal/cntlm"
)

var (
	cntlmFile = "/usr/local/etc/cntlm.conf"
	port      int
	setAll    bool
)

func SetCommand() cli.Command {
	return cli.Command{
		Name:        "set",
		Aliases:     []string{""},
		Usage:       "proxy set",
		Description: "Set CNTLM Proxy Config",
		Action: func(_ *cli.Context) {
			fmt.Println("Set command invoked")
		},
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
				Action: func(_ *cli.Context) {
					p := setProxyURL(port)
					cmds := []execute.Command{}
					cmds = append(cmds, execute.Command{Cmd: "npm", Args: []string{"config", "set", "proxy", p}})
					_, err := execute.RunCommands(cmds)
					if err != nil {
						log.Fatal(err)
					}
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
				Action: func(_ *cli.Context) {
					p := setProxyURL(port)
					cmds := []execute.Command{}
					http := execute.Command{Cmd: "git", Args: []string{"config", "--global", "http.proxy", p}}
					https := execute.Command{Cmd: "git", Args: []string{"config", "--global", "https.proxy", p}}
					cmds = append(cmds, http, https)
					_, err := execute.RunCommands(cmds)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println("Set npm config successfully")
				},
			},
			{
				Name:        "username",
				Usage:       "proxy set username",
				Description: "additional description?",
				Action: func(_ *cli.Context) {
					fmt.Printf("Enter Username: ")
					output, err := prompt.GetInput()
					if err != nil {
						log.Fatal(err)
					}
					update := fmt.Sprintln("Username\t", output)
					cntlm.UpdateFile(cntlmFile, update)
					fmt.Println("Set CNTLM username successfully")
				},
			},
			{
				Name:        "password",
				Usage:       "proxy set password",
				Description: "additional description?",
				Action: func(_ *cli.Context) {
					fmt.Printf("Enter Password: ")
					e := execute.Command{Cmd: "cntlm", Args: []string{"-H"}}
					output, err := execute.RunCommand(e)
					if err != nil {
						log.Fatal(err)
					}
					cntlm.UpdateFile(cntlmFile, output)
					fmt.Println("Set CNTLM password successfully")
				},
			},
			{
				Name:        "domain",
				Usage:       "proxy set domain",
				Description: "additional description?",
				Action: func(_ *cli.Context) {
					fmt.Printf("Enter Proxy Domain: ")
					output, err := prompt.GetInput()
					if err != nil {
						log.Fatal(err)
					}
					update := fmt.Sprintln("Domain\t", output)
					cntlm.UpdateFile(cntlmFile, update)
					fmt.Println("Set CNTLM domain successfully")
				},
			},
		},
	}
}

func setProxyURL(port int) string {
	return fmt.Sprintf("http://localhost:%d", port)
}
