package csv2sql_test

import (
	"bytes"
	"strings"

	. "github.com/biggernoise/csv2sql"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MsSqlWriter", func() {
	var records []*Record
	var config *MsSqlWriterConfig

	BeforeEach(func() {
		records = make([]*Record, 0, 3)
		in := `"first_name","last_name","username"
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
		source := strings.NewReader(in)
		reader := NewCsvReader(&CsvReaderConfig{Delimeter: ','}, source)

		for record, err := reader.Read(); err == nil; record, err = reader.Read() {
			records = append(records, record)
		}
		config = &MsSqlWriterConfig{
			TableName:       "People",
			ValuesPerInsert: 2,
			Columns: []MsSqlWriterConfigColumn{
				{ColumnName: "first_name", ColumnType: ColumnTypeString},
				{ColumnName: "last_name", ColumnType: ColumnTypeString},
				{ColumnName: "username", ColumnType: ColumnTypeString},
			},
		}
	})
	Context("Simple Operations", func() {
		It("Constructs without exploding", func() {

			var output bytes.Buffer
			writer := NewMsSqlWriter(config, &output)

			Expect(writer).ToNot(BeNil())
		})

		It("Inserts a single Record", func() {
			var output bytes.Buffer
			writer := NewMsSqlWriter(config, &output)
			err := writer.Write(records[0])
			Expect(err).To(BeNil())
			writer.Done()

			Expect(output.String()).To(Equal(`INSERT INTO [People] ([first_name], [last_name], [username])
VALUES
('Rob', 'Pike', 'rob')
GO

`))

		})
	})
})
