package colexec

import (
	types "colexecdb/pkg/query_engine/a_types"
	vector "colexecdb/pkg/query_engine/b_vector"
	batch "colexecdb/pkg/query_engine/c_batch"
	process "colexecdb/pkg/query_engine/e_process"
)

type ColumnExpressionExecutor struct {
	typ    types.Type
	colIdx int32
}

var _ ExpressionExecutor = new(ColumnExpressionExecutor)

func (expr *ColumnExpressionExecutor) Eval(_ *process.Process, batches []*batch.Batch) (*vector.Vector, error) {
	vec := batches[0].Vecs[expr.colIdx]
	return vec, nil
}

func (expr *ColumnExpressionExecutor) Free() {
}
