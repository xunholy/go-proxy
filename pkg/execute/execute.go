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

func Command(e NewCommand) (output string) {
	cmd := exec.Command(e.Cmd, e.Args...)
	cmd.Stdin = os.Stdin
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	output = string(out)
	return
}

func Commands(cmds []NewCommand) {
	for i := 0; i < len(cmds); i++ {
		cmd := exec.Command(cmds[i].Cmd, cmds[i].Args...)
		cmd.Stdin = os.Stdin
		err := cmd.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
