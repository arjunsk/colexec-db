package physicalplan

import (
	batch "colexecdb/pkg/query_engine/c_batch"
	"colexecdb/pkg/query_engine/d_parser"
	process "colexecdb/pkg/query_engine/e_process"
	logicalplan "colexecdb/pkg/query_engine/g_logical_plan"
	operators "colexecdb/pkg/query_engine/j_operators"
	"colexecdb/pkg/query_engine/j_operators/output"
	"colexecdb/pkg/query_engine/j_operators/projection"
	"colexecdb/pkg/storage_engine"
	"context"
	"errors"
	"sync/atomic"
)

// PhysicalPlan contains all the information needed for compilation.
type PhysicalPlan struct {
	scope      []*Scope
	pn         logicalplan.Plan
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
func (c *PhysicalPlan) Compile(ctx context.Context, pn logicalplan.Plan, fill func(any, *batch.Batch) error) (err error) {

	c.Ctx = c.Process.Ctx
	c.pn = pn
	c.fill = fill
	c.scope, err = c.compileScope(ctx, pn)
	return nil
}

func (c *PhysicalPlan) compileScope(ctx context.Context, pn logicalplan.Plan) ([]*Scope, error) {
	switch qry := pn.(type) {
	case *logicalplan.QueryPlan:
		rs := Scope{
			Magic:        Normal,
			Plan:         pn,
			Instructions: make(operators.Operators, 0),
		}
		rs.Instructions = append(rs.Instructions, operators.Operator{
			Op: operators.Projection,
			Arg: &projection.Argument{
				Es: qry.Params,
			},
		})

		// For returning the final result
		rs.Instructions = append(rs.Instructions, operators.Operator{
			Op: operators.Output,
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

	case *logicalplan.DDLPlan:
		switch qry.Type {
		case logicalplan.DdlCreateTable:
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
