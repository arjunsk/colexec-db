package storage_engine

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMergeReader_Read(t *testing.T) {

	r := NewMergeReader()

	b1, _ := r.Read(context.Background(), []string{"mock_0", "mock_1"})
	require.Equal(t, 2, len(b1.Vecs))
	require.Equal(t, 3, b1.Vecs[0].Length())
}
