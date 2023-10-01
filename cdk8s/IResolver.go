package cdk8s

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Contract for resolver objects.
type IResolver interface {
	// This function is invoked on every property during cdk8s synthesis.
	//
	// To replace a value, implementations must invoke `context.replaceValue`.
	Resolve(context ResolutionContext)
}

// The jsii proxy for IResolver
type jsiiProxy_IResolver struct {
	_ byte // padding
}

func (i *jsiiProxy_IResolver) Resolve(context ResolutionContext) {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"resolve",
		[]interface{}{context},
	)
}

