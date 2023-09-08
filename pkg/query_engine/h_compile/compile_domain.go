package compile

import (
	batch "colexecdb/pkg/query_engine/c_batch"
	"colexecdb/pkg/query_engine/d_parser"
	process "colexecdb/pkg/query_engine/e_process"
	planner "colexecdb/pkg/query_engine/g_planner"
	"colexecdb/pkg/storage_engine"
	"context"
	"sync"
	"sync/atomic"
)

// Compile contains all the information needed for compilation.
type Compile struct {
	scope      []*Scope
	pn         planner.Plan
	affectRows atomic.Uint64
	sql        string

	Engine  storage_engine.Engine
	Ctx     context.Context
	Process *process.Process
	stmt    parser.Statement

	//fill is a result writer runs a callback function.
	//fill will be called when result data is ready.
	fill func(any, *batch.Batch) error

	lock sync.RWMutex
}

func (c *Compile) getAffectedRows() uint64 {
	affectRows := c.affectRows.Load()
	return affectRows
}

type RunResult struct {
	AffectedRows uint64
	Batches      []*batch.Batch
}

type ScopeType int

const (
	Normal ScopeType = iota
	CreateTable
)

type Source struct {
	SchemaName   string
	RelationName string
	Attributes   []string
	Reader       storage_engine.Reader
	Bat          *batch.Batch
}
