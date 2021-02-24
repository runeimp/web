package web

import (
	"net"
	"net/http"
	"testing"
	"time"
)

func TestWebClient(t *testing.T) {
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).Dial,
			MaxIdleConns:        100,
			MaxConnsPerHost:     100,
			MaxIdleConnsPerHost: 100,
			TLSHandshakeTimeout: 5 * time.Second,
		},
	}
	webClient := NewClient()

	if httpClient.Timeout != webClient.Timeout {
		t.Fatalf("expected: %v, got: %v", httpClient, webClient)
	}
	// Untestable? Dialer Timeout is inaccessable in the Dial function?
	// if httpClient.Transport.(*http.Transport).DialContext.(*net.Dialer).Timeout != webClient.Transport.(*http.Transport).DialContext.(*net.Dialer).Timeout {
	// 	t.Fatalf("expected: %v, got: %v", httpClient, webClient)
	// }
	if httpClient.Transport.(*http.Transport).TLSHandshakeTimeout != webClient.Transport.(*http.Transport).TLSHandshakeTimeout {
		t.Fatalf("expected: %v, got: %v", httpClient, webClient)
	}
	if httpClient.Transport.(*http.Transport).MaxConnsPerHost != webClient.Transport.(*http.Transport).MaxConnsPerHost {
		t.Fatalf("expected: %v, got: %v", httpClient, webClient)
	}
	if httpClient.Transport.(*http.Transport).MaxIdleConns != webClient.Transport.(*http.Transport).MaxIdleConns {
		t.Fatalf("expected: %v, got: %v", httpClient, webClient)
	}
	if httpClient.Transport.(*http.Transport).MaxIdleConnsPerHost != webClient.Transport.(*http.Transport).MaxIdleConnsPerHost {
		t.Fatalf("expected: %v, got: %v", httpClient, webClient)
	}
}

func TestWebDefaults(t *testing.T) {
	if Config.ClientTimeout != 10 {
		t.Fatalf("expected: %d, got: %v", 10, Config.ClientTimeout)
	}
	if Config.DialerTimeout != 5 {
		t.Fatalf("expected: %d, got: %v", 5, Config.DialerTimeout)
	}
	if Config.HandshakeTimeout != 5 {
		t.Fatalf("expected: %d, got: %v", 5, Config.HandshakeTimeout)
	}
	if Config.MaxConnsPerHost != 100 {
		t.Fatalf("expected: %d, got: %v", 100, Config.MaxConnsPerHost)
	}
	if Config.MaxIdleConns != 100 {
		t.Fatalf("expected: %d, got: %v", 100, Config.MaxIdleConns)
	}
	if Config.MaxIdleConnsPerHost != 100 {
		t.Fatalf("expected: %d, got: %v", 100, Config.MaxIdleConnsPerHost)
	}
}
