//go:build !no_runtime_type_checking

package cdk8s

import (
	"fmt"
)

func (i *jsiiProxy_IResolver) validateResolveParameters(context ResolutionContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

