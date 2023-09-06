package planner

type Optimizer interface {
	Optimize()
}

type Plan interface {
	Optimize([]Optimizer) Plan
}

var _ Plan = new(QueryPlan)
var _ Plan = new(DataDefinitionPlan)

type StatementType uint8

const (
	SELECT StatementType = iota
	INSERT
)

type QueryPlan struct {
	statementType StatementType
	params        Expr
}

func (q *QueryPlan) Optimize(_ []Optimizer) Plan {
	return nil
}

type DataDefinitionPlan struct {
}

func (d *DataDefinitionPlan) Optimize(_ []Optimizer) Plan {
	return nil
}
