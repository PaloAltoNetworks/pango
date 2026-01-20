package commit_test

import (
	"encoding/xml"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/PaloAltoNetworks/pango/commit"
)

var _ = Describe("FirewallCommit", func() {
	Describe("Normal Commit", func() {
		It("should marshal a normal commit with description", func() {
			expected := `<commit><description>Hello</description></commit>`

			c := commit.FirewallCommit{
				Description: "Hello",
			}

			b, err := xml.Marshal(c.Element())
			Expect(err).ToNot(HaveOccurred())
			Expect(string(b)).To(Equal(expected))
		})
	})

	Describe("Partial Commits", func() {
		It("should marshal partial commit with admins", func() {
			s := []string{
				"<commit>",
				"<description>example</description>",
				"<partial><admin>",
				"<member>a1</member>",
				"<member>a2</member>",
				"</admin></partial>",
				"</commit>",
			}
			expected := strings.Join(s, "")

			c := commit.FirewallCommit{
				Description: "example",
				Admins:      []string{"a1", "a2"},
			}

			b, err := xml.Marshal(c.Element())
			Expect(err).ToNot(HaveOccurred())
			Expect(string(b)).To(Equal(expected))
		})

		It("should marshal partial commit excluding device and network", func() {
			s := []string{
				"<commit>",
				"<description>example</description>",
				"<partial>",
				"<device-and-network>excluded</device-and-network>",
				"</partial>",
				"</commit>",
			}
			expected := strings.Join(s, "")

			c := commit.FirewallCommit{
				Description:             "example",
				ExcludeDeviceAndNetwork: true,
			}

			b, err := xml.Marshal(c.Element())
			Expect(err).ToNot(HaveOccurred())
			Expect(string(b)).To(Equal(expected))
		})

		It("should marshal partial commit excluding policy and objects", func() {
			s := []string{
				"<commit>",
				"<description>example</description>",
				"<partial>",
				"<policy-and-objects>excluded</policy-and-objects>",
				"</partial>",
				"</commit>",
			}
			expected := strings.Join(s, "")

			c := commit.FirewallCommit{
				Description:             "example",
				ExcludePolicyAndObjects: true,
			}

			b, err := xml.Marshal(c.Element())
			Expect(err).ToNot(HaveOccurred())
			Expect(string(b)).To(Equal(expected))
		})

		It("should marshal partial commit excluding shared objects", func() {
			s := []string{
				"<commit>",
				"<description>example</description>",
				"<partial>",
				"<shared-object>excluded</shared-object>",
				"</partial>",
				"</commit>",
			}
			expected := strings.Join(s, "")

			c := commit.FirewallCommit{
				Description:          "example",
				ExcludeSharedObjects: true,
			}

			b, err := xml.Marshal(c.Element())
			Expect(err).ToNot(HaveOccurred())
			Expect(string(b)).To(Equal(expected))
		})

		It("should marshal full partial commit with all options", func() {
			s := []string{
				"<commit>",
				"<description>check</description>",
				"<partial>",
				"<admin>",
				"<member>user1</member>",
				"<member>user2</member>",
				"</admin>",
				"<device-and-network>excluded</device-and-network>",
				"<shared-object>excluded</shared-object>",
				"<policy-and-objects>excluded</policy-and-objects>",
				"</partial>",
				"</commit>",
			}
			expected := strings.Join(s, "")

			c := commit.FirewallCommit{
				Description:             "check",
				Admins:                  []string{"user1", "user2"},
				ExcludeDeviceAndNetwork: true,
				ExcludeSharedObjects:    true,
				ExcludePolicyAndObjects: true,
			}

			b, err := xml.Marshal(c.Element())
			Expect(err).ToNot(HaveOccurred())
			Expect(string(b)).To(Equal(expected))
		})
	})

	Describe("Action", func() {
		It("should return empty string for action", func() {
			c := commit.FirewallCommit{}
			Expect(c.Action()).To(Equal(""))
		})
	})
})
