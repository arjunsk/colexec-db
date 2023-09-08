package process

import (
	batch "colexecdb/pkg/query_engine/c_batch"
	"context"
)

type Register struct {
	// InputBatch, stores the result of the previous operator.
	InputBatch *batch.Batch
}

// Process is a heavyweight unit of execution that has its own memory space (Register)
type Process struct {
	Reg Register

	Ctx    context.Context
	Cancel context.CancelFunc
}

type ExecStatus int

const (
	ExecStop ExecStatus = iota
	ExecNext
)
