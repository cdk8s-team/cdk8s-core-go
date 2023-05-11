//go:build no_runtime_type_checking

package cdk8s

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DependencyVertex) validateAddChildParameters(dep DependencyVertex) error {
	return nil
}

