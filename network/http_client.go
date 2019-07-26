package network

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

func NewClient(skipTLSVerification bool) *http.Client {
	return &http.Client{
		Timeout: 1 * time.Second,
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: skipTLSVerification,
				MinVersion:         tls.VersionTLS12,
			},
			DialContext: (&net.Dialer{
				Timeout:   5 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
		},
	}
}
