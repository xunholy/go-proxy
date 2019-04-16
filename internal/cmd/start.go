package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/xUnholy/go-proxy/pkg/execute"

	"github.com/xUnholy/go-proxy/internal/cntlm"
	env "github.com/xUnholy/go-proxy/internal/environment"
)

func SetupStartCli() *cobra.Command {
	// Top-level command
	var startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start CNTLM Proxy",
		Run:   startCmd,
	}
	startCmd.Flags().IntVar(&port, "port", 3128, "set custom CNTLM `PORT`")
	return startCmd
}

func startCmd(cmd *cobra.Command, args []string) {
	proxyURL := makeProxyURL(port)
	env.UpdateGlobalEnvironmentVariables(proxyURL)
	update := fmt.Sprintf("Listen\t%v", port)
	if err := cntlm.UpdateFile(update); err != nil {
		log.Fatal(err)
	}
	cmds := execute.Command{Cmd: "cntlm", Args: []string{"-g"}}
	_, err := execute.RunCommand(cmds)
	if err != nil {
		log.Fatalf("CNTLM Proxy couldn't be started. Is it already running? %q", err)
	}
	fmt.Println("CNTLM Proxy Started On", proxyURL)
}
