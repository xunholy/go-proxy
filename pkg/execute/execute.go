package execute

import (
	"os"
	"os/exec"
)

type Command struct {
	Cmd  string
	Args []string
}

type CommandOutput struct {
	Output string
}

func RunCommand(e Command) (string, error) {
	cmd := exec.Command(e.Cmd, e.Args...)
	cmd.Stdin = os.Stdin
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func RunCommands(cmds []Command) ([]CommandOutput, error) {
	output := []CommandOutput{}
	for _, c := range cmds {
		cmd := Command{Cmd: c.Cmd, Args: c.Args}
		out, err := RunCommand(cmd)
		if err != nil {
			return output, err
		}
		output = append(output, CommandOutput{Output: out})
	}
	return output, nil
}
