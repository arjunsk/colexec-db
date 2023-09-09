package function

import (
	types "colexecdb/pkg/query_engine/a_types"
	vector "colexecdb/pkg/query_engine/b_vector"
	process "colexecdb/pkg/query_engine/e_process"
)

var supportedOperators = []FuncNew{
	{
		functionName: "sqrt",
		overloadFn: Overload{
			overloadId: 0,
			args:       []types.T{types.T_int32},
			retType: func(parameters []types.Type) types.Type {
				return types.T_int32.ToType()
			},
			newOp: func() ExecuteLogicOfOverload {
				return func(parameters []*vector.Vector, result *vector.Vector, proc *process.Process, length int) error {
					return sqrt(parameters, result, proc, length)
				}
			},
		},
	},
}
