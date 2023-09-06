package process

import (
	batch "colexecdb/pkg/query_engine/b_batch"
	"context"
)

// New creates a new Process.
// A process stores the execution context.
func New(ctx context.Context) *Process {
	return &Process{Ctx: ctx}
}

func (proc *Process) SetInputBatch(bat *batch.Batch) {
	proc.Reg.InputBatch = bat
}

func (proc *Process) GetInputBatch() *batch.Batch {
	return proc.Reg.InputBatch
}
