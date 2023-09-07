package function

import (
	types "colexecdb/pkg/query_engine/a_types"
	vector "colexecdb/pkg/query_engine/a_vector"
	process "colexecdb/pkg/query_engine/c_process"
)

var supportedOperators = []FuncNew{
	{
		functionName: "SQRT",
		overloadFn: Overload{
			overloadId: 1,
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
