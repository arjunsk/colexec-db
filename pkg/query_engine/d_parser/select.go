package parser

type SelectStatement struct {
	Fields         Fields
	Sources        Sources
	WhereCondition *ExprNode
}

func (*SelectStatement) stmt() {}

func (stmt *SelectStatement) travel() {

}
