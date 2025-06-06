package tests

import (
	"encoding/xml"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/PaloAltoNetworks/pango/generic"
	"github.com/PaloAltoNetworks/pango/objects/application"
	"github.com/PaloAltoNetworks/pango/version"
)

var _ = Describe("XML Marshalling and Unmarshalling tests", func() {
	version, _ := version.New("11.0.0")
	var specifier application.Specifier
	var normalizer application.Normalizer

	JustBeforeEach(func() {
		specifier, normalizer, _ = application.Versioning(version)
	})
	Context("When unmarshalling XML document with no unknown nodes", func() {
		Context("within top-level object", func() {
			It("should return an Entry object with empty Misc field", func() {
				data := []byte(`<result><entry name="app-1"></entry></result>`)
				err := xml.Unmarshal(data, &normalizer)
				Expect(err).To(Not(HaveOccurred()))

				entries, err := normalizer.Normalize()
				Expect(err).To(Not(HaveOccurred()))
				Expect(entries).To(HaveLen(1))

				Expect(entries[0].Name).To(Equal("app-1"))
				Expect(entries[0].Misc).To(HaveLen(0))
			})
		})
		Context("within child object", func() {
			It("should return an Entry object with empty Misc field within child", func() {
				data := []byte(`<entry name="app-1"><default><ident-by-ip-protocol>test-protocol</ident-by-ip-protocol></default></entry>`)

				marshalled := []byte(fmt.Sprintf("<result>%s</result>", string(data)))
				err := xml.Unmarshal(marshalled, &normalizer)
				Expect(err).To(Not(HaveOccurred()))

				entries, err := normalizer.Normalize()
				Expect(err).To(Not(HaveOccurred()))
				Expect(entries).To(HaveLen(1))

				Expect(entries[0].Name).To(Equal("app-1"))
				Expect(entries[0].Misc).To(HaveLen(0))

				Expect(entries[0].Default.IdentByIpProtocol).To(Not(BeNil()))
				Expect(*entries[0].Default.IdentByIpProtocol).To(Equal("test-protocol"))
				Expect(entries[0].Default.Misc).To(HaveLen(0))

				specified, err := specifier(entries[0])
				Expect(err).To(Not(HaveOccurred()))

				marshalled, err = xml.Marshal(specified)
				Expect(err).To(Not(HaveOccurred()))
				Expect(marshalled).To(Equal(data))
			})
		})
		Context("within a list child object", func() {
			It("should return an Entry object with empty Misc field within a list child", func() {
				data := []byte(`<entry name="app-1"><signature><entry name="signature-1"><order-free>yes</order-free></entry></signature></entry>`)

				marshalled := []byte(fmt.Sprintf("<result>%s</result>", string(data)))
				err := xml.Unmarshal(marshalled, &normalizer)
				Expect(err).To(Not(HaveOccurred()))

				entries, err := normalizer.Normalize()
				Expect(err).To(Not(HaveOccurred()))
				Expect(entries).To(HaveLen(1))

				Expect(entries[0].Name).To(Equal("app-1"))
				Expect(entries[0].Misc).To(HaveLen(0))

				Expect(entries[0].Signature).To(HaveLen(1))
				Expect(entries[0].Signature[0].OrderFree).To(Not(BeNil()))
				Expect(*entries[0].Signature[0].OrderFree).To(BeTrue())
				Expect(entries[0].Signature[0].Misc).To(HaveLen(0))

				specified, err := specifier(entries[0])
				Expect(err).To(Not(HaveOccurred()))

				marshalled, err = xml.Marshal(specified)
				Expect(err).To(Not(HaveOccurred()))
				Expect(marshalled).To(Equal(data))
			})
		})
	})
	Context("When unmarshalling XML document with no unknown nodes", func() {
		var expectedXmlNodes []generic.Xml

		JustBeforeEach(func() {
			fakeTextString := "fake-text"
			expectedXmlNodes = []generic.Xml{
				{
					XMLName:     xml.Name{Local: "fake-node"},
					Text:        []byte("fake-text"),
					TrimmedText: &fakeTextString,
				},
			}
		})
		Context("within top-level object", func() {
			It("should return an Entry object with a non-empty Misc field", func() {
				data := []byte(`<result><entry name="address-1"><fake-node>fake-text</fake-node></entry></result>`)
				err := xml.Unmarshal(data, &normalizer)
				Expect(err).To(Not(HaveOccurred()))

				entries, err := normalizer.Normalize()
				Expect(err).To(Not(HaveOccurred()))
				Expect(entries).To(HaveLen(1))

				Expect(entries[0].Name).To(Equal("address-1"))
				Expect(entries[0].Misc).To(HaveExactElements(expectedXmlNodes))
			})
		})
		Context("within child object", func() {
			It("should return an Entry object with a non-empty Misc field within child", func() {
				data := []byte(`<entry name="app-1"><default><ident-by-ip-protocol>test-protocol</ident-by-ip-protocol><fake-node>fake-text</fake-node></default></entry>`)

				marshalled := []byte(fmt.Sprintf("<result>%s</result>", string(data)))
				err := xml.Unmarshal(marshalled, &normalizer)
				Expect(err).To(Not(HaveOccurred()))

				entries, err := normalizer.Normalize()
				Expect(err).To(Not(HaveOccurred()))
				Expect(entries).To(HaveLen(1))

				Expect(entries[0].Name).To(Equal("app-1"))
				Expect(entries[0].Misc).To(HaveLen(0))

				Expect(entries[0].Default.IdentByIpProtocol).To(Not(BeNil()))
				Expect(*entries[0].Default.IdentByIpProtocol).To(Equal("test-protocol"))
				Expect(entries[0].Default.Misc).To(HaveExactElements(expectedXmlNodes))

				specified, err := specifier(entries[0])
				Expect(err).To(Not(HaveOccurred()))

				marshalled, err = xml.Marshal(specified)
				Expect(err).To(Not(HaveOccurred()))
				Expect(marshalled).To(Equal(data))
			})
		})
		Context("within a list child object", func() {
			It("should return an Entry object with a non-empty Misc field within a list child, one for each element", func() {
				data := []byte(`<entry name="app-1"><signature><entry name="signature-1"><order-free>yes</order-free><fake-node>fake-text1</fake-node></entry><entry name="signature-2"><order-free>no</order-free><fake-node>fake-text2</fake-node></entry></signature></entry>`)

				marshalled := []byte(fmt.Sprintf("<result>%s</result>", string(data)))
				err := xml.Unmarshal(marshalled, &normalizer)
				Expect(err).To(Not(HaveOccurred()))

				entries, err := normalizer.Normalize()
				Expect(err).To(Not(HaveOccurred()))
				Expect(entries).To(HaveLen(1))

				Expect(entries[0].Name).To(Equal("app-1"))
				Expect(entries[0].Misc).To(HaveLen(0))

				Expect(entries[0].Signature).To(HaveLen(2))
				Expect(entries[0].Signature[0].OrderFree).To(Not(BeNil()))
				Expect(*entries[0].Signature[0].OrderFree).To(BeTrue())

				fakeTextString := "fake-text1"
				expectedXmlNodes = []generic.Xml{
					{
						XMLName:     xml.Name{Local: "fake-node"},
						Text:        []byte(fakeTextString),
						TrimmedText: &fakeTextString,
					},
				}
				Expect(entries[0].Signature[0].Misc).To(HaveExactElements())

				fakeTextString = "fake-text2"
				expectedXmlNodes = []generic.Xml{
					{
						XMLName:     xml.Name{Local: "fake-node"},
						Text:        []byte(fakeTextString),
						TrimmedText: &fakeTextString,
					},
				}
				Expect(entries[0].Signature[1].Misc).To(HaveExactElements())

				specified, err := specifier(entries[0])
				Expect(err).To(Not(HaveOccurred()))

				marshalled, err = xml.Marshal(specified)
				Expect(err).To(Not(HaveOccurred()))
				Expect(marshalled).To(Equal(data))
			})
		})
	})
})
