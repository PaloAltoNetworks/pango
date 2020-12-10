package pango_test

import (
	"log"

	"github.com/PaloAltoNetworks/pango"
	"github.com/PaloAltoNetworks/pango/commit"
	"github.com/PaloAltoNetworks/pango/objs/addr"
	"github.com/PaloAltoNetworks/pango/objs/addrgrp"
	"github.com/PaloAltoNetworks/pango/poli/security"
	"github.com/PaloAltoNetworks/pango/util"
)

// Example_createAddressGroup is a Panorama example on how to create/delete a
// security policy with the associated address group and addresses
func Example_createAddressGroup() {
	var deviceGroup = "MyDeviceGroup"
	var tags = []string{"sometag"}
	var err error

	pan := &pango.Panorama{Client: pango.Client{
		Hostname: "192.168.1.1",
		Username: "admin",
		Password: "admin",
		Logging:  pango.LogAction | pango.LogOp,
	}}
	if err = pan.Initialize(); err != nil {
		log.Panic(err)
		return
	}

	// Create the addresses, address group and security policy
	addr1 := addr.Entry{
		Name:        "SampleAddress1",
		Value:       "10.192.226.101/32",
		Type:        addr.IpNetmask,
		Description: "First address of a sample address group",
		Tags:        tags,
	}
	if err = pan.Objects.Address.Set(deviceGroup, addr1); err != nil {
		log.Panic(err)
	}

	addr2 := addr.Entry{
		Name:        "SampleAddress2",
		Value:       "10.192.226.102/32",
		Type:        addr.IpNetmask,
		Description: "Second address of a sample address group",
		Tags:        tags,
	}
	if err = pan.Objects.Address.Set(deviceGroup, addr2); err != nil {
		log.Panic(err)
	}

	ag := addrgrp.Entry{
		Name:            "SampleAddressGroup",
		Description:     "This in an example on how to use address groups",
		StaticAddresses: []string{addr1.Name, addr2.Name},
		Tags:            tags,
	}
	if err = pan.Objects.AddressGroup.Set(deviceGroup, ag); err != nil {
		log.Panic(err)
	}

	securityPolicy := security.Entry{
		Name:                 "SamplePolicy",
		Description:          "This is where the request number goes",
		Tags:                 tags,
		SourceZones:          []string{"ICORPEXT"},
		SourceAddresses:      []string{"any"},
		DestinationZones:     []string{"ICORPDMZ"},
		DestinationAddresses: []string{ag.Name},
		Applications:         []string{"ssl"},
		Services:             []string{"application-default"},
		LogSetting:           "APAC-Dell-Standard-Logging",
		Group:                "Dell_Corp_Default",
	}
	securityPolicy.Defaults()

	if err = pan.Policies.Security.VerifiableSet(deviceGroup, util.PreRulebase, securityPolicy); err != nil {
		log.Panic(err)
	}

	panCommit := commit.PanoramaCommit{
		Description:  "Created example address group",
		Admins:       nil,
		DeviceGroups: []string{deviceGroup},
	}

	resp, bytes, err := pan.Commit(panCommit, "", nil)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Job ID: %v\n", resp)
	log.Printf("Response XML: %v\n", string(bytes))

	// Delete the addresses, address group and security policy
	securityPolicy, err = pan.Policies.Security.Get(deviceGroup, util.PreRulebase, "SamplePolicy")
	if err != nil {
		log.Panic(err)
	}
	if err = pan.Policies.Security.Delete(deviceGroup, util.PreRulebase, securityPolicy); err != nil {
		log.Panic(err)
	}

	for _, agn := range securityPolicy.DestinationAddresses {
		ag, err := pan.Objects.AddressGroup.Get(deviceGroup, agn)
		if err != nil {
			log.Panic(err)
		}
		if err = pan.Objects.AddressGroup.Delete(deviceGroup, ag); err != nil {
			log.Panic(err)
		}
		for _, an := range ag.StaticAddresses {
			a, err := pan.Objects.Address.Get(deviceGroup, an)
			if err = pan.Objects.Address.Delete(deviceGroup, a); err != nil {
				log.Panic(err)
			}
		}
	}

	panCommit = commit.PanoramaCommit{
		Description:  "Deleted sample address group",
		Admins:       nil,
		DeviceGroups: []string{deviceGroup},
	}

	resp, bytes, err = pan.Commit(panCommit, "", nil)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Job ID: %v\n", resp)
	log.Printf("Response XML: %v\n", string(bytes))

}
