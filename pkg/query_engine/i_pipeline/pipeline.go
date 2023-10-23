package pipeline

import (
	batch "colexecdb/pkg/query_engine/c_batch"
	process "colexecdb/pkg/query_engine/e_process"
	operators "colexecdb/pkg/query_engine/j_operators"
	"colexecdb/pkg/storage_engine"
)

type Pipeline struct {
	// attrs, column list.
	attrs []string
	// orders to be executed
	instructions operators.Operators
}

func New(attrs []string, ins operators.Operators) *Pipeline {
	return &Pipeline{
		instructions: ins,
		attrs:        attrs,
	}
}

func (p *Pipeline) Run(r storage_engine.Reader, proc *process.Process) (end bool, err error) {

	var bat *batch.Batch
	if err = operators.Prepare(p.instructions, proc); err != nil {
		return false, err
	}

	for {
		select {
		case <-proc.Ctx.Done():
			proc.SetInputBatch(nil)
			return true, nil
		default:
		}
		// read data from storage engine
		if bat, err = r.Read(proc.Ctx, p.attrs); err != nil {
			return false, err
		}

		proc.SetInputBatch(bat)
		end, err = operators.Run(p.instructions, proc)
		if err != nil {
			return end, err
		}
		if end {
			return end, nil
		}
	}
}
