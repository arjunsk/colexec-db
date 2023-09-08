package function

import (
	vector "colexecdb/pkg/query_engine/b_vector"
	process "colexecdb/pkg/query_engine/e_process"
	"math"
)

func sqrt(parameters []*vector.Vector, result *vector.Vector, proc *process.Process, length int) error {

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

	return nil
}
