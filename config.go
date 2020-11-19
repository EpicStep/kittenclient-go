package kittenclient

type ClientConfig struct {
	KittenURL string
	RowBinary bool
	UDP bool
	Table string
}

func (config *ClientConfig) setup() {
	if config.KittenURL == "" {
		config.KittenURL = "http://localhost:13338"
	}
}

