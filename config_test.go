package kittenclient_test

import (
	"github.com/EpicStep/kittenclient-go"
	"testing"
)

func TestConfig(t *testing.T) {
	t.Parallel()

	config := kittenclient.ClientConfig{
		Addr: "",
		UDP:  false,
	}

	_ = kittenclient.NewLogger(&config)

	if config.Addr != "localhost:13338" {
		t.Error("Expected localhost:13338, got", config.Addr)
	}
}
