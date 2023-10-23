package operators

import (
	"bytes"
	process "colexecdb/pkg/query_engine/e_process"
	"colexecdb/pkg/query_engine/j_operators/output"
	"colexecdb/pkg/query_engine/j_operators/projection"
)

var stringFunc = [...]func(any, *bytes.Buffer){
	Projection: projection.String,
	Output:     output.String,
}

var prepareFunc = [...]func(*process.Process, any) error{
	Projection: projection.Prepare,
	Output:     output.Prepare,
}

var execFunc = [...]func(*process.Process, any) (process.ExecStatus, error){
	Projection: projection.Call,
	Output:     output.Call,
}
