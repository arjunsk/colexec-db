package function

import (
	types "colexecdb/pkg/query_engine/a_types"
	vector "colexecdb/pkg/query_engine/b_vector"
	process "colexecdb/pkg/query_engine/e_process"
	"colexecdb/pkg/query_engine/l_vectorize/vmath"
	"math"
)

func sqrt(parameters []*vector.Vector, result *vector.Vector, proc *process.Process, length int) error {

	switch parameters[0].GetType().Oid {
	case types.T_int32:
		if parameters[0].GetNsp().GetCardinality() == 0 {
			vecRes := vmath.Sqrt[int32](vector.MustFixedCol[int32](parameters[0]))
			_ = vector.AppendList[int32](result, vecRes)
		} else {
			for i := 0; i < length; i++ {
				v, null := vector.Get[int32](parameters[0], uint32(i))
				if null {
					if err := vector.Append[int32](result, 0, true); err != nil {
						return err
					}
				} else {
					ans := math.Sqrt(float64(v))
					if err := vector.Append[int32](result, int32(ans), true); err != nil {
						return err
					}
				}
			}
		}

	case types.T_int64:
		if parameters[0].GetNsp().GetCardinality() == 0 {
			vecRes := vmath.Sqrt[int64](vector.MustFixedCol[int64](parameters[0]))
			_ = vector.AppendList[int64](result, vecRes)
		} else {
			for i := 0; i < length; i++ {
				v, null := vector.Get[int64](parameters[0], uint32(i))
				if null {
					if err := vector.Append[int64](result, 0, true); err != nil {
						return err
					}
				} else {
					ans := math.Sqrt(float64(v))
					if err := vector.Append[int64](result, int64(ans), true); err != nil {
						return err
					}
				}
			}
		}

	}
	return nil
}
