package physicalplan

import (
	batch "colexecdb/pkg/query_engine/c_batch"
	"colexecdb/pkg/query_engine/d_parser"
	process "colexecdb/pkg/query_engine/e_process"
	queryplan "colexecdb/pkg/query_engine/g_query_plan"
	relalgebra "colexecdb/pkg/query_engine/j_rel_algebra"
	"colexecdb/pkg/query_engine/j_rel_algebra/output"
	"colexecdb/pkg/query_engine/j_rel_algebra/projection"
	"colexecdb/pkg/storage_engine"
	"context"
	"errors"
	"sync/atomic"
)

// PhysicalPlan contains all the information needed for compilation.
type PhysicalPlan struct {
	scope      []*Scope
	pn         queryplan.Plan
	affectRows atomic.Uint64
	sql        string

	Engine  storage_engine.Engine
	Ctx     context.Context
	Process *process.Process
	stmt    parser.Statement

	//fill is a result writer runs a callback function.
	fill func(any, *batch.Batch) error
}

// New is used to new an object of compile
func New(sql string, ctx context.Context, proc *process.Process, stmt parser.Statement) *PhysicalPlan {
	c := &PhysicalPlan{}
	c.Ctx = ctx
	c.sql = sql
	c.Process = proc
	c.stmt = stmt
	return c
}

// Compile is the entrance of the compute-execute-layer.
// It generates a scope (logic pipeline) for a query plan.
func (c *PhysicalPlan) Compile(ctx context.Context, pn queryplan.Plan, fill func(any, *batch.Batch) error) (err error) {

	c.Ctx = c.Process.Ctx
	c.pn = pn
	c.fill = fill
	c.scope, err = c.compileScope(ctx, pn)
	return nil
}

func (c *PhysicalPlan) compileScope(ctx context.Context, pn queryplan.Plan) ([]*Scope, error) {
	switch qry := pn.(type) {
	case *queryplan.QueryPlan:
		rs := Scope{
			Magic:        Normal,
			Plan:         pn,
			Instructions: make(relalgebra.Instructions, 0),
		}
		rs.Instructions = append(rs.Instructions, relalgebra.Instruction{
			Op: relalgebra.Projection,
			Arg: &projection.Argument{
				Es: qry.Params,
			},
		})

		// For returning the final result
		rs.Instructions = append(rs.Instructions, relalgebra.Instruction{
			Op: relalgebra.Output,
			Arg: &output.Argument{
				Func: c.fill,
			},
		})

		rs.DataSource = &Source{
			Reader:     storage_engine.NewMergeReader(),
			Attributes: []string{"mock_0", "mock_1"},
		}

		rs.Process = c.Process

		return []*Scope{&rs}, nil

	case *queryplan.DDLPlan:
		switch qry.Type {
		case queryplan.DdlCreateTable:
			rs := Scope{
				Magic:   CreateTable,
				Plan:    pn,
				Process: c.Process,
			}
			return []*Scope{&rs}, nil
		}
	}
	return nil, errors.New("unimplemented")
}

func (c *PhysicalPlan) setAffectedRows(i uint64) {
	c.affectRows.Store(i)
}

func (c *PhysicalPlan) getAffectedRows() uint64 {
	return c.affectRows.Load()
}

func (c *PhysicalPlan) addAffectedRows(i uint64) {
	c.affectRows.Add(i)
}
