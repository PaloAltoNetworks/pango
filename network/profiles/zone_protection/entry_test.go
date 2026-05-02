package zone_protection_test

import (
	"encoding/xml"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	zp "github.com/PaloAltoNetworks/pango/network/profiles/zone_protection"
	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

func TestZoneProtection(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ZoneProtection Suite")
}

var _ = Describe("Zone Protection Entry", func() {
	var (
		specifier  zp.Specifier
		normalizer zp.Normalizer
	)

	BeforeEach(func() {
		vn, _ := version.New("10.2.0")
		var err error
		specifier, normalizer, err = zp.Versioning(vn)
		Expect(err).ToNot(HaveOccurred())
	})

	Describe("XML tag fixes", func() {
		Context("other-ip flood field", func() {
			It("marshals Other flood as <other-ip>, not <other>", func() {
				enable := true
				entry := &zp.Entry{
					Name: "test",
					Flood: &zp.Flood{
						Other: &zp.FloodProtocol{Enable: &enable},
					},
				}
				spec, err := specifier(entry)
				Expect(err).ToNot(HaveOccurred())

				b, err := xml.Marshal(spec)
				Expect(err).ToNot(HaveOccurred())
				Expect(string(b)).To(ContainSubstring("<other-ip>"))
				Expect(string(b)).ToNot(ContainSubstring("<other>"))
			})

			It("unmarshals <other-ip> back into Other field", func() {
				data := []byte(`<result><entry name="test"><flood><other-ip><enable>yes</enable></other-ip></flood></entry></result>`)
				err := xml.Unmarshal(data, &normalizer)
				Expect(err).ToNot(HaveOccurred())

				entries, err := normalizer.Normalize()
				Expect(err).ToNot(HaveOccurred())
				Expect(entries).To(HaveLen(1))
				Expect(entries[0].Flood).ToNot(BeNil())
				Expect(entries[0].Flood.Other).ToNot(BeNil())
				Expect(entries[0].Flood.Other.Enable).ToNot(BeNil())
				Expect(*entries[0].Flood.Other.Enable).To(BeTrue())
			})
		})

		Context("packet-based wrapper", func() {
			It("marshals TcpSynWithData under <packet-based>", func() {
				val := true
				entry := &zp.Entry{
					Name:           "test",
					TcpSynWithData: &val,
				}
				spec, err := specifier(entry)
				Expect(err).ToNot(HaveOccurred())

				b, err := xml.Marshal(spec)
				Expect(err).ToNot(HaveOccurred())
				Expect(string(b)).To(ContainSubstring("<packet-based>"))
				Expect(string(b)).To(ContainSubstring("<tcp-syn-with-data>yes</tcp-syn-with-data>"))
			})

			It("does not emit <packet-based> when all three fields are nil", func() {
				entry := &zp.Entry{Name: "test"}
				spec, err := specifier(entry)
				Expect(err).ToNot(HaveOccurred())

				b, err := xml.Marshal(spec)
				Expect(err).ToNot(HaveOccurred())
				Expect(string(b)).ToNot(ContainSubstring("<packet-based>"))
			})

			It("unmarshals <packet-based> fields back into Entry", func() {
				data := []byte(`<result><entry name="test"><packet-based><tcp-syn-with-data>yes</tcp-syn-with-data><strip-tcp-fast-open-and-data>no</strip-tcp-fast-open-and-data></packet-based></entry></result>`)
				err := xml.Unmarshal(data, &normalizer)
				Expect(err).ToNot(HaveOccurred())

				entries, err := normalizer.Normalize()
				Expect(err).ToNot(HaveOccurred())
				Expect(entries).To(HaveLen(1))
				e := entries[0]
				Expect(e.TcpSynWithData).ToNot(BeNil())
				Expect(*e.TcpSynWithData).To(BeTrue())
				Expect(e.StripTcpFastOpenAndData).ToNot(BeNil())
				Expect(*e.StripTcpFastOpenAndData).To(BeFalse())
			})
		})
	})

	Describe("New fields", func() {
		Context("discard-strict-source-routing and discard-loose-source-routing", func() {
			It("marshals both fields at top level", func() {
				t := true
				entry := &zp.Entry{
					Name:                       "test",
					DiscardStrictSourceRouting: &t,
					DiscardLooseSourceRouting:  &t,
				}
				spec, err := specifier(entry)
				Expect(err).ToNot(HaveOccurred())

				b, err := xml.Marshal(spec)
				Expect(err).ToNot(HaveOccurred())
				Expect(string(b)).To(ContainSubstring("<discard-strict-source-routing>yes</discard-strict-source-routing>"))
				Expect(string(b)).To(ContainSubstring("<discard-loose-source-routing>yes</discard-loose-source-routing>"))
			})

			It("unmarshals both fields", func() {
				data := []byte(`<result><entry name="test"><discard-strict-source-routing>yes</discard-strict-source-routing><discard-loose-source-routing>no</discard-loose-source-routing></entry></result>`)
				err := xml.Unmarshal(data, &normalizer)
				Expect(err).ToNot(HaveOccurred())

				entries, err := normalizer.Normalize()
				Expect(err).ToNot(HaveOccurred())
				Expect(entries).To(HaveLen(1))
				e := entries[0]
				Expect(e.DiscardStrictSourceRouting).ToNot(BeNil())
				Expect(*e.DiscardStrictSourceRouting).To(BeTrue())
				Expect(e.DiscardLooseSourceRouting).ToNot(BeNil())
				Expect(*e.DiscardLooseSourceRouting).To(BeFalse())
			})
		})
	})

	Describe("Round-trip", func() {
		It("marshals and unmarshals a full entry without data loss", func() {
			alarmRate := int64(10000)
			activateRate := int64(10000)
			maximalRate := int64(40000)
			enable := true
			discard := true

			original := &zp.Entry{
				Name:                       "full-test",
				Description:                util.String("Managed by Terraform"),
				DiscardIpSpoof:             &discard,
				DiscardStrictSourceRouting: &discard,
				DiscardLooseSourceRouting:  &discard,
				Flood: &zp.Flood{
					Syn: &zp.FloodSyn{
						Enable: &enable,
						Red: &zp.FloodRates{
							AlarmRate:    &alarmRate,
							ActivateRate: &activateRate,
							MaximalRate:  &maximalRate,
						},
					},
					Icmp: &zp.FloodProtocol{
						Enable: &enable,
						Red: &zp.FloodRates{
							AlarmRate:    &alarmRate,
							ActivateRate: &activateRate,
							MaximalRate:  &maximalRate,
						},
					},
					Other: &zp.FloodProtocol{
						Enable: &enable,
					},
				},
			}

			spec, err := specifier(original)
			Expect(err).ToNot(HaveOccurred())

			b, err := xml.Marshal(spec)
			Expect(err).ToNot(HaveOccurred())

			wrapped := append([]byte("<result>"), append(b, []byte("</result>")...)...)
			err = xml.Unmarshal(wrapped, &normalizer)
			Expect(err).ToNot(HaveOccurred())

			entries, err := normalizer.Normalize()
			Expect(err).ToNot(HaveOccurred())
			Expect(entries).To(HaveLen(1))

			got := entries[0]
			Expect(got.Name).To(Equal(original.Name))
			Expect(got.Description).ToNot(BeNil())
			Expect(*got.Description).To(Equal(*original.Description))
			Expect(got.DiscardIpSpoof).ToNot(BeNil())
			Expect(*got.DiscardIpSpoof).To(BeTrue())
			Expect(got.DiscardStrictSourceRouting).ToNot(BeNil())
			Expect(*got.DiscardStrictSourceRouting).To(BeTrue())
			Expect(got.Flood).ToNot(BeNil())
			Expect(got.Flood.Other).ToNot(BeNil())
			Expect(*got.Flood.Other.Enable).To(BeTrue())

			Expect(zp.SpecMatches(original, got)).To(BeTrue())
		})
	})
})
