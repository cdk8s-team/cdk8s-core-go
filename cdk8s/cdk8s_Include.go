// This is the core library of Cloud Development Kit (CDK) for Kubernetes (cdk8s). cdk8s apps synthesize into standard Kubernetes manifests which can be applied to any Kubernetes cluster.
package cdk8s

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-core-go/cdk8s/jsii"

	"github.com/aws/constructs-go/constructs/v3"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/internal"
)

// Reads a YAML manifest from a file or a URL and defines all resources as API objects within the defined scope.
//
// The names (`metadata.name`) of imported resources will be preserved as-is
// from the manifest.
type Include interface {
	constructs.Construct
	// Returns all the included API objects.
	ApiObjects() *[]ApiObject
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if there the construct is valid.
	// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
	// instead of overriding this method.
	OnValidate() *[]*string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Include
type jsiiProxy_Include struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_Include) ApiObjects() *[]ApiObject {
	var returns *[]ApiObject
	_jsii_.Get(
		j,
		"apiObjects",
		&returns,
	)
	return returns
}


func NewInclude(scope constructs.Construct, id *string, props *IncludeProps) Include {
	_init_.Initialize()

	if err := validateNewIncludeParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Include{}

	_jsii_.Create(
		"cdk8s.Include",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewInclude_Override(i Include, scope constructs.Construct, id *string, props *IncludeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s.Include",
		[]interface{}{scope, id, props},
		i,
	)
}

func (i *jsiiProxy_Include) OnPrepare() {
	_jsii_.InvokeVoid(
		i,
		"onPrepare",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Include) OnSynthesize(session constructs.ISynthesisSession) {
	if err := i.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (i *jsiiProxy_Include) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Include) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

