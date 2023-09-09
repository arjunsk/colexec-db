package catalog

import (
	types "colexecdb/pkg/query_engine/a_types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMockSchemaAll(t *testing.T) {
	a := MockTableDef(2)

	require.Equal(t, 2, len(a.ColDefs))

	require.Equal(t, 0, a.ColDefs[0].Idx)
	require.Equal(t, "mock_0", a.ColDefs[0].Name)
	require.Equal(t, types.T_int32.ToType(), a.ColDefs[0].Type)

	require.Equal(t, 1, a.ColDefs[1].Idx)
	require.Equal(t, "mock_1", a.ColDefs[1].Name)
	require.Equal(t, types.T_int64.ToType(), a.ColDefs[1].Type)

}
