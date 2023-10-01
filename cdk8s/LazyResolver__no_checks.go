//go:build no_runtime_type_checking

package cdk8s

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LazyResolver) validateResolveParameters(context ResolutionContext) error {
	return nil
}

