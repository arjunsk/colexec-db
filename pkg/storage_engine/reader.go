package storage_engine

import (
	batch "colexecdb/pkg/query_engine/b_batch"
	"context"
)

type MergeReader struct {
}

func NewMergeReader() *MergeReader {
	return &MergeReader{}
}

func (m *MergeReader) Read(ctx context.Context, strings []string) (*batch.Batch, error) {
	panic("implement me")
}

func (m *MergeReader) Close() error {
	panic("implement me")
}
