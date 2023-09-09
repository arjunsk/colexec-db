package rel_algebra

import process "colexecdb/pkg/query_engine/e_process"

type OpType int

const (
	Projection OpType = iota
	Output
	Join
)

type Instructions []Instruction

// Instruction contains relational algebra
type Instruction struct {
	Op  OpType
	Arg InstructionArgument
}

type InstructionArgument interface {
	Free(proc *process.Process, pipelineFailed bool)
}
