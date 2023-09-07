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
			colName := selExpr.Args[0].(*ast.ColumnNameExpr).Name.Name.L
			q.Params = []Expr{
				&ExprFunc{
					typ:  types.T_int32.ToType(),
					Name: selExpr.FnName.L,
					Args: []Expr{
						&ExprCol{
							typ:    ctx.ResolveColType("", "", colName),
							ColIdx: ctx.ResolveColIdx("", "", colName),
						},
					},
				},
			}

		case *ast.ColumnNameExpr:
			colName := selExpr.Name.Name.L
			q.Params = []Expr{
				&ExprCol{
					typ:    ctx.ResolveColType("", "", colName),
					ColIdx: ctx.ResolveColIdx("", "", colName),
				},
			}
		default:
			return nil, errors.New("not supported")
		}
	}

	return q, nil
}
