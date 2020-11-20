package kittenclient

import (
	"fmt"
)

type Client struct {
	config    *ClientConfig
	transport *transport
}

func NewLogger(config *ClientConfig) *Client {
	config.setup()

	client := Client{config: config, transport: &transport{config: config}}

	return &client
}

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

