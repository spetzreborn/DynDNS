package main

import (
	"errors"

	"github.com/BurntSushi/toml"
)

// Config should be populated from an TOML configuration file.
type Config struct {
	Items []Item
}

// Item is the coupled configuration for each record that shall be updated.
type Item struct {
	Client   Client
	Provider Provider
	Record   string
}

// Client is the configuration for each client. "How to get the current IP"
type Client struct {
	ClientType   string
	ClientConfig interface{}
}

// Provider is the configuration for each DNS provider. "Where to update DNS records"
type Provider struct {
	ProviderType   string
	ProviderConfig interface{}
}

// NewConfig returns default configuration with consideration to configuration file.
func NewConfig(configFile *string) (*Config, error) {
	config := &Config{}
	if _, err := toml.DecodeFile(*configFile, config); err != nil {
		return nil, errors.New("toml decoding failed: " + err.Error())
	}

	for _, item := range config.Items {
		//TODO Use lookup table for ClientTypes
		switch item.Client.ClientType {
		case "AtlasProbe":
		case "IP":
		default:
			return nil, errors.New("not an correct ClientType: " + item.Client.ClientType)
		}
	}

	return config, nil
}
