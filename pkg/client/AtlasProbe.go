package client

/*

Simple implementation to get IPv4 addresses from Ripe Atlas Probe project
https://atlas.ripe.net/

Uses the unauthorized API to get infromation from probes with known ID.


Possible enchancment:
* Use https://github.com/keltia/ripe-atlas instead?

*/

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"
)

// atlasProbe is an stuct of the status for an probe. This comes from Ripe and the Atlas Probe project.
// https://atlas.ripe.net/docs/api/v2/reference/#!/probes/Type_0
type atlasProbe struct {
	AddressV4      string `json:"address_v4"`
	AddressV6      string `json:"address_v6"`
	AsnV4          int    `json:"asn_v4"`
	AsnV6          int    `json:"asn_v6"`
	CountryCode    string `json:"country_code"`
	Description    string `json:"description"`
	FirstConnected string `json:"first_connected"`
	Geometry       struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"geometry"`
	ID            int    `json:"id"`
	IsAnchor      bool   `json:"is_anchor"`
	IsPublic      bool   `json:"is_public"`
	LastConnected string `json:"last_connected"`
	PrefixV4      string `json:"prefix_v4"`
	PrefixV6      string `json:"prefix_v6"`
	Status        struct {
		ID    int       `json:"id"`
		Name  string    `json:"name"`
		Since time.Time `json:"since"`
	} `json:"status"`
	StatusSince string `json:"status_since"`
	Tags        []struct {
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"tags"`
	TotalUptime int    `json:"total_uptime"`
	Type        string `json:"type"`
}

type IHTTPClient interface {
	Get(url string) (resp *http.Response, err error)
}

type AtlasClient struct {
	httpClient IHTTPClient
	addressV4  net.IP
}

// Init initializes the probe with response from Atlas Probe project
func (c *AtlasClient) Init(probeID string) error {
	if _, err := strconv.Atoi(fmt.Sprintf("%v", probeID)); err != nil {
		return errors.New("argument must only be digits")
	}
	c.httpClient = &http.Client{Timeout: 5 * time.Second}
	ipv4, err := c.getIPFromProbe(probeID)
	if err != nil {
		return err
	}
	ip := net.ParseIP(*ipv4)
	c.addressV4 = ip
	return nil
}

func (c *AtlasClient) getIPFromProbe(probeID string) (*string, error) {
	resp, err := c.httpClient.Get(fmt.Sprintf("https://atlas.ripe.net/api/v2/probes/%s", probeID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ap := atlasProbe{}
	json.NewDecoder(resp.Body).Decode(&ap)
	if ap.AddressV4 == "" {
		return nil, errors.New("could not get IPv4 address from Atlas Probe")
	}
	return &ap.AddressV4, nil
}

// GetIPv4 returns the IPv4 address for the probe
func (c *AtlasClient) GetIPv4() net.IP {
	return c.addressV4
}
