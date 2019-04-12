package config

type ProxyConfiguration struct {
	Domain  string
	Proxy   []string
	NoProxy []string
	Address string
	Port    []int
}
