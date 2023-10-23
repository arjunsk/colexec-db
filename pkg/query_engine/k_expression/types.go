package expression

import (
	vector "colexecdb/pkg/query_engine/b_vector"
	batch "colexecdb/pkg/query_engine/c_batch"
	process "colexecdb/pkg/query_engine/e_process"
)

type ExpressionExecutor interface {
	Eval(proc *process.Process, batches []*batch.Batch) (*vector.Vector, error)
	Free()
}
