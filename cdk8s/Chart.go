package cdk8s

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-core-go/cdk8s/jsii"

	"github.com/aws/constructs-go/constructs/v3"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/internal"
)

type Chart interface {
	constructs.Construct
	// Returns all the included API objects.
	ApiObjects() *[]ApiObject
	// Labels applied to all resources in this chart.
	//
	// This is an immutable copy.
	Labels() *map[string]*string
	// The default namespace for all objects in this chart.
	Namespace() *string
	// Create a dependency between this Chart and other constructs.
	//
	// These can be other ApiObjects, Charts, or custom.
	AddDependency(dependencies ...constructs.IConstruct)
	// Generates a app-unique name for an object given it's construct node path.
	//
	// Different resource types may have different constraints on names
	// (`metadata.name`). The previous version of the name generator was
	// compatible with DNS_SUBDOMAIN but not with DNS_LABEL.
	//
	// For example, `Deployment` names must comply with DNS_SUBDOMAIN while
	// `Service` names must comply with DNS_LABEL.
	//
	// Since there is no formal specification for this, the default name
	// generation scheme for kubernetes objects in cdk8s was changed to DNS_LABEL,
	// since it’s the common denominator for all kubernetes resources
	// (supposedly).
	//
	// You can override this method if you wish to customize object names at the
	// chart level.
	GenerateObjectName(apiObject ApiObject) *string
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
	// Renders this chart to a set of Kubernetes JSON resources.
	//
	// Returns: array of resource manifests.
	ToJson() *[]interface{}
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Chart
type jsiiProxy_Chart struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_Chart) ApiObjects() *[]ApiObject {
	var returns *[]ApiObject
	_jsii_.Get(
		j,
		"apiObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Chart) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Chart) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}


func NewChart(scope constructs.Construct, id *string, props *ChartProps) Chart {
	_init_.Initialize()

	if err := validateNewChartParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Chart{}

	_jsii_.Create(
		"cdk8s.Chart",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewChart_Override(c Chart, scope constructs.Construct, id *string, props *ChartProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s.Chart",
		[]interface{}{scope, id, props},
		c,
	)
}

// Return whether the given object is a Chart.
//
// We do attribute detection since we can't reliably use 'instanceof'.
func Chart_IsChart(x interface{}) *bool {
	_init_.Initialize()

	if err := validateChart_IsChartParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s.Chart",
		"isChart",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Finds the chart in which a node is defined.
func Chart_Of(c constructs.IConstruct) Chart {
	_init_.Initialize()

	if err := validateChart_OfParameters(c); err != nil {
		panic(err)
	}
	var returns Chart

	_jsii_.StaticInvoke(
		"cdk8s.Chart",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Chart) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addDependency",
		args,
	)
}

func (c *jsiiProxy_Chart) GenerateObjectName(apiObject ApiObject) *string {
	if err := c.validateGenerateObjectNameParameters(apiObject); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"generateObjectName",
		[]interface{}{apiObject},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Chart) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Chart) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_Chart) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Chart) ToJson() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Chart) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

