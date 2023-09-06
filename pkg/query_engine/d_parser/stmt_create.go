package parser

type CreateTableStatement struct {
}

var _ Statement = new(CreateTableStatement)

func (i *CreateTableStatement) stmt() {
}
