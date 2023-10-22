package client

import (
	batch "colexecdb/pkg/query_engine/c_batch"
	parser "colexecdb/pkg/query_engine/d_parser"
	process "colexecdb/pkg/query_engine/e_process"
	catalog "colexecdb/pkg/query_engine/f_catalog"
	logicalplan "colexecdb/pkg/query_engine/g_logical_plan"
	physicalplan "colexecdb/pkg/query_engine/h_phyical_plan"
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

	// get table def from catalog
	tblDef := catalog.MockTableDef(2)
	ctx := catalog.NewMockSchemaContext()
	ctx.AppendTableDef("tbl1", tblDef)

	// create logical plan (plan builder)
	lp, err := logicalplan.BuildPlan(stmt, ctx)
	if err != nil {
		return Result{}, err
	}

	// TODO: implement later
	lp.Optimize(nil)

	// init physical_plan
	p := process.New(exec.ctx)
	pp := physicalplan.New(sql, exec.ctx, p, stmt)

	// compiles query plan to physical plan
	var batches []*batch.Batch
	fillFn := func(a any, bat *batch.Batch) error {
		if bat != nil {
			rows, _ := bat.Dup()
			batches = append(batches, rows)
		}
		return nil
	}
	err = pp.Compile(exec.ctx, lp, fillFn)
	if err != nil {
		return Result{}, err
	}

	// run the physical plan
	runResult, err := pp.Run()
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
