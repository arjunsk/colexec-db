package planner

import (
	types "colexecdb/pkg/query_engine/a_types"
	catalog "colexecdb/pkg/query_engine/e_catalog"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMockCompilerContext_Resolve(t *testing.T) {
	table := catalog.MockTableDef(2)

	a := NewMockCompilerContext()
	a.AppendTableDef("tbl1", table)

	require.Equal(t, int32(0), a.ResolveColIdx("", "tbl1", "mock_0"))
	require.Equal(t, types.T_int32.ToType(), a.ResolveColType("", "tbl1", "mock_0"))

	require.Equal(t, int32(1), a.ResolveColIdx("", "tbl1", "mock_1"))
	require.Equal(t, types.T_int64.ToType(), a.ResolveColType("", "tbl1", "mock_1"))

}
