package main

import (
	"colexecdb/pkg/client"
	vector "colexecdb/pkg/query_engine/b_vector"
	"context"
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {

	exec := client.NewSQLExecutor()
	res, _ := exec.Exec(context.Background(), "select mock_0, sqrt(mock_1) from tbl1;")

	var col1 []int32
	var col2 []int64
	res.ReadRows(func(cols []*vector.Vector) bool {
		col1 = append(col1, client.GetFixedRows[int32](cols[0])...)
		col2 = append(col2, client.GetFixedRows[int64](cols[1])...)
		return true
	})

	Print(col1, col2)
}

func Print(col1 []int32, col2 []int64) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	for i := 0; i < len(col1); i++ {
		_, _ = fmt.Fprintf(w, "%d\t%d\n", col1[i], col2[i])
	}
	_ = w.Flush()
}
