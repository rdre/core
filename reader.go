package core

import (
	"encoding/csv"
	"io"
)

type csvRCReader struct{}

func (c csvRCReader) Read(r io.Reader) RecordContext {
	csvRecordTxt := csv.NewReader(r)
	csvRecordTxt.LazyQuotes = true
	csvRecordTxt.TrimLeadingSpace = true
	data, _ := csvRecordTxt.ReadAll()
	size := len(data)
	recordSet := RecordList{}
	for _, r := range data {
		recordSet = append(recordSet, r)
	}
	return RecordContext{recordSet, size}
}
