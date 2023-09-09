package catalog

import types "colexecdb/pkg/query_engine/a_types"

type TableDef struct {
	Name    string
	ColDefs []*ColDef
}

type ColDef struct {
	Name        string
	Type        types.Type
	Idx         int
	NullAbility bool
}

func (s *TableDef) AppendCol(name string, typ types.Type) error {
	def := &ColDef{
		Name:        name,
		Type:        typ,
		NullAbility: true,
	}
	return s.appendColDef(def)
}

func (s *TableDef) appendColDef(def *ColDef) (err error) {
	def.Idx = len(s.ColDefs)
	s.ColDefs = append(s.ColDefs, def)
	return
}
