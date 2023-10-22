package logicalplan

import (
	types "colexecdb/pkg/query_engine/a_types"
	parser "colexecdb/pkg/query_engine/d_parser"
	catalog "colexecdb/pkg/query_engine/f_catalog"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBuildPlan(t *testing.T) {

	// ast stmt
	stmt, _ := parser.Parse("select mock_0, abs(mock_1) from tbl1;")

	// plan_ctx
	schema := catalog.MockTableDef(2)
	ctx := catalog.NewMockSchemaContext()
	ctx.AppendTableDef("tbl1", schema)

	execPlan, err := BuildPlan(stmt, ctx)
	require.Nil(t, err)

	require.Equal(t, SELECT, execPlan.(*QueryPlan).StatementType)
	require.Equal(t, 2, len(execPlan.(*QueryPlan).Params))

	require.Equal(t, types.T_int32.ToType(), execPlan.(*QueryPlan).Params[0].(*ExprCol).Type)
	require.Equal(t, int32(0), execPlan.(*QueryPlan).Params[0].(*ExprCol).ColIdx)

	require.Equal(t, 1, len(execPlan.(*QueryPlan).Params[1].(*ExprFunc).Args))
	require.Equal(t, types.T_int64.ToType(), execPlan.(*QueryPlan).Params[1].(*ExprFunc).Args[0].(*ExprCol).Type)
	require.Equal(t, int32(1), execPlan.(*QueryPlan).Params[1].(*ExprFunc).Args[0].(*ExprCol).ColIdx)
}
