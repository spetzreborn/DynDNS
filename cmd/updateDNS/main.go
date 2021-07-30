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

	for _, item := range config.Items {
		c := client.ClientTypes[item.Client.ClientType]
		p := provider.ProviderTypes[item.Provider.ProviderType]

		err := c.Init(item.Client.ClientConfig)
		if err != nil {
			log.Fatalf("could not initiate client for %s: %s", item.Record, err.Error())
		}

		err = p.Init(item.Provider.ProviderConfig)
		if err != nil {
			log.Fatalf("could not initiate provider for %s: %s", item.Record, err.Error())
		}

		currentIP, err := p.GetARecord(item.Record)
		if err != nil {
			log.Printf("could not get A record for %s... skipping\n", item.Record)
			continue
		}
		if currentIP == nil {
			//TODO Print if verbose?
			fmt.Printf("no current A record for %s setting new: %s\n", item.Record, c.GetIPv4().String())
			p.SetARecord(item.Record, c.GetIPv4())
		} else if currentIP.String() != c.GetIPv4().String() {
			fmt.Printf("updating %s\n", item.Record)
			p.SetARecord(item.Record, c.GetIPv4())
		} else {
			fmt.Printf("record %s is already correct\n", item.Record)
		}
	}

}
