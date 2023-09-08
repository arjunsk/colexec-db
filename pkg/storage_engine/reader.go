package storage_engine

import (
	batch "colexecdb/pkg/query_engine/c_batch"
	"context"
)

type MergeReader struct {
	iteratorIdx int
	level1      []*batch.Batch
}

func NewMergeReader() *MergeReader {
	reader := MergeReader{
		iteratorIdx: 0,
		level1:      make([]*batch.Batch, 3),
	}
	colCnt := 2
	for i, _ := range reader.level1 {
		reader.level1[i] = batch.MockBatch(colCnt, 3)
	}

	return &reader
}

func (m *MergeReader) Read(ctx context.Context, attrs []string) (*batch.Batch, error) {
	if m.iteratorIdx >= len(m.level1) {
		return nil, nil
	}
	defer func() {
		m.iteratorIdx++
	}()

	return m.level1[m.iteratorIdx], nil
}

func (m *MergeReader) Close() error {
	panic("implement me")
}
