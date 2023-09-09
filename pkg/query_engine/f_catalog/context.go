package catalog

import (
	types "colexecdb/pkg/query_engine/a_types"
	"context"
)

type Context interface {
	AppendTableDef(tableName string, schema *TableDef)
	ResolveColType(dbName string, tableName string, colName string) types.Type
	ResolveColIdx(dbName string, tableName string, colName string) int32
}

var _ Context = new(SchemaContext)

type SchemaContext struct {
	ctx    context.Context
	Tables map[string]*TableDef
}

func NewMockSchemaContext() *SchemaContext {
	return &SchemaContext{
		Tables: make(map[string]*TableDef),
	}
}

func (m *SchemaContext) AppendTableDef(tableName string, schema *TableDef) {
	m.Tables[tableName] = schema
}

func (m *SchemaContext) ResolveColType(dbName string, tableName string, colName string) types.Type {
	for _, col := range m.Tables[tableName].ColDefs {
		if col.Name == colName {
			return col.Type
		}
	}

	panic("no column found")
}

func (m *SchemaContext) ResolveColIdx(dbName string, tableName string, colName string) int32 {
	for i, col := range m.Tables[tableName].ColDefs {
		if col.Name == colName {
			return int32(i)
		}
	}

	panic("no column found")
}
