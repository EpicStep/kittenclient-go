package kittenclient

// ClientConfig save config to kittenhouse connections.
type ClientConfig struct {
	// Address of kittenhouse
	Addr string
	// UDP transport instead TCP
	UDP bool
}

func (config *ClientConfig) setup() {
	if config.Addr == "" {
		config.Addr = "localhost:13338"
	}
}
