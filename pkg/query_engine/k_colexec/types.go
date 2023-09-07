package colexec

import (
	vector "colexecdb/pkg/query_engine/a_vector"
	batch "colexecdb/pkg/query_engine/b_batch"
	process "colexecdb/pkg/query_engine/c_process"
)

type ExpressionExecutor interface {
	Eval(proc *process.Process, batches []*batch.Batch) (*vector.Vector, error)
	Free()
}
