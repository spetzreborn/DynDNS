package provider

/*

Simple implementation to get IPv4 addresses from Ripe Atlas Probe project
https://atlas.ripe.net/

Uses the unauthorized API to get infromation from probes with known ID.


Possible enchancment:
* Use https://github.com/keltia/ripe-atlas instead?

*/

import (
	"errors"
	"fmt"
	"net"
	"unsafe"

	"github.com/glesys/glesys-go"
)

// Glesys is an stuct of the status for an probe. This comes from Ripe and the Atlas Probe project.
// https://atlas.ripe.net/docs/api/v2/reference/#!/probes/Type_0
type Glesys struct {
	AddressV4    string
	GlesysClient *glesys.Client
}

type GlesysConfig struct {
	Project         string
	APIKey          string
	ApplicationName string
}

// Init initializes the client with glesys.
func (g *Glesys) Init(config interface{}) (err error) {
	zzz := *(*GlesysConfig)(unsafe.Pointer(&config))
	fmt.Printf("%#v\n", zzz)
	fmt.Printf("%T\n", zzz)
	fmt.Printf("%#v\n", config)
	fmt.Printf("%T\n", config)
	gc, ok := config.(GlesysConfig)
	if ok == false {
		return errors.New("argument must be mappable to an GlesysConfig")
	}

	g.GlesysClient = glesys.NewClient(gc.Project, "API-KEY", "Application name")

	if g.AddressV4 == "" {
		return errors.New("could not get IPv4 address from Atlas Probe")
	}
	return nil
}

// GetARecord returns the IP address for the FQDN
func (g Glesys) GetARecord(fqdn string) (net.IP, error) {
	return net.ParseIP(g.AddressV4), nil
}

// SetARecord returns the IP address for the FQDN
func (g *Glesys) SetARecord(fqdn string, ip net.IP) error {
	return nil
}
