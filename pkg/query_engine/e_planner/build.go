package planner

import (
	types "colexecdb/pkg/query_engine/a_types"
	"errors"
	"github.com/blastrain/vitess-sqlparser/tidbparser/ast"
)

func BuildPlan(stmt ast.StmtNode) (Plan, error) {
	switch s := stmt.(type) {
	case *ast.SelectStmt:
		return buildSelect(s)
	case *ast.InsertStmt:
		return buildInsert(s)
	case *ast.CreateTableStmt:
		return buildCreateTable(s)
	default:
		return nil, errors.New("plan not defined")
	}
}

func buildSelect(sel *ast.SelectStmt) (*QueryPlan, error) {
	q := &QueryPlan{
		StatementType: SELECT,
	}

	for _, selExpr := range sel.Fields.Fields {
		switch selExpr := (selExpr.Expr).(type) {
		case *ast.FuncCallExpr:
			q.Params = []Expr{
				&ExprFunc{
					typ:  types.T_int32.ToType(),
					Name: selExpr.FnName.L,
					Args: []Expr{
						&ExprCol{
							typ:     types.T_int32.ToType(),
							ColName: selExpr.Args[0].(*ast.ColumnNameExpr).Name.Name.L,
						},
					},
				},
			}

		case *ast.ColumnNameExpr:
			q.Params = []Expr{
				&ExprCol{
					typ:     types.T_int32.ToType(),
					ColName: selExpr.Name.Name.L,
				},
			}
		default:
			return nil, errors.New("not supported")
		}
	}

	return q, nil
}

func buildInsert(stmt *ast.InsertStmt) (*QueryPlan, error) {
	return &QueryPlan{
		StatementType: INSERT,
	}, nil
}

func buildCreateTable(stmt *ast.CreateTableStmt) (*DDLPlan, error) {
	return &DDLPlan{}, nil
}
