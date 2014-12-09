package core

import (
	"encoding/csv"
	"io"
)

type csvRCWriter struct{}

func (c csvRCWriter) Write(w io.Writer, rc RecordContext) {
	csvResultTxt := csv.NewWriter(w)
	for _, r := range rc.RL {
		csvResultTxt.Write(r)
	}
	csvResultTxt.Flush()
}
