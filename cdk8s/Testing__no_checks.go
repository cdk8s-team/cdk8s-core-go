//go:build no_runtime_type_checking

package cdk8s

// Building without runtime type checking enabled, so all the below just return nil

func validateTesting_AppParameters(props *AppProps) error {
	return nil
}

func validateTesting_SynthParameters(chart Chart) error {
	return nil
}

