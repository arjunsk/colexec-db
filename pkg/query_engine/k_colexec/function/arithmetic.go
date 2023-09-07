package function

import (
	vector "colexecdb/pkg/query_engine/a_vector"
	process "colexecdb/pkg/query_engine/c_process"
	"math"
)

func sqrt(parameters []*vector.Vector, result *vector.Vector, proc *process.Process, length int) error {

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

	return nil
}
