package executor

import (
	batch "colexecdb/pkg/query_engine/b_batch"
	process "colexecdb/pkg/query_engine/c_process"
	parser "colexecdb/pkg/query_engine/d_parser"
	catalog "colexecdb/pkg/query_engine/e_catalog"
	planner "colexecdb/pkg/query_engine/f_planner"
	compile "colexecdb/pkg/query_engine/g_compile"
	"context"
)

type txnExecutor struct {
	s   *sqlExecutor
	ctx context.Context
}

func newTxnExecutor(
	ctx context.Context,
	s *sqlExecutor) (*txnExecutor, error) {
	return &txnExecutor{s: s, ctx: ctx}, nil
}

func (exec *txnExecutor) Exec(sql string) (result Result, err error) {

	// parse sql to ast statements
	stmt, err := parser.Parse(sql)
	if err != nil {
		return Result{}, err
	}
	// get table def
	schema := catalog.MockSchemaAll(3)
	ctx := planner.NewMockCompilerContext()
	ctx.AppendSchema("tbl1", schema)

	// create plan
	execPlan, err := planner.BuildPlan(stmt, ctx)
	if err != nil {
		return Result{}, err
	}

	// init compile object
	process := process.New(exec.ctx)
	c := compile.New(sql, exec.ctx, process, stmt)

	// compile()/prepare() compile object
	var batches []*batch.Batch
	fillFn := func(a any, bat *batch.Batch) error {
		if bat != nil {
			batches = append(batches, bat)
		}
		return nil
	}
	err = c.Compile(exec.ctx, execPlan, fillFn)
	if err != nil {
		return Result{}, err
	}

	// run() compile object
	runResult, err := c.Run(0)
	if err != nil {
		return Result{}, err
	}

	// set output
	result.Batches = batches
	result.AffectedRows = runResult.AffectRows

	return
}

func (exec *txnExecutor) commit() error {
	return nil
}

func (exec *txnExecutor) rollback(err error) error {
	return nil
}
