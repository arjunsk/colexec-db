package vmath

import (
	types "colexecdb/pkg/query_engine/a_types"
	"math"
)

// Abs Can may be run on GPU using C vector library. But right now it is just a simple loop.
func Abs[T types.FixedSizeT](in []T) (out []T) {
	out = make([]T, len(in))

	for i, item := range in {
		out[i] = T(math.Abs(float64(item)))
	}
	return out
}
