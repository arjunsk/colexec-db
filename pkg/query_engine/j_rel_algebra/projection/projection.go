package projection

import (
	"bytes"
	vector "colexecdb/pkg/query_engine/b_vector"
	batch "colexecdb/pkg/query_engine/c_batch"
	process "colexecdb/pkg/query_engine/e_process"
	colexec "colexecdb/pkg/query_engine/k_colexec"
)

func String(arg any, buf *bytes.Buffer) {
	buf.WriteString("projection()")
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
	//if bat.Last() {
	//	proc.SetInputBatch(bat)
	//	return process.ExecNext, nil
	//}
	if bat.IsEmpty() {
		return process.ExecNext, nil
	}

	ap := arg.(*Argument)
	resultBat := batch.NewWithSize(len(ap.Es))

	// do projection.
	for i := range ap.ctr.projExecutors {
		vec, err := ap.ctr.projExecutors[i].Eval(proc, []*batch.Batch{bat})
		if err != nil {
			return process.ExecNext, err
		}
		resultBat.Vecs[i] = vec
	}

	//_ = FixProjectionResult(ap.ctr.projExecutors, resultBat)

	resultBat.SetRowCount(bat.GetRowCount())

	proc.SetInputBatch(resultBat)
	return process.ExecNext, nil
}

func FixProjectionResult(executors []colexec.ExpressionExecutor, rbat *batch.Batch) (err error) {

	//TODO: Understand why we need this code.

	alreadySet := make([]int, len(rbat.Vecs))
	for i := range alreadySet {
		alreadySet[i] = -1
	}

	finalVectors := make([]*vector.Vector, 0, len(rbat.Vecs))
	for i, oldVec := range rbat.Vecs {
		if alreadySet[i] < 0 {
			newVec := (*vector.Vector)(nil)
			if _, ok := executors[i].(*colexec.ColumnExpressionExecutor); ok {
				newVec, _ = oldVec.Dup()
			} else if functionExpr, ok := executors[i].(*colexec.FunctionExpressionExecutor); ok {
				newVec = functionExpr.ResultVector
				functionExpr.ResultVector = vector.NewVec(*functionExpr.ResultVector.GetType())
			} else {
				newVec, _ = oldVec.Dup()
			}

			finalVectors = append(finalVectors, newVec)
			indexOfNewVec := len(finalVectors) - 1
			for j := range rbat.Vecs {
				if rbat.Vecs[j] == oldVec {
					alreadySet[j] = indexOfNewVec
				}
			}
		}
	}

	for i, idx := range alreadySet {
		rbat.Vecs[i] = finalVectors[idx]
	}
	return nil
}
