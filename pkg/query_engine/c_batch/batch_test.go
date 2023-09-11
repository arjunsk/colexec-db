package batch

import (
	types "colexecdb/pkg/query_engine/a_types"
	vector "colexecdb/pkg/query_engine/b_vector"
	"testing"
)

func Test1(t *testing.T) {
	colCount := 2
	rowCount := 3

	bat := NewWithSize(colCount) // 2 columns
	bat.SetRowCount(rowCount)    // 3 rows

	col1 := vector.NewVec(types.T_int32.ToType())
	for i := 0; i < rowCount; i++ {
		_ = col1.Append(int32(i), false)
	}
	bat.Vecs[0] = col1

	col2 := vector.NewVec(types.T_int32.ToType())
	for i := 0; i < rowCount; i++ {
		_ = col2.Append(int32(i), false)
	}
	bat.Vecs[1] = col2

}
