package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/golang/crypto/ssh/terminal"
)

var (
	port        int
	proxy       string
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

func credentials() string {
	fmt.Print("Enter Password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err == nil {
		fmt.Println("\nPassword typed: " + string(bytePassword))
	}
	password := string(bytePassword)
	return strings.TrimSpace(password)
}
