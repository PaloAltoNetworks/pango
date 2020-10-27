package spyware

// Valid values for WhiteList.LogLevel.
const (
	LogLevelDefault       = "default"
	LogLevelNone          = "none"
	LogLevelLow           = "low"
	LogLevelInformational = "informational"
	LogLevelMedium        = "medium"
	LogLevelHigh          = "high"
	LogLevelCritical      = "critical"
)

// Valid values for PacketCapture params.
const (
	Disable         = "disable"
	SinglePacket    = "single-packet"
	ExtendedCapture = "extended-capture"
)

// Valid values for Action params.
const (
	// Valid for Entry.Action and BlockList.Action only.
	ActionAlert = "alert"

	// Valid for DnsCategory.Action only.
	ActionDefault = "default"

	ActionAllow    = "allow"
	ActionBlock    = "block"
	ActionSinkhole = "sinkhole"
)

const (
	singular = "anti-spyware security profile"
	plural   = "anti-spyware security profiles"
)
