package output

import (
	batch "colexecdb/pkg/query_engine/c_batch"
	process "colexecdb/pkg/query_engine/e_process"
)

type Argument struct {
	Data any
	Func func(any, *batch.Batch) error
}

func (arg *Argument) Free(proc *process.Process, pipelineFailed bool) {
	if !pipelineFailed {
		_ = arg.Func(arg.Data, nil)
	}
}
