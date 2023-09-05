package thread

import (
	batch "colexecdb/pkg/query_engine/b_batch"
	"context"
)

// New creates a new Process.
// A process stores the execution context.
func New(ctx context.Context) *Thread {
	return &Thread{Ctx: ctx}
}

func (proc *Thread) SetInputBatch(bat *batch.Batch) {
	proc.Reg.InputBatch = bat
}

func (proc *Thread) GetInputBatch() *batch.Batch {
	return proc.Reg.InputBatch
}
