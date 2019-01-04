package aggregate


// Valid values for Origin.
const (
    OriginIgp = "igp"
    OriginEgp = "egp"
    OriginIncomplete = "incomplete"
)

// Valid values for CommunityType.
const (
    CommunityTypeNone = "none"
    CommunityTypeRemoveAll = "remove-all"
    CommunityTypeRemoveRegex = "remove-regex"
    CommunityTypeAppend = "append"
    CommunityTypeOverwrite = "overwrite"
)

// Valid values for CommunityValue when CommunityType is "append" or
// "overwrite".
const (
    AppendNoExport = "no-export"
    AppendNoAdvertise = "no-advertise"
    AppendLocalAs = "local-as"
    AppendNoPeer = "nopeer"
)

// Valid values for AsPathType.
const (
    AsPathTypeNone = "none"
    AsPathTypePrepend = "prepend"
    // Both of these are disabled as of PAN-OS 8.1, so commenting out.
    //AsPathTypeRemove = "remove"
    //AsPathTypeRemoveAndPrepend = "remove-and-prepend"
)

const (
    singular = "bgp aggregation policy"
    plural = "bgp aggregation policies"
)
