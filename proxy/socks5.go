package proxy

import (
	"context"
	"golang.org/x/net/proxy"
	"log"
	"net"
	"net/http"
)

// SOCKS5Client ...
func SOCKS5Client(socks5Proxy string) *http.Client {
	// socks5Proxy := "127.0.0.1:1080"
	dialer, err := proxy.SOCKS5("tcp", socks5Proxy, nil, proxy.Direct)
	if err != nil {
		log.Fatal("Error creating dialer, aborting.")
	}

	dialFunc := func(ctx context.Context, network, addr string) (net.Conn, error) {
		return dialer.Dial(network, addr)
	}

	tr := &http.Transport{DialContext: dialFunc} // Dial: dialer.Dial,
	httpClient := &http.Client{Transport: tr}
	return httpClient
}
