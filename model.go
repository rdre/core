package core

var (
	ruleHandler = make(map[string]Rule)
)

// RecordContext represents the data the rules will be applied on
type RecordContext struct {
	RL   RecordList
	Size int
}

// A Record represents a single data record / header in the system
type Record []string

// RecordList represents a set of records
type RecordList []Record

// SetVal sets the value of a column giving an index
func (r *Record) SetVal(index int, val string) {
	(*r)[index] = val
}

// Insert adds a field with the value at the specified index
func (r *Record) Insert(index int, val string) {
	switch {
	case index == 0:
		infoLogger.Print("1st case ...", index, "->", val)
		*r = append([]string{val}, (*r)...)
	case index >= len((*r)):
		infoLogger.Print("2nd case ...", index, "->", val)
		*r = append((*r), val)
	default:
		infoLogger.Print("default case ...", index, "->", val)
		(*r) = append((*r), "")
		copy((*r)[index+1:], (*r)[index:]) // R->L
		(*r)[index] = val
	}
}

// Delete deletes a field at the specified index
func (r *Record) Delete(index int) {
	if index >= len(*r) || index < 0 {
		// error here
	}
	copy((*r)[index:], (*r)[index+1:])
	*r = (*r)[0 : len(*r)-1]
}

// Work represents a single unit of work to be processed by a rule
type Work struct {
	R        *Record // records
	IsHeader bool    // is this the header record
}

type csvReadWriter struct {
	csvRCWriter
	csvRCReader
}
