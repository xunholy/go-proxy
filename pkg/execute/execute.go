package execute

import (
	"io"
	"os/exec"
)

type Commander interface {
	ExecuteCommand() ([]byte, error)
}

type Command struct {
	Cmd   string
	Args  []string
	Dir   string
	Stdin io.Reader
}

func (c Command) ExecuteCommand() ([]byte, error) {
	cmd := exec.Command(c.Cmd, c.Args...)
	cmd.Dir = c.Dir
	cmd.Stdin = c.Stdin
	out, err := cmd.CombinedOutput()
	return out, err
}

func RunCommand(c Commander) ([]byte, error) {
	return c.ExecuteCommand()
}
