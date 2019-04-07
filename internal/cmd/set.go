package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/xUnholy/go-proxy/pkg/execute"
	"github.com/xUnholy/go-proxy/pkg/prompt"

	"github.com/xUnholy/go-proxy/internal/cntlm"
	git "github.com/xUnholy/go-proxy/internal/git"
	npm "github.com/xUnholy/go-proxy/internal/npm"
)

var (
	cntlmFile = "/usr/local/etc/cntlm.conf"
	port      int
)

func SetupSetCli() *cobra.Command {
	// Top-level command
	var setCmd = &cobra.Command{
		Use:   "proxy set",
		Short: "Set CNTLM Proxy Config",
	}

	var setNpmCmd = &cobra.Command{
		Use:   "set npm proxy config",
		Short: "This command will set the NPM proxy values. Both https-proxy and proxy will be set",
		Run:   setNpmCmd,
	}
	setNpmCmd.Flags().IntVar(&port, "port", 3128, "set custom CNTLM `PORT`")

	var setGitCmd = &cobra.Command{
		Use:   "set git proxy config",
		Short: "This command will set the GIT global proxy values. Both http.proxy and https.proxy will be set",
		Run:   setGitCmd,
	}
	setGitCmd.Flags().IntVar(&port, "port", 3128, "set custom CNTLM `PORT`")

	var setUsernameCmd = &cobra.Command{
		Use:   "proxy set username",
		Short: "This command will update the Username value in your CNTLM.conf file",
		Run:   setUsernameCmd,
	}

	var setPasswordCmd = &cobra.Command{
		Use:   "proxy set password",
		Short: "This command will update the Password value in your CNTLM.conf file",
		Run:   setPasswordCmd,
	}

	var setDomainCmd = &cobra.Command{
		Use:   "proxy set domain",
		Short: "This command will update the domain value in your CNTLM.conf file",
		Run:   setDomainCmd,
	}

	// add subcommands to set
	setCmd.AddCommand(setNpmCmd, setGitCmd, setUsernameCmd, setPasswordCmd, setDomainCmd)
	return setCmd
}

func setNpmCmd(cmd *cobra.Command, args []string) {
	p := makeProxyURL(port)
	cmds := npm.EnableProxyConfiguration(p)
	_, err := execute.RunCommands(cmds)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Set npm config successfully")
}

func setGitCmd(cmd *cobra.Command, args []string) {
	p := makeProxyURL(port)
	cmds := git.EnableProxyConfiguration(p)
	_, err := execute.RunCommands(cmds)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Set git config successfully")
}

func setUsernameCmd(cmd *cobra.Command, args []string) {
	fmt.Printf("Enter Username: ")
	output, err := prompt.GetInput()
	if err != nil {
		log.Fatal(err)
	}
	update := fmt.Sprintln("Username\t", output)
	cntlm.UpdateFile(cntlmFile, update)
	fmt.Println("Set CNTLM username successfully")
}

func setPasswordCmd(cmd *cobra.Command, args []string) {
	// TODO: changing the password should restart cntlm to re-auth [go-proxy/#47]
	fmt.Printf("Enter Password: ")
	e := execute.Command{Cmd: "cntlm", Args: []string{"-H"}, Stdin: os.Stdin}
	output, err := execute.RunCommand(e)
	if err != nil {
		log.Fatal(err)
	}
	cntlm.UpdateFile(cntlmFile, output)
	fmt.Println("Set CNTLM password successfully")
}

func setDomainCmd(cmd *cobra.Command, args []string) {
	fmt.Printf("Enter Proxy Domain: ")
	output, err := prompt.GetInput()
	if err != nil {
		log.Fatal(err)
	}
	update := fmt.Sprintln("Domain\t", output)
	cntlm.UpdateFile(cntlmFile, update)
	fmt.Println("Set CNTLM domain successfully")
}
