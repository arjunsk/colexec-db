package thread

import (
	batch "colexecdb/pkg/query_engine/b_batch"
	"context"
)

type Register struct {
	// InputBatch, stores the result of the previous operator.
	InputBatch *batch.Batch
}

type Thread struct {
	Reg Register

	Ctx context.Context

	Cancel context.CancelFunc
}

type ExecStatus int

const (
	ExecStop ExecStatus = iota
	ExecNext
	ExecHasMore
)
