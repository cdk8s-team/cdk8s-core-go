//go:build no_runtime_type_checking

package cdk8s

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_Include) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func validateNewIncludeParameters(scope constructs.Construct, id *string, props *IncludeProps) error {
	return nil
}

