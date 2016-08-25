package csv2sql

import (
	"encoding/csv"
	"io"
)

type CsvReaderConfig struct {
	Delimeter rune
}

type CsvReader struct {
	config   *CsvReaderConfig
	source   io.Reader
	csv      *csv.Reader
	template *RecordTemplate
}

func NewCsvReader(config *CsvReaderConfig, r io.Reader) (parser *CsvReader) {
	return &CsvReader{
		config: config,
		source: r,
	}
}

func (r *CsvReader) Read() (rec *Record, err error) {
	if r.csv == nil {
		if err = r.open(); err != nil {
			return
		}
	}
	var values []string
	if values, err = r.csv.Read(); err != nil {
		return
	}
	rec, err = r.template.NewRecord(values...)
	return
}

func (r *CsvReader) open() (err error) {
	r.csv = csv.NewReader(r.source)
	r.csv.Comment = '#'
	r.csv.LazyQuotes = true
	r.csv.Comma = r.config.Delimeter
	if err != nil {
		return
	}

	var headers []string

	if headers, err = r.csv.Read(); err != nil {
		return
	}
	r.template = NewRecordTemplate(headers...)

	return
}
