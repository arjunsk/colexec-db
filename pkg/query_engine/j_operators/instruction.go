package operators

import process "colexecdb/pkg/query_engine/e_process"

// Prepare range instructions and do init work for each operator's argument by calling its prepare function
func Prepare(ins Operators, proc *process.Process) error {
	for _, in := range ins {
		if err := prepareFunc[in.Op](proc, in.Arg); err != nil {
			return err
		}
	}
	return nil
}

func Run(ins Operators, proc *process.Process) (end bool, err error) {
	var ok process.ExecStatus
	for i := 0; i < len(ins); i++ {
		if ok, err = execFunc[ins[i].Op](proc, ins[i].Arg); err != nil {
			return ok == process.ExecStop || end, err
		}

		if ok == process.ExecStop {
			end = true
		}
	}
	return end, err
}
