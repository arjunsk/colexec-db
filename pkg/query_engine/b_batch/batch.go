package batch

import vector "colexecdb/pkg/query_engine/a_vector"

type Batch struct {
	Recursive int32 // For recursive CTE, 1 is last batch, 2 is end of batch
	Attrs     []string
	Vecs      []*vector.Vector
	rowCount  int
}

func (bat *Batch) Last() bool {
	return bat.Recursive > 0
}

func (bat *Batch) IsEmpty() bool {
	return bat.rowCount == 0
}

func (bat *Batch) SetRowCount(rowCount int) {
	bat.rowCount = rowCount
}

func (bat *Batch) GetRowCount() int {
	return bat.rowCount
}

func NewWithSize(n int) *Batch {
	return &Batch{
		Vecs:     make([]*vector.Vector, n),
		rowCount: 0,
	}
}
