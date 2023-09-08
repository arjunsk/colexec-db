package compile

import (
	batch "colexecdb/pkg/query_engine/b_batch"
	process "colexecdb/pkg/query_engine/c_process"
	"colexecdb/pkg/query_engine/d_parser"
	planner "colexecdb/pkg/query_engine/f_planner"
	"colexecdb/pkg/storage_engine"
	"context"
	"sync"
	"sync/atomic"
)

// Compile contains all the information needed for compilation.
type Compile struct {
	scope      []*Scope
	pn         planner.Plan
	fill       func(any, *batch.Batch) error
	affectRows atomic.Uint64
	sql        string

	Engine  storage_engine.Engine
	Ctx     context.Context
	Process *process.Process
	stmt    parser.Statement

	lock sync.RWMutex
}

type RunResult struct {
	AffectRows uint64
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
