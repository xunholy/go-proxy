package environment

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
)

var proxyVariables = []string{
	"ALL_PROXY",
	"HTTP_PROXY",
	"HTTPS_PROXY",
	"FTP_PROXY",
	"http_proxy",
	"https_proxy",
	"ftp_proxy",
}

func UpdateGlobalEnvironmentVariables(proxyURL *url.URL) error {
	dirPath := os.Getenv("HOME")
	filename := fmt.Sprintf("%v/.proxyrc", dirPath)
	data := []byte{}
	for _, v := range proxyVariables {
		a := fmt.Sprintf("export %v=%v\n", v, proxyURL)
		data = append(data, a...)
	}
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	output := `To set global proxy Environment variables you must do one of the following:

	Run:
		"source %v"

	OR
		Open new terminal

`
	fmt.Printf(output, filename)
	return nil
}
