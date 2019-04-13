package config

type ProxyConfiguration struct {
	Domain       string
	ProxyAddress []string
	NoProxy      []string
	Address      string
	Port         []int
	Credentials  CredentialConfiguration
}
