package client

/*

Return given IP from string as net.IP

*/

import (
	"errors"
	"fmt"
	"net"
)

// IP hold an IPv4 address
type IP struct {
	AddressV4 net.IP
}

// Init initializes the probe with a given IPv4 addess in string format
func (i *IP) Init(ip interface{}) (err error) {
	if _, ok := ip.(string); ok == false {
		return errors.New("argument must be a string")

	}
	ipaddress := net.ParseIP(fmt.Sprintf("%v", ip))
	if ipaddress == nil {
		return fmt.Errorf("%s is not an valid IP address", ip)
	}
	if ipaddress.To4() == nil {
		return fmt.Errorf("%s is not an valid IPv4 address", ip)
	}

	i.AddressV4 = ipaddress
	return nil
}

// GetIPv4 returns the IPv4 address for the probe
func (i IP) GetIPv4() net.IP {
	return i.AddressV4
}
