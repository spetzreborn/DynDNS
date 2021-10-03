package client

import "net"

// Client is an interface for various way to get an IP address for an device (client).
// This way it is possible to add new techniques to get an IP address from without changing the main code.
type IClient interface {
	// Init is using empty interface because I don't know any other way to allow different clients to use their own configurations.
	Init(param string) error
	GetIPv4() net.IP
}

type Factory struct {
	client IClient
}

func New(c string) *Factory {
	return &Factory{client: ClientTypes[c]}
}

// ClientTypes is a lookup table for avaiable types of clients, and which stuct they refer to.
// New clients must be added to the init function, as it is used both to verify the configuration file but also to return the correct structs in the main program loop.
var ClientTypes = map[string]IClient{}

func init() {
	ClientTypes["AtlasProbe"] = &AtlasClient{}
	ClientTypes["IP"] = &IP{}
}

func (c *Factory) Init(params string) error {
	err := c.client.Init(params)
	if err != nil {
		return err
	}
	return nil
}

func (c *Factory) GetIPv4() net.IP {
	return c.client.GetIPv4()
}
