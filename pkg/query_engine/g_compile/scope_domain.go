package compile

import (
	process "colexecdb/pkg/query_engine/c_process"
	planner "colexecdb/pkg/query_engine/f_planner"
	rel_algebra "colexecdb/pkg/query_engine/j_rel_algebra"
)

type Scope struct {
	Magic ScopeType
	Plan  planner.Plan

	DataSource   *Source
	Process      *process.Process
	Instructions rel_algebra.Instructions
}
