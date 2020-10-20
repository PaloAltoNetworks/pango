package virus

import (
	"github.com/PaloAltoNetworks/pango/version"
)

type tc struct {
	desc    string
	version version.Number
	conf    Entry
}

func getTests() []tc {
	return []tc{
		{"v1 blank", version.Number{6, 1, 0, ""}, Entry{
			Name:          "t1",
			Description:   "foobar",
			PacketCapture: false,
		}},
		{"v1 with decoders", version.Number{6, 1, 0, ""}, Entry{
			Name:          "t1",
			Description:   "foobar",
			PacketCapture: true,
			Decoders: []Decoder{
				Decoder{
					Name: "smtp",
				},
				Decoder{
					Name:   "smb",
					Action: Allow,
				},
				Decoder{
					Name:           "pop3",
					Action:         ResetClient,
					WildfireAction: Alert,
				},
				Decoder{
					Name:           "http",
					WildfireAction: ResetServer,
				},
			},
		}},
		{"v1 with app exceptions", version.Number{6, 1, 0, ""}, Entry{
			Name:        "t1",
			Description: "foobar",
			ApplicationExceptions: []ApplicationException{
				ApplicationException{
					Application: "amazon-aws-console",
				},
				ApplicationException{
					Application: "dell-update",
					Action:      Allow,
				},
				ApplicationException{
					Application: "hotmail",
					Action:      ResetClient,
				},
			},
		}},
		{"v2 blank", version.Number{10, 0, 0, ""}, Entry{
			Name:          "t1",
			Description:   "foobar",
			PacketCapture: false,
		}},
		{"v2 with decoders", version.Number{10, 0, 0, ""}, Entry{
			Name:          "t1",
			Description:   "foobar",
			PacketCapture: true,
			Decoders: []Decoder{
				Decoder{
					Name: "smtp",
				},
				Decoder{
					Name:   "smb",
					Action: Allow,
				},
				Decoder{
					Name:           "pop3",
					Action:         ResetClient,
					WildfireAction: Alert,
				},
				Decoder{
					Name:           "http",
					WildfireAction: ResetServer,
				},
			},
		}},
		{"v2 with ml decoder", version.Number{10, 0, 0, ""}, Entry{
			Name:          "t1",
			Description:   "foobar",
			PacketCapture: true,
			Decoders: []Decoder{
				Decoder{
					Name:                  "smtp",
					WildfireAction:        ResetBoth,
					MachineLearningAction: Default,
				},
			},
		}},
		{"v2 with app exceptions", version.Number{10, 0, 0, ""}, Entry{
			Name:        "t1",
			Description: "foobar",
			ApplicationExceptions: []ApplicationException{
				ApplicationException{
					Application: "amazon-aws-console",
				},
				ApplicationException{
					Application: "dell-update",
					Action:      Allow,
				},
				ApplicationException{
					Application: "hotmail",
					Action:      ResetClient,
				},
			},
		}},
		{"v2 with ml config", version.Number{10, 0, 0, ""}, Entry{
			Name:        "t1",
			Description: "foobar",
			MachineLearningModels: []MachineLearningModel{
				MachineLearningModel{
					Model:  "model1",
					Action: Drop,
				},
				MachineLearningModel{
					Model:  "model2",
					Action: Alert,
				},
			},
		}},
		{"v2 with ml config and exceptions", version.Number{10, 0, 0, ""}, Entry{
			Name: "t1",
			MachineLearningModels: []MachineLearningModel{
				MachineLearningModel{
					Model:  "model1",
					Action: Drop,
				},
			},
			MachineLearningExceptions: []MachineLearningException{
				MachineLearningException{
					Name:        "exception1",
					Description: "foo",
					Filename:    "myfilename",
				},
			},
		}},
	}
}
