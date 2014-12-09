package core

import (
	"testing"
)

func TestRenameCol(t *testing.T) {
	rule := renameCol{"ID", "PERSON ID"}
	r := CopyTestModels()
	//	TestTrace.Println("Header before", h)
	work := &Work{
		&r[0],
		true,
	}
	rule.Apply(work)
	//	TestTrace.Println("Header after", h)
	if r[0][0] != "PERSON ID" {
		t.Errorf("RenameCol{previous: %s, newcol: %s} failed. Expected: %s, Actual: %s}", rule.Previous, rule.Next, rule.Next, r[0][0])
	}
}

func TestFindReplace(t *testing.T) {
	rule := findReplace{"John", "Jane"}
	r := CopyTestModels()
	//	TestTrace.Println("Record before", r[0])
	work := &Work{
		&r[1],
		true,
	}
	rule.Apply(work)
	//	TestTrace.Println("Record after", r[0])
	if r[1][1] != "Jane" {
		t.Errorf("FindReplace{previous: %s, newcol: %s} failed. Expected: %s, Actual: %s}", rule.Find, rule.Replace, rule.Replace, r[1][1])
	}
}
