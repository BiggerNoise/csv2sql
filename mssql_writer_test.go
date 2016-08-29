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

	Context("Simple Operations", func() {
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

		It("Inserts a single Record with a field/column name mismatch", func() {
			in := `"first_name","last_name","user_name"
"Rob","Pike",rob`

			source := strings.NewReader(in)
			reader := NewCsvReader(&CsvReaderConfig{Delimeter: ','}, source)
			var output bytes.Buffer
			config.Columns[2].FieldName = "user_name"
			record, _ := reader.Read()

			writer := NewMsSqlWriter(config, &output)
			err := writer.Write(record)
			Expect(err).To(BeNil())
			writer.Done()

			Expect(output.String()).To(Equal(`INSERT INTO [People] ([first_name], [last_name], [username])
VALUES
('Rob', 'Pike', 'rob')
GO

`))

		})
		It("Inserts two Records", func() {
			var output bytes.Buffer
			writer := NewMsSqlWriter(config, &output)
			err := writer.Write(records[0])
			Expect(err).To(BeNil())
			err = writer.Write(records[1])
			Expect(err).To(BeNil())
			writer.Done()

			Expect(output.String()).To(Equal(`INSERT INTO [People] ([first_name], [last_name], [username])
VALUES
('Rob', 'Pike', 'rob'),
('Ken', 'Thompson', 'ken')
GO

`))

		})

		It("Automatically Competes insert sections", func() {
			var output bytes.Buffer
			writer := NewMsSqlWriter(config, &output)
			err := writer.Write(records[0])
			Expect(err).To(BeNil())
			err = writer.Write(records[1])
			Expect(err).To(BeNil())
			err = writer.Write(records[2])
			Expect(err).To(BeNil())
			writer.Done()

			Expect(output.String()).To(Equal(`INSERT INTO [People] ([first_name], [last_name], [username])
VALUES
('Rob', 'Pike', 'rob'),
('Ken', 'Thompson', 'ken')
GO

INSERT INTO [People] ([first_name], [last_name], [username])
VALUES
('Robert', 'Griesemer', 'gri')
GO

`))

		})
	})

	Context("Operations with Numeric type", func() {

		BeforeEach(func() {
			records = make([]*Record, 0, 3)
			in := `"first_name","last_name","username",uid
"Rob","Pike",rob,1
Ken,Thompson,ken,2
"Robert","Griesemer","gri",42
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
					{ColumnName: "uid", ColumnType: ColumnTypeInteger},
				},
			}
		})

		It("Makes the Insert Correctly", func() {
			var output bytes.Buffer
			writer := NewMsSqlWriter(config, &output)
			err := writer.Write(records[0])
			Expect(err).To(BeNil())
			writer.Done()

			Expect(output.String()).To(Equal(`INSERT INTO [People] ([first_name], [last_name], [username], [uid])
VALUES
('Rob', 'Pike', 'rob', 1)
GO

`))
		})
	})
})
