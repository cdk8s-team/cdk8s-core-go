//go:build no_runtime_type_checking

package cdk8s

// Building without runtime type checking enabled, so all the below just return nil

func validateNames_ToDnsLabelParameters(scope constructs.Construct, options *NameOptions) error {
	return nil
}

func validateNames_ToLabelValueParameters(scope constructs.Construct, options *NameOptions) error {
	return nil
}

