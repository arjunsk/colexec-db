package parser

import (
	"github.com/blastrain/vitess-sqlparser/tidbparser/ast"
	"github.com/blastrain/vitess-sqlparser/tidbparser/parser"
)

func Parse(sql string) (ast.StmtNode, error) {
	p := parser.New()
	return p.ParseOneStmt(sql, "", "")
}
