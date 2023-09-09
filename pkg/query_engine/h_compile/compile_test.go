package compile

import (
	batch "colexecdb/pkg/query_engine/c_batch"
	parser "colexecdb/pkg/query_engine/d_parser"
	process "colexecdb/pkg/query_engine/e_process"
	catalog "colexecdb/pkg/query_engine/f_catalog"
	planner "colexecdb/pkg/query_engine/g_planner"
	rel_algebra "colexecdb/pkg/query_engine/j_rel_algebra"
	"colexecdb/pkg/query_engine/j_rel_algebra/projection"
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCompile_Compile(t *testing.T) {

	sql := "select mock_0, sqrt(mock_1) from tbl1;"
	ctx := context.Background()

	stmt, _ := parser.Parse(sql)
	schema := catalog.MockTableDef(2)
	compilerCtx := catalog.NewMockSchemaContext()
	compilerCtx.AppendTableDef("tbl1", schema)
	execPlan, _ := planner.BuildPlan(stmt, compilerCtx)

	process := process.New(context.Background())
	c := New(sql, ctx, process, stmt)

	var batches []*batch.Batch
	fillFn := func(a any, bat *batch.Batch) error {
		if bat != nil {
			rows, _ := bat.Dup()
			batches = append(batches, rows)
		}
		return nil
	}
	_ = c.Compile(ctx, execPlan, fillFn)

	require.Equal(t, 1, len(c.scope))
	require.Equal(t, rel_algebra.Projection, c.scope[0].Instructions[0].Op)
	require.Equal(t, 2, len(c.scope[0].Instructions[0].Arg.(*projection.Argument).Es))

	require.Equal(t, Normal, c.scope[0].Magic)

	runResult, _ := c.Run(0)
	runResult.Batches = batches

	require.Equal(t, 3, len(runResult.Batches))

	for _, bat := range runResult.Batches {
		fmt.Println(bat.String())
	}
}
