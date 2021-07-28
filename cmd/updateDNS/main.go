package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/stockholmuniversity/goversionflag"

	client "github.com/spetzreborn/DynDNS/pkg/client"
)

func main() {

	configFile := flag.String("configfile", "", "Configuation file")
	goversionflag.PrintVersionAndExit()

	if *configFile == "" {
		log.Fatalln("Need configuration file.")
	}

	config, err := NewConfig(configFile)
	if err != nil {
		log.Fatalln("Got error when parsing configuration file: " + err.Error())
	}

	var c client.Client

	//TODO Use lookup table for ClientTypes
	for _, item := range config.Items {
		switch item.Client.ClientType {
		case "AtlasProbe":
			c = new(client.AtlasProbe)
		case "IP":
			c = new(client.IP)
		default:
			log.Fatalf("No such client type: %s\n", item.Client.ClientType)
		}
		err = c.Init(item.Client.ClientConfig)
		if err != nil {
			log.Fatalf("could not initiate client for %s: %s", item.Record, err.Error())
		}
		fmt.Println(c.GetIPv4().String())
	}

}
