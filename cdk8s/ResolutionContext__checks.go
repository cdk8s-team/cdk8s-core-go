//go:build !no_runtime_type_checking

package cdk8s

import (
	"fmt"
)

func (r *jsiiProxy_ResolutionContext) validateReplaceValueParameters(newValue interface{}) error {
	if newValue == nil {
		return fmt.Errorf("parameter newValue is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ResolutionContext) validateSetReplacedParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ResolutionContext) validateSetReplacedValueParameters(val interface{}) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewResolutionContextParameters(obj ApiObject, key *[]*string, value interface{}) error {
	if obj == nil {
		return fmt.Errorf("parameter obj is required, but nil was provided")
	}

	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

