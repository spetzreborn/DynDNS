package main

import (
	"errors"

	"github.com/BurntSushi/toml"
	client "github.com/spetzreborn/DynDNS/pkg/client"
	provider "github.com/spetzreborn/DynDNS/pkg/provider"
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
	ClientConfig map[string]string
}

// Provider is the configuration for each DNS provider. "Where to update DNS records"
type Provider struct {
	ProviderType   string
	ProviderConfig map[string]string
}

// NewConfig returns default configuration with consideration to configuration file.
func NewConfig(configFile *string) (*Config, error) {
	config := &Config{}
	if _, err := toml.DecodeFile(*configFile, config); err != nil {
		return nil, errors.New("toml decoding failed: " + err.Error())
	}

	// Verifiy that all config items have client and provider types that exists.
	for _, item := range config.Items {
		if _, keyExists := client.ClientTypes[item.Client.ClientType]; !keyExists {
			return nil, errors.New("not an correct ClientType: " + item.Client.ClientType)
		}
		if _, keyExists := provider.ProviderTypes[item.Provider.ProviderType]; !keyExists {
			return nil, errors.New("not an correct ProviderType: " + item.Provider.ProviderType)
		}
	}

	return config, nil
}
