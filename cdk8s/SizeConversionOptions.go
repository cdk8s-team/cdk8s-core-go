package cdk8s


// Options for how to convert time to a different unit.
type SizeConversionOptions struct {
	// How conversions should behave when it encounters a non-integer result.
	Rounding SizeRoundingBehavior `field:"optional" json:"rounding" yaml:"rounding"`
}

