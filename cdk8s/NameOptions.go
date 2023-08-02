package cdk8s


// Options for name generation.
type NameOptions struct {
	// Delimiter to use between components.
	// Default: "-".
	//
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// Extra components to include in the name.
	// Default: [] use the construct path components.
	//
	Extra *[]*string `field:"optional" json:"extra" yaml:"extra"`
	// Include a short hash as last part of the name.
	// Default: true.
	//
	IncludeHash *bool `field:"optional" json:"includeHash" yaml:"includeHash"`
	// Maximum allowed length for the name.
	// Default: 63.
	//
	MaxLen *float64 `field:"optional" json:"maxLen" yaml:"maxLen"`
}

