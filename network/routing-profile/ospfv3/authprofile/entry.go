package authprofile

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
	suffix = []string{"network", "routing-profile", "ospfv3", "auth-profile", "$name"}
)

type Entry struct {
	Name           string
	Spi            *string
	Ah             *Ah
	Esp            *Esp
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ah struct {
	Md5            *AhMd5
	Sha1           *AhSha1
	Sha256         *AhSha256
	Sha384         *AhSha384
	Sha512         *AhSha512
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type AhMd5 struct {
	Key            *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type AhSha1 struct {
	Key            *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type AhSha256 struct {
	Key            *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type AhSha384 struct {
	Key            *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type AhSha512 struct {
	Key            *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Esp struct {
	Authentication *EspAuthentication
	Encryption     *EspEncryption
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type EspAuthentication struct {
	Md5            *EspAuthenticationMd5
	None           *EspAuthenticationNone
	Sha1           *EspAuthenticationSha1
	Sha256         *EspAuthenticationSha256
	Sha384         *EspAuthenticationSha384
	Sha512         *EspAuthenticationSha512
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type EspAuthenticationMd5 struct {
	Key            *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type EspAuthenticationNone struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type EspAuthenticationSha1 struct {
	Key            *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type EspAuthenticationSha256 struct {
	Key            *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type EspAuthenticationSha384 struct {
	Key            *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type EspAuthenticationSha512 struct {
	Key            *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type EspEncryption struct {
	Algorithm      *string
	Key            *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
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
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Spi            *string       `xml:"spi,omitempty"`
	Ah             *ahXml        `xml:"ah,omitempty"`
	Esp            *espXml       `xml:"esp,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ahXml struct {
	Md5            *ahMd5Xml     `xml:"md5,omitempty"`
	Sha1           *ahSha1Xml    `xml:"sha1,omitempty"`
	Sha256         *ahSha256Xml  `xml:"sha256,omitempty"`
	Sha384         *ahSha384Xml  `xml:"sha384,omitempty"`
	Sha512         *ahSha512Xml  `xml:"sha512,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ahMd5Xml struct {
	Key            *string       `xml:"key,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ahSha1Xml struct {
	Key            *string       `xml:"key,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ahSha256Xml struct {
	Key            *string       `xml:"key,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ahSha384Xml struct {
	Key            *string       `xml:"key,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ahSha512Xml struct {
	Key            *string       `xml:"key,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type espXml struct {
	Authentication *espAuthenticationXml `xml:"authentication,omitempty"`
	Encryption     *espEncryptionXml     `xml:"encryption,omitempty"`
	Misc           []generic.Xml         `xml:",any"`
	MiscAttributes []xml.Attr            `xml:",any,attr"`
}
type espAuthenticationXml struct {
	Md5            *espAuthenticationMd5Xml    `xml:"md5,omitempty"`
	None           *espAuthenticationNoneXml   `xml:"none,omitempty"`
	Sha1           *espAuthenticationSha1Xml   `xml:"sha1,omitempty"`
	Sha256         *espAuthenticationSha256Xml `xml:"sha256,omitempty"`
	Sha384         *espAuthenticationSha384Xml `xml:"sha384,omitempty"`
	Sha512         *espAuthenticationSha512Xml `xml:"sha512,omitempty"`
	Misc           []generic.Xml               `xml:",any"`
	MiscAttributes []xml.Attr                  `xml:",any,attr"`
}
type espAuthenticationMd5Xml struct {
	Key            *string       `xml:"key,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type espAuthenticationNoneXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type espAuthenticationSha1Xml struct {
	Key            *string       `xml:"key,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type espAuthenticationSha256Xml struct {
	Key            *string       `xml:"key,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type espAuthenticationSha384Xml struct {
	Key            *string       `xml:"key,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type espAuthenticationSha512Xml struct {
	Key            *string       `xml:"key,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type espEncryptionXml struct {
	Algorithm      *string       `xml:"algorithm,omitempty"`
	Key            *string       `xml:"key,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.Spi = s.Spi
	if s.Ah != nil {
		var obj ahXml
		obj.MarshalFromObject(*s.Ah)
		o.Ah = &obj
	}
	if s.Esp != nil {
		var obj espXml
		obj.MarshalFromObject(*s.Esp)
		o.Esp = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var ahVal *Ah
	if o.Ah != nil {
		obj, err := o.Ah.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ahVal = obj
	}
	var espVal *Esp
	if o.Esp != nil {
		obj, err := o.Esp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		espVal = obj
	}

	result := &Entry{
		Name:           o.Name,
		Spi:            o.Spi,
		Ah:             ahVal,
		Esp:            espVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ahXml) MarshalFromObject(s Ah) {
	if s.Md5 != nil {
		var obj ahMd5Xml
		obj.MarshalFromObject(*s.Md5)
		o.Md5 = &obj
	}
	if s.Sha1 != nil {
		var obj ahSha1Xml
		obj.MarshalFromObject(*s.Sha1)
		o.Sha1 = &obj
	}
	if s.Sha256 != nil {
		var obj ahSha256Xml
		obj.MarshalFromObject(*s.Sha256)
		o.Sha256 = &obj
	}
	if s.Sha384 != nil {
		var obj ahSha384Xml
		obj.MarshalFromObject(*s.Sha384)
		o.Sha384 = &obj
	}
	if s.Sha512 != nil {
		var obj ahSha512Xml
		obj.MarshalFromObject(*s.Sha512)
		o.Sha512 = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ahXml) UnmarshalToObject() (*Ah, error) {
	var md5Val *AhMd5
	if o.Md5 != nil {
		obj, err := o.Md5.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		md5Val = obj
	}
	var sha1Val *AhSha1
	if o.Sha1 != nil {
		obj, err := o.Sha1.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sha1Val = obj
	}
	var sha256Val *AhSha256
	if o.Sha256 != nil {
		obj, err := o.Sha256.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sha256Val = obj
	}
	var sha384Val *AhSha384
	if o.Sha384 != nil {
		obj, err := o.Sha384.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sha384Val = obj
	}
	var sha512Val *AhSha512
	if o.Sha512 != nil {
		obj, err := o.Sha512.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sha512Val = obj
	}

	result := &Ah{
		Md5:            md5Val,
		Sha1:           sha1Val,
		Sha256:         sha256Val,
		Sha384:         sha384Val,
		Sha512:         sha512Val,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ahMd5Xml) MarshalFromObject(s AhMd5) {
	o.Key = s.Key
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ahMd5Xml) UnmarshalToObject() (*AhMd5, error) {

	result := &AhMd5{
		Key:            o.Key,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ahSha1Xml) MarshalFromObject(s AhSha1) {
	o.Key = s.Key
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ahSha1Xml) UnmarshalToObject() (*AhSha1, error) {

	result := &AhSha1{
		Key:            o.Key,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ahSha256Xml) MarshalFromObject(s AhSha256) {
	o.Key = s.Key
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ahSha256Xml) UnmarshalToObject() (*AhSha256, error) {

	result := &AhSha256{
		Key:            o.Key,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ahSha384Xml) MarshalFromObject(s AhSha384) {
	o.Key = s.Key
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ahSha384Xml) UnmarshalToObject() (*AhSha384, error) {

	result := &AhSha384{
		Key:            o.Key,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ahSha512Xml) MarshalFromObject(s AhSha512) {
	o.Key = s.Key
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ahSha512Xml) UnmarshalToObject() (*AhSha512, error) {

	result := &AhSha512{
		Key:            o.Key,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *espXml) MarshalFromObject(s Esp) {
	if s.Authentication != nil {
		var obj espAuthenticationXml
		obj.MarshalFromObject(*s.Authentication)
		o.Authentication = &obj
	}
	if s.Encryption != nil {
		var obj espEncryptionXml
		obj.MarshalFromObject(*s.Encryption)
		o.Encryption = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o espXml) UnmarshalToObject() (*Esp, error) {
	var authenticationVal *EspAuthentication
	if o.Authentication != nil {
		obj, err := o.Authentication.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		authenticationVal = obj
	}
	var encryptionVal *EspEncryption
	if o.Encryption != nil {
		obj, err := o.Encryption.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		encryptionVal = obj
	}

	result := &Esp{
		Authentication: authenticationVal,
		Encryption:     encryptionVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *espAuthenticationXml) MarshalFromObject(s EspAuthentication) {
	if s.Md5 != nil {
		var obj espAuthenticationMd5Xml
		obj.MarshalFromObject(*s.Md5)
		o.Md5 = &obj
	}
	if s.None != nil {
		var obj espAuthenticationNoneXml
		obj.MarshalFromObject(*s.None)
		o.None = &obj
	}
	if s.Sha1 != nil {
		var obj espAuthenticationSha1Xml
		obj.MarshalFromObject(*s.Sha1)
		o.Sha1 = &obj
	}
	if s.Sha256 != nil {
		var obj espAuthenticationSha256Xml
		obj.MarshalFromObject(*s.Sha256)
		o.Sha256 = &obj
	}
	if s.Sha384 != nil {
		var obj espAuthenticationSha384Xml
		obj.MarshalFromObject(*s.Sha384)
		o.Sha384 = &obj
	}
	if s.Sha512 != nil {
		var obj espAuthenticationSha512Xml
		obj.MarshalFromObject(*s.Sha512)
		o.Sha512 = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o espAuthenticationXml) UnmarshalToObject() (*EspAuthentication, error) {
	var md5Val *EspAuthenticationMd5
	if o.Md5 != nil {
		obj, err := o.Md5.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		md5Val = obj
	}
	var noneVal *EspAuthenticationNone
	if o.None != nil {
		obj, err := o.None.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noneVal = obj
	}
	var sha1Val *EspAuthenticationSha1
	if o.Sha1 != nil {
		obj, err := o.Sha1.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sha1Val = obj
	}
	var sha256Val *EspAuthenticationSha256
	if o.Sha256 != nil {
		obj, err := o.Sha256.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sha256Val = obj
	}
	var sha384Val *EspAuthenticationSha384
	if o.Sha384 != nil {
		obj, err := o.Sha384.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sha384Val = obj
	}
	var sha512Val *EspAuthenticationSha512
	if o.Sha512 != nil {
		obj, err := o.Sha512.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sha512Val = obj
	}

	result := &EspAuthentication{
		Md5:            md5Val,
		None:           noneVal,
		Sha1:           sha1Val,
		Sha256:         sha256Val,
		Sha384:         sha384Val,
		Sha512:         sha512Val,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *espAuthenticationMd5Xml) MarshalFromObject(s EspAuthenticationMd5) {
	o.Key = s.Key
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o espAuthenticationMd5Xml) UnmarshalToObject() (*EspAuthenticationMd5, error) {

	result := &EspAuthenticationMd5{
		Key:            o.Key,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *espAuthenticationNoneXml) MarshalFromObject(s EspAuthenticationNone) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o espAuthenticationNoneXml) UnmarshalToObject() (*EspAuthenticationNone, error) {

	result := &EspAuthenticationNone{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *espAuthenticationSha1Xml) MarshalFromObject(s EspAuthenticationSha1) {
	o.Key = s.Key
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o espAuthenticationSha1Xml) UnmarshalToObject() (*EspAuthenticationSha1, error) {

	result := &EspAuthenticationSha1{
		Key:            o.Key,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *espAuthenticationSha256Xml) MarshalFromObject(s EspAuthenticationSha256) {
	o.Key = s.Key
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o espAuthenticationSha256Xml) UnmarshalToObject() (*EspAuthenticationSha256, error) {

	result := &EspAuthenticationSha256{
		Key:            o.Key,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *espAuthenticationSha384Xml) MarshalFromObject(s EspAuthenticationSha384) {
	o.Key = s.Key
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o espAuthenticationSha384Xml) UnmarshalToObject() (*EspAuthenticationSha384, error) {

	result := &EspAuthenticationSha384{
		Key:            o.Key,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *espAuthenticationSha512Xml) MarshalFromObject(s EspAuthenticationSha512) {
	o.Key = s.Key
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o espAuthenticationSha512Xml) UnmarshalToObject() (*EspAuthenticationSha512, error) {

	result := &EspAuthenticationSha512{
		Key:            o.Key,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *espEncryptionXml) MarshalFromObject(s EspEncryption) {
	o.Algorithm = s.Algorithm
	o.Key = s.Key
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o espEncryptionXml) UnmarshalToObject() (*EspEncryption, error) {

	result := &EspEncryption{
		Algorithm:      o.Algorithm,
		Key:            o.Key,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "spi" || v == "Spi" {
		return e.Spi, nil
	}
	if v == "ah" || v == "Ah" {
		return e.Ah, nil
	}
	if v == "esp" || v == "Esp" {
		return e.Esp, nil
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
	if !util.StringsMatch(o.Spi, other.Spi) {
		return false
	}
	if !o.Ah.matches(other.Ah) {
		return false
	}
	if !o.Esp.matches(other.Esp) {
		return false
	}

	return true
}

func (o *Ah) matches(other *Ah) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Md5.matches(other.Md5) {
		return false
	}
	if !o.Sha1.matches(other.Sha1) {
		return false
	}
	if !o.Sha256.matches(other.Sha256) {
		return false
	}
	if !o.Sha384.matches(other.Sha384) {
		return false
	}
	if !o.Sha512.matches(other.Sha512) {
		return false
	}

	return true
}

func (o *AhMd5) matches(other *AhMd5) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
		return false
	}

	return true
}

func (o *AhSha1) matches(other *AhSha1) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
		return false
	}

	return true
}

func (o *AhSha256) matches(other *AhSha256) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
		return false
	}

	return true
}

func (o *AhSha384) matches(other *AhSha384) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
		return false
	}

	return true
}

func (o *AhSha512) matches(other *AhSha512) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
		return false
	}

	return true
}

func (o *Esp) matches(other *Esp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Authentication.matches(other.Authentication) {
		return false
	}
	if !o.Encryption.matches(other.Encryption) {
		return false
	}

	return true
}

func (o *EspAuthentication) matches(other *EspAuthentication) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Md5.matches(other.Md5) {
		return false
	}
	if !o.None.matches(other.None) {
		return false
	}
	if !o.Sha1.matches(other.Sha1) {
		return false
	}
	if !o.Sha256.matches(other.Sha256) {
		return false
	}
	if !o.Sha384.matches(other.Sha384) {
		return false
	}
	if !o.Sha512.matches(other.Sha512) {
		return false
	}

	return true
}

func (o *EspAuthenticationMd5) matches(other *EspAuthenticationMd5) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
		return false
	}

	return true
}

func (o *EspAuthenticationNone) matches(other *EspAuthenticationNone) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *EspAuthenticationSha1) matches(other *EspAuthenticationSha1) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
		return false
	}

	return true
}

func (o *EspAuthenticationSha256) matches(other *EspAuthenticationSha256) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
		return false
	}

	return true
}

func (o *EspAuthenticationSha384) matches(other *EspAuthenticationSha384) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
		return false
	}

	return true
}

func (o *EspAuthenticationSha512) matches(other *EspAuthenticationSha512) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
		return false
	}

	return true
}

func (o *EspEncryption) matches(other *EspEncryption) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Algorithm, other.Algorithm) {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
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
