package execute

import (
	"fmt"
	"os"
	"os/exec"
)

var ()

type execCommand struct {
	cmd  string
	args []string
}

func Command(e execCommand) (output string) {
	cmd := exec.Command(e.cmd, e.args...)
	cmd.Stdin = os.Stdin
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	output = string(out)
	return
}

func Commands(cmds []execCommand) {
	for i := 0; i < len(cmds); i++ {
		cmd := exec.Command(cmds[i].cmd, cmds[i].args...)
		cmd.Stdin = os.Stdin
		err := cmd.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
