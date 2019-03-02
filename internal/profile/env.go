package profile

import (
	"fmt"
	"io/ioutil"
	"log"
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

func UpdateGlobalEnvironmentVariables(proxyURL string) {
	dirPath := fmt.Sprintf("%v/.proxy", os.Getenv("HOME"))
	file := fmt.Sprintf("%v/proxy.sh", dirPath)
	content := []byte{}
	for _, v := range proxyVariables {
		a := fmt.Sprintf("export %v=%v\n", v, proxyURL)
		content = append(content, a...)
	}
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		log.Fatalf("Dir Not Found: %v", dirPath)
	}
	err := ioutil.WriteFile(file, content, 0644)
	if err != nil {
		log.Fatal(err)
	}
	output := `To set global proxy Environment variables you must do one of the following:

	Run:
		"source %v"

	OR
		Open new terminal

`
	fmt.Printf(output, file)
}
