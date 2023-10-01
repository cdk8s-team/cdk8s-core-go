package cdk8s

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-core-go/cdk8s/jsii"
)

// Resolves implicit tokens.
type ImplicitTokenResolver interface {
	IResolver
	// This function is invoked on every property during cdk8s synthesis.
	//
	// To replace a value, implementations must invoke `context.replaceValue`.
	Resolve(context ResolutionContext)
}

// The jsii proxy struct for ImplicitTokenResolver
type jsiiProxy_ImplicitTokenResolver struct {
	jsiiProxy_IResolver
}

func NewImplicitTokenResolver() ImplicitTokenResolver {
	_init_.Initialize()

	j := jsiiProxy_ImplicitTokenResolver{}

	_jsii_.Create(
		"cdk8s.ImplicitTokenResolver",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewImplicitTokenResolver_Override(i ImplicitTokenResolver) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s.ImplicitTokenResolver",
		nil, // no parameters
		i,
	)
}

func (i *jsiiProxy_ImplicitTokenResolver) Resolve(context ResolutionContext) {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"resolve",
		[]interface{}{context},
	)
}

