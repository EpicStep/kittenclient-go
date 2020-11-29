package kittenclient_test

import (
	"github.com/EpicStep/kittenclient-go"
	"os"
	"testing"
)

func TestNewLogger(t *testing.T) {
	t.Parallel()

	kittenAddr := os.Getenv("KITTENHOUSE_ADDR")

	config := kittenclient.ClientConfig{
		Addr: kittenAddr,
		UDP:  false,
	}

	logger := kittenclient.NewLogger(&config)
	err := logger.Log("internal_logs_buffer(time,server,port)", "('2020-11-21 12:20:00','testing', 8080)")
	if err != nil {
		t.Error(err)
	}
}
