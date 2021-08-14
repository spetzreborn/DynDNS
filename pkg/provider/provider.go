package provider

import "net"

// Provider is an interface for various way to get and set DNS A records from an DNS provider.
// This way it is possible to add new providers without changing the main code.
type Provider interface {
	// Init is using empty interface because I don't know any other way to allow different provider to use their own configurations.
	Init(params map[string]string) (err error)
	GetARecord(string) (net.IP, error)
	SetARecord(string, net.IP) (err error)
}

// ProviderTypes is a lookup table for avaiable types of providers, and which stuct they refer to.
// New provider must be added to the init function, as it is used both to verify the configuration file but also to return the correct structs in the main program loop.
var ProviderTypes = map[string]Provider{}

func init() {
	ProviderTypes["Fake"] = &Fake{}
	ProviderTypes["Glesys"] = &Glesys{}
}
