//go:build !no_runtime_type_checking

package cdk8s

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateTesting_AppParameters(props *AppProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateTesting_SynthParameters(chart Chart) error {
	if chart == nil {
		return fmt.Errorf("parameter chart is required, but nil was provided")
	}

	return nil
}

