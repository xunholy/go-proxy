package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

var (
	port        int
	setAll      bool
	password    string
	cntlmFile   = "/usr/local/etc/cntlm.conf"
	bashProfile = "~/.bash_profile"
)

type execCommand struct {
	cmd  string
	args []string
}

func executeCommand(e execCommand) (output string) {
	cmd := exec.Command(e.cmd, e.args...)
	cmd.Stdin = os.Stdin
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	output = string(out)
	return output
}
