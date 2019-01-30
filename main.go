package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/go-proxy/cmd"
	"github.com/urfave/cli"
)

type execCommand struct {
	cmd  string
	args []string
}

func executeCommand(e execCommand) {
	cmd := e.cmd
	args := e.args
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

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
