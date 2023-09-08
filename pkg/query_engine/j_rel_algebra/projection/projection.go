package projection

import (
	"bytes"
	batch "colexecdb/pkg/query_engine/c_batch"
	process "colexecdb/pkg/query_engine/e_process"
	colexec "colexecdb/pkg/query_engine/k_colexec"
)

func String(arg any, buf *bytes.Buffer) {
	buf.WriteString("projection(")
	buf.WriteString(")")
}

func Prepare(proc *process.Process, arg any) (err error) {
	ap := arg.(*Argument)
	ap.ctr = new(container)
	ap.ctr.projExecutors, err = colexec.NewExpressionExecutorsFromPlanExpressions(proc, ap.Es)

	return err
}

func Call(proc *process.Process, arg any) (process.ExecStatus, error) {

	bat := proc.GetInputBatch()
	if bat == nil {
		proc.SetInputBatch(nil)
		return process.ExecStop, nil
	}
	if bat.Last() {
		proc.SetInputBatch(bat)
		return process.ExecNext, nil
	}
	if bat.IsEmpty() {
		return process.ExecNext, nil
	}

	ap := arg.(*Argument)
	rbat := batch.NewWithSize(len(ap.Es))

	// do projection.
	for i := range ap.ctr.projExecutors {
		vec, err := ap.ctr.projExecutors[i].Eval(proc, []*batch.Batch{bat})
		if err != nil {
			return process.ExecNext, err
		}
		rbat.Vecs[i] = vec
	}

	rbat.SetRowCount(bat.GetRowCount())

	proc.SetInputBatch(rbat)
	return process.ExecNext, nil
}
