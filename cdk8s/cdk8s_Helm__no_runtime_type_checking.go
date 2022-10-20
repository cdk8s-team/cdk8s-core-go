//go:build no_runtime_type_checking

// This is the core library of Cloud Development Kit (CDK) for Kubernetes (cdk8s). cdk8s apps synthesize into standard Kubernetes manifests which can be applied to any Kubernetes cluster.
package cdk8s

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_Helm) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func validateNewHelmParameters(scope constructs.Construct, id *string, props *HelmProps) error {
	return nil
}

