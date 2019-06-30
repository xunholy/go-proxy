package os

import "fmt"

// TODO: Locate windows drive prefix to defaultWindowsPath [go-proxy/#69]
const (
	defaultLinuxPath   = "/etc/cntlm.conf"
	defaultWindowsPath = "\\Program Files\\Cntlm\\cntlm.ini"
	defaultMacOSPath   = "/usr/local/etc/cntlm.conf"
)

func GetConfigurationPath(runtimeOS string) (string, error) {
	if runtimeOS == "linux" {
		return defaultLinuxPath, nil
	}
	if runtimeOS == "windows" {
		return defaultWindowsPath, nil
	}
	if runtimeOS == "darwin" {
		return defaultMacOSPath, nil
	}
	return "", fmt.Errorf("unsupported OS distribution")
}
