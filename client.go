package kittenclient

import (
	"fmt"
	"strings"
)

type Client struct {
	config *ClientConfig
	http   *transport
}

func NewLogger(config *ClientConfig) *Client {
	config.setup()

	client := Client{config: config}

	return &client
}

func (client *Client) Log(format string, v ...interface{}) error {
	formatted := fmt.Sprintf(format, v...)
	url := fmt.Sprintf("%s?table=%s", client.config.KittenURL, client.config.Table)

	err := client.http.sendHttp(url, "application/json", strings.NewReader(formatted))
	if err != nil {
		return err
	}

	return nil
}

