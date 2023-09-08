package compile

import (
	pipeline "colexecdb/pkg/query_engine/i_pipeline"
)

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
