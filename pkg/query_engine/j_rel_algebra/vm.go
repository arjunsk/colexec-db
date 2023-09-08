package rel_algebra

import process "colexecdb/pkg/query_engine/e_process"

// Prepare range instructions and do init work for each operator's argument by calling its prepare function
func Prepare(ins Instructions, proc *process.Process) error {
	for _, in := range ins {
		if err := prepareFunc[in.Op](proc, in.Arg); err != nil {
			return err
		}
	}
	return nil
}

func Run(ins Instructions, proc *process.Process) (end bool, err error) {
	return fubarRun(ins, proc, 0)
}

func fubarRun(ins Instructions, proc *process.Process, start int) (end bool, err error) {
	var fubarStack []int
	var ok process.ExecStatus

	for i := start; i < len(ins); i++ {
		if ok, err = execFunc[ins[i].Op](proc, ins[i].Arg); err != nil {
			return ok == process.ExecStop || end, err
		}

		if ok == process.ExecStop {
			end = true
		} else if ok == process.ExecHasMore {
			fubarStack = append(fubarStack, i)
		}
	}

	// run the stack backwards.
	// Only executed for process.ExecHasMore
	for i := len(fubarStack) - 1; i >= 0; i-- {
		// Note that, we are passing the start argument as the idx, where the process execution returned ExecHasMore
		end, err = fubarRun(ins, proc, fubarStack[i])
		if end || err != nil {
			return end, err
		}
	}
	return end, err
}
