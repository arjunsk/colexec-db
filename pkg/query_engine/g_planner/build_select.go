package planner

import (
	types "colexecdb/pkg/query_engine/a_types"
	"errors"
	"github.com/blastrain/vitess-sqlparser/tidbparser/ast"
)

func buildSelect(sel *ast.SelectStmt, ctx CompilerContext) (*QueryPlan, error) {
	q := &QueryPlan{
		StatementType: SELECT,
		Params:        make([]Expr, 0),
	}

	tblName := sel.From.TableRefs.Left.(*ast.TableSource).Source.(*ast.TableName).Name.L
	for _, selExpr := range sel.Fields.Fields {
		switch selExpr := (selExpr.Expr).(type) {
		case *ast.FuncCallExpr:

			param := ExprFunc{
				// return type of function.
				// TODO: Later modify it to be dynamic
				Type: types.T_int32.ToType(),
				Name: selExpr.FnName.L,
			}
			for _, arg := range selExpr.Args {
				colName := arg.(*ast.ColumnNameExpr).Name.Name.L
				param.Args = append(param.Args, &ExprCol{
					Type:   ctx.ResolveColType("", tblName, colName),
					ColIdx: ctx.ResolveColIdx("", tblName, colName),
				})
			}

			q.Params = append(q.Params, &param)

		case *ast.ColumnNameExpr:
			colName := selExpr.Name.Name.L
			param := ExprCol{
				Type:   ctx.ResolveColType("", tblName, colName),
				ColIdx: ctx.ResolveColIdx("", tblName, colName),
			}
			q.Params = append(q.Params, &param)
		default:
			return nil, errors.New("not supported")
		}
	}

	return q, nil
}
