//go:build !no_runtime_type_checking

package cdk8s

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateCron_ScheduleParameters(options *CronOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateNewCronParameters(cronOptions *CronOptions) error {
	if err := _jsii_.ValidateStruct(cronOptions, func() string { return "parameter cronOptions" }); err != nil {
		return err
	}

	return nil
}

