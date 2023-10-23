package operators

import process "colexecdb/pkg/query_engine/e_process"

type OpType int

const (
	Projection OpType = iota
	Output
	Join
)

type Operators []Operator

// Operator contains relational algebra like Projection, Output, Join.
type Operator struct {
	Op  OpType
	Arg OperatorArgument
}

type OperatorArgument interface {
	Free(proc *process.Process, pipelineFailed bool)
}
