package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
	"github.com/xUnholy/go-proxy/cmd"
)

func main() {

	app := cli.NewApp()
	app.Name = "proxy"
	app.Version = "0.0.1"
	app.Usage = "executing and configuring cntlm"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Michael Fornaro",
			Email: "Michael.Fornaro@anz.com",
		},
	}

	app.Commands = []cli.Command{
		cmd.StartCommand(),
		cmd.StopCommand(),
		cmd.SetCommand(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
