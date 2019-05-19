package execute

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockedExitStatus = 0

const testResult = "World!"

func mockExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	es := strconv.Itoa(mockedExitStatus)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1", "EXIT_STATUS=" + es}
	return cmd
}

func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	fmt.Fprintf(os.Stdout, "%v", testResult)
	i, _ := strconv.Atoi(os.Getenv("EXIT_STATUS"))
	os.Exit(i)
}

func TestRunCommand(t *testing.T) {
	tests := []struct {
		expected bool
		exitCode int
		command  Command
	}{
		{expected: false, exitCode: 0, command: Command{Cmd: "echo", Args: []string{"Hello"}}},
		{expected: true, exitCode: 1, command: Command{Cmd: "echo", Args: []string{"Hello"}}},
	}
	execCommand = mockExecCommand
	defer func() { execCommand = exec.Command }()
	for _, i := range tests {
		mockedExitStatus = i.exitCode
		output, _, err := RunCommand(i.command)
		assert.Equal(t, i.expected, err != nil)
		if mockedExitStatus == 0 {
			assert.Equal(t, testResult, string(output))
		}
	}
}
