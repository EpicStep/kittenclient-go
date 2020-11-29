package kittenclient

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"net"
	"net/http"
	"strings"
)

type transport struct {
	config *ClientConfig
}

// Send HTTP request to kittenhouse
func (transport *transport) sendHttp(table string, data string) error {
	url := fmt.Sprintf("http://%s/?table=%s", transport.config.Addr, table)

	_, err := http.Post(url, "application/json", strings.NewReader(data))

	if err != nil {
		return err
	}

	return nil
}

// sendUDP send packet in the following format:
// | table_name | \0 | data | crc32 (4 byte) |
func (transport *transport) sendUDP(table string, data string) error {
	conn, err := net.Dial("udp", transport.config.Addr)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer([]byte{})

	buf.Write([]byte(table))
	buf.WriteByte(0)
	buf.Write([]byte(data))

	hash := crc32.ChecksumIEEE(buf.Bytes())
	if err = binary.Write(buf, binary.LittleEndian, hash); err != nil {
		return err
	}

	if _, err = conn.Write(buf.Bytes()); err != nil {
		return err
	}

	if err = conn.Close(); err != nil {
		return err
	}

	return nil
}
