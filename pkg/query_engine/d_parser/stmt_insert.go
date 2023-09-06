package parser

type InsertStatement struct {
}

var _ Statement = new(InsertStatement)

func (i *InsertStatement) stmt() {
}
