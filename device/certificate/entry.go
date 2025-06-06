package certificate

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
	suffix = []string{"certificate", "$name"}
)

type Entry struct {
	Name            string
	Algorithm       *string
	Ca              *bool
	ExpiryEpoch     *string
	Issuer          *string
	IssuerHash      *string
	NotValidAfter   *string
	NotValidBefore  *string
	RevokeDateEpoch *string
	Status          *string
	Subject         *string
	SubjectHash     *string
	CloudResourceId *CloudResourceId
	CommonName      *string
	Csr             *string
	PrivateKey      *string
	PrivateKeyOnHsm *bool
	PublicKey       *string
	Misc            []generic.Xml
}
type CloudResourceId struct {
	Aws   *CloudResourceIdAws
	Azure *CloudResourceIdAzure
	Misc  []generic.Xml
}
type CloudResourceIdAws struct {
	Secret *string
	Misc   []generic.Xml
}
type CloudResourceIdAzure struct {
	KeyVaultUri *string
	Secret      *string
	Misc        []generic.Xml
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXmlContainer_11_0_2 struct {
	Answer []entryXml_11_0_2 `xml:"entry"`
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
func (o *entryXmlContainer_11_0_2) Normalize() ([]*Entry, error) {
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
func specifyEntry_11_0_2(source *Entry) (any, error) {
	var obj entryXml_11_0_2
	obj.MarshalFromObject(*source)
	return obj, nil
}

type entryXml struct {
	XMLName         xml.Name            `xml:"entry"`
	Name            string              `xml:"name,attr"`
	Algorithm       *string             `xml:"algorithm,omitempty"`
	Ca              *string             `xml:"ca,omitempty"`
	ExpiryEpoch     *string             `xml:"expiry-epoch,omitempty"`
	Issuer          *string             `xml:"issuer,omitempty"`
	IssuerHash      *string             `xml:"issuer-hash,omitempty"`
	NotValidAfter   *string             `xml:"not-valid-after,omitempty"`
	NotValidBefore  *string             `xml:"not-valid-before,omitempty"`
	RevokeDateEpoch *string             `xml:"revoke-date-epoch,omitempty"`
	Status          *string             `xml:"status,omitempty"`
	Subject         *string             `xml:"subject,omitempty"`
	SubjectHash     *string             `xml:"subject-hash,omitempty"`
	CloudResourceId *cloudResourceIdXml `xml:"cloud-resource-id,omitempty"`
	CommonName      *string             `xml:"common-name,omitempty"`
	Csr             *string             `xml:"csr,omitempty"`
	PrivateKey      *string             `xml:"private-key,omitempty"`
	PrivateKeyOnHsm *string             `xml:"private-key-on-hsm,omitempty"`
	PublicKey       *string             `xml:"public-key,omitempty"`
	Misc            []generic.Xml       `xml:",any"`
}
type cloudResourceIdXml struct {
	Aws   *cloudResourceIdAwsXml   `xml:"aws,omitempty"`
	Azure *cloudResourceIdAzureXml `xml:"azure,omitempty"`
	Misc  []generic.Xml            `xml:",any"`
}
type cloudResourceIdAwsXml struct {
	Secret *string       `xml:"secret,omitempty"`
	Misc   []generic.Xml `xml:",any"`
}
type cloudResourceIdAzureXml struct {
	KeyVaultUri *string       `xml:"key-vault-uri,omitempty"`
	Secret      *string       `xml:"secret,omitempty"`
	Misc        []generic.Xml `xml:",any"`
}
type entryXml_11_0_2 struct {
	XMLName         xml.Name                   `xml:"entry"`
	Name            string                     `xml:"name,attr"`
	Algorithm       *string                    `xml:"algorithm,omitempty"`
	Ca              *string                    `xml:"ca,omitempty"`
	ExpiryEpoch     *string                    `xml:"expiry-epoch,omitempty"`
	Issuer          *string                    `xml:"issuer,omitempty"`
	IssuerHash      *string                    `xml:"issuer-hash,omitempty"`
	NotValidAfter   *string                    `xml:"not-valid-after,omitempty"`
	NotValidBefore  *string                    `xml:"not-valid-before,omitempty"`
	RevokeDateEpoch *string                    `xml:"revoke-date-epoch,omitempty"`
	Status          *string                    `xml:"status,omitempty"`
	Subject         *string                    `xml:"subject,omitempty"`
	SubjectHash     *string                    `xml:"subject-hash,omitempty"`
	CloudResourceId *cloudResourceIdXml_11_0_2 `xml:"cloud-resource-id,omitempty"`
	CommonName      *string                    `xml:"common-name,omitempty"`
	Csr             *string                    `xml:"csr,omitempty"`
	PrivateKey      *string                    `xml:"private-key,omitempty"`
	PrivateKeyOnHsm *string                    `xml:"private-key-on-hsm,omitempty"`
	PublicKey       *string                    `xml:"public-key,omitempty"`
	Misc            []generic.Xml              `xml:",any"`
}
type cloudResourceIdXml_11_0_2 struct {
	Aws   *cloudResourceIdAwsXml_11_0_2   `xml:"aws,omitempty"`
	Azure *cloudResourceIdAzureXml_11_0_2 `xml:"azure,omitempty"`
	Misc  []generic.Xml                   `xml:",any"`
}
type cloudResourceIdAwsXml_11_0_2 struct {
	Secret *string       `xml:"secret,omitempty"`
	Misc   []generic.Xml `xml:",any"`
}
type cloudResourceIdAzureXml_11_0_2 struct {
	KeyVaultUri *string       `xml:"key-vault-uri,omitempty"`
	Secret      *string       `xml:"secret,omitempty"`
	Misc        []generic.Xml `xml:",any"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.Algorithm = s.Algorithm
	o.Ca = util.YesNo(s.Ca, nil)
	o.ExpiryEpoch = s.ExpiryEpoch
	o.Issuer = s.Issuer
	o.IssuerHash = s.IssuerHash
	o.NotValidAfter = s.NotValidAfter
	o.NotValidBefore = s.NotValidBefore
	o.RevokeDateEpoch = s.RevokeDateEpoch
	o.Status = s.Status
	o.Subject = s.Subject
	o.SubjectHash = s.SubjectHash
	if s.CloudResourceId != nil {
		var obj cloudResourceIdXml
		obj.MarshalFromObject(*s.CloudResourceId)
		o.CloudResourceId = &obj
	}
	o.CommonName = s.CommonName
	o.Csr = s.Csr
	o.PrivateKey = s.PrivateKey
	o.PrivateKeyOnHsm = util.YesNo(s.PrivateKeyOnHsm, nil)
	o.PublicKey = s.PublicKey
	o.Misc = s.Misc
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var cloudResourceIdVal *CloudResourceId
	if o.CloudResourceId != nil {
		obj, err := o.CloudResourceId.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		cloudResourceIdVal = obj
	}

	result := &Entry{
		Name:            o.Name,
		Algorithm:       o.Algorithm,
		Ca:              util.AsBool(o.Ca, nil),
		ExpiryEpoch:     o.ExpiryEpoch,
		Issuer:          o.Issuer,
		IssuerHash:      o.IssuerHash,
		NotValidAfter:   o.NotValidAfter,
		NotValidBefore:  o.NotValidBefore,
		RevokeDateEpoch: o.RevokeDateEpoch,
		Status:          o.Status,
		Subject:         o.Subject,
		SubjectHash:     o.SubjectHash,
		CloudResourceId: cloudResourceIdVal,
		CommonName:      o.CommonName,
		Csr:             o.Csr,
		PrivateKey:      o.PrivateKey,
		PrivateKeyOnHsm: util.AsBool(o.PrivateKeyOnHsm, nil),
		PublicKey:       o.PublicKey,
		Misc:            o.Misc,
	}
	return result, nil
}
func (o *cloudResourceIdXml) MarshalFromObject(s CloudResourceId) {
	if s.Aws != nil {
		var obj cloudResourceIdAwsXml
		obj.MarshalFromObject(*s.Aws)
		o.Aws = &obj
	}
	if s.Azure != nil {
		var obj cloudResourceIdAzureXml
		obj.MarshalFromObject(*s.Azure)
		o.Azure = &obj
	}
	o.Misc = s.Misc
}

func (o cloudResourceIdXml) UnmarshalToObject() (*CloudResourceId, error) {
	var awsVal *CloudResourceIdAws
	if o.Aws != nil {
		obj, err := o.Aws.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		awsVal = obj
	}
	var azureVal *CloudResourceIdAzure
	if o.Azure != nil {
		obj, err := o.Azure.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		azureVal = obj
	}

	result := &CloudResourceId{
		Aws:   awsVal,
		Azure: azureVal,
		Misc:  o.Misc,
	}
	return result, nil
}
func (o *cloudResourceIdAwsXml) MarshalFromObject(s CloudResourceIdAws) {
	o.Secret = s.Secret
	o.Misc = s.Misc
}

func (o cloudResourceIdAwsXml) UnmarshalToObject() (*CloudResourceIdAws, error) {

	result := &CloudResourceIdAws{
		Secret: o.Secret,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *cloudResourceIdAzureXml) MarshalFromObject(s CloudResourceIdAzure) {
	o.KeyVaultUri = s.KeyVaultUri
	o.Secret = s.Secret
	o.Misc = s.Misc
}

func (o cloudResourceIdAzureXml) UnmarshalToObject() (*CloudResourceIdAzure, error) {

	result := &CloudResourceIdAzure{
		KeyVaultUri: o.KeyVaultUri,
		Secret:      o.Secret,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *entryXml_11_0_2) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.Algorithm = s.Algorithm
	o.Ca = util.YesNo(s.Ca, nil)
	o.ExpiryEpoch = s.ExpiryEpoch
	o.Issuer = s.Issuer
	o.IssuerHash = s.IssuerHash
	o.NotValidAfter = s.NotValidAfter
	o.NotValidBefore = s.NotValidBefore
	o.RevokeDateEpoch = s.RevokeDateEpoch
	o.Status = s.Status
	o.Subject = s.Subject
	o.SubjectHash = s.SubjectHash
	if s.CloudResourceId != nil {
		var obj cloudResourceIdXml_11_0_2
		obj.MarshalFromObject(*s.CloudResourceId)
		o.CloudResourceId = &obj
	}
	o.CommonName = s.CommonName
	o.Csr = s.Csr
	o.PrivateKey = s.PrivateKey
	o.PrivateKeyOnHsm = util.YesNo(s.PrivateKeyOnHsm, nil)
	o.PublicKey = s.PublicKey
	o.Misc = s.Misc
}

func (o entryXml_11_0_2) UnmarshalToObject() (*Entry, error) {
	var cloudResourceIdVal *CloudResourceId
	if o.CloudResourceId != nil {
		obj, err := o.CloudResourceId.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		cloudResourceIdVal = obj
	}

	result := &Entry{
		Name:            o.Name,
		Algorithm:       o.Algorithm,
		Ca:              util.AsBool(o.Ca, nil),
		ExpiryEpoch:     o.ExpiryEpoch,
		Issuer:          o.Issuer,
		IssuerHash:      o.IssuerHash,
		NotValidAfter:   o.NotValidAfter,
		NotValidBefore:  o.NotValidBefore,
		RevokeDateEpoch: o.RevokeDateEpoch,
		Status:          o.Status,
		Subject:         o.Subject,
		SubjectHash:     o.SubjectHash,
		CloudResourceId: cloudResourceIdVal,
		CommonName:      o.CommonName,
		Csr:             o.Csr,
		PrivateKey:      o.PrivateKey,
		PrivateKeyOnHsm: util.AsBool(o.PrivateKeyOnHsm, nil),
		PublicKey:       o.PublicKey,
		Misc:            o.Misc,
	}
	return result, nil
}
func (o *cloudResourceIdXml_11_0_2) MarshalFromObject(s CloudResourceId) {
	if s.Aws != nil {
		var obj cloudResourceIdAwsXml_11_0_2
		obj.MarshalFromObject(*s.Aws)
		o.Aws = &obj
	}
	if s.Azure != nil {
		var obj cloudResourceIdAzureXml_11_0_2
		obj.MarshalFromObject(*s.Azure)
		o.Azure = &obj
	}
	o.Misc = s.Misc
}

func (o cloudResourceIdXml_11_0_2) UnmarshalToObject() (*CloudResourceId, error) {
	var awsVal *CloudResourceIdAws
	if o.Aws != nil {
		obj, err := o.Aws.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		awsVal = obj
	}
	var azureVal *CloudResourceIdAzure
	if o.Azure != nil {
		obj, err := o.Azure.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		azureVal = obj
	}

	result := &CloudResourceId{
		Aws:   awsVal,
		Azure: azureVal,
		Misc:  o.Misc,
	}
	return result, nil
}
func (o *cloudResourceIdAwsXml_11_0_2) MarshalFromObject(s CloudResourceIdAws) {
	o.Secret = s.Secret
	o.Misc = s.Misc
}

func (o cloudResourceIdAwsXml_11_0_2) UnmarshalToObject() (*CloudResourceIdAws, error) {

	result := &CloudResourceIdAws{
		Secret: o.Secret,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *cloudResourceIdAzureXml_11_0_2) MarshalFromObject(s CloudResourceIdAzure) {
	o.KeyVaultUri = s.KeyVaultUri
	o.Secret = s.Secret
	o.Misc = s.Misc
}

func (o cloudResourceIdAzureXml_11_0_2) UnmarshalToObject() (*CloudResourceIdAzure, error) {

	result := &CloudResourceIdAzure{
		KeyVaultUri: o.KeyVaultUri,
		Secret:      o.Secret,
		Misc:        o.Misc,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "algorithm" || v == "Algorithm" {
		return e.Algorithm, nil
	}
	if v == "ca" || v == "Ca" {
		return e.Ca, nil
	}
	if v == "expiry_epoch" || v == "ExpiryEpoch" {
		return e.ExpiryEpoch, nil
	}
	if v == "issuer" || v == "Issuer" {
		return e.Issuer, nil
	}
	if v == "issuer_hash" || v == "IssuerHash" {
		return e.IssuerHash, nil
	}
	if v == "not_valid_after" || v == "NotValidAfter" {
		return e.NotValidAfter, nil
	}
	if v == "not_valid_before" || v == "NotValidBefore" {
		return e.NotValidBefore, nil
	}
	if v == "revoke_date_epoch" || v == "RevokeDateEpoch" {
		return e.RevokeDateEpoch, nil
	}
	if v == "status" || v == "Status" {
		return e.Status, nil
	}
	if v == "subject" || v == "Subject" {
		return e.Subject, nil
	}
	if v == "subject_hash" || v == "SubjectHash" {
		return e.SubjectHash, nil
	}
	if v == "cloud_resource_id" || v == "CloudResourceId" {
		return e.CloudResourceId, nil
	}
	if v == "common_name" || v == "CommonName" {
		return e.CommonName, nil
	}
	if v == "csr" || v == "Csr" {
		return e.Csr, nil
	}
	if v == "private_key" || v == "PrivateKey" {
		return e.PrivateKey, nil
	}
	if v == "private_key_on_hsm" || v == "PrivateKeyOnHsm" {
		return e.PrivateKeyOnHsm, nil
	}
	if v == "public_key" || v == "PublicKey" {
		return e.PublicKey, nil
	}

	return nil, fmt.Errorf("unknown field")
}

func Versioning(vn version.Number) (Specifier, Normalizer, error) {
	version_11_0_2, _ := version.New("11.0.2")
	version_11_1_0, _ := version.New("11.1.0")
	if vn.Gte(version_11_0_2) && vn.Lt(version_11_1_0) {
		return specifyEntry_11_0_2, &entryXmlContainer_11_0_2{}, nil
	}

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
	if !util.StringsMatch(o.Algorithm, other.Algorithm) {
		return false
	}
	if !util.BoolsMatch(o.Ca, other.Ca) {
		return false
	}
	if !util.StringsMatch(o.ExpiryEpoch, other.ExpiryEpoch) {
		return false
	}
	if !util.StringsMatch(o.Issuer, other.Issuer) {
		return false
	}
	if !util.StringsMatch(o.IssuerHash, other.IssuerHash) {
		return false
	}
	if !util.StringsMatch(o.NotValidAfter, other.NotValidAfter) {
		return false
	}
	if !util.StringsMatch(o.NotValidBefore, other.NotValidBefore) {
		return false
	}
	if !util.StringsMatch(o.RevokeDateEpoch, other.RevokeDateEpoch) {
		return false
	}
	if !util.StringsMatch(o.Status, other.Status) {
		return false
	}
	if !util.StringsMatch(o.Subject, other.Subject) {
		return false
	}
	if !util.StringsMatch(o.SubjectHash, other.SubjectHash) {
		return false
	}
	if !o.CloudResourceId.matches(other.CloudResourceId) {
		return false
	}
	if !util.StringsMatch(o.CommonName, other.CommonName) {
		return false
	}
	if !util.StringsMatch(o.Csr, other.Csr) {
		return false
	}
	if !util.StringsMatch(o.PrivateKey, other.PrivateKey) {
		return false
	}
	if !util.BoolsMatch(o.PrivateKeyOnHsm, other.PrivateKeyOnHsm) {
		return false
	}
	if !util.StringsMatch(o.PublicKey, other.PublicKey) {
		return false
	}

	return true
}

func (o *CloudResourceId) matches(other *CloudResourceId) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Aws.matches(other.Aws) {
		return false
	}
	if !o.Azure.matches(other.Azure) {
		return false
	}

	return true
}

func (o *CloudResourceIdAws) matches(other *CloudResourceIdAws) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Secret, other.Secret) {
		return false
	}

	return true
}

func (o *CloudResourceIdAzure) matches(other *CloudResourceIdAzure) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.KeyVaultUri, other.KeyVaultUri) {
		return false
	}
	if !util.StringsMatch(o.Secret, other.Secret) {
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
