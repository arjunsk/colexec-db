package parser

import (
	"github.com/blastrain/vitess-sqlparser/tidbparser/ast"
	"github.com/blastrain/vitess-sqlparser/tidbparser/parser"
)

type Statement = ast.StmtNode

func Parse(sql string) (Statement, error) {
	p := parser.New()
	return p.ParseOneStmt(sql, "", "")
}
