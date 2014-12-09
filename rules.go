package core

type renameCol struct {
	Previous, Next string
}

func (rc *renameCol) Apply(w *Work) (result *Record, hasResult bool) {
	if w.IsHeader {
		headers := w.R
		for i, col := range *headers {
			if col == rc.Previous {
				headers.SetVal(i, rc.Next)
			}
		}
	}
	return w.R, true
}

type findReplace struct {
	Find, Replace string
}

func (fr *findReplace) Apply(w *Work) (result *Record, hasResult bool) {
	record := w.R
	for i, field := range *record {
		if field == fr.Find {
			record.SetVal(i, fr.Replace)
		}
	}
	return record, true
}
