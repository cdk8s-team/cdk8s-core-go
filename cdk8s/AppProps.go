package cdk8s


type AppProps struct {
	// The directory to output Kubernetes manifests.
	//
	// If you synthesize your application using `cdk8s synth`, you must
	// also pass this value to the CLI using the `--output` option or
	// the `output` property in the `cdk8s.yaml` configuration file.
	// Otherwise, the CLI will not know about the output directory,
	// and synthesis will fail.
	//
	// This property is intended for internal and testing use.
	// Default: - CDK8S_OUTDIR if defined, otherwise "dist".
	//
	Outdir *string `field:"optional" json:"outdir" yaml:"outdir"`
	// The file extension to use for rendered YAML files.
	// Default: .k8s.yaml
	//
	OutputFileExtension *string `field:"optional" json:"outputFileExtension" yaml:"outputFileExtension"`
	// When set to true, the output directory will contain a `construct-metadata.json` file that holds construct related metadata on every resource in the app.
	// Default: false.
	//
	RecordConstructMetadata *bool `field:"optional" json:"recordConstructMetadata" yaml:"recordConstructMetadata"`
	// A list of resolvers that can be used to replace property values before they are written to the manifest file.
	//
	// When multiple resolvers are passed,
	// they are invoked by order in the list, and only the first one that applies
	// (e.g calls `context.replaceValue`) is invoked.
	// See: https://cdk8s.io/docs/latest/basics/app/#resolvers
	//
	// Default: - no resolvers.
	//
	Resolvers *[]IResolver `field:"optional" json:"resolvers" yaml:"resolvers"`
	// How to divide the YAML output into files.
	// Default: YamlOutputType.FILE_PER_CHART
	//
	YamlOutputType YamlOutputType `field:"optional" json:"yamlOutputType" yaml:"yamlOutputType"`
}

