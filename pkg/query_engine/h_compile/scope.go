package compile

import (
	process "colexecdb/pkg/query_engine/e_process"
	planner "colexecdb/pkg/query_engine/g_planner"
	pipeline "colexecdb/pkg/query_engine/i_pipeline"
	rel_algebra "colexecdb/pkg/query_engine/j_rel_algebra"
	"colexecdb/pkg/storage_engine"
)

type ScopeType int

const (
	Normal ScopeType = iota
	CreateTable
)

type Scope struct {
	Magic ScopeType
	Plan  planner.Plan

	DataSource   *Source
	Process      *process.Process
	Instructions rel_algebra.Instructions
	affectedRows uint64
}

type Source struct {
	SchemaName   string
	RelationName string
	Attributes   []string
	Reader       storage_engine.Reader
}

func (s *Scope) CreateTable(c *Compile) error {
	dbName := ""                               //s.Plan.GetDdl().GetCreateDatabase().GetDatabase()
	return c.Engine.Create(c.Ctx, dbName, nil) //c.Proc.TxnOperator)
}

func (s *Scope) Run(c *Compile) error {
	p := pipeline.New(s.DataSource.Attributes, s.Instructions)
	if _, err := p.Run(s.DataSource.Reader, s.Process); err != nil {
		return err
	}
	return nil
}

func (s *Scope) AffectedRows() uint64 {
	return s.affectedRows
}
