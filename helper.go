package core

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

var (
	infoLogger  = log.New(os.Stdout, "[Context] INFO:", log.Ldate|log.Ltime|log.Lshortfile)
	traceLogger = log.New(ioutil.Discard, "[Context] TRACE:", log.Ldate|log.Ltime|log.Lshortfile)
	errLogger   = log.New(os.Stderr, "[Context] ERROR:", log.Ldate|log.Ltime|log.Lshortfile)
)

// NewContext creates a new App context from optionally defined context options
func NewContext(opts ...ContextOpt) *AppContext {
	ac := &AppContext{RuleHandler: ruleHandler}
	createContext(ac, opts...)

	ruleConfig, _ := ioutil.ReadFile(ac.RuleFile) // TODO: handle errors

	rules := ac.createRules(ruleConfig)
	ac.rw = createWorkers(rules)

	defaultRecordImpExp := csvReadWriter{}
	if ac.Imp == nil {
		// set default csv importer
		ac.Imp = defaultRecordImpExp
	}

	if ac.Exp == nil {
		// set default csv exporter
		ac.Exp = defaultRecordImpExp
	}

	return ac
}

func createContext(appContext *AppContext, opts ...ContextOpt) {
	for _, opt := range opts {
		opt(appContext)
	}
}

func createWorkers(rules []Rule) []rdreWorker {
	numOfRules := len(rules)
	workers := make([]rdreWorker, 0, numOfRules)
	for id, rule := range rules {
		r := rdreWorker{workerID: id, Rule: rule}
		workers = append(workers, r)
	}
	return workers
}

func createPrototype(r Rule, config []byte) Rule {
	if reflect.TypeOf(r).Name() == "RuleFunc" {
		return r
	}
	v := reflect.New(reflect.TypeOf(r).Elem())
	iv := v.Interface()
	rule := iv.(Rule)
	if config != nil {
		yaml.Unmarshal(config, rule)
	}
	return rule
}
