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

const testResult = "foo!"

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

func TestRunCommands(t *testing.T) {
	cmds := append([]Command{}, Command{Cmd: "npm", Args: []string{"config", "set", "proxy", "3128"}})

	tests := []struct {
		expected bool
		exitCode int
		commands []Command
	}{
		{expected: false, exitCode: 0, commands: cmds},
		{expected: true, exitCode: 1, commands: cmds},
	}

	execCommand = mockExecCommand
	defer func() { execCommand = exec.Command }()

	for _, i := range tests {
		mockedExitStatus = i.exitCode
		_, err := RunCommands(i.commands)
		assert.Equal(t, i.expected, err != nil)
	}

}

func TestRunCommand(t *testing.T) {
	cmd := Command{Cmd: "npm", Args: []string{"config", "set", "proxy", "3128"}}

	tests := []struct {
		expected bool
		exitCode int
		command  Command
	}{
		{expected: false, exitCode: 0, command: cmd},
		{expected: true, exitCode: 1, command: cmd},
	}

	execCommand = mockExecCommand
	defer func() { execCommand = exec.Command }()

	for _, i := range tests {
		mockedExitStatus = i.exitCode
		out, err := RunCommand(i.command)
		assert.Equal(t, i.expected, err != nil)
		if i.exitCode == 0 {
			assert.Equal(t, testResult, out)
		}

	}
}
