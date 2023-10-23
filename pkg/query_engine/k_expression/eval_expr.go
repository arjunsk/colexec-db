package expression

import (
	process "colexecdb/pkg/query_engine/e_process"
	logicalplan "colexecdb/pkg/query_engine/g_logical_plan"
	"colexecdb/pkg/query_engine/k_expression/function"
	"errors"
)

func NewExpressionExecutorsFromPlanExpressions(proc *process.Process, planExprs []logicalplan.Expr) (executors []ExpressionExecutor, err error) {
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

func NewExpressionExecutor(proc *process.Process, planExpr logicalplan.Expr) (ExpressionExecutor, error) {
	switch t := planExpr.(type) {
	case *logicalplan.ExprCol:
		typ := planExpr.(*logicalplan.ExprCol).Type
		return &ColumnExpressionExecutor{
			colIdx: t.ColIdx,
			typ:    typ,
		}, nil

	case *logicalplan.ExprFunc:
		overload, err := function.GetFunctionById(proc.Ctx, t.Name)
		if err != nil {
			return nil, err
		}

		executor := &FunctionExpressionExecutor{}
		typ := planExpr.(*logicalplan.ExprFunc).Type
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
