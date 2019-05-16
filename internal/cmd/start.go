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
	var cntlmProps = map[string]string{
		"Listen":     fmt.Sprint(cfg.Proxy.Port),
		"Username":   cfg.Proxy.Credentials.Username,
		"Domain":     cfg.Proxy.Domain,
		"PassLM":     cfg.Proxy.Credentials.PassLM,
		"PassNT":     cfg.Proxy.Credentials.PassNT,
		"PassNTLMv2": cfg.Proxy.Credentials.PassNTLMv2,
	}
	if err := cntlm.UpdateFile(cntlmProps); err != nil {
		return err
	}
	return nil
}
