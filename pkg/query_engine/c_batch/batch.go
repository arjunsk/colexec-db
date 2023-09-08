package batch

import vector "colexecdb/pkg/query_engine/b_vector"

type Batch struct {
	Recursive int32 // For recursive CTE, 1 is last batch, 2 is end of batch
	Attrs     []string
	Vecs      []*vector.Vector
	rowCount  int
}

func NewWithSize(n int) *Batch {
	return &Batch{
		Vecs:     make([]*vector.Vector, n),
		Attrs:    make([]string, n),
		rowCount: 0,
	}
}

func (bat *Batch) SetRowCount(rowCount int) {
	bat.rowCount = rowCount
}

func (bat *Batch) GetRowCount() int {
	return bat.rowCount
}

func (bat *Batch) GetVector(pos uint32) *vector.Vector {
	return bat.Vecs[pos]
}

func (bat *Batch) SetVector(pos uint32, vec *vector.Vector) {
	bat.Vecs[pos] = vec
}

func (bat *Batch) Last() bool {
	return bat.Recursive > 0
}

func (bat *Batch) IsEmpty() bool {
	return bat.rowCount == 0
}
