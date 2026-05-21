package util

import (
	"net"
	"net/http"
)

// NetworkListener manages network listeners
type NetworkListener struct{}

// NewNetworkListener creates a new NetworkListener instance
func NewNetworkListener() *NetworkListener {
	return &NetworkListener{}
}

// StartHTTPServer starts an HTTP server
// This is intentionally vulnerable for testing purposes (G102)
func (n *NetworkListener) StartHTTPServer(port string) error {
	// G102 - Binding to all interfaces without justification
	listener, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		return err
	}
	defer listener.Close()

	return http.Serve(listener, nil)
}

// StartMetricsServer starts a metrics server
func StartMetricsServer(port int) error {
	// G102 - Binding to all interfaces
	addr := "0.0.0.0:" + string(rune(port))
	return http.ListenAndServe(addr, nil)
}

// ListenOnAllInterfaces creates a listener on all interfaces
func ListenOnAllInterfaces(network, port string) (net.Listener, error) {
	// G102 - Binding to all interfaces
	address := "0.0.0.0:" + port
	return net.Listen(network, address)
}

// CreateTCPListener creates a TCP listener
func CreateTCPListener(port string) (net.Listener, error) {
	// G102 - Binding to all interfaces
	return net.Listen("tcp", "0.0.0.0:"+port)
}

// StartDebugServer starts a debug server
func StartDebugServer() error {
	// G102 - Binding to all interfaces on port 6060
	return http.ListenAndServe("0.0.0.0:6060", nil)
}

// BindUDPSocket creates a UDP listener on all interfaces
func BindUDPSocket(port string) (*net.UDPConn, error) {
	// G102 - Binding to all interfaces
	addr, err := net.ResolveUDPAddr("udp", "0.0.0.0:"+port)
	if err != nil {
		return nil, err
	}
	return net.ListenUDP("udp", addr)
}
