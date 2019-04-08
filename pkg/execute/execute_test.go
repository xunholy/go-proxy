package execute_test

import (
	"fmt"
	"os"
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
	i, _ := strconv.Atoi(os.Getenv("EXIT_STATUS"))
	os.Exit(i)
}

func TestRunCommand(t *testing.T) {

	tests := []struct {
		expected bool
		command  execute.TestCommand
	}{
		{expected: false, command: execute.TestCommand{Cmd: "echo", Args: []string{"Hello"}, ExitCode: 0}},
		{expected: true, command: execute.TestCommand{Cmd: "echo", Args: []string{"Hello"}, ExitCode: 1}},
	}

	for _, i := range tests {
		r := execute.RunCommand(i.command)
		assert.Equal(t, i.expected, r.Err != nil)
		if i.command.ExitCode == 0 {
			assert.Equal(t, testResult, string(r.Output))
		}

	}
}
