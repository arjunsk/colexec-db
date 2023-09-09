package output

import (
	"bytes"
	process "colexecdb/pkg/query_engine/e_process"
)

func String(arg any, buf *bytes.Buffer) {
	buf.WriteString("sql output")
}

func Prepare(_ *process.Process, _ any) error {
	return nil
}

func Call(proc *process.Process, arg any) (process.ExecStatus, error) {
	ap := arg.(*Argument)
	bat := proc.GetInputBatch()
	if bat == nil {
		return process.ExecStop, nil
	}
	if bat.IsEmpty() {
		proc.SetInputBatch(bat)
		return process.ExecNext, nil
	}
	if err := ap.Func(ap.Data, bat); err != nil {
		return process.ExecStop, err
	}
	return process.ExecNext, nil
}
