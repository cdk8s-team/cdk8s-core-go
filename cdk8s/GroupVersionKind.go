package cdk8s


type GroupVersionKind struct {
	// The object's API version (e.g. `authorization.k8s.io/v1`).
	ApiVersion *string `field:"required" json:"apiVersion" yaml:"apiVersion"`
	// The object kind.
	Kind *string `field:"required" json:"kind" yaml:"kind"`
}

