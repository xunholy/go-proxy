package npmproxy

import "github.com/xUnholy/go-proxy/pkg/execute"

func EnableProxyConfiguration(port string) []execute.Command {
	http := execute.Command{Cmd: "npm", Args: []string{"config", "set", "proxy", port}}
	https := execute.Command{Cmd: "npm", Args: []string{"config", "set", "https-proxy", port}}
	cmds := append([]execute.Command{}, http, https)
	return cmds
}

func DisableProxyConfiguration() []execute.Command {
	http := execute.Command{Cmd: "npm", Args: []string{"config", "delete", "proxy"}}
	https := execute.Command{Cmd: "npm", Args: []string{"config", "delete", "https-proxy"}}
	cmds := append([]execute.Command{}, http, https)
	return cmds
}
