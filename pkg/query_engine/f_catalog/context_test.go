package catalog

import (
	types "colexecdb/pkg/query_engine/a_types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMockSchemaContext_Resolve(t *testing.T) {
	table := MockTableDef(2)

	a := NewMockSchemaContext()
	a.AppendTableDef("tbl1", table)

	require.Equal(t, int32(0), a.ResolveColIdx("", "tbl1", "mock_0"))
	require.Equal(t, types.T_int32.ToType(), a.ResolveColType("", "tbl1", "mock_0"))

	require.Equal(t, int32(1), a.ResolveColIdx("", "tbl1", "mock_1"))
	require.Equal(t, types.T_int64.ToType(), a.ResolveColType("", "tbl1", "mock_1"))

}
