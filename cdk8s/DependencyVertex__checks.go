//go:build !no_runtime_type_checking

package cdk8s

import (
	"fmt"
)

func (d *jsiiProxy_DependencyVertex) validateAddChildParameters(dep DependencyVertex) error {
	if dep == nil {
		return fmt.Errorf("parameter dep is required, but nil was provided")
	}

	return nil
}

