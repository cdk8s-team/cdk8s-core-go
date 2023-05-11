package cdk8s


type IncludeProps struct {
	// Local file path or URL which includes a Kubernetes YAML manifest.
	//
	// Example:
	//   mymanifest.yaml
	//
	Url *string `field:"required" json:"url" yaml:"url"`
}

