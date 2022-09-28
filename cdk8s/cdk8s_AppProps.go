// This is the core library of Cloud Development Kit (CDK) for Kubernetes (cdk8s). cdk8s apps synthesize into standard Kubernetes manifests which can be applied to any Kubernetes cluster.
package cdk8s


type AppProps struct {
	// The directory to output Kubernetes manifests.
	Outdir *string `field:"optional" json:"outdir" yaml:"outdir"`
	// The file extension to use for rendered YAML files.
	OutputFileExtension *string `field:"optional" json:"outputFileExtension" yaml:"outputFileExtension"`
	// When set to true, the output directory will contain a `construct-metadata.json` file that holds construct related metadata on every resource in the app.
	RecordConstructMetadata *bool `field:"optional" json:"recordConstructMetadata" yaml:"recordConstructMetadata"`
	// How to divide the YAML output into files.
	YamlOutputType YamlOutputType `field:"optional" json:"yamlOutputType" yaml:"yamlOutputType"`
}

