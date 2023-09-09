package vmath

import (
	types "colexecdb/pkg/query_engine/a_types"
	"math"
)

// Abs returns the absolute value.
// TODO: Implement vectorized execution using SIMD.
func Abs[T types.FixedSizeT](in []T) (out []T) {
	out = make([]T, len(in))

	for i, item := range in {
		out[i] = T(math.Abs(float64(item)))
	}
	return out
}
