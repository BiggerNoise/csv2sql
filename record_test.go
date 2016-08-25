package csv2sql_test

import (
	. "github.com/biggernoise/csv2sql"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Record", func() {
	template := NewRecordTemplate("one", "two", "three")
	values := []string{"this", "that", "the other"}

	Context("Basic Operations", func() {
		It("Constructs a Record without Exploding", func() {
			record, err := template.NewRecord(values...)

			Expect(record).ToNot(BeNil())
			Expect(err).To(BeNil())
		})

		It("if arguments are mismatched in size return error", func() {
			_, err := template.NewRecord("value")
			Expect(err).NotTo(BeNil())
		})
	})

	Context("Access Values", func() {
		It("Get string values", func() {
			record, err := template.NewRecord(values...)
			Expect(err).To(BeNil())
			Expect(record.GetValue("one")).To(Equal("this"))
			Expect(record.GetValue("two")).To(Equal("that"))
			Expect(record.GetValue("three")).To(Equal("the other"))

		})
	})
})
