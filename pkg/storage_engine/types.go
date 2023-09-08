package storage_engine

import (
	batch "colexecdb/pkg/query_engine/c_batch"
	"context"
)

type Engine interface {
	Create(ctx context.Context, name string, operator interface{}) error
}

type Reader interface {
	Close() error
	Read(context.Context, []string) (*batch.Batch, error)
}
