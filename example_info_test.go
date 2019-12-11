package pango_test

import (
	"encoding/json"
	"fmt"

	"github.com/PaloAltoNetworks/pango"
)

// About is a struct to hold information about the given PAN-OS device.
type About struct {
	Hostname string `json:"hostname"`
	Type     string `json:"type"`
	Model    string `json:"model"`
	Version  string `json:"version"`
	Serial   string `json:"serial"`
}

// ExamplePanosInfo outputs various info about a PAN-OS device as
// JSON.
func Example_outputApiKey() {
	var out About

	conInfo := pango.Client{
		Hostname: "192.168.1.1",
		Username: "admin",
		Password: "admin",
		Logging:  pango.LogQuiet,
	}

	con, err := pango.Connect(conInfo)
	if err != nil {
		return
	}

	switch x := con.(type) {
	case *pango.Firewall:
		out = About{
			Hostname: x.Hostname,
			Type:     "NGFW",
			Model:    x.SystemInfo["model"],
			Version:  x.Version.String(),
			Serial:   x.SystemInfo["serial"],
		}
	case *pango.Panorama:
		out = About{
			Hostname: x.Hostname,
			Type:     "Panorama",
			Model:    x.SystemInfo["model"],
			Version:  x.Version.String(),
			Serial:   x.SystemInfo["serial"],
		}
	}

	b, err := json.Marshal(out)
	if err != nil {
		return
	}

	fmt.Printf("%s\n", b)
}
