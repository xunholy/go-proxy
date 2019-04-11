package gitproxy

import (
	"fmt"

	"github.com/xUnholy/go-proxy/pkg/execute"
)

func EnableProxyConfiguration(port string) error {
	http := execute.Command{Cmd: "git", Args: []string{"config", "--global", "http.proxy", port}}
	https := execute.Command{Cmd: "git", Args: []string{"config", "--global", "https.proxy", port}}
	_, err := execute.RunCommand(http)
	if err != nil {
		return fmt.Errorf("failed to enable git http command %q", err)
	}
	_, err = execute.RunCommand(https)
	if err != nil {
		return fmt.Errorf("failed to enable git https command %q", err)
	}
	return nil
}

func DisableProxyConfiguration() error {
	http := execute.Command{Cmd: "git", Args: []string{"config", "--global", "--unset", "http.proxy"}}
	https := execute.Command{Cmd: "git", Args: []string{"config", "--global", "--unset", "https.proxy"}}
	_, err := execute.RunCommand(http)
	if err != nil {
		return fmt.Errorf("failed to disable git http command %q", err)
	}
	_, err = execute.RunCommand(https)
	if err != nil {
		return fmt.Errorf("failed to disable git https command %q", err)
	}
	return nil
}
