package pango

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("While configuring logging", func() {
	When("converting category strings into LogCategory bitmask", func() {
		Context("with unknown string category", func() {
			It("should fail with error message", func() {
				categories := []string{"invalid"}
				_, err := LogCategoryFromStrings(categories)
				Expect(err).Should(HaveOccurred())
			})
		})

		Context("with \"all\" string present", func() {
			It("should return LogCategoryMask without LogCategorySensitive bit set", func() {
				categories := []string{"all"}
				categoriesMask, err := LogCategoryFromStrings(categories)
				Expect(err).Should(Succeed())
				Expect(categoriesMask & LogCategorySensitive).ShouldNot(Equal(LogCategorySensitive))
			})
		})

		Context("with \"sensitive\" string present", func() {
			It("should return LogCategoryMask with LogCategorySensitive bit set", func() {
				categories := []string{"receive", "sensitive"}
				categoriesMask, err := LogCategoryFromStrings(categories)
				Expect(err).Should(Succeed())
				expectedMask := LogCategoryReceive | LogCategorySensitive
				Expect(categoriesMask).To(Equal(expectedMask))
			})
		})
	})

	When("converting LogCategory bitmask into category strings", func() {
		Context("with invalid bitmask set", func() {
			It("should return error about LogCategoty bitmask being invalid", func() {
				categories := LogCategory(1 << 31)
				_, err := LogCategoryAsStrings(categories)
				Expect(err).Should(HaveOccurred())
			})
		})

		Context("with valid bitmask without LogCategorySensitive", func() {
			It("the list should not \"sensitive\" category", func() {
				categories := LogCategoryReceive | LogCategorySend
				result, err := LogCategoryAsStrings(categories)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(result).Should(ContainElements([]string{"send", "receive"}))
				Expect(result).ShouldNot(ContainElement("sensitive"))
			})
		})

		Context("with bitmask set to LogCategoryAll", func() {
			var categories []string
			BeforeEach(func() {
				for _, v := range logCategoryToString {
					if v != "all" && v != "sensitive" {
						categories = append(categories, v)
					}

				}
			})

			It("the list should not contain \"all\" category", func() {
				result, err := LogCategoryAsStrings(LogCategoryAll)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(result).ShouldNot(ContainElement("all"))
			})

			It("the list should not contain \"sensitive\" category", func() {
				result, err := LogCategoryAsStrings(LogCategoryAll)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(result).ShouldNot(ContainElement("sensitive"))
			})
		})

		Context("with explicitly added LogCategorySensitive", func() {
			It("should have \"sensitive\" element", func() {
				result, err := LogCategoryAsStrings(LogCategoryCurl | LogCategorySensitive)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(result).To(HaveLen(2))
				Expect(result).Should(ContainElements([]string{"curl", "sensitive"}))
			})
		})
	})
})
