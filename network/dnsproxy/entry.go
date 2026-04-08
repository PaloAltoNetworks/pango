package dnsproxy

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
	suffix = []string{"network", "dns-proxy", "$name"}
)

type Entry struct {
	Name           string
	Cache          *Cache
	Default        *Default
	DomainServers  []DomainServers
	Enabled        *bool
	Interface      []string
	StaticEntries  []StaticEntries
	TcpQueries     *TcpQueries
	UdpQueries     *UdpQueries
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Cache struct {
	CacheEdns      *bool
	Enabled        *bool
	MaxTtl         *CacheMaxTtl
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type CacheMaxTtl struct {
	Enabled        *bool
	TimeToLive     *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Default struct {
	Inheritance    *DefaultInheritance
	Primary        *string
	Secondary      *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type DefaultInheritance struct {
	Source         *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type DomainServers struct {
	Name           string
	Cacheable      *bool
	Primary        *string
	Secondary      *string
	DomainName     []string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type StaticEntries struct {
	Name           string
	Domain         *string
	Address        []string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type TcpQueries struct {
	Enabled            *bool
	MaxPendingRequests *int64
	Misc               []generic.Xml
	MiscAttributes     []xml.Attr
}
type UdpQueries struct {
	Retries        *UdpQueriesRetries
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type UdpQueriesRetries struct {
	Attempts       *int64
	Interval       *int64
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
	XMLName        xml.Name                   `xml:"entry"`
	Name           string                     `xml:"name,attr"`
	Cache          *cacheXml                  `xml:"cache,omitempty"`
	Default        *defaultXml                `xml:"default,omitempty"`
	DomainServers  *domainServersContainerXml `xml:"domain-servers,omitempty"`
	Enabled        *string                    `xml:"enabled,omitempty"`
	Interface      *util.MemberType           `xml:"interface,omitempty"`
	StaticEntries  *staticEntriesContainerXml `xml:"static-entries,omitempty"`
	TcpQueries     *tcpQueriesXml             `xml:"tcp-queries,omitempty"`
	UdpQueries     *udpQueriesXml             `xml:"udp-queries,omitempty"`
	Misc           []generic.Xml              `xml:",any"`
	MiscAttributes []xml.Attr                 `xml:",any,attr"`
}
type cacheXml struct {
	CacheEdns      *string         `xml:"cache-edns,omitempty"`
	Enabled        *string         `xml:"enabled,omitempty"`
	MaxTtl         *cacheMaxTtlXml `xml:"max-ttl,omitempty"`
	Misc           []generic.Xml   `xml:",any"`
	MiscAttributes []xml.Attr      `xml:",any,attr"`
}
type cacheMaxTtlXml struct {
	Enabled        *string       `xml:"enabled,omitempty"`
	TimeToLive     *int64        `xml:"time-to-live,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type defaultXml struct {
	Inheritance    *defaultInheritanceXml `xml:"inheritance,omitempty"`
	Primary        *string                `xml:"primary,omitempty"`
	Secondary      *string                `xml:"secondary,omitempty"`
	Misc           []generic.Xml          `xml:",any"`
	MiscAttributes []xml.Attr             `xml:",any,attr"`
}
type defaultInheritanceXml struct {
	Source         *string       `xml:"source,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type domainServersContainerXml struct {
	Entries []domainServersXml `xml:"entry"`
}
type domainServersXml struct {
	XMLName        xml.Name         `xml:"entry"`
	Name           string           `xml:"name,attr"`
	Cacheable      *string          `xml:"cacheable,omitempty"`
	Primary        *string          `xml:"primary,omitempty"`
	Secondary      *string          `xml:"secondary,omitempty"`
	DomainName     *util.MemberType `xml:"domain-name,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
	MiscAttributes []xml.Attr       `xml:",any,attr"`
}
type staticEntriesContainerXml struct {
	Entries []staticEntriesXml `xml:"entry"`
}
type staticEntriesXml struct {
	XMLName        xml.Name         `xml:"entry"`
	Name           string           `xml:"name,attr"`
	Domain         *string          `xml:"domain,omitempty"`
	Address        *util.MemberType `xml:"address,omitempty"`
	Misc           []generic.Xml    `xml:",any"`
	MiscAttributes []xml.Attr       `xml:",any,attr"`
}
type tcpQueriesXml struct {
	Enabled            *string       `xml:"enabled,omitempty"`
	MaxPendingRequests *int64        `xml:"max-pending-requests,omitempty"`
	Misc               []generic.Xml `xml:",any"`
	MiscAttributes     []xml.Attr    `xml:",any,attr"`
}
type udpQueriesXml struct {
	Retries        *udpQueriesRetriesXml `xml:"retries,omitempty"`
	Misc           []generic.Xml         `xml:",any"`
	MiscAttributes []xml.Attr            `xml:",any,attr"`
}
type udpQueriesRetriesXml struct {
	Attempts       *int64        `xml:"attempts,omitempty"`
	Interval       *int64        `xml:"interval,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	if s.Cache != nil {
		var obj cacheXml
		obj.MarshalFromObject(*s.Cache)
		o.Cache = &obj
	}
	if s.Default != nil {
		var obj defaultXml
		obj.MarshalFromObject(*s.Default)
		o.Default = &obj
	}
	if s.DomainServers != nil {
		var objs []domainServersXml
		for _, elt := range s.DomainServers {
			var obj domainServersXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.DomainServers = &domainServersContainerXml{Entries: objs}
	}
	o.Enabled = util.YesNo(s.Enabled, nil)
	if s.Interface != nil {
		o.Interface = util.StrToMem(s.Interface)
	}
	if s.StaticEntries != nil {
		var objs []staticEntriesXml
		for _, elt := range s.StaticEntries {
			var obj staticEntriesXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.StaticEntries = &staticEntriesContainerXml{Entries: objs}
	}
	if s.TcpQueries != nil {
		var obj tcpQueriesXml
		obj.MarshalFromObject(*s.TcpQueries)
		o.TcpQueries = &obj
	}
	if s.UdpQueries != nil {
		var obj udpQueriesXml
		obj.MarshalFromObject(*s.UdpQueries)
		o.UdpQueries = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var cacheVal *Cache
	if o.Cache != nil {
		obj, err := o.Cache.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		cacheVal = obj
	}
	var defaultVal *Default
	if o.Default != nil {
		obj, err := o.Default.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		defaultVal = obj
	}
	var domainServersVal []DomainServers
	if o.DomainServers != nil {
		for _, elt := range o.DomainServers.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			domainServersVal = append(domainServersVal, *obj)
		}
	}
	var interfaceVal []string
	if o.Interface != nil {
		interfaceVal = util.MemToStr(o.Interface)
	}
	var staticEntriesVal []StaticEntries
	if o.StaticEntries != nil {
		for _, elt := range o.StaticEntries.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			staticEntriesVal = append(staticEntriesVal, *obj)
		}
	}
	var tcpQueriesVal *TcpQueries
	if o.TcpQueries != nil {
		obj, err := o.TcpQueries.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		tcpQueriesVal = obj
	}
	var udpQueriesVal *UdpQueries
	if o.UdpQueries != nil {
		obj, err := o.UdpQueries.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		udpQueriesVal = obj
	}

	result := &Entry{
		Name:           o.Name,
		Cache:          cacheVal,
		Default:        defaultVal,
		DomainServers:  domainServersVal,
		Enabled:        util.AsBool(o.Enabled, nil),
		Interface:      interfaceVal,
		StaticEntries:  staticEntriesVal,
		TcpQueries:     tcpQueriesVal,
		UdpQueries:     udpQueriesVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *cacheXml) MarshalFromObject(s Cache) {
	o.CacheEdns = util.YesNo(s.CacheEdns, nil)
	o.Enabled = util.YesNo(s.Enabled, nil)
	if s.MaxTtl != nil {
		var obj cacheMaxTtlXml
		obj.MarshalFromObject(*s.MaxTtl)
		o.MaxTtl = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o cacheXml) UnmarshalToObject() (*Cache, error) {
	var maxTtlVal *CacheMaxTtl
	if o.MaxTtl != nil {
		obj, err := o.MaxTtl.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		maxTtlVal = obj
	}

	result := &Cache{
		CacheEdns:      util.AsBool(o.CacheEdns, nil),
		Enabled:        util.AsBool(o.Enabled, nil),
		MaxTtl:         maxTtlVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *cacheMaxTtlXml) MarshalFromObject(s CacheMaxTtl) {
	o.Enabled = util.YesNo(s.Enabled, nil)
	o.TimeToLive = s.TimeToLive
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o cacheMaxTtlXml) UnmarshalToObject() (*CacheMaxTtl, error) {

	result := &CacheMaxTtl{
		Enabled:        util.AsBool(o.Enabled, nil),
		TimeToLive:     o.TimeToLive,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *defaultXml) MarshalFromObject(s Default) {
	if s.Inheritance != nil {
		var obj defaultInheritanceXml
		obj.MarshalFromObject(*s.Inheritance)
		o.Inheritance = &obj
	}
	o.Primary = s.Primary
	o.Secondary = s.Secondary
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o defaultXml) UnmarshalToObject() (*Default, error) {
	var inheritanceVal *DefaultInheritance
	if o.Inheritance != nil {
		obj, err := o.Inheritance.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		inheritanceVal = obj
	}

	result := &Default{
		Inheritance:    inheritanceVal,
		Primary:        o.Primary,
		Secondary:      o.Secondary,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *defaultInheritanceXml) MarshalFromObject(s DefaultInheritance) {
	o.Source = s.Source
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o defaultInheritanceXml) UnmarshalToObject() (*DefaultInheritance, error) {

	result := &DefaultInheritance{
		Source:         o.Source,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *domainServersXml) MarshalFromObject(s DomainServers) {
	o.Name = s.Name
	o.Cacheable = util.YesNo(s.Cacheable, nil)
	o.Primary = s.Primary
	o.Secondary = s.Secondary
	if s.DomainName != nil {
		o.DomainName = util.StrToMem(s.DomainName)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o domainServersXml) UnmarshalToObject() (*DomainServers, error) {
	var domainNameVal []string
	if o.DomainName != nil {
		domainNameVal = util.MemToStr(o.DomainName)
	}

	result := &DomainServers{
		Name:           o.Name,
		Cacheable:      util.AsBool(o.Cacheable, nil),
		Primary:        o.Primary,
		Secondary:      o.Secondary,
		DomainName:     domainNameVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *staticEntriesXml) MarshalFromObject(s StaticEntries) {
	o.Name = s.Name
	o.Domain = s.Domain
	if s.Address != nil {
		o.Address = util.StrToMem(s.Address)
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o staticEntriesXml) UnmarshalToObject() (*StaticEntries, error) {
	var addressVal []string
	if o.Address != nil {
		addressVal = util.MemToStr(o.Address)
	}

	result := &StaticEntries{
		Name:           o.Name,
		Domain:         o.Domain,
		Address:        addressVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *tcpQueriesXml) MarshalFromObject(s TcpQueries) {
	o.Enabled = util.YesNo(s.Enabled, nil)
	o.MaxPendingRequests = s.MaxPendingRequests
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o tcpQueriesXml) UnmarshalToObject() (*TcpQueries, error) {

	result := &TcpQueries{
		Enabled:            util.AsBool(o.Enabled, nil),
		MaxPendingRequests: o.MaxPendingRequests,
		Misc:               o.Misc,
		MiscAttributes:     o.MiscAttributes,
	}
	return result, nil
}
func (o *udpQueriesXml) MarshalFromObject(s UdpQueries) {
	if s.Retries != nil {
		var obj udpQueriesRetriesXml
		obj.MarshalFromObject(*s.Retries)
		o.Retries = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o udpQueriesXml) UnmarshalToObject() (*UdpQueries, error) {
	var retriesVal *UdpQueriesRetries
	if o.Retries != nil {
		obj, err := o.Retries.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		retriesVal = obj
	}

	result := &UdpQueries{
		Retries:        retriesVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *udpQueriesRetriesXml) MarshalFromObject(s UdpQueriesRetries) {
	o.Attempts = s.Attempts
	o.Interval = s.Interval
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o udpQueriesRetriesXml) UnmarshalToObject() (*UdpQueriesRetries, error) {

	result := &UdpQueriesRetries{
		Attempts:       o.Attempts,
		Interval:       o.Interval,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "cache" || v == "Cache" {
		return e.Cache, nil
	}
	if v == "default" || v == "Default" {
		return e.Default, nil
	}
	if v == "domain_servers" || v == "DomainServers" {
		return e.DomainServers, nil
	}
	if v == "domain_servers|LENGTH" || v == "DomainServers|LENGTH" {
		return int64(len(e.DomainServers)), nil
	}
	if v == "enabled" || v == "Enabled" {
		return e.Enabled, nil
	}
	if v == "interface" || v == "Interface" {
		return e.Interface, nil
	}
	if v == "interface|LENGTH" || v == "Interface|LENGTH" {
		return int64(len(e.Interface)), nil
	}
	if v == "static_entries" || v == "StaticEntries" {
		return e.StaticEntries, nil
	}
	if v == "static_entries|LENGTH" || v == "StaticEntries|LENGTH" {
		return int64(len(e.StaticEntries)), nil
	}
	if v == "tcp_queries" || v == "TcpQueries" {
		return e.TcpQueries, nil
	}
	if v == "udp_queries" || v == "UdpQueries" {
		return e.UdpQueries, nil
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
	if !o.Cache.matches(other.Cache) {
		return false
	}
	if !o.Default.matches(other.Default) {
		return false
	}
	if len(o.DomainServers) != len(other.DomainServers) {
		return false
	}
	for idx := range o.DomainServers {
		if !o.DomainServers[idx].matches(&other.DomainServers[idx]) {
			return false
		}
	}
	if !util.BoolsMatch(o.Enabled, other.Enabled) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Interface, other.Interface) {
		return false
	}
	if len(o.StaticEntries) != len(other.StaticEntries) {
		return false
	}
	for idx := range o.StaticEntries {
		if !o.StaticEntries[idx].matches(&other.StaticEntries[idx]) {
			return false
		}
	}
	if !o.TcpQueries.matches(other.TcpQueries) {
		return false
	}
	if !o.UdpQueries.matches(other.UdpQueries) {
		return false
	}

	return true
}

func (o *Cache) matches(other *Cache) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.CacheEdns, other.CacheEdns) {
		return false
	}
	if !util.BoolsMatch(o.Enabled, other.Enabled) {
		return false
	}
	if !o.MaxTtl.matches(other.MaxTtl) {
		return false
	}

	return true
}

func (o *CacheMaxTtl) matches(other *CacheMaxTtl) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enabled, other.Enabled) {
		return false
	}
	if !util.Ints64Match(o.TimeToLive, other.TimeToLive) {
		return false
	}

	return true
}

func (o *Default) matches(other *Default) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Inheritance.matches(other.Inheritance) {
		return false
	}
	if !util.StringsMatch(o.Primary, other.Primary) {
		return false
	}
	if !util.StringsMatch(o.Secondary, other.Secondary) {
		return false
	}

	return true
}

func (o *DefaultInheritance) matches(other *DefaultInheritance) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Source, other.Source) {
		return false
	}

	return true
}

func (o *DomainServers) matches(other *DomainServers) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.BoolsMatch(o.Cacheable, other.Cacheable) {
		return false
	}
	if !util.StringsMatch(o.Primary, other.Primary) {
		return false
	}
	if !util.StringsMatch(o.Secondary, other.Secondary) {
		return false
	}
	if !util.OrderedListsMatch[string](o.DomainName, other.DomainName) {
		return false
	}

	return true
}

func (o *StaticEntries) matches(other *StaticEntries) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Domain, other.Domain) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Address, other.Address) {
		return false
	}

	return true
}

func (o *TcpQueries) matches(other *TcpQueries) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enabled, other.Enabled) {
		return false
	}
	if !util.Ints64Match(o.MaxPendingRequests, other.MaxPendingRequests) {
		return false
	}

	return true
}

func (o *UdpQueries) matches(other *UdpQueries) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Retries.matches(other.Retries) {
		return false
	}

	return true
}

func (o *UdpQueriesRetries) matches(other *UdpQueriesRetries) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.Attempts, other.Attempts) {
		return false
	}
	if !util.Ints64Match(o.Interval, other.Interval) {
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
