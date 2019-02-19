package execute

import (
	"fmt"
	"os"
	"os/exec"
)

type NewCommand struct {
	Cmd  string
	Args []string
}

type CommandsOutput struct {
	Cmd    string
	Output string
}

func Command(e NewCommand) (string, error) {
	cmd := exec.Command(e.Cmd, e.Args...)
	cmd.Stdin = os.Stdin
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return "", err
	}
	return string(out), nil
}

func Commands(cmds []NewCommand) ([]CommandsOutput, error) {
	output := []CommandsOutput{}
	for i := 0; i < len(cmds); i++ {
		cmd := NewCommand{Cmd: cmds[i].Cmd, Args: cmds[i].Args}
		out, err := Command(cmd)
		if err != nil {
			return output, err
		}
		output = append(output, CommandsOutput{Cmd: cmds[i].Cmd, Output: out})
	}
	return output, nil
}
