package config

type Configuration struct {
	ProxyConfig ProxyConfiguration
}

type ProxyConfiguration struct {
	Running      bool
	Domain       string
	ProxyAddress []string
	NoProxy      []string
	Address      string
	Port         []int
	Credentials  CredentialConfiguration
}

type CredentialConfiguration struct {
	Username   string
	Password   string
	PassLM     string
	PassNT     string
	PassNTLMv2 string
}
