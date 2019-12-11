package pango_test

import (
	"log"

	"github.com/PaloAltoNetworks/pango"
	"github.com/PaloAltoNetworks/pango/netw/interface/eth"
)

// ExampleCreateInterface demonstrates how to use pango to create an interface
// if the interface is not already configured.
func Example_createInterface() {
	var err error

	// Connect to the firewall.
	fw := pango.Firewall{Client: pango.Client{
		Hostname: "192.168.1.1",
		Username: "admin",
		Password: "admin",
	}}

	// Connect to the firewall and verify authentication params.
	if err = fw.Initialize(); err != nil {
		log.Fatalf("Failed to connect to %s: %s", fw.Hostname, err)
	}

	// Define the ethernet interface we want to configure.
	e := eth.Entry{
		Name:      "ethernet1/7",
		Mode:      "layer3",
		Comment:   "Made by pango",
		StaticIps: []string{"10.1.1.1/24", "10.2.1.1/24"},
	}

	// If the interface is already present, leave it alone.
	ethList, err := fw.Network.EthernetInterface.GetList()
	if err != nil {
		log.Fatalf("Failed to get interface listing: %s", err)
	}
	for i := range ethList {
		if ethList[i] == e.Name {
			log.Printf("Interface %q already exists, quitting.", e.Name)
			return
		}
	}

	// Since the interface is not present, configure it.
	if err = fw.Network.EthernetInterface.Set("vsys1", e); err != nil {
		log.Fatalf("Failed to create %q: %s", e.Name, err)
	}
	log.Printf("Created %q ok", e.Name)
}
