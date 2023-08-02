package cdk8s


// Options for how to convert time to a different unit.
type TimeConversionOptions struct {
	// If `true`, conversions into a larger time unit (e.g. `Seconds` to `Minutes`) will fail if the result is not an integer.
	// Default: true.
	//
	Integral *bool `field:"optional" json:"integral" yaml:"integral"`
}

