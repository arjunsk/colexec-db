package planner

import (
	"errors"
	"github.com/blastrain/vitess-sqlparser/tidbparser/ast"
)

func BuildPlan(stmt ast.StmtNode, ctx CompilerContext) (Plan, error) {
	switch s := stmt.(type) {
	case *ast.SelectStmt:
		return buildSelect(s, ctx)
	case *ast.InsertStmt:
		return buildInsert(s, ctx)
	case *ast.CreateTableStmt:
		return buildCreateTable(s, ctx)
	default:
		return nil, errors.New("plan not defined")
	}
}

func buildInsert(stmt *ast.InsertStmt, ctx CompilerContext) (*QueryPlan, error) {
	return &QueryPlan{
		StatementType: INSERT,
	}, nil
}

func buildCreateTable(stmt *ast.CreateTableStmt, ctx CompilerContext) (*DDLPlan, error) {
	return &DDLPlan{}, nil
}
