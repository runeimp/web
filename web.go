package web

import (
	"net"
	"net/http"
	"time"
)

// Configuration defines the settings for new *http.Clients
type Configuration struct {
	ClientTimeout       time.Duration // Timeout in seconds
	DialerTimeout       time.Duration // Timeout in seconds
	HandshakeTimeout    time.Duration // Timeout in seconds
	MaxConnsPerHost     int
	MaxIdleConns        int
	MaxIdleConnsPerHost int
}

// Version defines the library version
const Version = "1.0.0"

// Config is the current settings configuration for the web package
var Config Configuration

func init() {
	Config = Configuration{
		ClientTimeout:       10, // Timeout in seconds
		DialerTimeout:       5,  // Timeout in seconds
		HandshakeTimeout:    5,  // Timeout in seconds
		MaxConnsPerHost:     100,
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
	}
}

// NewClient creates a new http.Client based on this packages current Web*
func NewClient() *http.Client {
	return &http.Client{
		Timeout: Config.ClientTimeout * time.Second,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: Config.DialerTimeout * time.Second,
			}).Dial,
			MaxIdleConns:        Config.MaxIdleConns,
			MaxConnsPerHost:     Config.MaxConnsPerHost,
			MaxIdleConnsPerHost: Config.MaxIdleConnsPerHost,
			TLSHandshakeTimeout: Config.HandshakeTimeout * time.Second,
		},
	}
}
