package config

import (
	tools "github.com/xUnholy/go-proxy/internal/tools"
)

func UnsetAllConfiguration(cfg *Configuration) error {
	// TODO: Consider a better way to loop through Tools as the list of supported CLI's may grow
	if !cfg.Proxy.Tools.Git {
		if err := tools.DisableGITProxyConfiguration(); err != nil {
			return err
		}
	}
	if !cfg.Proxy.Tools.Npm {
		if err := tools.DisableNPMProxyConfiguration(); err != nil {
			return err
		}
	}
	return nil
}
