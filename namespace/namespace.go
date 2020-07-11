package namespace

import (
	"encoding/xml"
	"fmt"

	"github.com/PaloAltoNetworks/pango/util"
)

// Namespace encapsulates all the copy/paste code from each
// namespace into a single location.
type Namespace struct {
	Singular string
	Plural   string
	con      util.XapiClient
}

// New returns a new namespace.
func New(s, p string, con util.XapiClient) *Namespace {
	return &Namespace{
		Singular: s,
		Plural:   p,
		con:      con,
	}
}

// Object returns a single object at the given xpath.
func (n *Namespace) Object(cmd string, path []string, desc string, ans interface{}) error {
	return n.retrieve(cmd, path, true, desc, false, false, ans)
}

// Objects returns multiple objects at the given xpath.
func (n *Namespace) Objects(cmd string, path []string, ans interface{}) error {
	return n.retrieve(cmd, path, false, "", true, false, ans)
}

// Listing returns a name listing at the given xpath.
func (n *Namespace) Listing(cmd string, path []string, ans Namer) ([]string, error) {
	if err := n.retrieve(cmd, path, false, "", true, true, ans); err != nil {
		return nil, err
	}

	return ans.Names(), nil
}

// Set performs a SET to create / update one or more objects.
func (n *Namespace) Set(names, path []string, data []interface{}) error {
	n.con.LogAction("(set) %s: %v", n.Plural, names)

	if len(data) == 0 {
		return nil
	}

	// Sanity check: verify name uniqueness.
	tally := make(map[string]int)
	for _, name := range names {
		tally[name] = tally[name] + 1
		if tally[name] > 1 {
			return fmt.Errorf("%s is defined multiple times: %q", n.Singular, name)
		}
	}

	elm := util.BulkElement{
		XMLName: xml.Name{Local: path[len(path)-2]},
		Data:    data,
	}

	if len(data) == 1 {
		path = path[:len(path)-1]
	} else {
		path = path[:len(path)-2]
	}

	_, err := n.con.Set(path, elm.Config(), nil, nil)
	return err
}

// Edit performs an EDIT to create / update a single object.
func (n *Namespace) Edit(name string, path []string, data interface{}) error {
	n.con.LogAction("(edit) %s: %q", n.Singular, name)

	_, err := n.con.Edit(path, data, nil, nil)
	return err
}

// Delete performs a DELETE to remove one or more objects.
func (n *Namespace) Delete(names, path []string) error {
	n.con.LogAction("(delete) %s: %v", n.Plural, names)

	if len(names) == 0 {
		return nil
	}

	_, err := n.con.Delete(path, nil, nil)
	return err
}

// MoveGroup places a logical group of objects in the desired location (rulebase
// objects).
//
// The `movement` param should be one of the Move constants in the util package.
//
// The `rule` param is the other rule the `movement` param is referencing.  If
// this is an empty string, then the first rule in the group isn't moved
// anywhere, but all other rules will still be moved to be grouped with the
// first one.
func (n *Namespace) MoveGroup(pather MovePather, lister MoveLister, movement int, rule string, grp []string) error {
	var err error

	fIdx, oIdx := -1, -1

	n.con.LogAction("(move) %s group", n.Singular)
	if len(grp) == 0 {
		return fmt.Errorf("Requires at least one %s", n.Singular)
	} else if rule == grp[0] {
		return fmt.Errorf("Can't position %q in relation to itself", rule)
	} else if !util.ValidMovement(movement) {
		return fmt.Errorf("Invalid movement specified: %d", movement)
	} else if util.RelativeMovement(movement) && rule == "" {
		return fmt.Errorf("Specify 'rule' in order to perform relative group positioning")
	}
	path := pather(grp[0])

	if movement != util.MoveSkip {
		// Get the list of current security policies.
		curList, err := lister()
		if err != nil {
			return err
		} else if len(curList) == 0 {
			return fmt.Errorf("No policies found")
		}

		switch movement {
		case util.MoveTop:
			_, em := n.con.Move(path, "top", "", nil, nil)
			if em != nil {
				if em.Error() != "already at the top" {
					err = em
				}
			}
		case util.MoveBottom:
			_, em := n.con.Move(path, "bottom", "", nil, nil)
			if em != nil {
				if em.Error() != "already at the bottom" {
					err = em
				}
			}
		default:
			// Find the indexes of the first security policy and the ref policy.
			for i, v := range curList {
				if v == grp[0] {
					fIdx = i
				} else if v == rule {
					oIdx = i
				}
				if fIdx != -1 && oIdx != -1 {
					break
				}
			}

			// Sanity check:  both rules should be present.
			if fIdx == -1 {
				return fmt.Errorf("First %s in group %q does not exist", n.Singular, grp[0])
			} else if oIdx == -1 {
				return fmt.Errorf("Reference %s %q does not exist", n.Singular, rule)
			}

			// Perform the move of the first security policy, if needed.
			if (movement == util.MoveBefore && fIdx > oIdx) || (movement == util.MoveDirectlyBefore && fIdx+1 != oIdx) {
				_, err = n.con.Move(path, "before", rule, nil, nil)
			} else if (movement == util.MoveAfter && fIdx < oIdx) || (movement == util.MoveDirectlyAfter && fIdx != oIdx+1) {
				_, err = n.con.Move(path, "after", rule, nil, nil)
			}
		}

		// If we moved something, make sure it worked.
		if err != nil {
			return err
		}
	}

	// Now move all other rules under the first.
	li := len(path) - 1
	for i := 1; i < len(grp); i++ {
		path[li] = util.AsEntryXpath([]string{grp[i]})
		_, err = n.con.Move(path, "after", grp[i-1], nil, nil)
		if err != nil {
			return err
		}
	}

	return nil
}

// Internal functions.

// retrieve does either a GET or SHOW to retrieve config.
func (n *Namespace) retrieve(cmd string, path []string, singular bool, singleDesc string, plural, namesOnly bool, ans interface{}) error {
	var err error
	var data []byte
	var tag string

	// Sanity check.
	if cmd != util.Get && cmd != util.Show {
		return fmt.Errorf("invalid cmd: %s", cmd)
	}

	// Do logging and determine the actual path to query.
	if singular {
		if singleDesc != "" {
			n.con.LogQuery("(%s) %s: %s", cmd, n.Singular, singleDesc)
		} else {
			n.con.LogQuery("(%s) %s", cmd, n.Singular)
		}
	} else if plural {
		tag = path[len(path)-2]
		if cmd == util.Show {
			path = path[:len(path)-1]
		}
		if namesOnly {
			if cmd == util.Get {
				path = append(path, "@name")
			}
			n.con.LogQuery("(%s) %s names", cmd, n.Singular)
		} else {
			if cmd == util.Get {
				path = path[:len(path)-1]
			}
			n.con.LogQuery("(%s) list of %s", cmd, n.Plural)
		}
	}

	// Perform the query.
	switch cmd {
	case util.Get:
		data, err = n.con.Get(path, nil, nil)
	case util.Show:
		data, err = n.con.Show(path, nil, nil)
	}
	if err != nil {
		return err
	}

	// Unmarshal the response into the given struct.
	data = util.StripPanosPackaging(data, tag)
	return UnpackageXmlInto(data, ans)
}
