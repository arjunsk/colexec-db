package logicalplan

import (
	"colexecdb/pkg/query_engine/f_catalog"
	"errors"
	"github.com/blastrain/vitess-sqlparser/tidbparser/ast"
)

func BuildPlan(stmt ast.StmtNode, ctx catalog.Context) (Plan, error) {
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
