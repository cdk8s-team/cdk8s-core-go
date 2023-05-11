//go:build no_runtime_type_checking

package cdk8s

// Building without runtime type checking enabled, so all the below just return nil

func validateApiObject_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateApiObject_IsConstructParameters(x interface{}) error {
	return nil
}

func validateApiObject_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewApiObjectParameters(scope constructs.Construct, id *string, props *ApiObjectProps) error {
	return nil
}

