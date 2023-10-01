package cdk8s

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-core-go/cdk8s/jsii"
)

// Resolvers instanecs of `Lazy`.
type LazyResolver interface {
	IResolver
	// This function is invoked on every property during cdk8s synthesis.
	//
	// To replace a value, implementations must invoke `context.replaceValue`.
	Resolve(context ResolutionContext)
}

// The jsii proxy struct for LazyResolver
type jsiiProxy_LazyResolver struct {
	jsiiProxy_IResolver
}

func NewLazyResolver() LazyResolver {
	_init_.Initialize()

	j := jsiiProxy_LazyResolver{}

	_jsii_.Create(
		"cdk8s.LazyResolver",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewLazyResolver_Override(l LazyResolver) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s.LazyResolver",
		nil, // no parameters
		l,
	)
}

func (l *jsiiProxy_LazyResolver) Resolve(context ResolutionContext) {
	if err := l.validateResolveParameters(context); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"resolve",
		[]interface{}{context},
	)
}

