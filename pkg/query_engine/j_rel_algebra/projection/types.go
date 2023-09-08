package projection

import (
	process "colexecdb/pkg/query_engine/e_process"
	planner "colexecdb/pkg/query_engine/g_planner"
	colexec "colexecdb/pkg/query_engine/k_colexec"
)

type container struct {
	projExecutors []colexec.ExpressionExecutor
}

type Argument struct {
	ctr *container
	Es  []planner.Expr
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
