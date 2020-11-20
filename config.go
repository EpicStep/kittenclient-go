package kittenclient

// Config for logger
type ClientConfig struct {
	// Addres of kittenhouse
	Addr      string
	// UDP transoprt instand TCP
	UDP       bool
}

func (config *ClientConfig) setup() {
	if config.Addr == "" {
		config.Addr = "localhost:13338"
	}
}

