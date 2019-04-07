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
		Use:   "proxy stop",
		Short: "Stop CNTLM Proxy",
		Run:   stopCmd,
	}
	return stopCmd
}

func stopCmd(cmd *cobra.Command, args []string) {
	profile.UpdateGlobalEnvironmentVariables("")
	cmds := execute.Command{Cmd: "pkill", Args: []string{"cntlm"}}
	_, err := execute.RunCommand(cmds)
	if err != nil {
		fmt.Println("Couldn't kill CNTLM process")
		log.Fatal(err)
	}
	fmt.Println("CNTLM proxy stopped")
}
