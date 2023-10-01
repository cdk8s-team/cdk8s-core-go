//go:build no_runtime_type_checking

package cdk8s

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ResolutionContext) validateReplaceValueParameters(newValue interface{}) error {
	return nil
}

func (j *jsiiProxy_ResolutionContext) validateSetReplacedParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_ResolutionContext) validateSetReplacedValueParameters(val interface{}) error {
	return nil
}

func validateNewResolutionContextParameters(obj ApiObject, key *[]*string, value interface{}) error {
	return nil
}

