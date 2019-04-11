package execute

import (
	"io"
	"os/exec"
)

var execCommand = exec.Command

type Commander interface {
	ExecuteCommand() *exec.Cmd
}

type Command struct {
	Cmd   string
	Args  []string
	Dir   string
	Stdin io.Reader
}

func (c Command) ExecuteCommand() *exec.Cmd {
	cmd := execCommand(c.Cmd, c.Args...)
	cmd.Dir = c.Dir
	cmd.Stdin = c.Stdin
	return cmd
}

func RunCommand(c Commander) ([]byte, error) {
	cmd := c.ExecuteCommand()
	out, err := cmd.CombinedOutput()
	return out, err
}
