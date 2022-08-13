// This is the core library of Cloud Development Kit (CDK) for Kubernetes (cdk8s). cdk8s apps synthesize into standard Kubernetes manifests which can be applied to any Kubernetes cluster.
package cdk8s

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-core-go/cdk8s/jsii"

	"github.com/aws/constructs-go/constructs/v3"
)

// Represents a Helm deployment.
//
// Use this construct to import an existing Helm chart and incorporate it into your constructs.
type Helm interface {
	Include
	// Returns all the included API objects.
	ApiObjects() *[]ApiObject
	// The helm release name.
	ReleaseName() *string
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

// The jsii proxy struct for Helm
type jsiiProxy_Helm struct {
	jsiiProxy_Include
}

func (j *jsiiProxy_Helm) ApiObjects() *[]ApiObject {
	var returns *[]ApiObject
	_jsii_.Get(
		j,
		"apiObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Helm) ReleaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseName",
		&returns,
	)
	return returns
}


func NewHelm(scope constructs.Construct, id *string, props *HelmProps) Helm {
	_init_.Initialize()

	j := jsiiProxy_Helm{}

	_jsii_.Create(
		"cdk8s.Helm",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewHelm_Override(h Helm, scope constructs.Construct, id *string, props *HelmProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s.Helm",
		[]interface{}{scope, id, props},
		h,
	)
}

func (h *jsiiProxy_Helm) OnPrepare() {
	_jsii_.InvokeVoid(
		h,
		"onPrepare",
		nil, // no parameters
	)
}

func (h *jsiiProxy_Helm) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		h,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (h *jsiiProxy_Helm) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_Helm) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

