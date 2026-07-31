// Package util contains various shared structs and functions used across
// the pango package.
package util

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"regexp"
	"strings"

	"github.com/PaloAltoNetworks/pango/version"
)

var FixedPanosVersionForMultiConfigMove = version.Number{99, 99, 99, ""}

// VsysEntryType defines an entry config node with vsys entries underneath.
type VsysEntryType struct {
	Entries []VsysEntry `xml:"entry"`
}

// VsysEntry defines the "vsys" xpath node under a VsysEntryType config node.
type VsysEntry struct {
	XMLName xml.Name   `xml:"entry"`
	Serial  string     `xml:"name,attr"`
	Vsys    *EntryType `xml:"vsys"`
}

// VsysEntToMap normalizes a VsysEntryType pointer into a map.
func VsysEntToMap(ve *VsysEntryType) map[string][]string {
	if ve == nil {
		return nil
	}

	ans := make(map[string][]string)
	for i := range ve.Entries {
		ans[ve.Entries[i].Serial] = EntToStr(ve.Entries[i].Vsys)
	}

	return ans
}

// MapToVsysEnt converts a map into a VsysEntryType pointer.
//
// This struct is used for "Target" information on Panorama when dealing with
// various policies.  Maps are unordered, but FWICT Panorama doesn't seem to
// order anything anyways when doing things in the GUI, so hopefully this is
// ok...?
func MapToVsysEnt(e map[string][]string) *VsysEntryType {
	if len(e) == 0 {
		return nil
	}

	i := 0
	ve := make([]VsysEntry, len(e))
	for key := range e {
		ve[i].Serial = key
		ve[i].Vsys = StrToEnt(e[key])
		i++
	}

	return &VsysEntryType{ve}
}

// YesNo returns "yes" on true, "no" on false.
func YesNo(val *bool, defaultVal *bool) *string {
	if val == nil && defaultVal == nil {
		return nil
	}

	result := "no"
	if val != nil {
		if *val {
			result = "yes"
		}
	} else if *defaultVal {
		result = "yes"
	}
	return &result
}

// AsBool returns true on yes, else false.
func AsBool(val *string, defaultVal *string) *bool {
	if val == nil && defaultVal == nil {
		return nil
	}

	result := false
	if val != nil {
		if *val == "yes" {
			result = true
		}
	} else if *defaultVal == "yes" {
		result = true
	}
	return &result
}

// AsXpath makes an xpath out of the given interface.
func AsXpath(i interface{}) string {
	switch val := i.(type) {
	case string:
		return val
	case []string:
		return fmt.Sprintf("/%s", strings.Join(val, "/"))
	default:
		return ""
	}
}

// xpathSafe returns val as an XPath string expression that is safe to embed
// in a predicate, guarding against XPath injection (CWE-643) when val
// contains single-quote characters.
//
// XPath 1.0 provides no escape mechanism for a quote character inside a
// string literal, so values containing a single quote are emitted using the
// concat() function (e.g. a'b becomes concat('a',"'",'b')). Values without a
// single quote are simply wrapped in single quotes. The returned expression
// always evaluates to val as a literal string and can never alter the
// surrounding predicate structure.
func xpathSafe(val string) string {
	if !strings.Contains(val, "'") {
		return "'" + val + "'"
	}

	parts := strings.Split(val, "'")
	quoted := make([]string, len(parts))
	for i, p := range parts {
		quoted[i] = "'" + p + "'"
	}

	return "concat(" + strings.Join(quoted, `,"'",`) + ")"
}

// AsEntryXpath returns the given values as an entry xpath segment.
func AsEntryXpath(vals ...string) string {
	if len(vals) == 0 || (len(vals) == 1 && vals[0] == "") {
		return "entry"
	}

	var buf bytes.Buffer

	buf.WriteString("entry[")
	for i := range vals {
		if i != 0 {
			buf.WriteString(" or ")
		}
		buf.WriteString("@name=")
		buf.WriteString(xpathSafe(vals[i]))
	}
	buf.WriteString("]")

	return buf.String()
}

// AsUuidXpath returns an xpath segment as a UUID location.
func AsUuidXpath(v string) string {
	return "entry[@uuid=" + xpathSafe(v) + "]"
}

// EntryName is the inverse of AsEntryXpath for a single value: given an entry
// xpath segment such as entry[@name='foo'] or the injection-safe
// entry[@name=concat('a',"'",'b')] form produced by xpathSafe, it returns the
// literal name. Any input that is not a recognizable entry-name predicate
// (including the bare "entry" listing segment) is returned unchanged.
func EntryName(component string) string {
	const prefix = "entry[@name="
	const suffix = "]"

	expr := component
	if strings.HasPrefix(expr, prefix) && strings.HasSuffix(expr, suffix) {
		expr = expr[len(prefix) : len(expr)-len(suffix)]
	}

	return xpathUnquote(expr)
}

// xpathUnquote decodes an XPath string expression produced by xpathSafe back to
// its literal value. It handles both a single-quoted literal ('foo') and the
// concat('a',"'",'b') form. Unrecognized input is returned unchanged.
func xpathUnquote(expr string) string {
	// Injection-safe concat(...) form.
	if strings.HasPrefix(expr, "concat(") && strings.HasSuffix(expr, ")") {
		body := expr[len("concat(") : len(expr)-1]

		var buf bytes.Buffer
		for i := 0; i < len(body); {
			switch body[i] {
			case '\'':
				// A single-quoted segment: '...'. Segments never contain a
				// single quote because that is xpathSafe's split delimiter.
				j := strings.IndexByte(body[i+1:], '\'')
				if j < 0 {
					return expr // malformed; leave as-is.
				}
				buf.WriteString(body[i+1 : i+1+j])
				i += j + 2
			case '"':
				// The literal single quote is always emitted as "'".
				buf.WriteByte('\'')
				i += 3
			case ',':
				i++
			default:
				return expr // unexpected token; leave as-is.
			}
		}

		return buf.String()
	}

	// Plain single-quoted literal.
	if len(expr) >= 2 && expr[0] == '\'' && expr[len(expr)-1] == '\'' {
		return expr[1 : len(expr)-1]
	}

	return expr
}

// AsMemberXpath returns the given values as a member xpath segment.
func AsMemberXpath(vals []string) string {
	var buf bytes.Buffer

	buf.WriteString("member[")
	for i := range vals {
		if i != 0 {
			buf.WriteString(" or ")
		}
		buf.WriteString("text()=")
		buf.WriteString(xpathSafe(vals[i]))
	}

	buf.WriteString("]")

	return buf.String()
}

// TemplateXpathPrefix returns the template xpath prefix of the given template name.
func TemplateXpathPrefix(tmpl, ts string) []string {
	if tmpl != "" {
		return []string{
			"config",
			"devices",
			AsEntryXpath("localhost.localdomain"),
			"template",
			AsEntryXpath(tmpl),
		}
	}

	return []string{
		"config",
		"devices",
		AsEntryXpath("localhost.localdomain"),
		"template-stack",
		AsEntryXpath(ts),
	}
}

// DeviceGroupXpathPrefix returns a device group xpath prefix.
// If the device group is empty, then the default is "shared".
func DeviceGroupXpathPrefix(dg string) []string {
	if dg == "" || dg == "shared" {
		return []string{"config", "shared"}
	}

	return []string{
		"config",
		"devices",
		AsEntryXpath("localhost.localdomain"),
		"device-group",
		AsEntryXpath(dg),
	}
}

// VsysXpathPrefix returns a vsys xpath prefix.
func VsysXpathPrefix(vsys string) []string {
	if vsys == "" {
		vsys = "vsys1"
	} else if vsys == "shared" {
		return []string{"config", "shared"}
	}

	return []string{
		"config",
		"devices",
		AsEntryXpath("localhost.localdomain"),
		"vsys",
		AsEntryXpath(vsys),
	}
}

// PanoramaXpathPrefix returns the panorama xpath prefix.
func PanoramaXpathPrefix() []string {
	return []string{
		"config",
		"panorama",
	}
}

// StripPanosPackaging removes the response / result and an optional third
// containing XML tag from the given byte slice.
func StripPanosPackaging(input []byte, tag string) []byte {
	var index int
	gt := []byte(">")
	lt := []byte("<")

	// Remove response.
	index = bytes.Index(input, gt)
	ans := input[index+1:]
	index = bytes.LastIndex(ans, lt)
	ans = ans[:index]

	// Remove result.
	index = bytes.Index(ans, gt)
	ans = ans[index+1:]
	index = bytes.LastIndex(ans, lt)
	ans = ans[:index]

	ans = bytes.TrimSpace(ans)

	if tag != "" {
		if bytes.HasPrefix(ans, []byte("<"+tag+" ")) || bytes.HasPrefix(ans, []byte("<"+tag+">")) {
			index = bytes.Index(ans, gt)
			ans = ans[index+1:]
			if len(ans) > 0 {
				index = bytes.LastIndex(ans, lt)
				ans = ans[:index]
				ans = bytes.TrimSpace(ans)
			}
		}
	}

	return ans
}

// CdataText is for getting CDATA contents of XML docs.
type CdataText struct {
	Text string `xml:",cdata"`
}

// RawXml is what allows the use of Edit commands on a XPATH without
// truncating any other child objects that may be attached to it.
type RawXml struct {
	Text string `xml:",innerxml"`
}

// CleanRawXml removes extra XML attributes from RawXml objects without
// requiring us to have to parse everything.
func CleanRawXml(v string) string {
	re := regexp.MustCompile(` admin="\S+" dirtyId="\d+" time="\d{4}/\d{2}/\d{2} \d{2}:\d{2}:\d{2}"`)
	return re.ReplaceAllString(v, "")
}

// ValidMovement returns if the movement constant is valid or not.
func ValidMovement(v int) bool {
	switch v {
	case MoveSkip, MoveBefore, MoveDirectlyBefore, MoveAfter, MoveDirectlyAfter, MoveTop, MoveBottom:
		return true
	}

	return false
}

// RelativeMovement returns if the movement constant is a relative movement.
func RelativeMovement(v int) bool {
	switch v {
	case MoveBefore, MoveDirectlyBefore, MoveAfter, MoveDirectlyAfter:
		return true
	}

	return false
}

// ValidateRulebase validates the device group and rulebase pairing for
// Panorama policies.
func ValidateRulebase(dg, base string) error {
	switch base {
	case "":
		return fmt.Errorf("rulebase must be specified")
	case Rulebase:
		if dg != "shared" {
			return fmt.Errorf("rulebase %q requires \"shared\" device group", base)
		}
	case PreRulebase, PostRulebase:
	default:
		return fmt.Errorf("unknown rulebase %q", base)
	}

	return nil
}
