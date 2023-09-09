package queryplan

import (
	"colexecdb/pkg/query_engine/f_catalog"
	"github.com/blastrain/vitess-sqlparser/tidbparser/ast"
)

func buildCreateTable(stmt *ast.CreateTableStmt, ctx catalog.Context) (*DDLPlan, error) {
	return &DDLPlan{}, nil
}
