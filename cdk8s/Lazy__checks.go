//go:build !no_runtime_type_checking

package cdk8s

import (
	"fmt"
)

func validateLazy_AnyParameters(producer IAnyProducer) error {
	if producer == nil {
		return fmt.Errorf("parameter producer is required, but nil was provided")
	}

	return nil
}

