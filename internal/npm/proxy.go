package npmproxy

import (
	"fmt"

	"github.com/xUnholy/go-proxy/pkg/execute"
)

func EnableProxyConfiguration(port string) error {
	http := execute.Command{Cmd: "npm", Args: []string{"config", "set", "proxy", port}}
	https := execute.Command{Cmd: "npm", Args: []string{"config", "set", "https-proxy", port}}
	_, err := execute.RunCommand(http)
	if err != nil {
		return fmt.Errorf("failed to enable npm http command %q", err)
	}
	_, err = execute.RunCommand(https)
	if err != nil {
		return fmt.Errorf("failed to enable npm https command %q", err)
	}
	return nil
}

func DisableProxyConfiguration() error {
	http := execute.Command{Cmd: "npm", Args: []string{"config", "delete", "proxy"}}
	https := execute.Command{Cmd: "npm", Args: []string{"config", "delete", "https-proxy"}}
	_, err := execute.RunCommand(http)
	if err != nil {
		return fmt.Errorf("failed to disable npm http command %q", err)
	}
	_, err = execute.RunCommand(https)
	if err != nil {
		return fmt.Errorf("failed to disable npm https command %q", err)
	}
	return nil
}
