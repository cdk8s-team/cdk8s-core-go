//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// This is the core library of Cloud Development Kit (CDK) for Kubernetes (cdk8s). cdk8s apps synthesize into standard Kubernetes manifests which can be applied to any Kubernetes cluster.
package cdk8s

// Building without runtime type checking enabled, so all the below just return nil

func validateApiObject_IsConstructParameters(x interface{}) error {
	return nil
}

func validateApiObject_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewApiObjectParameters(scope constructs.Construct, id *string, props *ApiObjectProps) error {
	return nil
}
