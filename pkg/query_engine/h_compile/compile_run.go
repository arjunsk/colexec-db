package compile

import (
	"sync"
)
import "github.com/panjf2000/ants/v2"

func (c *Compile) Run(_ uint64) (*RunResult, error) {
	result := &RunResult{
		AffectRows: 0,
	}
	if err := c.runOnce(); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Compile) runOnce() error {
	var wg sync.WaitGroup
	errC := make(chan error, len(c.scope))
	for i := range c.scope {
		wg.Add(1)
		s := c.scope[i]
		ants.Submit(func() {
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

func (c *Compile) run(s *Scope) error {
	if s == nil {
		return nil
	}
	switch s.Magic {
	case Normal:
		err := s.Run(c)
		if err != nil {
			return err
		}
		//c.addAffectedRows(s.AffectedRows())
		return nil
	case CreateTable:
		err := s.CreateTable(c)
		if err != nil {
			return err
		}
		//c.setAffectedRows(1)
		return nil
	}
	return nil
}
