package function

func init() {

	allSupportedFunctions = make(map[string]FuncNew)

	for _, fn := range supportedOperators {
		allSupportedFunctions[fn.functionName] = fn
	}
}
