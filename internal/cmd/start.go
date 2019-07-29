package cmd

import (
	"fmt"
	"log"
	"net/url"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xUnholy/go-proxy/pkg/execute"

	"github.com/xUnholy/go-proxy/internal/cntlm"
	"github.com/xUnholy/go-proxy/internal/config"
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
	cfg, err := config.LoadConfiguration(proxyProfile)
	if err != nil {
		log.Fatal(err)
	}
	if cfg.Proxy.Running {
		// TODO: A user may want to start another CNTLM instance, logic should handle if its intended or not
		log.Fatalf("CNTLM Proxy is already running")
	}
	proxyURL, err := url.Parse(fmt.Sprintf("http://%s:%v", cfg.Proxy.Address, port))
	if err != nil {
		log.Fatal(err)
	}
	err = env.UpdateGlobalEnvironmentVariables(proxyURL)
	if err != nil {
		log.Fatal(err)
	}
	cfg.Proxy.Port = port
	viper.Set("Proxy.Port", port)
	viper.Set("Proxy.Running", true)
	viper.Set("Proxy.ProxyURL", proxyURL)
	err = config.SaveConfiguration(proxyProfile)
	if err != nil {
		log.Fatal(err)
	}
	err = UpdateCNTLMConfig(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	cmds := execute.Command{Cmd: "cntlm", Args: []string{"-g"}}
	_, err = execute.RunCommand(cmds)
	if err != nil {
		log.Fatalf("CNTLM Proxy couldn't be started. Is it already running? %q", err)
	}
	fmt.Println("CNTLM Proxy Started On", proxyURL)
}

func UpdateCNTLMConfig(cfg *config.Configuration) error {
	update := fmt.Sprintf("Listen\t%v\n", cfg.Proxy.Port)
	update += fmt.Sprintf("Username\t%v\n", cfg.Proxy.Credentials.Username)
	update += fmt.Sprintf("Domain\t%v\n", cfg.Proxy.Domain)
	update += fmt.Sprintf("PassLM\t%v\n", cfg.Proxy.Credentials.PassLM)
	update += fmt.Sprintf("PassNT\t%v\n", cfg.Proxy.Credentials.PassNT)
	update += fmt.Sprintf("PassNTLMv2\t%v\n", cfg.Proxy.Credentials.PassNTLMv2)
	// TODO: Set the proxy address
	if err := cntlm.UpdateFile(update); err != nil {
		return err
	}
	return nil
}
