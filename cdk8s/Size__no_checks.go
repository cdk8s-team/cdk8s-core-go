//go:build no_runtime_type_checking

package cdk8s

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Size) validateToGibibytesParameters(opts *SizeConversionOptions) error {
	return nil
}

func (s *jsiiProxy_Size) validateToKibibytesParameters(opts *SizeConversionOptions) error {
	return nil
}

func (s *jsiiProxy_Size) validateToMebibytesParameters(opts *SizeConversionOptions) error {
	return nil
}

func (s *jsiiProxy_Size) validateToPebibytesParameters(opts *SizeConversionOptions) error {
	return nil
}

func (s *jsiiProxy_Size) validateToTebibytesParameters(opts *SizeConversionOptions) error {
	return nil
}

func validateSize_GibibytesParameters(amount *float64) error {
	return nil
}

func validateSize_KibibytesParameters(amount *float64) error {
	return nil
}

func validateSize_MebibytesParameters(amount *float64) error {
	return nil
}

func validateSize_PebibyteParameters(amount *float64) error {
	return nil
}

func validateSize_TebibytesParameters(amount *float64) error {
	return nil
}

