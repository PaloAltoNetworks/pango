package texception

// Valid values for BlockIpTrackBy.
const (
	TrackBySource               = "source"
	TrackBySourceAndDestination = "source-and-destination"
)

// Valid values for PacketCapture params.
const (
	Disable         = "disable"
	SinglePacket    = "single-packet"
	ExtendedCapture = "extended-capture"
)

// Valid values for Action params.
const (
	ActionDefault     = "default"
	ActionAllow       = "allow"
	ActionAlert       = "alert"
	ActionDrop        = "drop"
	ActionResetClient = "reset-client"
	ActionResetServer = "reset-server"
	ActionResetBoth   = "reset-both"
	ActionBlockIp     = "block-ip"
)

const (
	singular = "anti-spyware security profile"
	plural   = "anti-spyware security profiles"
)
