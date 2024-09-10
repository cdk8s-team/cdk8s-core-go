package cdk8s


// Options for how to convert size to a different unit.
type SizeConversionOptions struct {
	// How conversions should behave when it encounters a non-integer result.
	// Default: SizeRoundingBehavior.FAIL
	//
	Rounding SizeRoundingBehavior `field:"optional" json:"rounding" yaml:"rounding"`
}

