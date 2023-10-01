//go:build no_runtime_type_checking

package cdk8s

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IResolver) validateResolveParameters(context ResolutionContext) error {
	return nil
}

