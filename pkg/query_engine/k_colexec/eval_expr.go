package colexec

import (
	process "colexecdb/pkg/query_engine/c_process"
	planner "colexecdb/pkg/query_engine/f_planner"
	"colexecdb/pkg/query_engine/k_colexec/function"
	"errors"
)

func NewExpressionExecutorsFromPlanExpressions(proc *process.Process, planExprs []planner.Expr) (executors []ExpressionExecutor, err error) {
	executors = make([]ExpressionExecutor, len(planExprs))
	for i := range executors {
		executors[i], err = NewExpressionExecutor(proc, planExprs[i])
		if err != nil {
			for j := 0; j < i; j++ {
				executors[j].Free()
			}
			return nil, err
		}
	}
	return executors, err
}

func NewExpressionExecutor(proc *process.Process, planExpr planner.Expr) (ExpressionExecutor, error) {
	switch t := planExpr.(type) {
	case *planner.ExprCol:
		typ := planExpr.Typ()
		return &ColumnExpressionExecutor{
			colName: t.ColName,
			typ:     typ,
		}, nil

	case *planner.ExprFunc:
		overload, err := function.GetFunctionById(proc.Ctx, t.Name)
		if err != nil {
			return nil, err
		}

		executor := &FunctionExpressionExecutor{}
		typ := planExpr.Typ()
		if err = executor.Init(proc, len(t.Args), typ, overload.GetExecuteMethod()); err != nil {
			return nil, err
		}

		for i := range executor.parameterExecutor {
			subExecutor, paramErr := NewExpressionExecutor(proc, t.Args[i])
			if paramErr != nil {
				for j := 0; j < i; j++ {
					executor.parameterExecutor[j].Free()
				}
				return nil, paramErr
			}
			executor.SetParameter(i, subExecutor)
		}

		return executor, nil
	}

	return nil, errors.New("unsupported executor")
}
