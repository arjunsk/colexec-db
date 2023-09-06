package pipeline

import (
	batch "colexecdb/pkg/query_engine/b_batch"
	process "colexecdb/pkg/query_engine/c_process"
	relalgebra "colexecdb/pkg/query_engine/i_rel_algebra"
	"colexecdb/pkg/storage_engine"
)

func New(attrs []string, ins relalgebra.Instructions) *Pipeline {
	return &Pipeline{
		instructions: ins,
		attrs:        attrs,
	}
}

func (p *Pipeline) Run(r storage_engine.Reader, proc *process.Process) (end bool, err error) {

	var bat *batch.Batch
	if err = relalgebra.Prepare(p.instructions, proc); err != nil {
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
		end, err = relalgebra.Run(p.instructions, proc)
		if err != nil {
			return end, err
		}
		if end {
			return end, nil
		}
	}
}
