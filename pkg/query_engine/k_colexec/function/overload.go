package function

import (
	types "colexecdb/pkg/query_engine/a_types"
	vector "colexecdb/pkg/query_engine/b_vector"
	process "colexecdb/pkg/query_engine/e_process"
)

type Overload struct {
	overloadId int
	args       []types.T

	retType func(parameters []types.Type) types.Type
	newOp   func() ExecuteLogicOfOverload
}

type ExecuteLogicOfOverload func(
	parameters []*vector.Vector,
	result *vector.Vector,
	proc *process.Process, length int) error

func (ov *Overload) GetExecuteMethod() ExecuteLogicOfOverload {
	f := ov.newOp
	return f()
}
