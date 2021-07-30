package provider

/*

Fake is a fake DNS provider that returns nilfor all records if not set otherwise with SetARecord().

*/

import (
	"net"
)

// Fake is an provider that is 'fake', as it does not communicate with any real DNS provider.
type Fake struct {
	AddressV4 map[string]net.IP
}

// Init initializes the fake provider, it does not validate the config or returns any errors.
func (f *Fake) Init(config interface{}) (err error) {
	f.AddressV4 = make(map[string]net.IP)
	return nil
}

// GetARecord returns the IP address for the FQDN
func (f Fake) GetARecord(fqdn string) (net.IP, error) {
	if _, exists := f.AddressV4[fqdn]; !exists {
		return nil, nil
	}
	return f.AddressV4[fqdn], nil
}

// SetARecord returns the IP address for the FQDN
func (f *Fake) SetARecord(fqdn string, ip net.IP) error {
	f.AddressV4[fqdn] = ip
	return nil
}
