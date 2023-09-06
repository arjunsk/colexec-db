package scope

import (
	process "colexecdb/pkg/query_engine/c_process"
	planner "colexecdb/pkg/query_engine/e_planner"
	compile "colexecdb/pkg/query_engine/f_compile"
	rel_algebra "colexecdb/pkg/query_engine/i_rel_algebra"
)

type Scope struct {
	Magic compile.ScopeType
	Plan  planner.Plan

	DataSource   *compile.Source
	Process      *process.Process
	Instructions rel_algebra.Instructions
}
