package core

import (
	"io/ioutil"
	"log"
	_ "os"
	"testing"
)

var (
	//	TestTrace = log.New(os.Stdout, "[CoreTest] INFO:", log.Ldate|log.Ltime|log.Lshortfile)

	TestTrace = log.New(ioutil.Discard, "[CoreTest] INFO:", log.Ldate|log.Ltime|log.Lshortfile)
)

func _TestChain01(t *testing.T) {
	rule := &renameCol{"ID", "PERSON ID"}
	rules := make([]Rule, 0, 1)
	rules = append(rules, rule)
	ruleWorkers := createWorkers(rules)
	r := CopyTestModels()
	w := Work{&r[0], true}
	out := make(chan Work)
	go func() {
		out <- w
		close(out)
	}()

	//	TestTrace.Println("Header before", h)
	var nxt <-chan Work = out
	for _, w := range ruleWorkers {
		nxt = w.Work(nxt)
	}
	<-nxt
	//	TestTrace.Println("Header after", h)

	if r[0][0] != "PERSON ID" {
		t.Errorf("RenameCol{previous: %s, newcol: %s} failed. Expected: %s, Actual: %s}", rule.Previous, rule.Next, rule.Next, r[0][0])
	}
}

func TestChain02(t *testing.T) {
	rule0 := &renameCol{"ID", "PERSON ID"}
	rule1 := &renameCol{"FIRST NAME", "FORENAME"}

	rules := make([]Rule, 0, 2)
	rules = append(rules, rule0)
	rules = append(rules, rule1)
	ruleWorkers := createWorkers(rules)

	r := CopyTestModels()
	w := Work{&r[0], true}
	out := make(chan Work)
	go func() {
		out <- w
		close(out)
	}()

	//	TestTrace.Println("Header before", h)
	var nxt <-chan Work = out
	for i := range ruleWorkers {
		nxt = ruleWorkers[i].Work(nxt)
	}
	<-nxt
	//	TestTrace.Println("Header after", h)

	if r[0][0] != "PERSON ID" {
		t.Errorf("RenameCol{previous: %s, newcol: %s} failed. Expected: %s, Actual: %s}", rule0.Previous, rule0.Next, rule0.Next, r[0][0])
	}
}

func TestChain03(t *testing.T) {
	// unimplemented
	rule0 := &renameCol{"ID", "PERSON ID"}
	rule1 := &renameCol{"FIRST NAME", "FORENAME"}
	rule2 := &renameCol{"LAST NAME", "SURNAME"}
	rules := []Rule{rule0, rule1, rule2}
	ruleWorkers := createWorkers(rules)

	if l := len(ruleWorkers); l != 3 {
		t.Errorf("Expected %d rule workers, created %d instead", 3, l)
	}

	r := CopyTestModels()

	out := make(chan Work)
	go func() {
		for i, record := range r {
			record := record
			//			out <- Work{&h, &record}
			if i == 0 {
				out <- Work{&record, true}
				continue
			}
			out <- Work{&record, false}
		}
		close(out)
	}()
	var nxt <-chan Work = out
	for i := range ruleWorkers {
		nxt = ruleWorkers[i].Work(nxt)
	}

	for j := range nxt {
		TestTrace.Println(*j.R)
	}

	if r[0][0] != "PERSON ID" {
		t.Errorf("RenameCol{previous: %s, newcol: %s} failed. Expected: %s, Actual: %s}", rule0.Previous, rule0.Next, rule0.Next, r[0][0])
	}
	if r[0][1] != "FORENAME" {
		t.Errorf("RenameCol{previous: %s, newcol: %s} failed. Expected: %s, Actual: %s}", rule1.Previous, rule1.Next, rule1.Next, r[0][1])
	}
	if r[0][2] != "SURNAME" {
		t.Errorf("RenameCol{previous: %s, newcol: %s} failed. Expected: %s, Actual: %s}", rule2.Previous, rule2.Next, rule2.Next, r[0][2])
	}
}

func TestChain10(t *testing.T) {

}

func TestCreateRules(t *testing.T) {
	// unimplemented
}

func TestCreateContext(t *testing.T) {
	// unimplemented
}

// tests the full lifecycle of a work
func TestStart(t *testing.T) {
	// unimplemented
	ruleFile, _ := ioutil.TempFile("", "ruleFile.txt")
	ruleFile.WriteString(TestRuleConfig)
	defer ruleFile.Close()
	opts := func(ac *AppContext) {
		impExp := MockRecordImpExp{}
		ac.RuleFile = ruleFile.Name()
		ac.Imp = impExp
		ac.Exp = impExp
	}
	ac := NewContext(opts)
	ac.Start()
}
