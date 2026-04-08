package password

import (
	"encoding/xml"
	"fmt"

	"github.com/PaloAltoNetworks/pango/filtering"
	"github.com/PaloAltoNetworks/pango/generic"
	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

var (
	_ filtering.Fielder = &Entry{}
)

var (
	suffix = []string{"password-profile", "$name"}
)

type Entry struct {
	Name           string
	PasswordChange *PasswordChange
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type PasswordChange struct {
	ExpirationPeriod              *int64
	ExpirationWarningPeriod       *int64
	PostExpirationAdminLoginCount *int64
	PostExpirationGracePeriod     *int64
	Misc                          []generic.Xml
	MiscAttributes                []xml.Attr
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

func (o *entryXmlContainer) Normalize() ([]*Entry, error) {
	entries := make([]*Entry, 0, len(o.Answer))
	for _, elt := range o.Answer {
		obj, err := elt.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		entries = append(entries, obj)
	}

	return entries, nil
}

func specifyEntry(source *Entry) (any, error) {
	var obj entryXml
	obj.MarshalFromObject(*source)
	return obj, nil
}

type entryXml struct {
	XMLName        xml.Name           `xml:"entry"`
	Name           string             `xml:"name,attr"`
	PasswordChange *passwordChangeXml `xml:"password-change,omitempty"`
	Misc           []generic.Xml      `xml:",any"`
	MiscAttributes []xml.Attr         `xml:",any,attr"`
}
type passwordChangeXml struct {
	ExpirationPeriod              *int64        `xml:"expiration-period,omitempty"`
	ExpirationWarningPeriod       *int64        `xml:"expiration-warning-period,omitempty"`
	PostExpirationAdminLoginCount *int64        `xml:"post-expiration-admin-login-count,omitempty"`
	PostExpirationGracePeriod     *int64        `xml:"post-expiration-grace-period,omitempty"`
	Misc                          []generic.Xml `xml:",any"`
	MiscAttributes                []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	if s.PasswordChange != nil {
		var obj passwordChangeXml
		obj.MarshalFromObject(*s.PasswordChange)
		o.PasswordChange = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var passwordChangeVal *PasswordChange
	if o.PasswordChange != nil {
		obj, err := o.PasswordChange.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		passwordChangeVal = obj
	}

	result := &Entry{
		Name:           o.Name,
		PasswordChange: passwordChangeVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *passwordChangeXml) MarshalFromObject(s PasswordChange) {
	o.ExpirationPeriod = s.ExpirationPeriod
	o.ExpirationWarningPeriod = s.ExpirationWarningPeriod
	o.PostExpirationAdminLoginCount = s.PostExpirationAdminLoginCount
	o.PostExpirationGracePeriod = s.PostExpirationGracePeriod
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o passwordChangeXml) UnmarshalToObject() (*PasswordChange, error) {

	result := &PasswordChange{
		ExpirationPeriod:              o.ExpirationPeriod,
		ExpirationWarningPeriod:       o.ExpirationWarningPeriod,
		PostExpirationAdminLoginCount: o.PostExpirationAdminLoginCount,
		PostExpirationGracePeriod:     o.PostExpirationGracePeriod,
		Misc:                          o.Misc,
		MiscAttributes:                o.MiscAttributes,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "password_change" || v == "PasswordChange" {
		return e.PasswordChange, nil
	}

	return nil, fmt.Errorf("unknown field")
}

func Versioning(vn version.Number) (Specifier, Normalizer, error) {

	return specifyEntry, &entryXmlContainer{}, nil
}
func SpecMatches(a, b *Entry) bool {
	if a == nil && b == nil {
		return true
	}

	if (a == nil && b != nil) || (a != nil && b == nil) {
		return false
	}

	return a.matches(b)
}

func (o *Entry) matches(other *Entry) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.PasswordChange.matches(other.PasswordChange) {
		return false
	}

	return true
}

func (o *PasswordChange) matches(other *PasswordChange) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.ExpirationPeriod, other.ExpirationPeriod) {
		return false
	}
	if !util.Ints64Match(o.ExpirationWarningPeriod, other.ExpirationWarningPeriod) {
		return false
	}
	if !util.Ints64Match(o.PostExpirationAdminLoginCount, other.PostExpirationAdminLoginCount) {
		return false
	}
	if !util.Ints64Match(o.PostExpirationGracePeriod, other.PostExpirationGracePeriod) {
		return false
	}

	return true
}

func (o *Entry) EntryName() string {
	return o.Name
}

func (o *Entry) SetEntryName(name string) {
	o.Name = name
}

func (o *Entry) GetMiscAttributes() []xml.Attr {
	return o.MiscAttributes
}

func (o *Entry) SetMiscAttributes(attrs []xml.Attr) {
	o.MiscAttributes = attrs
}
