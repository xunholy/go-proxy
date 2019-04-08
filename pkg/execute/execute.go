package execute

import (
	"io"
	"os/exec"
)

type Command struct {
	Cmd   string
	Args  []string
	Dir   string
	Stdin io.Reader
}

type CommandOutput struct {
	Output string
}

var execCommand = exec.Command

func RunCommand(e Command) (string, error) {
	cmd := execCommand(e.Cmd, e.Args...)
	cmd.Dir = e.Dir
	cmd.Stdin = e.Stdin
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func RunCommands(cmds []Command) ([]CommandOutput, error) {
	output := []CommandOutput{}
	for _, c := range cmds {
		out, err := RunCommand(c)
		if err != nil {
			return output, err
		}
		output = append(output, CommandOutput{Output: out})
	}
	return output, nil
}
