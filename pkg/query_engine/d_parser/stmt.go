package parser

type Statement interface {
	stmt()
	travel()
}
