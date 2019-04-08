package execute

import (
	"io"
	"os"
	"os/exec"
	"strconv"
)

type Commander interface {
	ExecuteCommand() CmdOutput
}

type CmdOutput struct {
	Err    error
	Output []byte
	Cmd    *exec.Cmd
}

type Command struct {
	Cmd   string
	Args  []string
	Dir   string
	Stdin io.Reader
}

type TestCommand struct {
	Cmd      string
	Args     []string
	ExitCode int
}

func (c Command) ExecuteCommand() CmdOutput {
	cmd := exec.Command(c.Cmd, c.Args...)
	cmd.Dir = c.Dir
	cmd.Stdin = c.Stdin
	out, err := cmd.CombinedOutput()
	return CmdOutput{Output: out, Err: err}
}

func (tc TestCommand) ExecuteCommand() CmdOutput {
	cs := []string{"-test.run=TestHelperProcess", "--", tc.Cmd}
	cs = append(cs, tc.Args...)
	cmd := exec.Command(os.Args[0], cs...)
	es := strconv.Itoa(tc.ExitCode)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1", "EXIT_STATUS=" + es}
	return CmdOutput{Cmd: cmd}
}

func RunCommand(c Commander) CmdOutput {
	return c.ExecuteCommand()
}
