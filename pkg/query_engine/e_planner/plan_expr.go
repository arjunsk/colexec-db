package planner

import types "colexecdb/pkg/query_engine/a_types"

type Expr interface {
	Typ() types.Type
}

var _ Expr = new(ExprCol)
var _ Expr = new(ExprFunc)

type ExprCol struct {
	typ     types.Type
	ColName string
}

func (e *ExprCol) Typ() types.Type {
	return e.typ
}

type ExprFunc struct {
	typ types.Type

	Name string
	Args []Expr
}

func (f *ExprFunc) Typ() types.Type {
	return f.typ
}
