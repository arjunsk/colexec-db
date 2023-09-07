package planner

import (
	types "colexecdb/pkg/query_engine/a_types"
	"errors"
	"github.com/blastrain/vitess-sqlparser/tidbparser/ast"
)

func buildSelect(sel *ast.SelectStmt, ctx CompilerContext) (*QueryPlan, error) {
	q := &QueryPlan{
		StatementType: SELECT,
	}

	for _, selExpr := range sel.Fields.Fields {
		switch selExpr := (selExpr.Expr).(type) {
		case *ast.FuncCallExpr:

			param := ExprFunc{
				Type: types.T_int32.ToType(),
				Name: selExpr.FnName.L,
			}
			for _, arg := range selExpr.Args {
				colName := arg.(*ast.ColumnNameExpr).Name.Name.L
				param.Args = append(param.Args, &ExprCol{
					Type:   ctx.ResolveColType("", "", colName),
					ColIdx: ctx.ResolveColIdx("", "", colName),
				})
			}

			q.Params = []Expr{&param}

		case *ast.ColumnNameExpr:
			colName := selExpr.Name.Name.L
			q.Params = []Expr{
				&ExprCol{
					Type:   ctx.ResolveColType("", "", colName),
					ColIdx: ctx.ResolveColIdx("", "", colName),
				},
			}
		default:
			return nil, errors.New("not supported")
		}
	}

	return q, nil
}
