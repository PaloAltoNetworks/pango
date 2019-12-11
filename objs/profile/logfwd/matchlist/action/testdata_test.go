package action

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
		{"v1 tagging add-tag with local", version.Number{8, 0, 0, ""}, Entry{
			Name:         "t1",
			ActionType:   ActionTypeTagging,
			Action:       ActionAddTag,
			Target:       TargetSource,
			Registration: RegistrationLocal,
			Tags:         []string{"tag3", "tag1", "tag2"},
		}},
		{"v1 tagging remove-tag with panorama", version.Number{8, 0, 0, ""}, Entry{
			Name:         "t2",
			ActionType:   ActionTypeTagging,
			Action:       ActionRemoveTag,
			Target:       TargetDestination,
			Registration: RegistrationPanorama,
		}},
		{"v1 tagging add-tag with remote", version.Number{8, 0, 0, ""}, Entry{
			Name:         "t3",
			ActionType:   ActionTypeTagging,
			Action:       ActionRemoveTag,
			Target:       TargetDestination,
			Registration: RegistrationRemote,
			HttpProfile:  "my profile",
			Tags:         []string{"mytag"},
		}},
		{"v2 tagging add-tag with local", version.Number{8, 1, 0, ""}, Entry{
			Name:         "t1",
			ActionType:   ActionTypeTagging,
			Action:       ActionAddTag,
			Target:       TargetSource,
			Registration: RegistrationLocal,
			Tags:         []string{"tag3", "tag1", "tag2"},
		}},
		{"v2 tagging remove-tag with panorama", version.Number{8, 1, 0, ""}, Entry{
			Name:         "t2",
			ActionType:   ActionTypeTagging,
			Action:       ActionRemoveTag,
			Target:       TargetDestination,
			Registration: RegistrationPanorama,
		}},
		{"v2 tagging add-tag with remote", version.Number{8, 1, 0, ""}, Entry{
			Name:         "t3",
			ActionType:   ActionTypeTagging,
			Action:       ActionRemoveTag,
			Target:       TargetDestination,
			Registration: RegistrationRemote,
			HttpProfile:  "my profile",
			Tags:         []string{"mytag"},
		}},
		{"v2 integration", version.Number{8, 1, 0, ""}, Entry{
			Name:       "t4",
			ActionType: ActionTypeIntegration,
			Action:     ActionAzure,
		}},
		{"v3 tagging add-tag with local", version.Number{9, 0, 0, ""}, Entry{
			Name:         "t1",
			ActionType:   ActionTypeTagging,
			Action:       ActionAddTag,
			Target:       TargetSource,
			Registration: RegistrationLocal,
			Tags:         []string{"tag3", "tag1", "tag2"},
		}},
		{"v3 tagging remove-tag with panorama", version.Number{9, 0, 0, ""}, Entry{
			Name:         "t2",
			ActionType:   ActionTypeTagging,
			Action:       ActionRemoveTag,
			Target:       TargetDestination,
			Registration: RegistrationPanorama,
		}},
		{"v3 tagging add-tag with remote", version.Number{9, 0, 0, ""}, Entry{
			Name:         "t3",
			ActionType:   ActionTypeTagging,
			Action:       ActionRemoveTag,
			Target:       TargetDestination,
			Registration: RegistrationRemote,
			HttpProfile:  "my profile",
			Tags:         []string{"mytag"},
		}},
		{"v3 integration", version.Number{9, 0, 0, ""}, Entry{
			Name:       "t4",
			ActionType: ActionTypeIntegration,
			Action:     ActionAzure,
		}},
		{"v3 tagging add-tag with local and timeout", version.Number{9, 0, 0, ""}, Entry{
			Name:         "t5",
			ActionType:   ActionTypeTagging,
			Action:       ActionAddTag,
			Target:       TargetSource,
			Registration: RegistrationLocal,
			Timeout:      42,
			Tags:         []string{"tag3", "tag1", "tag2"},
		}},
	}
}
