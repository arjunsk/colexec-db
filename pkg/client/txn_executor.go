package executor

import (
	batch "colexecdb/pkg/query_engine/c_batch"
	parser "colexecdb/pkg/query_engine/d_parser"
	process "colexecdb/pkg/query_engine/e_process"
	catalog "colexecdb/pkg/query_engine/f_catalog"
	planner "colexecdb/pkg/query_engine/g_planner"
	compile "colexecdb/pkg/query_engine/h_compile"
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
	schema := catalog.MockTableDef(2)
	ctx := planner.NewMockCompilerContext()
	ctx.AppendTableDef("tbl1", schema)

	// create plan
	execPlan, err := planner.BuildPlan(stmt, ctx)
	if err != nil {
		return Result{}, err
	}

	// init compile object
	p := process.New(exec.ctx)
	c := compile.New(sql, exec.ctx, p, stmt)

	var batches []*batch.Batch
	fillFn := func(a any, bat *batch.Batch) error {
		if bat != nil {
			rows, _ := bat.Dup()
			batches = append(batches, rows)
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
	result.AffectedRows = runResult.AffectedRows

	return
}

func (exec *txnExecutor) commit() error {
	return nil
}

func (exec *txnExecutor) rollback(err error) error {
	return nil
}
