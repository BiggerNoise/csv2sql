package csv2sql_test

import (
	"encoding/csv"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CSV Parsing", func() {
	var PlainCsvFile = filepath.Join(FilesDir, "csv_file.csv")
	Context("With a CSV File", func() {
		It("Doesn't suck ass", func() {
			Expect(true).To(Equal(true))
		})

		It("Reads Records from a plain CSV file", func() {
			file, err := os.Open(PlainCsvFile)
			Expect(err).To(BeNil())
			csv.NewReader(file)
		})

	})

})
