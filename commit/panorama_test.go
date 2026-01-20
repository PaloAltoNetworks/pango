package commit_test

import (
	"encoding/xml"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/PaloAltoNetworks/pango/commit"
)

var _ = Describe("PanoramaCommit", func() {
	Describe("Normal Commit", func() {
		It("should marshal a normal commit with description", func() {
			expected := `<commit><description>Hello</description></commit>`

			c := commit.PanoramaCommit{
				Description: "Hello",
			}

			b, err := xml.Marshal(c.Element())
			Expect(err).ToNot(HaveOccurred())
			Expect(string(b)).To(Equal(expected))
		})
	})

	Describe("Force Partial Commit", func() {
		It("should marshal a forced partial commit", func() {
			s := []string{
				"<commit>",
				"<description>forced partial commit</description>",
				"<force>",
				"<partial>",
				"<device-group><member>dg1</member></device-group>",
				"<shared-object>excluded</shared-object>",
				"</partial>",
				"</force>",
				"</commit>",
			}
			expected := strings.Join(s, "")

			c := commit.PanoramaCommit{
				Description:          "forced partial commit",
				DeviceGroups:         []string{"dg1"},
				ExcludeSharedObjects: true,
				Force:                true,
			}

			b, err := xml.Marshal(c.Element())
			Expect(err).ToNot(HaveOccurred())
			Expect(string(b)).To(Equal(expected))
		})
	})

	Describe("Full Commit", func() {
		It("should marshal a full partial commit with all options", func() {
			s := []string{
				"<commit>",
				"<description>full</description>",
				"<partial>",
				"<admin><member>admin1</member><member>admin2</member></admin>",
				"<device-group><member>dg1</member></device-group>",
				"<template><member>tmpl1</member></template>",
				"<template-stack><member>ts1</member></template-stack>",
				"<wildfire-appliance><member>wfa1</member></wildfire-appliance>",
				"<wildfire-appliance-cluster><member>wfc1</member></wildfire-appliance-cluster>",
				"<log-collector><member>lc1</member></log-collector>",
				"<log-collector-group><member>lcg1</member></log-collector-group>",
				"<device-and-network>excluded</device-and-network>",
				"<shared-object>excluded</shared-object>",
				"</partial>",
				"</commit>",
			}
			expected := strings.Join(s, "")

			c := commit.PanoramaCommit{
				Description:             "full",
				Admins:                  []string{"admin1", "admin2"},
				DeviceGroups:            []string{"dg1"},
				Templates:               []string{"tmpl1"},
				TemplateStacks:          []string{"ts1"},
				WildfireAppliances:      []string{"wfa1"},
				WildfireClusters:        []string{"wfc1"},
				LogCollectors:           []string{"lc1"},
				LogCollectorGroups:      []string{"lcg1"},
				ExcludeDeviceAndNetwork: true,
				ExcludeSharedObjects:    true,
			}

			b, err := xml.Marshal(c.Element())
			Expect(err).ToNot(HaveOccurred())
			Expect(string(b)).To(Equal(expected))
		})
	})

	Describe("Action", func() {
		It("should return empty string for action", func() {
			c := commit.PanoramaCommit{}
			Expect(c.Action()).To(Equal(""))
		})
	})
})

var _ = Describe("PanoramaCommitAll", func() {
	Describe("Blank Commit All", func() {
		It("should marshal a blank commit-all", func() {
			s := []string{
				"<commit-all>",
				"</commit-all>",
			}
			expected := strings.Join(s, "")

			c := commit.PanoramaCommitAll{}

			b, err := xml.Marshal(c.Element())
			Expect(err).ToNot(HaveOccurred())
			Expect(string(b)).To(Equal(expected))
		})
	})

	Describe("Device Group", func() {
		It("should marshal commit-all for device group with IncludeTemplate=true", func() {
			s := []string{
				"<commit-all>",
				"<shared-policy>",
				"<device-group>",
				"<entry name=\"foo\"><devices><entry name=\"d1\"></entry></devices></entry>",
				"</device-group>",
				"<description>device group</description>",
				"<include-template>yes</include-template>",
				"<force-template-values>no</force-template-values>",
				"</shared-policy>",
				"</commit-all>",
			}
			expected := strings.Join(s, "")

			c := commit.PanoramaCommitAll{
				Type:            commit.TypeDeviceGroup,
				Name:            "foo",
				Description:     "device group",
				Devices:         []string{"d1"},
				IncludeTemplate: true,
			}

			b, err := xml.Marshal(c.Element())
			Expect(err).ToNot(HaveOccurred())
			Expect(string(b)).To(Equal(expected))
		})

		It("should marshal commit-all for device group with IncludeTemplate=false and ForceTemplateValues=false", func() {
			s := []string{
				"<commit-all>",
				"<shared-policy>",
				"<device-group>",
				"<entry name=\"foo\"><devices><entry name=\"d1\"></entry></devices></entry>",
				"</device-group>",
				"<description>device group</description>",
				"<include-template>no</include-template>",
				"<force-template-values>no</force-template-values>",
				"</shared-policy>",
				"</commit-all>",
			}
			expected := strings.Join(s, "")

			c := commit.PanoramaCommitAll{
				Type:        commit.TypeDeviceGroup,
				Name:        "foo",
				Description: "device group",
				Devices:     []string{"d1"},
				// IncludeTemplate and ForceTemplateValues default to false
			}

			b, err := xml.Marshal(c.Element())
			Expect(err).ToNot(HaveOccurred())
			Expect(string(b)).To(Equal(expected))
		})
	})

	Describe("Template", func() {
		It("should marshal commit-all for template with ForceTemplateValues=true", func() {
			s := []string{
				"<commit-all>",
				"<template>",
				"<name>tmpl1</name>",
				"<description>template</description>",
				"<device><member>12321</member></device>",
				"<force-template-values>yes</force-template-values>",
				"</template>",
				"</commit-all>",
			}
			expected := strings.Join(s, "")

			c := commit.PanoramaCommitAll{
				Type:                commit.TypeTemplate,
				Name:                "tmpl1",
				Description:         "template",
				Devices:             []string{"12321"},
				ForceTemplateValues: true,
			}

			b, err := xml.Marshal(c.Element())
			Expect(err).ToNot(HaveOccurred())
			Expect(string(b)).To(Equal(expected))
		})

		It("should marshal commit-all for template with ForceTemplateValues=false", func() {
			s := []string{
				"<commit-all>",
				"<template>",
				"<name>tmpl1</name>",
				"<description>template</description>",
				"<device><member>12321</member></device>",
				"<force-template-values>no</force-template-values>",
				"</template>",
				"</commit-all>",
			}
			expected := strings.Join(s, "")

			c := commit.PanoramaCommitAll{
				Type:        commit.TypeTemplate,
				Name:        "tmpl1",
				Description: "template",
				Devices:     []string{"12321"},
				// ForceTemplateValues defaults to false
			}

			b, err := xml.Marshal(c.Element())
			Expect(err).ToNot(HaveOccurred())
			Expect(string(b)).To(Equal(expected))
		})
	})

	Describe("Template Stack", func() {
		It("should marshal commit-all for template stack with ForceTemplateValues=true", func() {
			s := []string{
				"<commit-all>",
				"<template-stack>",
				"<name>ts1</name>",
				"<description>template stack</description>",
				"<device><member>12321</member></device>",
				"<force-template-values>yes</force-template-values>",
				"</template-stack>",
				"</commit-all>",
			}
			expected := strings.Join(s, "")

			c := commit.PanoramaCommitAll{
				Type:                commit.TypeTemplateStack,
				Name:                "ts1",
				Description:         "template stack",
				Devices:             []string{"12321"},
				ForceTemplateValues: true,
			}

			b, err := xml.Marshal(c.Element())
			Expect(err).ToNot(HaveOccurred())
			Expect(string(b)).To(Equal(expected))
		})

		It("should marshal commit-all for template stack with ForceTemplateValues=false", func() {
			s := []string{
				"<commit-all>",
				"<template-stack>",
				"<name>ts1</name>",
				"<description>template stack</description>",
				"<device><member>12321</member></device>",
				"<force-template-values>no</force-template-values>",
				"</template-stack>",
				"</commit-all>",
			}
			expected := strings.Join(s, "")

			c := commit.PanoramaCommitAll{
				Type:        commit.TypeTemplateStack,
				Name:        "ts1",
				Description: "template stack",
				Devices:     []string{"12321"},
				// ForceTemplateValues defaults to false
			}

			b, err := xml.Marshal(c.Element())
			Expect(err).ToNot(HaveOccurred())
			Expect(string(b)).To(Equal(expected))
		})
	})

	Describe("Log Collector Group", func() {
		It("should marshal commit-all for log collector group", func() {
			s := []string{
				"<commit-all>",
				"<log-collector-config>",
				"<log-collector-group>lcg1</log-collector-group>",
				"<description>log collector</description>",
				"</log-collector-config>",
				"</commit-all>",
			}
			expected := strings.Join(s, "")

			c := commit.PanoramaCommitAll{
				Type:        commit.TypeLogCollectorGroup,
				Name:        "lcg1",
				Description: "log collector",
			}

			b, err := xml.Marshal(c.Element())
			Expect(err).ToNot(HaveOccurred())
			Expect(string(b)).To(Equal(expected))
		})
	})

	Describe("Wildfire Appliance", func() {
		It("should marshal commit-all for wildfire appliance", func() {
			s := []string{
				"<commit-all>",
				"<wildfire-appliance-config>",
				"<description>wf check</description>",
				"<wildfire-appliance>wfa1</wildfire-appliance>",
				"</wildfire-appliance-config>",
				"</commit-all>",
			}
			expected := strings.Join(s, "")

			c := commit.PanoramaCommitAll{
				Type:        commit.TypeWildfireAppliance,
				Name:        "wfa1",
				Description: "wf check",
			}

			b, err := xml.Marshal(c.Element())
			Expect(err).ToNot(HaveOccurred())
			Expect(string(b)).To(Equal(expected))
		})
	})

	Describe("Wildfire Cluster", func() {
		It("should marshal commit-all for wildfire cluster", func() {
			s := []string{
				"<commit-all>",
				"<wildfire-appliance-config>",
				"<description>wf check</description>",
				"<wildfire-appliance-cluster>wfc1</wildfire-appliance-cluster>",
				"</wildfire-appliance-config>",
				"</commit-all>",
			}
			expected := strings.Join(s, "")

			c := commit.PanoramaCommitAll{
				Type:        commit.TypeWildfireCluster,
				Name:        "wfc1",
				Description: "wf check",
			}

			b, err := xml.Marshal(c.Element())
			Expect(err).ToNot(HaveOccurred())
			Expect(string(b)).To(Equal(expected))
		})
	})

	Describe("Action", func() {
		It("should return 'all' for action", func() {
			c := commit.PanoramaCommitAll{}
			expected := "all"

			Expect(c.Action()).To(Equal(expected))
		})
	})
})
