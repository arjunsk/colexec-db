package catalog

import (
	types "colexecdb/pkg/query_engine/a_types"
	"fmt"
	"time"
)

func MockSchemaAll(colCnt int) *Schema {
	schema := Schema{
		Name:    time.Now().String(),
		ColDefs: make([]*ColDef, colCnt),
	}

	prefix := "mock_"
	var typ types.Type

	for i := 0; i < colCnt; i++ {
		switch i % 20 {
		default:
			typ = types.T_int32.ToType()
		}
		name := fmt.Sprintf("%s%d", prefix, i)
		_ = schema.AppendCol(name, typ)
	}
	return &schema
}
