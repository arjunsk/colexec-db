package function

import (
	types "colexecdb/pkg/query_engine/a_types"
	vector "colexecdb/pkg/query_engine/b_vector"
	process "colexecdb/pkg/query_engine/e_process"
	"colexecdb/pkg/query_engine/l_vectorize/vmath"
	"math"
)

func abs(parameters []*vector.Vector, result *vector.Vector, proc *process.Process, length int) error {

	switch parameters[0].GetType().Oid {
	case types.T_int32:
		err := absGeneric[int32](parameters, result, length)
		if err != nil {
			return err
		}

	case types.T_int64:
		err := absGeneric[int64](parameters, result, length)
		if err != nil {
			return err
		}

	}
	return nil
}

func absGeneric[T types.FixedSizeT](parameters []*vector.Vector, result *vector.Vector, length int) error {
	if parameters[0].GetNsp().GetCardinality() == 0 {
		vecRes := vmath.Abs[T](vector.MustFixedCol[T](parameters[0]))
		_ = vector.AppendList[T](result, vecRes)
	} else {
		for i := 0; i < length; i++ {
			v, null := vector.Get[T](parameters[0], uint32(i))
			if null {
				if err := result.Append(0, true); err != nil {
					return err
				}
			} else {
				ans := math.Abs(float64(v))
				if err := result.Append(T(ans), false); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
