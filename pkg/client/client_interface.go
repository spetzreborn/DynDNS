package client

import "net"

// Client is an interface for various way to get an IP address for an device (client).
// This way it is possible to add new techniques to get an IP address from without changing the main code.
type Client interface {
	// Init is using empty interface because I don't know any other way to allow different clients to use their own configurations.
	Init(params interface{}) (err error)
	GetIPv4() net.IP
}
