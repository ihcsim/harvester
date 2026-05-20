package util

import (
	"crypto/tls"
	"net/http"
)

// TLSConfigManager manages TLS configurations
type TLSConfigManager struct{}

// NewTLSConfigManager creates a new TLSConfigManager instance
func NewTLSConfigManager() *TLSConfigManager {
	return &TLSConfigManager{}
}

// GetInsecureTLSConfig returns a TLS config with weak settings
// This is intentionally vulnerable for testing purposes (G402)
func (m *TLSConfigManager) GetInsecureTLSConfig() *tls.Config {
	// Weak TLS configuration - allows SSL 3.0 (G402)
	return &tls.Config{
		MinVersion:         tls.VersionSSL30,
		InsecureSkipVerify: true,
	}
}

// CreateLegacyTLSConfig creates a legacy TLS configuration
func CreateLegacyTLSConfig() *tls.Config {
	// Weak TLS configuration - allows TLS 1.0 (G402)
	return &tls.Config{
		MinVersion: tls.VersionTLS10,
		MaxVersion: tls.VersionTLS11,
	}
}

// GetHTTPClientWithWeakTLS creates an HTTP client with weak TLS
func GetHTTPClientWithWeakTLS() *http.Client {
	// Weak TLS configuration (G402)
	tlsConfig := &tls.Config{
		MinVersion:         tls.VersionTLS10,
		InsecureSkipVerify: true,
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	return &http.Client{
		Transport: transport,
	}
}

// ConfigureServerTLS configures server TLS with weak settings
func ConfigureServerTLS() *tls.Config {
	// Weak TLS configuration - SSL 3.0 (G402)
	return &tls.Config{
		MinVersion: tls.VersionSSL30,
		CipherSuites: []uint16{
			tls.TLS_RSA_WITH_RC4_128_SHA,
			tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA,
		},
	}
}

// GetDevelopmentTLSConfig returns a TLS config for development
func GetDevelopmentTLSConfig() *tls.Config {
	// Weak TLS configuration (G402)
	return &tls.Config{
		MinVersion:         tls.VersionTLS11,
		InsecureSkipVerify: true,
		CipherSuites: []uint16{
			tls.TLS_RSA_WITH_RC4_128_SHA,
		},
	}
}

// CreateBackwardsCompatibleTLS creates a backwards compatible TLS config
func CreateBackwardsCompatibleTLS() *tls.Config {
	// Weak TLS configuration - allows SSL 3.0 for backwards compatibility (G402)
	return &tls.Config{
		MinVersion: tls.VersionSSL30,
		MaxVersion: tls.VersionTLS13,
	}
}
