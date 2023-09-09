package planner

import types "colexecdb/pkg/query_engine/a_types"

type Expr interface {
	IsExpr() // only used to establish parent relationship
}

var _ Expr = new(ExprCol)
var _ Expr = new(ExprFunc)

type ExprCol struct {
	Type   types.Type
	ColIdx int32
}

func (e *ExprCol) IsExpr() {
	return
}

type ExprFunc struct {
	Type types.Type

	Name string
	Args []Expr
}

func (f *ExprFunc) IsExpr() {
	return
}
