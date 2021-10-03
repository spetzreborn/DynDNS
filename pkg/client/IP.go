package client

/*

Return given IP from string as net.IP

*/

import (
	"fmt"
	"net"
)

// IP hold an IPv4 address
type IP struct {
	addressV4 net.IP
}

// Init initializes the probe with a given IPv4 addess in string format
func (i *IP) Init(param string) error {
	ipaddress := net.ParseIP(param)
	if ipaddress == nil {
		return fmt.Errorf("%s is not an valid IP address", param)
	}
	if ipaddress.To4() == nil {
		return fmt.Errorf("%s is not an valid IPv4 address", param)
	}
	i.addressV4 = ipaddress
	return nil
}

// GetIPv4 returns the IPv4 address for the probe
func (i IP) GetIPv4() *net.IP {
	return &i.addressV4
}
