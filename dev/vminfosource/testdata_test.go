package vminfosource

import (
	"github.com/PaloAltoNetworks/pango/version"
)

type testCase struct {
	desc    string
	version version.Number
	conf    Entry
}

func getTests() []testCase {
	return []testCase{
		{"v1 vpc disabled", version.Number{6, 1, 0, ""}, Entry{
			Name: "one",
			AwsVpc: &AwsVpc{
				Description:     "description",
				Disabled:        true,
				Source:          "source",
				AccessKeyId:     "accesskeyid",
				SecretAccessKey: "secret",
				UpdateInterval:  5,
				EnableTimeout:   false,
				Timeout:         6,
				VpcId:           "vpcid",
			},
		}},
		{"v1 vpc enable timeout", version.Number{6, 1, 0, ""}, Entry{
			Name: "one",
			AwsVpc: &AwsVpc{
				Description:     "description",
				Disabled:        false,
				Source:          "source",
				AccessKeyId:     "accesskeyid",
				SecretAccessKey: "secret",
				UpdateInterval:  5,
				EnableTimeout:   true,
				Timeout:         6,
				VpcId:           "vpcid",
			},
		}},
		{"v1 esxi disabled", version.Number{6, 1, 0, ""}, Entry{
			Name: "name",
			Esxi: &Esxi{
				Description:    "desc",
				Port:           8443,
				Disabled:       true,
				EnableTimeout:  false,
				Timeout:        5,
				Source:         "src",
				Username:       "user",
				Password:       "passwd",
				UpdateInterval: 6,
			},
		}},
		{"v1 esxi enable timeout", version.Number{6, 1, 0, ""}, Entry{
			Name: "name",
			Esxi: &Esxi{
				Description:    "desc",
				Port:           8443,
				Disabled:       false,
				EnableTimeout:  true,
				Timeout:        5,
				Source:         "src",
				Username:       "user",
				Password:       "passwd",
				UpdateInterval: 6,
			},
		}},
		{"v1 vcenter disabled", version.Number{6, 1, 0, ""}, Entry{
			Name: "name",
			Vcenter: &Vcenter{
				Description:    "desc",
				Port:           8443,
				Disabled:       true,
				EnableTimeout:  false,
				Timeout:        5,
				Source:         "src",
				Username:       "user",
				Password:       "passwd",
				UpdateInterval: 6,
			},
		}},
		{"v1 vcenter enable timeout", version.Number{6, 1, 0, ""}, Entry{
			Name: "name",
			Vcenter: &Vcenter{
				Description:    "desc",
				Port:           8443,
				Disabled:       false,
				EnableTimeout:  true,
				Timeout:        5,
				Source:         "src",
				Username:       "user",
				Password:       "passwd",
				UpdateInterval: 6,
			},
		}},
		{"v2 vpc disabled", version.Number{8, 1, 0, ""}, Entry{
			Name: "one",
			AwsVpc: &AwsVpc{
				Description:     "description",
				Disabled:        true,
				Source:          "source",
				AccessKeyId:     "accesskeyid",
				SecretAccessKey: "secret",
				UpdateInterval:  5,
				EnableTimeout:   false,
				Timeout:         6,
				VpcId:           "vpcid",
			},
		}},
		{"v2 vpc enable timeout", version.Number{8, 1, 0, ""}, Entry{
			Name: "one",
			AwsVpc: &AwsVpc{
				Description:     "description",
				Disabled:        false,
				Source:          "source",
				AccessKeyId:     "accesskeyid",
				SecretAccessKey: "secret",
				UpdateInterval:  5,
				EnableTimeout:   true,
				Timeout:         6,
				VpcId:           "vpcid",
			},
		}},
		{"v2 esxi disabled", version.Number{8, 1, 0, ""}, Entry{
			Name: "name",
			Esxi: &Esxi{
				Description:    "desc",
				Port:           8443,
				Disabled:       true,
				EnableTimeout:  false,
				Timeout:        5,
				Source:         "src",
				Username:       "user",
				Password:       "passwd",
				UpdateInterval: 6,
			},
		}},
		{"v2 esxi enable timeout", version.Number{8, 1, 0, ""}, Entry{
			Name: "name",
			Esxi: &Esxi{
				Description:    "desc",
				Port:           8443,
				Disabled:       false,
				EnableTimeout:  true,
				Timeout:        5,
				Source:         "src",
				Username:       "user",
				Password:       "passwd",
				UpdateInterval: 6,
			},
		}},
		{"v2 vcenter disabled", version.Number{8, 1, 0, ""}, Entry{
			Name: "name",
			Vcenter: &Vcenter{
				Description:    "desc",
				Port:           8443,
				Disabled:       true,
				EnableTimeout:  false,
				Timeout:        5,
				Source:         "src",
				Username:       "user",
				Password:       "passwd",
				UpdateInterval: 6,
			},
		}},
		{"v2 vcenter enable timeout", version.Number{8, 1, 0, ""}, Entry{
			Name: "name",
			Vcenter: &Vcenter{
				Description:    "desc",
				Port:           8443,
				Disabled:       false,
				EnableTimeout:  true,
				Timeout:        5,
				Source:         "src",
				Username:       "user",
				Password:       "passwd",
				UpdateInterval: 6,
			},
		}},
		{"v2 google compute disabled service in gce", version.Number{8, 1, 0, ""}, Entry{
			Name: "name",
			GoogleCompute: &GoogleCompute{
				Description:    "desc",
				Disabled:       true,
				AuthType:       AuthTypeServiceInGce,
				ProjectId:      "proj",
				ZoneName:       "zone",
				UpdateInterval: 6,
				EnableTimeout:  false,
				Timeout:        5,
			},
		}},
		{"v2 google compute enable timeout service account", version.Number{8, 1, 0, ""}, Entry{
			Name: "name",
			GoogleCompute: &GoogleCompute{
				Description:              "desc",
				Disabled:                 false,
				AuthType:                 AuthTypeServiceAccount,
				ServiceAccountCredential: "creds json",
				ProjectId:                "proj",
				ZoneName:                 "zone",
				UpdateInterval:           6,
				EnableTimeout:            true,
				Timeout:                  5,
			},
		}},
	}
}
