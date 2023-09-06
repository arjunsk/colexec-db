package rel_algebra

import process "colexecdb/pkg/query_engine/c_process"

type OpType int

const (
	Top OpType = iota
	Order
	Group
	Projection
	Join
)

// Instruction contains relational algebra
type Instruction struct {
	Op  OpType
	Arg InstructionArgument
}

type Instructions []Instruction

type InstructionArgument interface {
	Free(proc *process.Process, pipelineFailed bool)
}
