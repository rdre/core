package core

import (
	"testing"
)

func _TestWork(t *testing.T) {
	r := CopyTestModels()
	w := Work{
		&r[0],
		true,
	}

	rule := &renameCol{"ID", "PERSON ID"}
	rw := rdreWorker{1, rule}

	out := make(chan Work, 1)
	go func() {
		out <- w
		close(out)
	}()

	var nxt <-chan Work = out
	nxt = rw.Work(nxt)
	<-nxt

	if r[0][0] != "PERSON ID" {
		t.Errorf("RenameCol{previous: %s, newcol: %s} failed. Expected: %s, Actual: %s}", rule.Previous, rule.Next, rule.Next, r[0][0])
	}
}
