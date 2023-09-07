package colexec

import (
	types "colexecdb/pkg/query_engine/a_types"
	vector "colexecdb/pkg/query_engine/a_vector"
	batch "colexecdb/pkg/query_engine/b_batch"
	process "colexecdb/pkg/query_engine/c_process"
)

type ColumnExpressionExecutor struct {
	colIdx int32
	typ    types.Type
}

var _ ExpressionExecutor = new(ColumnExpressionExecutor)

func (expr *ColumnExpressionExecutor) Eval(_ *process.Process, batches []*batch.Batch) (*vector.Vector, error) {
	vec := batches[0].Vecs[expr.colIdx]
	return vec, nil
}

func (expr *ColumnExpressionExecutor) Free() {
}
