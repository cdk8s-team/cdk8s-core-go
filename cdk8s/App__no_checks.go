//go:build no_runtime_type_checking

package cdk8s

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_App) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func validateNewAppParameters(props *AppProps) error {
	return nil
}

