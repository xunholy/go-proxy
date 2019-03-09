package execute

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func mockExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}

func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	_, err := fmt.Fprintf(os.Stdout, testResult)
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}

const testResult = "foo!"

func TestRunCommands(t *testing.T) {
	tests := []Command{}
	tests = append(tests, Command{Cmd: "npm", Args: []string{"config", "set", "proxy", "3128"}})

	execCommand = mockExecCommand
	defer func() { execCommand = exec.Command }()

	_, err := RunCommands(tests)
	if err != nil {
		t.Errorf("Expected nil error, got %#v", err)
	}

}

func TestRunCommand(t *testing.T) {
	tests := []struct {
		command Command
	}{
		{command: Command{Cmd: "npm", Args: []string{"config", "set", "proxy", "3128"}}},
	}

	execCommand = mockExecCommand
	defer func() { execCommand = exec.Command }()

	for _, i := range tests {
		out, err := RunCommand(i.command)
		if err != nil {
			t.Errorf("Expected nil error, got %#v", err)
		}
		if out != testResult {
			t.Errorf("Expected %q, got %q", testResult, out)
		}
	}
}
