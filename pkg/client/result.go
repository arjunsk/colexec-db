package executor

import (
	vector "colexecdb/pkg/query_engine/b_vector"
	batch "colexecdb/pkg/query_engine/c_batch"
)

// Result exec sql result
type Result struct {
	AffectedRows uint64
	Batches      []*batch.Batch
}

func (r *Result) ReadRows(apply func(cols []*vector.Vector) bool) {
	for _, rows := range r.Batches {
		if !apply(rows.Vecs) {
			return
		}
	}
}

// GetFixedRows get fixed rows, int, float, etc.
func GetFixedRows[T any](vec *vector.Vector) []T {
	return vector.MustFixedCol[T](vec)
}
