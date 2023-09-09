package main

import (
	executor "colexecdb/pkg/client"
	vector "colexecdb/pkg/query_engine/b_vector"
	"context"
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {

	exec := executor.NewSQLExecutor()
	res, _ := exec.Exec(context.Background(), "select mock_0, sqrt(mock_1) from tbl1;")

	var mock0 []int32
	var mock1 []int64
	res.ReadRows(func(cols []*vector.Vector) bool {
		mock0 = append(mock0, executor.GetFixedRows[int32](cols[0])...)
		mock1 = append(mock1, executor.GetFixedRows[int64](cols[1])...)
		return true
	})

	Print(mock0, mock1)
}

func Print(mock0 []int32, mock1 []int64) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	for i := 0; i < len(mock0); i++ {
		_, _ = fmt.Fprintf(w, "%d\t%d\n", mock0[i], mock1[i])
	}
	_ = w.Flush()
}
