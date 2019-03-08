package execute

import (
	"os"
	"os/exec"
	"testing"
)

var testCase string

func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	tc := "TEST_CASE=" + testCase
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1", tc}
	return cmd
}

func TestRunCommands(t *testing.T) {
	testCase = "case1"
	tests := []Command{}
	tests = append(tests, Command{Cmd: "npm", Args: []string{"config", "set", "proxy", "3128"}})
	execCommand = fakeExecCommand
	defer func() { execCommand = exec.Command }()
	out, err := RunCommands(tests)
	if err != nil {
		t.Fatal(out)
	}
}
