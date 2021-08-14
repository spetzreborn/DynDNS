package provider

/*

Implementation to get and set DNS addresses from Glesys:
https://glesys.se

Uses their API:
https://github.com/glesys/glesys-go

*/

import (
	"errors"
	"fmt"
	"net"

	"github.com/glesys/glesys-go"
)

// Glesys is an struct containing connection information to Glesys API, and response information
// https://github.com/glesys/glesys-go
type Glesys struct {
	AddressV4    string
	GlesysClient *glesys.Client
	Initiated    bool
}

// Init initializes the client with glesys.
func (g *Glesys) Init(config map[string]string) (err error) {
	validate := []string{"Project", "APIKey", "ApplicationName"}
	for _, v := range validate {
		if _, ok := config[v]; !ok {
			return fmt.Errorf("need argument \"%s\"", v)
		}
	}

	g.GlesysClient = glesys.NewClient("Project", "API-KEY", "Application name")

	return nil
}

// GetARecord returns the IP address for the FQDN
func (g Glesys) GetARecord(fqdn string) (net.IP, error) {
	if !g.Initiated {
		return nil, errors.New("Glesys not properly initiated, cant get A record")
	}
	return net.ParseIP(g.AddressV4), nil
}

// SetARecord returns the IP address for the FQDN
func (g *Glesys) SetARecord(fqdn string, ip net.IP) error {
	return errors.New("Cant set A record")
}
