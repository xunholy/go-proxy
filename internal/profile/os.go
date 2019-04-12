package profile

import (
	"fmt"
	"runtime"
)

// TODO: Locate windows drive prefix to defaultWindowsPath [go-proxy/#69]
const (
	defaultLinuxPath   = "/etc/cntlm.conf"
	defaultWindowsPath = "\\Program Files\\Cntlm\\cntlm.ini"
	defaultMacOSPath   = "/usr/local/etc/cntlm.conf"
)

func GetConfigurationPath() (string, error) {
	if runtime.GOOS == "linux" {
		return defaultLinuxPath, nil
	}
	if runtime.GOOS == "windows" {
		return defaultWindowsPath, nil
	}
	if runtime.GOOS == "darwin" {
		return defaultMacOSPath, nil
	}
	return "", fmt.Errorf("unsupported OS distribution")
}
