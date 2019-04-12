package config

type ProxyConfiguration struct {
	Username string
	Domain   string
	Proxy    string
	NoProxy  []string
	Port     int
}
