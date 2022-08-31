// This is the core library of Cloud Development Kit (CDK) for Kubernetes (cdk8s). cdk8s apps synthesize into standard Kubernetes manifests which can be applied to any Kubernetes cluster.
package cdk8s

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-core-go/cdk8s/jsii"
)

// Object metadata.
type ApiObjectMetadataDefinition interface {
	// The name of the API object.
	//
	// If a name is specified in `metadata.name` this will be the name returned.
	// Otherwise, a name will be generated by calling
	// `Chart.of(this).generatedObjectName(this)`, which by default uses the
	// construct path to generate a DNS-compatible name for the resource.
	Name() *string
	// The object's namespace.
	Namespace() *string
	// Adds an arbitrary key/value to the object metadata.
	Add(key *string, value interface{})
	// Add an annotation.
	AddAnnotation(key *string, value *string)
	// Add one or more finalizers.
	AddFinalizers(finalizers ...*string)
	// Add a label.
	AddLabel(key *string, value *string)
	// Add an owner.
	AddOwnerReference(owner *OwnerReference)
	// Returns: a value of a label or undefined.
	GetLabel(key *string) *string
	// Synthesizes a k8s ObjectMeta for this metadata set.
	ToJson() interface{}
}

// The jsii proxy struct for ApiObjectMetadataDefinition
type jsiiProxy_ApiObjectMetadataDefinition struct {
	_ byte // padding
}

func (j *jsiiProxy_ApiObjectMetadataDefinition) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiObjectMetadataDefinition) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}


func NewApiObjectMetadataDefinition(options *ApiObjectMetadata) ApiObjectMetadataDefinition {
	_init_.Initialize()

	if err := validateNewApiObjectMetadataDefinitionParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiObjectMetadataDefinition{}

	_jsii_.Create(
		"cdk8s.ApiObjectMetadataDefinition",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewApiObjectMetadataDefinition_Override(a ApiObjectMetadataDefinition, options *ApiObjectMetadata) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s.ApiObjectMetadataDefinition",
		[]interface{}{options},
		a,
	)
}

func (a *jsiiProxy_ApiObjectMetadataDefinition) Add(key *string, value interface{}) {
	if err := a.validateAddParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"add",
		[]interface{}{key, value},
	)
}

func (a *jsiiProxy_ApiObjectMetadataDefinition) AddAnnotation(key *string, value *string) {
	if err := a.validateAddAnnotationParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addAnnotation",
		[]interface{}{key, value},
	)
}

func (a *jsiiProxy_ApiObjectMetadataDefinition) AddFinalizers(finalizers ...*string) {
	args := []interface{}{}
	for _, a := range finalizers {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addFinalizers",
		args,
	)
}

func (a *jsiiProxy_ApiObjectMetadataDefinition) AddLabel(key *string, value *string) {
	if err := a.validateAddLabelParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addLabel",
		[]interface{}{key, value},
	)
}

func (a *jsiiProxy_ApiObjectMetadataDefinition) AddOwnerReference(owner *OwnerReference) {
	if err := a.validateAddOwnerReferenceParameters(owner); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOwnerReference",
		[]interface{}{owner},
	)
}

func (a *jsiiProxy_ApiObjectMetadataDefinition) GetLabel(key *string) *string {
	if err := a.validateGetLabelParameters(key); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getLabel",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiObjectMetadataDefinition) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

