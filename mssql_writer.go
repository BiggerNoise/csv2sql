package csv2sql

import (
	"fmt"
	"io"
)

const (
	ColumnTypeString  = 1
	ColumnTypeInteger = 2
	ColumnTypeDecimal = 3
)

type MsSqlWriterConfigColumn struct {
	ColumnName string
	ColumnType int
	FieldName  string
}

type MsSqlWriterConfig struct {
	TableName       string
	ValuesPerInsert int
	Columns         []MsSqlWriterConfigColumn
}

type MsSqlWriter struct {
	config     *MsSqlWriterConfig
	output     io.Writer
	totalCount int
	setCount   int
}

func NewMsSqlWriter(config *MsSqlWriterConfig, output io.Writer) (w *MsSqlWriter) {
	return &MsSqlWriter{config: config, output: output}
}

func (w *MsSqlWriter) Write(r *Record) (err error) {
	if w.setCount == 0 {
		w.open()
	}
	w.writeRecord(r)

	return
}

func (w *MsSqlWriter) Done() {
	if w.setCount > 0 {
		w.close()
	}
	return
}

func (w *MsSqlWriter) open() {
	fmt.Fprintf(w.output, "INSERT INTO [%s] (", w.config.TableName)

	for index, column := range w.config.Columns {
		if index > 0 {
			fmt.Fprintf(w.output, ", ")
		}
		fmt.Fprintf(w.output, "[%s]", column.ColumnName)
	}
	fmt.Fprintln(w.output, ")")
	fmt.Fprintln(w.output, "VALUES")

	return
}

func (w *MsSqlWriter) writeRecord(r *Record) (err error) {
	fmt.Fprintf(w.output, "(")

	for index, column := range w.config.Columns {
		if index > 0 {
			fmt.Fprintf(w.output, ", ")
		}
		val, _ := r.GetValue(column.ColumnName)
		fmt.Fprintf(w.output, "'%s'", val)
	}
	fmt.Fprintf(w.output, ")")
	w.setCount += 1
	w.totalCount += 1
	return

}

func (w *MsSqlWriter) close() {
	fmt.Fprintln(w.output, "")
	fmt.Fprintln(w.output, "GO")
	fmt.Fprintln(w.output, "")
}
