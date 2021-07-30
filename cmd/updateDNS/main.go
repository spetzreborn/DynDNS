package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/stockholmuniversity/goversionflag"

	client "github.com/spetzreborn/DynDNS/pkg/client"
	provider "github.com/spetzreborn/DynDNS/pkg/provider"
)

func main() {

	configFile := flag.String("configfile", "", "Configuation file")
	goversionflag.PrintVersionAndExit()

	if *configFile == "" {
		log.Fatalln("Need configuration file.")
	}

	config, err := NewConfig(configFile)
	if err != nil {
		log.Fatalln("error when parsing configuration file: " + err.Error())
	}

	var c client.Client
	var p provider.Provider

	for _, item := range config.Items {
		c = client.ClientTypes[item.Client.ClientType]
		p = provider.ProviderTypes[item.Provider.ProviderType]

		err := c.Init(item.Client.ClientConfig)
		if err != nil {
			log.Fatalf("could not initiate client for %s: %s", item.Record, err.Error())
		}

		err = p.Init(item.Provider.ProviderConfig)
		if err != nil {
			log.Fatalf("could not initiate provider for %s: %s", item.Record, err.Error())
		}

		fmt.Println(c.GetIPv4().String())
	}

}
