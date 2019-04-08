package npmproxy

import (
	"fmt"

	"github.com/xUnholy/go-proxy/pkg/execute"
)

func EnableProxyConfiguration(port string) error {
	http := execute.Command{Cmd: "npm", Args: []string{"config", "set", "proxy", port}}
	https := execute.Command{Cmd: "npm", Args: []string{"config", "set", "https-proxy", port}}
	r := execute.RunCommand(http)
	if r.Err != nil {
		return fmt.Errorf("failed to enable npm http command %q", r.Err)
	}
	r = execute.RunCommand(https)
	if r.Err != nil {
		return fmt.Errorf("failed to enable npm https command %q", r.Err)
	}
	return nil
}

func DisableProxyConfiguration() error {
	http := execute.Command{Cmd: "npm", Args: []string{"config", "delete", "proxy"}}
	https := execute.Command{Cmd: "npm", Args: []string{"config", "delete", "https-proxy"}}
	r := execute.RunCommand(http)
	if r.Err != nil {
		return fmt.Errorf("failed to disable npm http command %q", r.Err)
	}
	r = execute.RunCommand(https)
	if r.Err != nil {
		return fmt.Errorf("failed to disable npm https command %q", r.Err)
	}
	return nil
}
