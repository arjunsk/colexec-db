package physicalplan

import (
	batch "colexecdb/pkg/query_engine/c_batch"
	parser "colexecdb/pkg/query_engine/d_parser"
	process "colexecdb/pkg/query_engine/e_process"
	catalog "colexecdb/pkg/query_engine/f_catalog"
	logicalplan "colexecdb/pkg/query_engine/g_logical_plan"
	operators "colexecdb/pkg/query_engine/j_operators"
	"colexecdb/pkg/query_engine/j_operators/projection"
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRun(t *testing.T) {

	sql := "select mock_0, abs(mock_1) from tbl1;"
	ctx := context.Background()

	stmt, _ := parser.Parse(sql)
	schema := catalog.MockTableDef(2)
	sctx := catalog.NewMockSchemaContext()
	sctx.AppendTableDef("tbl1", schema)
	qp, _ := logicalplan.BuildPlan(stmt, sctx)

	process := process.New(context.Background())
	lp := New(sql, ctx, process, stmt)

	var batches []*batch.Batch
	fillFn := func(a any, bat *batch.Batch) error {
		if bat != nil {
			rows, _ := bat.Dup()
			batches = append(batches, rows)
		}
		return nil
	}
	_ = lp.Compile(ctx, qp, fillFn)

	require.Equal(t, 1, len(lp.scope))
	require.Equal(t, operators.Projection, lp.scope[0].Instructions[0].Op)
	require.Equal(t, 2, len(lp.scope[0].Instructions[0].Arg.(*projection.Argument).Es))

	require.Equal(t, Normal, lp.scope[0].Magic)

	runResult, _ := lp.Run()
	runResult.Batches = batches

	require.Equal(t, 3, len(runResult.Batches))

	for _, bat := range runResult.Batches {
		fmt.Println(bat.String())
	}
}
