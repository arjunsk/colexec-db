package function

import (
	"context"
)

func GetFunctionById(ctx context.Context, fid string) (f Overload, err error) {
	return allSupportedFunctions[fid].overloadFn, nil
}

var allSupportedFunctions map[string]FuncNew

type FuncNew struct {
	functionName string
	overloadFn   Overload
}
