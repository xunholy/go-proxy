package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "proxy",
	Short:   "Executing and configuring cntlm",
	Version: "0.2.0",
}

func Execute() {
	rootCmd.AddCommand(SetupSetCli(), SetupUnsetCli(), SetupStartCli(), SetupStopCli(), SetupConfigCli())
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
