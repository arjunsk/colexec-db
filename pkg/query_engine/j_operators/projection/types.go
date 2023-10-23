package projection

import (
	process "colexecdb/pkg/query_engine/e_process"
	logicalplan "colexecdb/pkg/query_engine/g_logical_plan"
	expression "colexecdb/pkg/query_engine/k_expression"
)

type container struct {
	projExecutors []expression.ExpressionExecutor
}

type Argument struct {
	ctr *container
	Es  []logicalplan.Expr
}

func (arg *Argument) Free(proc *process.Process, pipelineFailed bool) {
	if arg.ctr != nil {
		for i := range arg.ctr.projExecutors {
			if arg.ctr.projExecutors[i] != nil {
				arg.ctr.projExecutors[i].Free()
			}
		}
	}
}
