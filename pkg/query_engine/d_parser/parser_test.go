package parser

import (
	"github.com/blastrain/vitess-sqlparser/tidbparser/ast"
	"github.com/blastrain/vitess-sqlparser/tidbparser/parser"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParse1(t *testing.T) {
	p := parser.New()
	res1, _ := p.ParseOneStmt("select a from t1;", "", "")

	res2 := res1.(*ast.SelectStmt)
	require.Equal(t, 1, len(res2.Fields.Fields))

	res3 := res2.Fields.Fields[0].Expr.(*ast.ColumnNameExpr)
	require.Equal(t, "a", res3.Name.Name.L)

	require.Equal(t, "t1", res2.From.TableRefs.Left.(*ast.TableSource).Source.(*ast.TableName).Name.L)
}

func TestParse2(t *testing.T) {
	p := parser.New()
	res1, _ := p.ParseOneStmt("select a, sqrt(b) from t1;", "", "")

	res2 := res1.(*ast.SelectStmt)
	require.Equal(t, 2, len(res2.Fields.Fields))

	res3 := res2.Fields.Fields[0].Expr.(*ast.ColumnNameExpr)
	require.Equal(t, "a", res3.Name.Name.L)

	res4 := res2.Fields.Fields[1].Expr.(*ast.FuncCallExpr)
	require.Equal(t, "sqrt", res4.FnName.L)
	require.Equal(t, 1, len(res4.Args))
	require.Equal(t, "b", res4.Args[0].(*ast.ColumnNameExpr).Name.Name.L)

}
