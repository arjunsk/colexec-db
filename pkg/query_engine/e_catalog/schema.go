package catalog

import types "colexecdb/pkg/query_engine/a_types"

type Schema struct {
	Name    string
	ColDefs []*ColDef
}

type ColDef struct {
	Type        types.Type
	Name        string
	Idx         int
	NullAbility bool
}

func (s *Schema) AppendCol(name string, typ types.Type) error {
	def := &ColDef{
		Name:        name,
		Type:        typ,
		NullAbility: true,
	}
	return s.AppendColDef(def)
}

func (s *Schema) AppendColDef(def *ColDef) (err error) {
	def.Idx = len(s.ColDefs)
	s.ColDefs = append(s.ColDefs, def)
	return
}
