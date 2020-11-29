package kittenclient

import (
	"fmt"
)

// Client is a KittenHouse client.
type Client struct {
	config    *ClientConfig
	transport *transport
}

// NewLogger returning logging client.
func NewLogger(config *ClientConfig) *Client {
	config.setup()

	client := Client{config: config, transport: &transport{config: config}}

	return &client
}

// Log is a function to log something to KittenHouse.
func (client *Client) Log(table string, format string, v ...interface{}) error {
	formatted := fmt.Sprintf(format, v...)

	var err error

	if client.config.UDP {
		err = client.transport.sendUDP(table, formatted)
	} else {
		err = client.transport.sendHttp(table, formatted)
	}
	if err != nil {
		return err
	}

	return nil
}
