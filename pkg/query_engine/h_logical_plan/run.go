package logicalplan

import (
	batch "colexecdb/pkg/query_engine/c_batch"
	"sync"
)
import "github.com/panjf2000/ants/v2"

type RunResult struct {
	AffectedRows uint64
	Batches      []*batch.Batch
}

func (c *LogicalPlan) Run(_ uint64) (*RunResult, error) {
	result := &RunResult{
		AffectedRows: 0,
	}
	if err := c.runOnce(); err != nil {
		return nil, err
	}
	result.AffectedRows = c.getAffectedRows()
	return result, nil
}

func (c *LogicalPlan) runOnce() error {
	var wg sync.WaitGroup
	errC := make(chan error, len(c.scope))
	for i := range c.scope {
		wg.Add(1)
		s := c.scope[i]
		_ = ants.Submit(func() {
			errC <- c.run(s)
			wg.Done()
		})
	}
	wg.Wait()
	c.scope = nil
	close(errC)
	for e := range errC {
		if e != nil {
			return e
		}
	}
	return nil
}

func (c *LogicalPlan) run(s *Scope) error {
	if s == nil {
		return nil
	}
	switch s.Magic {
	case Normal:
		err := s.Run(c)
		if err != nil {
			return err
		}
		c.addAffectedRows(s.AffectedRows())
		return nil
	case CreateTable:
		err := s.CreateTable(c)
		if err != nil {
			return err
		}
		c.setAffectedRows(1)
		return nil
	}
	return nil
}
