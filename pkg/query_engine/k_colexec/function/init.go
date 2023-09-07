package function

func init() {

	for _, fn := range supportedOperators {
		allSupportedFunctions[fn.functionName] = fn
	}
}
