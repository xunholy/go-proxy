package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xUnholy/go-proxy/pkg/execute"
	"github.com/xUnholy/go-proxy/pkg/prompt"

	"github.com/xUnholy/go-proxy/internal/cntlm"
	"github.com/xUnholy/go-proxy/internal/config"
	tools "github.com/xUnholy/go-proxy/internal/tools"
)

var (
	port int
)

func SetupSetCli() *cobra.Command {
	// Top-level command
	var setCmd = &cobra.Command{
		Use:   "set",
		Short: "Set CNTLM Proxy Config",
	}

	var setNpmCmd = &cobra.Command{
		Use:   "npm",
		Short: "This command will set the NPM proxy values. Both https-proxy and proxy will be set",
		Run:   setNpmCmd,
	}

	var setGitCmd = &cobra.Command{
		Use:   "git",
		Short: "This command will set the GIT global proxy values. Both http.proxy and https.proxy will be set",
		Run:   setGitCmd,
	}

	var setUsernameCmd = &cobra.Command{
		Use:   "username",
		Short: "This command will update the Username value in your CNTLM.conf file",
		Run:   setUsernameCmd,
	}

	var setPasswordCmd = &cobra.Command{
		Use:   "password",
		Short: "This command will update the Password value in your CNTLM.conf file",
		Run:   setPasswordCmd,
	}

	var setDomainCmd = &cobra.Command{
		Use:   "domain",
		Short: "This command will update the domain value in your CNTLM.conf file",
		Run:   setDomainCmd,
	}

	// add subcommands to set
	setCmd.AddCommand(setNpmCmd, setGitCmd, setUsernameCmd, setPasswordCmd, setDomainCmd)
	return setCmd
}

func setNpmCmd(cmd *cobra.Command, args []string) {
	cfg, err := config.LoadConfiguration(proxyProfile)
	if err != nil {
		log.Fatal(err)
	}
	if err = tools.EnableNPMProxyConfiguration(cfg.Proxy.ProxyURL); err != nil {
		log.Fatalln(err)
	}
	viper.Set("Proxy.Tools.Npm", true)
	err = config.SaveConfiguration(proxyProfile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Set npm config successfully")
}

func setGitCmd(cmd *cobra.Command, args []string) {
	cfg, err := config.LoadConfiguration(proxyProfile)
	if err != nil {
		log.Fatal(err)
	}
	if err = tools.EnableGITProxyConfiguration(cfg.Proxy.ProxyURL); err != nil {
		log.Fatalln(err)
	}
	viper.Set("Proxy.Tools.Git", true)
	err = config.SaveConfiguration(proxyProfile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Set git config successfully")
}

// nolint
func setUsernameCmd(cmd *cobra.Command, args []string) {
	_, err := config.LoadConfiguration(proxyProfile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Enter Username: ")
	output, err := prompt.GetInput(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}
	update := fmt.Sprintln("Username\t", output)
	if err = cntlm.UpdateFile(update); err != nil {
		log.Fatal(err)
	}
	viper.Set("Proxy.Credentials.Username", output)
	err = config.SaveConfiguration(proxyProfile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Set CNTLM username successfully")
}

func setPasswordCmd(cmd *cobra.Command, args []string) {
	_, err := config.LoadConfiguration(proxyProfile)
	if err != nil {
		log.Fatal(err)
	}
	// TODO: changing the password should restart cntlm to re-auth [go-proxy/#47]
	fmt.Printf("Enter Password: ")
	// TODO: rewrite cntlm hash function as reusable pkg
	e := execute.Command{Cmd: "cntlm", Args: []string{"-H"}, Stdin: os.Stdin}
	out, err := execute.RunCommand(e)
	if err != nil {
		log.Fatalln(err)
	}
	if err = cntlm.UpdateFile(string(out)); err != nil {
		log.Fatal(err)
	}
	x := strings.Split(strings.TrimSpace(string(out)), "\n")
	for _, t := range x {
		fields := strings.Fields(t)
		if len(fields) >= 2 {
			viper.Set(fmt.Sprintf("Proxy.Credentials.%v", fields[0]), fields[1])
		}
	}
	err = config.SaveConfiguration(proxyProfile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Set CNTLM password successfully")
}

// nolint
func setDomainCmd(cmd *cobra.Command, args []string) {
	_, err := config.LoadConfiguration(proxyProfile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Enter Proxy Domain: ")
	output, err := prompt.GetInput(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}
	update := fmt.Sprintln("Domain\t", output)
	if err = cntlm.UpdateFile(update); err != nil {
		log.Fatal(err)
	}
	viper.Set("Proxy.Domain", output)
	err = config.SaveConfiguration(proxyProfile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Set CNTLM domain successfully")
}
