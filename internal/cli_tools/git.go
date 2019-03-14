package gitproxy

import "github.com/xUnholy/go-proxy/pkg/execute"

func EnableProxyConfiguration(port string) []execute.Command {
	http := execute.Command{Cmd: "git", Args: []string{"config", "--global", "http.proxy", port}}
	https := execute.Command{Cmd: "git", Args: []string{"config", "--global", "https.proxy", port}}
	cmds := append([]execute.Command{}, http, https)
	return cmds
}

func DisableProxyConfiguration() []execute.Command {
	http := execute.Command{Cmd: "git", Args: []string{"config", "--global", "--unset", "http.proxy"}}
	https := execute.Command{Cmd: "git", Args: []string{"config", "--global", "--unset", "https.proxy"}}
	cmds := append([]execute.Command{}, http, https)
	return cmds
}
