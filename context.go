package core

import (
	"gopkg.in/yaml.v2"
	"io"
)

// AppContext wraps together the necessary resources to apply rules to a record context.
// Responsible for creating rules, creating and chains together the necessary rule workers,
// and kicking them off.
type AppContext struct {
	RuleFile    string
	RuleHandler map[string]Rule
	rw          []rdreWorker

	rc RecordContext

	Imp RecordReader
	Exp RecordWriter
	RR  io.Reader
	RW  io.Writer
}

// Start kicks of the rule workers to concurrently process rules.
func (ac *AppContext) Start() {
	// set record context
	ac.rc = ac.Imp.Read(ac.RR)

	// pass first job rounds
	out := make(chan Work)

	go func() {
		for i := range ac.rc.RL {
			if i == 0 {
				out <- Work{&ac.rc.RL[i], true}
				continue
			}
			out <- Work{&ac.rc.RL[i], false}
		}
		close(out)
	}()

	// start workers
	var nxt <-chan Work = out
	for i := range ac.rw {
		nxt = ac.rw[i].Work(nxt)
	}

	finalResults := make(RecordList, 0, ac.rc.Size)
	for j := range nxt {
		finalResults = append(finalResults, *j.R)
	}

	ac.Exp.Write(ac.RW, RecordContext{
		finalResults,
		len(finalResults),
	})
}

func (ac *AppContext) createRules(ruleConfig []byte) []Rule {
	rules := make([]Rule, 0, 10)

	ruleDefs := yaml.MapSlice{}
	yaml.Unmarshal(ruleConfig, &ruleDefs)
	traceLogger.Println(ruleDefs)
	for _, ruleDef := range ruleDefs {
		key := ruleDef.Key.(string)
		var rule Rule
		rulePrototype := ac.RuleHandler[key]
		var def yaml.MapSlice
		if ruleDef.Value != nil {
			def = ruleDef.Value.(yaml.MapSlice)
		}
		rawDef, _ := yaml.Marshal(def) // TODO: handle error
		rule = createPrototype(rulePrototype, rawDef)
		traceLogger.Println("\n", string(rawDef), rulePrototype, rule)
		rules = append(rules, rule)
	}
	return rules
}
