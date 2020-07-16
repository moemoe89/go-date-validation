package validation_test

import (
	d "github.com/moemoe89/go-date-validation"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Validation", func() {

	Describe("Add", func() {

		Context("when string is number", func() {

			It("return an number", func() {
				s := "1"
				number, err := d.StringToInt(s)
				Expect(*number).To(Equal(1))
				Expect(err).NotTo(HaveOccurred())
			})

		})

		Context("when string isn't number", func() {

			It("return an error", func() {
				var xNum *int
				s := "s"
				number, err := d.StringToInt(s)
				Expect(number).To(Equal(xNum))
				Expect(err).To(HaveOccurred())
			})
		})

		Context("when string is date", func() {

			It("return an true", func() {
				date := "2017-02-28"
				ok, err := d.Validation(date)
				Expect(ok).To(Equal(true))
				Expect(err).NotTo(HaveOccurred())
			})

		})

		Context("when string isn't date", func() {

			It("return an error", func() {
				date := "lorem-epsum"
				ok, err := d.Validation(date)
				Expect(ok).To(Equal(false))
				Expect(err).To(HaveOccurred())
			})

		})

		Context("when year string isn't number", func() {

			It("return an error", func() {
				date := "s-12-12"
				ok, err := d.Validation(date)
				Expect(ok).To(Equal(false))
				Expect(err).To(HaveOccurred())
			})

		})

		Context("when month string isn't number", func() {

			It("return an error", func() {
				date := "2017-s-12"
				ok, err := d.Validation(date)
				Expect(ok).To(Equal(false))
				Expect(err).To(HaveOccurred())
			})

		})

		Context("when day string isn't number", func() {

			It("return an error", func() {
				date := "2017-12-s"
				ok, err := d.Validation(date)
				Expect(ok).To(Equal(false))
				Expect(err).To(HaveOccurred())
			})

		})

		Context("when year string is wrong number", func() {

			It("return an error", func() {
				date := "0-12-12"
				ok, err := d.Validation(date)
				Expect(ok).To(Equal(false))
				Expect(err).To(HaveOccurred())
			})

		})

		Context("when month string is wrong number", func() {

			It("return an error", func() {
				date := "2017-13-12"
				ok, err := d.Validation(date)
				Expect(ok).To(Equal(false))
				Expect(err).To(HaveOccurred())
			})

		})

		Context("when day string is wrong number", func() {

			It("return an error", func() {
				date := "2017-12-32"
				ok, err := d.Validation(date)
				Expect(ok).To(Equal(false))
				Expect(err).To(HaveOccurred())
			})

		})

		Context("when date string is out of range", func() {

			It("return an error", func() {
				date := "2017-02-29"
				ok, err := d.Validation(date)
				Expect(ok).To(Equal(false))
				Expect(err).To(HaveOccurred())
			})

		})

	})

})
