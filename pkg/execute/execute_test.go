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

// func TestHelperProcess(t *testing.T) {
// 	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
// 		return
// 	}
// 	defer os.Exit(0)
// 	args := os.Args
// 	for len(args) > 0 {
// 		if args[0] == "--" {
// 			args = args[1:]
// 			break
// 		}
// 		args = args[1:]
// 	}
// 	if len(args) == 0 {
// 		fmt.Fprintf(os.Stderr, "No command\n")
// 		os.Exit(2)
// 	}
// 	switch os.Getenv("TEST_CASE") {
// 	case "case1":
// 		t.Fatal()
// 	}
// }

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
