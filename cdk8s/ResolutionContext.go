package cdk8s

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-core-go/cdk8s/jsii"
)

// Context object for a specific resolution process.
type ResolutionContext interface {
	// Which key is currently being resolved.
	Key() *[]*string
	// Which ApiObject is currently being resolved.
	Obj() ApiObject
	// Whether or not the value was replaced by invoking the `replaceValue` method.
	Replaced() *bool
	SetReplaced(val *bool)
	// The replaced value that was set via the `replaceValue` method.
	ReplacedValue() interface{}
	SetReplacedValue(val interface{})
	// The value associated to the key currently being resolved.
	Value() interface{}
	// Replaces the original value in this resolution context with a new value.
	//
	// The new value is what will end up in the manifest.
	ReplaceValue(newValue interface{})
}

// The jsii proxy struct for ResolutionContext
type jsiiProxy_ResolutionContext struct {
	_ byte // padding
}

func (j *jsiiProxy_ResolutionContext) Key() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResolutionContext) Obj() ApiObject {
	var returns ApiObject
	_jsii_.Get(
		j,
		"obj",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResolutionContext) Replaced() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"replaced",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResolutionContext) ReplacedValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replacedValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResolutionContext) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func NewResolutionContext(obj ApiObject, key *[]*string, value interface{}) ResolutionContext {
	_init_.Initialize()

	if err := validateNewResolutionContextParameters(obj, key, value); err != nil {
		panic(err)
	}
	j := jsiiProxy_ResolutionContext{}

	_jsii_.Create(
		"cdk8s.ResolutionContext",
		[]interface{}{obj, key, value},
		&j,
	)

	return &j
}

func NewResolutionContext_Override(r ResolutionContext, obj ApiObject, key *[]*string, value interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s.ResolutionContext",
		[]interface{}{obj, key, value},
		r,
	)
}

func (j *jsiiProxy_ResolutionContext)SetReplaced(val *bool) {
	if err := j.validateSetReplacedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaced",
		val,
	)
}

func (j *jsiiProxy_ResolutionContext)SetReplacedValue(val interface{}) {
	if err := j.validateSetReplacedValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replacedValue",
		val,
	)
}

func (r *jsiiProxy_ResolutionContext) ReplaceValue(newValue interface{}) {
	if err := r.validateReplaceValueParameters(newValue); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"replaceValue",
		[]interface{}{newValue},
	)
}

