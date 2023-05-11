//go:build no_runtime_type_checking

package cdk8s

// Building without runtime type checking enabled, so all the below just return nil

func validateCron_ScheduleParameters(options *CronOptions) error {
	return nil
}

func validateNewCronParameters(cronOptions *CronOptions) error {
	return nil
}

