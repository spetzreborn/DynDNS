package main

import (
	"errors"

	"github.com/BurntSushi/toml"
	client "github.com/spetzreborn/DynDNS/pkg/client"
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

	// Verifiy that all config items have client type that exists.
	for _, item := range config.Items {
		if _, keyExists := client.ClientTypes[item.Client.ClientType]; !keyExists {
			return nil, errors.New("not an correct ClientType: " + item.Client.ClientType)
		}
	}

	return config, nil
}
