package csv2sql

import (
	"fmt"
)

type RecordTemplate struct {
	fields   []string
	fieldMap map[string]int
}

type Record struct {
	template *RecordTemplate
	values   []string
}

func NewRecordTemplate(fields ...string) *RecordTemplate {
	template := &RecordTemplate{fields: fields, fieldMap: make(map[string]int, len(fields))}

	for index, field := range fields {
		template.fieldMap[field] = index
	}
	return template
}

func (t *RecordTemplate) NewRecord(values ...string) (r *Record, err error) {
	if len(values) != len(t.fields) {
		err = fmt.Errorf("Mismatch in count between values and template. %d values expected and %d supplied", len(t.fields), len(values))
		return
	}

	r = &Record{template: t, values: values}
	return
}

func (r *Record) GetValue(key string) (v string, err error) {
	if index, found := r.template.fieldMap[key]; found {
		v = r.values[index]
	} else {
		err = fmt.Errorf("Key %s was not found in the record", key)
	}
	return
}
