package config

import "net/url"

type Configuration struct {
	Proxy ProxyConfiguration
}

type ProxyConfiguration struct {
	Pid          int
	Domain       string
	ProxyAddress []string
	ProxyURL     *url.URL
	NoProxy      []string
	Address      string
	Port         int
	Credentials  CredentialConfiguration
	Tools        ToolCLIProxyState
}

type CredentialConfiguration struct {
	Username   string
	Password   string
	PassLM     string
	PassNT     string
	PassNTLMv2 string
}

type ToolCLIProxyState struct {
	Git bool
	Npm bool
}
