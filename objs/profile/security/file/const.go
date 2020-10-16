package file

// Valid values for Rule.Direction.
const (
	DirectionUpload   = "upload"
	DirectionDownload = "download"
	DirectionBoth     = "both"
)

// Valid values for Rule.Action.
const (
	ActionAlert              = "alert"
	ActionBlock              = "block"
	ActionContinue           = "continue"
	ActionForward            = "forward"
	ActionContinueAndForward = "continue-and-forward"
)

const (
	singular = "file blocking security profile"
	plural   = "file blocking security profiles"
)
