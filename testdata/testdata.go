package testdata

import (
	"encoding/xml"
	"fmt"

	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

type Response struct {
	Raw   []byte
	Error error
}

type MockClient struct {
	// Variables for response.
	Resp          []Response
	Called        int
	Version       version.Number
	Plugin        []map[string]string
	PasswordHash  string
	UnimportError error

	// Variables saved from the mock client's invocation.
	Function      string
	Imports       []string
	Unimports     []string
	Path          string
	Elm           string
	Template      string
	TemplateStack string
	Vsys          string
	Extras        interface{}
}

func (c *MockClient) String() string                       { return "mock" }
func (c *MockClient) Versioning() version.Number           { return c.Version }
func (c *MockClient) Plugins() []map[string]string         { return c.Plugin }
func (c *MockClient) LogAction(f string, a ...interface{}) {}
func (c *MockClient) LogQuery(f string, a ...interface{})  {}
func (c *MockClient) LogOp(f string, a ...interface{})     {}
func (c *MockClient) LogUid(f string, a ...interface{})    {}
func (c *MockClient) Commit(d interface{}, e string, f interface{}) (uint, []byte, error) {
	return 0, nil, nil
}
func (c *MockClient) PositionFirstEntity(d int, e, f string, g, h []string) error { return nil }

func (c *MockClient) Op(req interface{}, vsys string, extras interface{}, ans interface{}) ([]byte, error) {
	c.Function = "op"
	if err := c.SetElm(req); err != nil {
		return nil, err
	}
	c.Vsys = vsys
	c.Extras = extras

	return c.finalize(ans)
}

func (c *MockClient) Show(path interface{}, extras interface{}, ans interface{}) ([]byte, error) {
	c.Function = "show"
	c.Path = util.AsXpath(path)
	c.Extras = extras

	return c.finalize(ans)
}

func (c *MockClient) Get(path interface{}, extras interface{}, ans interface{}) ([]byte, error) {
	c.Function = "get"
	c.Path = util.AsXpath(path)
	c.Extras = extras

	return c.finalize(ans)
}

func (c *MockClient) Delete(path interface{}, extras interface{}, ans interface{}) ([]byte, error) {
	c.Function = "delete"
	c.Path = util.AsXpath(path)
	c.Extras = extras

	return c.finalize(ans)
}

func (c *MockClient) Set(path, elm, extras, ans interface{}) ([]byte, error) {
	c.Function = "set"
	if err := c.SetElm(elm); err != nil {
		return nil, err
	}
	c.Path = util.AsXpath(path)
	c.Extras = extras

	return c.finalize(ans)
}

func (c *MockClient) Edit(path, elm, extras, ans interface{}) ([]byte, error) {
	c.Function = "edit"
	if err := c.SetElm(elm); err != nil {
		return nil, err
	}
	c.Path = util.AsXpath(path)
	c.Extras = extras

	return c.finalize(ans)
}

func (c *MockClient) Move(path interface{}, where, dst string, extras, ans interface{}) ([]byte, error) {
	return nil, nil
}

func (c *MockClient) Uid(cmd interface{}, vsys string, extras, resp interface{}) ([]byte, error) {
	c.Function = "uid"
	if err := c.SetElm(cmd); err != nil {
		return nil, err
	}
	c.Vsys = vsys
	c.Extras = extras

	return c.finalize(resp)
}

func (c *MockClient) EntryListUsing(fn util.Retriever, path []string) ([]string, error) {
	c.Path = util.AsXpath(path)
	return nil, nil
}

func (c *MockClient) MemberListUsing(fn util.Retriever, path []string) ([]string, error) {
	c.Path = util.AsXpath(path)
	return nil, nil
}

func (c *MockClient) RequestPasswordHash(a string) (string, error) { return c.PasswordHash, nil }

func (c *MockClient) finalize(resp interface{}) ([]byte, error) {
	ans := c.Resp[c.Called%len(c.Resp)]
	c.Called++

	if resp == nil {
		return ans.Raw, ans.Error
	}

	if err := xml.Unmarshal(ans.Raw, resp); err != nil {
		return nil, err
	}

	return ans.Raw, ans.Error
}

func (c *MockClient) SetElm(e interface{}) error {
	if e == nil {
		return nil
	}

	rb, err := xml.Marshal(e)
	if err != nil {
		return err
	}

	c.Elm = string(rb)
	return nil
}

func (c *MockClient) VsysImport(ns, tmpl, ts, vsys string, names []string) error {
	c.Template = tmpl
	c.TemplateStack = ts
	c.Vsys = vsys
	c.Imports = names

	return nil
}

func (c *MockClient) VsysUnimport(ns, tmpl, ts string, names []string) error {
	c.Template = tmpl
	c.TemplateStack = ts
	c.Unimports = names

	return c.UnimportError
}

func (c *MockClient) WaitForJob(a uint, resp interface{}) error {
	_, err := c.finalize(resp)
	return err
}

func (c *MockClient) AddResp(val string) {
	c.Resp = append(c.Resp, Response{
		[]byte(fmt.Sprintf("<response><result>%s</result></response>", val)), nil,
	})
}

func (c *MockClient) Reset() {
	c.Function = ""
	c.Imports = []string{}
	c.Unimports = []string{}
	c.Path = ""
	c.Elm = ""
	c.Template = ""
	c.TemplateStack = ""
	c.Vsys = ""
	c.Extras = nil
}

const (
	ApiKeyXml  = "<response><result><key>secret</key></result></response>"
	LicenseXml = `<response><result><licenses>
    <entry>
        <feature>Feature 1</feature>
        <description>License for feature</description>
        <serial>012345</serial>
        <issued>sometime</issued>
        <expires>never</expires>
        <expired>N/A</expired>
    </entry>
    <entry>
        <feature>Feature 2</feature>
        <description>License for another feature</description>
        <serial>012345</serial>
        <issued>sometime</issued>
        <expires>never</expires>
        <expired>N/A</expired>
    </entry>
    </licenses></result></response>`
	UserIdXml = `<response><result>
    <entry ip="10.1.1.1">
        <tag>
            <member>one</member>
            <member>five</member>
            <member>seven</member>
        </tag>
    </entry>
    <entry ip="10.2.2.2">
        <tag>
            <member>one</member>
            <member>two</member>
            <member>three</member>
            <member>four</member>
            <member>five</member>
            <member>six</member>
            <member>seven</member>
        </tag>
    </entry>
    </result></response>`
)
