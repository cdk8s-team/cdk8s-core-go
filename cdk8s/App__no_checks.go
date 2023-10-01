//go:build no_runtime_type_checking

package cdk8s

// Building without runtime type checking enabled, so all the below just return nil

func validateApp_IsConstructParameters(x interface{}) error {
	return nil
}

func validateApp_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewAppParameters(props *AppProps) error {
	return nil
}

