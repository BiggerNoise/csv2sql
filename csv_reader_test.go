package csv2sql_test

import (
	"io"
	"os"
	"path/filepath"

	. "github.com/biggernoise/csv2sql"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CSV Reader", func() {
	var PlainCsvFile = filepath.Join(FilesDir, "csv_file.csv")
	var PlainCsvFileReader io.Reader

	BeforeEach(func() {
		var err error
		PlainCsvFileReader, err = os.Open(PlainCsvFile)
		Expect(err).To(BeNil())
	})

	Context("With a CSV File", func() {
		It("Doesn't explode upon creation", func() {
			reader := NewCsvReader(&CsvReaderConfig{Delimeter: ','}, PlainCsvFileReader)
			Expect(reader).NotTo(BeNil())
		})

		It("Reads a Record", func() {
			reader := NewCsvReader(&CsvReaderConfig{Delimeter: ','}, PlainCsvFileReader)
			record, err := reader.Read()

			Expect(record).ToNot(BeNil())
			Expect(err).To(BeNil())

			Expect(record.GetValue("name")).To(Equal("Sir Lancelot"))
		})

		It("Reads multiple records", func() {
			reader := NewCsvReader(&CsvReaderConfig{Delimeter: ','}, PlainCsvFileReader)

			record, err := reader.Read()
			Expect(err).To(BeNil())
			Expect(record.GetValue("name")).To(Equal("Sir Lancelot"))

			record, err = reader.Read()
			Expect(err).To(BeNil())
			Expect(record.GetValue("name")).To(Equal("Sir Galahad"))

			_, err = reader.Read()
			Expect(err).To(Equal(io.EOF))
		})

	})

})
