package vector

import (
	types "colexecdb/pkg/query_engine/a_types"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test1(t *testing.T) {
	vec := NewVec(types.T_int32.ToType())
	err := Append[int32](vec, 1, false)
	require.NoError(t, err)

	err = Append[int32](vec, 2, false)
	require.NoError(t, err)

	err = Append[int32](vec, 0, true)
	require.NoError(t, err)

	v, null := Get[int32](vec, 0)
	require.Equal(t, int32(1), v)

	v, null = Get[int32](vec, 1)
	require.Equal(t, int32(2), v)

	v, null = Get[int32](vec, 2)
	require.Equal(t, true, null)
}
