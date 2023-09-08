package main

import (
	executor "colexecdb/pkg/client"
	"context"
	"fmt"
)

func main() {

	exec := executor.NewSQLExecutor()
	res, _ := exec.Exec(context.Background(), "select mock_0, sqrt(mock_1) from tbl1;")

	for _, bat := range res.Batches {
		fmt.Println(bat.String())
	}
}
