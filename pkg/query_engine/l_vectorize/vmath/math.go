package vmath

import (
	types "colexecdb/pkg/query_engine/a_types"
	"math"
)

// Sqrt TODO: Connect to rest of the code.
func Sqrt[T types.FixedSizeT](in []T) (out []T) {
	out = make([]T, len(in))

	for i, item := range in {
		out[i] = T(math.Sqrt(float64(item)))
	}
	return out
}
