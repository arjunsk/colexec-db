package planner

import (
	types "colexecdb/pkg/query_engine/a_types"
	catalog "colexecdb/pkg/query_engine/e_catalog"
	"context"
)

type CompilerContext interface {
	AppendTableDef(tableName string, schema *catalog.TableDef)
	ResolveColType(dbName string, tableName string, colName string) types.Type
	ResolveColIdx(dbName string, tableName string, colName string) int32
}

var _ CompilerContext = new(MockCompilerContext)

type MockCompilerContext struct {
	ctx    context.Context
	Tables map[string]*catalog.TableDef
}

func NewMockCompilerContext() *MockCompilerContext {
	return &MockCompilerContext{
		Tables: make(map[string]*catalog.TableDef),
	}
}

func (m *MockCompilerContext) AppendTableDef(tableName string, schema *catalog.TableDef) {
	m.Tables[tableName] = schema
}

func (m *MockCompilerContext) ResolveColType(dbName string, tableName string, colName string) types.Type {
	for _, col := range m.Tables[tableName].ColDefs {
		if col.Name == colName {
			return col.Type
		}
	}

	panic("no column found")
}

func (m *MockCompilerContext) ResolveColIdx(dbName string, tableName string, colName string) int32 {
	for i, col := range m.Tables[tableName].ColDefs {
		if col.Name == colName {
			return int32(i)
		}
	}

	panic("no column found")
}
