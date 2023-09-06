package storage_engine

import (
	batch "colexecdb/pkg/query_engine/b_batch"
	"context"
)

type mergeReader struct {
}

func NewMergeReader() *mergeReader {
	return &mergeReader{}
}

func (m *mergeReader) Read(ctx context.Context, strings []string) (*batch.Batch, error) {
	panic("implement me")
}

func (m *mergeReader) Close() error {
	panic("implement me")
}
