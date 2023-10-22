package logicalplan

import (
	"colexecdb/pkg/query_engine/f_catalog"
	"github.com/blastrain/vitess-sqlparser/tidbparser/ast"
)

func buildInsert(stmt *ast.InsertStmt, ctx catalog.Context) (*QueryPlan, error) {
	return &QueryPlan{
		StatementType: INSERT,
	}, nil
}
