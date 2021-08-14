package client

import "net"

// Client is an interface for various way to get an IP address for an device (client).
// This way it is possible to add new techniques to get an IP address from without changing the main code.
type Client interface {
	// Init is using empty interface because I don't know any other way to allow different clients to use their own configurations.
	Init(params map[string]string) (err error)
	GetIPv4() net.IP
}

// ClientTypes is a lookup table for avaiable types of clients, and which stuct they refer to.
// New clients must be added to the init function, as it is used both to verify the configuration file but also to return the correct structs in the main program loop.
var ClientTypes = map[string]Client{}

func init() {
	ClientTypes["AtlasProbe"] = &AtlasProbe{}
	ClientTypes["IP"] = &IP{}
}
