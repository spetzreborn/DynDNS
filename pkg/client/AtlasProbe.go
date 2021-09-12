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

// AtlasProbe is an stuct of the status for an probe. This comes from Ripe and the Atlas Probe project.
// https://atlas.ripe.net/docs/api/v2/reference/#!/probes/Type_0
type AtlasProbe struct {
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

// Init initializes the probe with response from Atlas Probe project
func (p *AtlasProbe) Init(param map[string]string) (err error) {
	if _, ok := param["probeID"]; !ok {
		return errors.New("argument must include a value \"probeID\"")

	}
	if _, err := strconv.Atoi(fmt.Sprintf("%v", param["probeID"])); err != nil {
		return errors.New("argument must only be digits")
	}

	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := httpClient.Get("https://atlas.ripe.net/api/v2/probes/" + fmt.Sprintf("%v", param["probeID"]))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&p)

	if p.AddressV4 == "" {
		return errors.New("could not get IPv4 address from Atlas Probe")
	}
	return nil
}

// GetIPv4 returns the IPv4 address for the probe
func (p AtlasProbe) GetIPv4() net.IP {
	return net.ParseIP(p.AddressV4)
}
