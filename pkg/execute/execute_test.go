package execute_test

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xUnholy/go-proxy/pkg/execute"
)

const testResult = "World!"

func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	fmt.Fprintf(os.Stdout, "%v", testResult)
	i, err := strconv.Atoi(os.Getenv("EXIT_STATUS"))
	if err != nil {
		os.Exit(1)
	}
	os.Exit(i)
}

type TestCommand struct {
	Cmd      string
	Args     []string
	ExitCode int
}

func (tc TestCommand) ExecuteCommand() *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", tc.Cmd}
	cs = append(cs, tc.Args...)
	cmd := exec.Command(os.Args[0], cs...)
	es := strconv.Itoa(tc.ExitCode)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1", "EXIT_STATUS=" + es}
	return cmd
}

func TestRunCommand(t *testing.T) {
	tests := []struct {
		expected bool
		command  TestCommand
	}{
		{expected: false, command: TestCommand{Cmd: "echo", Args: []string{"Hello"}, ExitCode: 0}},
		{expected: true, command: TestCommand{Cmd: "echo", Args: []string{"Hello"}, ExitCode: 1}},
	}
	for _, i := range tests {
		output, _, err := execute.RunCommand(i.command)
		assert.Equal(t, i.expected, err != nil)
		if i.command.ExitCode == 0 {
			assert.Equal(t, testResult, string(output))
		}
	}
}
