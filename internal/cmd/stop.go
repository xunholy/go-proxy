package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/xUnholy/go-proxy/pkg/execute"

	"github.com/xUnholy/go-proxy/internal/profile"
)

func SetupStopCli() *cobra.Command {
	// Top-level command
	var stopCmd = &cobra.Command{
		Use:   "stop",
		Short: "Stop CNTLM Proxy",
		Run:   stopCmd,
	}
	return stopCmd
}

func stopCmd(cmd *cobra.Command, args []string) {
	profile.UpdateGlobalEnvironmentVariables("")
	cmds := execute.Command{Cmd: "pkill", Args: []string{"cntlm"}}
	r := execute.RunCommand(cmds)
	if r.Err != nil {
		log.Fatal("Couldn't kill CNTLM process %q", r.Err)
	}
	fmt.Println("CNTLM proxy stopped")
}
