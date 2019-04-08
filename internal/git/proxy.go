package gitproxy

import (
	"fmt"

	"github.com/xUnholy/go-proxy/pkg/execute"
)

func EnableProxyConfiguration(port string) error {
	http := execute.Command{Cmd: "git", Args: []string{"config", "--global", "http.proxy", port}}
	https := execute.Command{Cmd: "git", Args: []string{"config", "--global", "https.proxy", port}}
	r := execute.RunCommand(http)
	if r.Err != nil {
		return fmt.Errorf("failed to enable git http command %q", r.Err)
	}
	r = execute.RunCommand(https)
	if r.Err != nil {
		return fmt.Errorf("failed to enable git https command %q", r.Err)
	}
	return nil
}

func DisableProxyConfiguration() error {
	http := execute.Command{Cmd: "git", Args: []string{"config", "--global", "--unset", "http.proxy"}}
	https := execute.Command{Cmd: "git", Args: []string{"config", "--global", "--unset", "https.proxy"}}
	r := execute.RunCommand(http)
	if r.Err != nil {
		return fmt.Errorf("failed to enable git http command %q", r.Err)
	}
	r = execute.RunCommand(https)
	if r.Err != nil {
		return fmt.Errorf("failed to enable git https command %q", r.Err)
	}
	return nil
}
