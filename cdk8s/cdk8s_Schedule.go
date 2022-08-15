// This is the core library of Cloud Development Kit (CDK) for Kubernetes (cdk8s). cdk8s apps synthesize into standard Kubernetes manifests which can be applied to any Kubernetes cluster.
package cdk8s

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-core-go/cdk8s/jsii"
)

// Represents a schedule.
//
// Note that rates cannot be defined in fractions of minutes.
type Schedule interface {
	// Retrieve the expression for this schedule.
	ExpressionString() *string
}

// The jsii proxy struct for Schedule
type jsiiProxy_Schedule struct {
	_ byte // padding
}

func (j *jsiiProxy_Schedule) ExpressionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expressionString",
		&returns,
	)
	return returns
}


func NewSchedule_Override(s Schedule) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s.Schedule",
		nil, // no parameters
		s,
	)
}

// Create a schedule from a set of cron fields.
func Schedule_Cron(options *CronOptions) Schedule {
	_init_.Initialize()

	var returns Schedule

	_jsii_.StaticInvoke(
		"cdk8s.Schedule",
		"cron",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Construct a schedule from a literal schedule expression.
func Schedule_Expression(expression *string) Schedule {
	_init_.Initialize()

	var returns Schedule

	_jsii_.StaticInvoke(
		"cdk8s.Schedule",
		"expression",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

// Construct a schedule from an interval and a time unit.
//
// Rates may be defined with any unit of time, but when converted into minutes, the duration must be a positive whole number of minutes.
func Schedule_Rate(duration Duration) Schedule {
	_init_.Initialize()

	var returns Schedule

	_jsii_.StaticInvoke(
		"cdk8s.Schedule",
		"rate",
		[]interface{}{duration},
		&returns,
	)

	return returns
}

