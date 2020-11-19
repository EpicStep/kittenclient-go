package kittenclient

import (
	"io"
	"net/http"
)

type transport struct {
	http *http.Client
	//udp  *net.Conn
}

//func newTransport(config *ClientConfig) *transport {
//	if !config.UDP {
//		transport := transport{http: http.Client{}, }
//	}
//}

func (transport *transport) sendHttp(url string, contentType string, body io.Reader) error {
	_, err := transport.http.Post(url, contentType, body)
	if err != nil {
		return err
	}

	return nil
}

//func (transport *transport) sendUDP() {
//
//}


