package core

import (
	"io"
)

// ContextOpt is a functional option providing the user the freedom to set resources for
// an AppContext.
// Options include setting the path to the rule file and record file
// Optionally
// - record context importer: allows record context to be imported in a custom way,
// from an unlimitted source of data. E.g. a custom record conext importer can import from a DB, yaml, json, etc
// - record context exporter: e.g. a custom record context exporter can export results to json, yaml, db, etc
// - rule handler: add factories for custom rules, override existing factories for default provided rules.
type ContextOpt func(*AppContext)

type RecordReadWriter interface {
	RecordReader
	RecordWriter
}

// RecordReader defines an interface for reading records from an io.Reader
type RecordReader interface {
	Read(r io.Reader) RecordContext
}

// RecordWriter defines an interface for writing records to an io.Writer
type RecordWriter interface {
	Write(w io.Writer, rc RecordContext)
}

// RecordImporter defines a basic interface for a custom record context importer.
type RecordImporter interface {
	ImportRecords() RecordContext
}

// RecordExporter defines a basic interface for exporting the result of processing the record context.
type RecordExporter interface {
	ExportRecords(rc RecordContext)
}

// Rule is an interface that defines the basic operation for a rule.
type Rule interface {
	// result may refer to the same record or a new record
	Apply(w *Work) (result *Record, hasResult bool)
}

// RuleFunc provides an interface for functional (stateless) rules
type RuleFunc func(w *Work) (result *Record, hasResult bool)

// Apply executes the functional rule
func (r RuleFunc) Apply(w *Work) (result *Record, hasResult bool) {
	return r(w)
}
