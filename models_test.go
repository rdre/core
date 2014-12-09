package core

import (
	"strings"
)

var TestRuleConfig = `
rename_col:
    curr: ID
    next: PERSON ID
    
find_replace:
    find: John
    replace: Jane
`

func CopyTestModels() RecordList {
	// test models
	tRecords := RecordList{
		Record{"ID", "FIRST NAME", "LAST NAME"},
		Record{"1", "John", "Doe"},
		Record{"2", "Jane", "Smith"},
	}
	return tRecords
}

type MockRecordImpExp struct{}

func (mie MockRecordImpExp) ImportRecords() RecordContext {
	r := CopyTestModels()
	rc := RecordContext{r, 3}
	return rc
}

func (mie MockRecordImpExp) ExportRecords(rc RecordContext) {
	for _, record := range rc.RL {
		traceRecord(record)
	}
}

func traceRecord(r Record) {
	line := strings.Join(r, ", ")
	TestTrace.Println(line)
}
