package planner

type Optimizer interface {
	Optimize()
}

type Plan interface {
	Optimize([]Optimizer) Plan
}

var _ Plan = new(QueryPlan)
var _ Plan = new(DDLPlan)

type StatementType uint8

const (
	SELECT StatementType = iota
	INSERT
)

type QueryPlan struct {
	StatementType StatementType
	Params        []Expr
}

func (q *QueryPlan) Optimize(_ []Optimizer) Plan {
	return nil
}

type DdlType uint8

const (
	DdlCreateTable DdlType = iota
)

type DDLPlan struct {
	Type DdlType
}

func (d *DDLPlan) Optimize(_ []Optimizer) Plan {
	return nil
}
