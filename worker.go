package core

type rdreWorker struct {
	workerID int
	Rule
}

func (w *rdreWorker) Work(in <-chan Work) <-chan Work {
	out := make(chan Work)
	go func() {
		i := 0
		for j := range in {
			i++
			traceLogger.Printf("id:%d; receiving record:%s; header:%t; rule:%s;\n",
				w.workerID, *j.R, j.IsHeader, w.Rule)
			result, hasResult := w.Apply(&j)
			if !hasResult {
				traceLogger.Printf("id:%d; filtering record:%s; header:%t;\n",
					w.workerID, *j.R, j.IsHeader)
				continue
			}
			j.R = result
			traceLogger.Printf("id:%d; sending record:%s; header:%t;\n",
				w.workerID, *j.R, j.IsHeader)
			out <- j
		}
		close(out)
	}()
	return out
}
