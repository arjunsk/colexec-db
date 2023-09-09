package function

import "context"

type FuncNew struct {
	functionName string
	overloadFn   Overload
}

var allSupportedFunctions map[string]FuncNew

func init() {

	allSupportedFunctions = make(map[string]FuncNew)

	for _, fn := range supportedOperators {
		allSupportedFunctions[fn.functionName] = fn
	}
}

func GetFunctionById(ctx context.Context, fid string) (f Overload, err error) {
	return allSupportedFunctions[fid].overloadFn, nil
}
