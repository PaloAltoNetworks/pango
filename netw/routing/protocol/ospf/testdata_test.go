package ospf

import (
	"github.com/PaloAltoNetworks/pango/version"
)

type testCase struct {
	desc    string
	version version.Number
	conf    Config
}

func getTests() []testCase {
	return []testCase{
		{"v1 ospf no raw", version.Number{10, 0, 0, ""}, Config{
			Enable:                        true,
			RouterId:                      "router id",
			RejectDefaultRoute:            true,
			AllowRedistributeDefaultRoute: true,
			Rfc1583:                       true,
			SpfCalculationDelay:           1.5,
			LsaInterval:                   2.5,
			EnableGracefulRestart:         true,
			GracePeriod:                   1000,
			HelperEnable:                  true,
			StrictLsaChecking:             true,
			MaxNeighborRestartTime:        500,
			BfdProfile:                    "bfd profile",
		}},
		{"v1 ospf with raw", version.Number{10, 0, 0, ""}, Config{
			Enable:                        true,
			RouterId:                      "router id",
			RejectDefaultRoute:            true,
			AllowRedistributeDefaultRoute: true,
			Rfc1583:                       true,
			SpfCalculationDelay:           1.5,
			LsaInterval:                   2.5,
			EnableGracefulRestart:         true,
			GracePeriod:                   1000,
			HelperEnable:                  true,
			StrictLsaChecking:             true,
			MaxNeighborRestartTime:        500,
			BfdProfile:                    "bfd profile",
			raw: map[string]string{
				"ap":   "auth profile",
				"area": "area",
				"exp":  "export rules",
			},
		}},
	}
}
