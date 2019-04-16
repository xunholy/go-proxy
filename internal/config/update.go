package config

import (
	"strings"

	tools "github.com/xUnholy/go-proxy/internal/tools"
)

func UnsetAllConfiguration(proxyProfile string) error {
	cfg, err := LoadConfiguration(proxyProfile)
	if err != nil {
		return err
	}
	for _, t := range cfg.Proxy.Tools {
		if strings.EqualFold(t, "git") {
			if err := tools.DisableGITProxyConfiguration(); err != nil {
				return err
			}
		}
		if strings.EqualFold(t, "npm") {
			if err := tools.DisableNPMProxyConfiguration(); err != nil {
				return err
			}
		}
	}
	return nil
}

// func SetAllConfiguration(proxyProfile string) error {
// 	cfg, err := LoadConfiguration(proxyProfile)
// 	if err != nil {
// 		return err
// 	}
// 	for _, t := range cfg.Proxy.Tools {
// 		if strings.ToLower(t) == "git" {
// 			tools.EnableGITProxyConfiguration()
// 		}
// 		if strings.ToLower(t) == "npm" {
// 			tools.EnableNPMProxyConfiguration()
// 		}
// 	}
// 	return nil
// }
