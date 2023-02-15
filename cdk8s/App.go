// This is the core library of Cloud Development Kit (CDK) for Kubernetes (cdk8s). cdk8s apps synthesize into standard Kubernetes manifests which can be applied to any Kubernetes cluster.
package cdk8s

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-core-go/cdk8s/jsii"

	"github.com/aws/constructs-go/constructs/v3"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/internal"
)

// Represents a cdk8s application.
type App interface {
	constructs.Construct
	// Returns all the charts in this app, sorted topologically.
	Charts() *[]Chart
	// The output directory into which manifests will be synthesized.
	Outdir() *string
	// The file extension to use for rendered YAML files.
	OutputFileExtension() *string
	// How to divide the YAML output into files.
	YamlOutputType() YamlOutputType
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
	// Synthesizes all manifests to the output directory.
	Synth()
	// Synthesizes the app into a YAML string.
	//
	// Returns: A string with all YAML objects across all charts in this app.
	SynthYaml() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for App
type jsiiProxy_App struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_App) Charts() *[]Chart {
	var returns *[]Chart
	_jsii_.Get(
		j,
		"charts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) OutputFileExtension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFileExtension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) YamlOutputType() YamlOutputType {
	var returns YamlOutputType
	_jsii_.Get(
		j,
		"yamlOutputType",
		&returns,
	)
	return returns
}


// Defines an app.
func NewApp(props *AppProps) App {
	_init_.Initialize()

	if err := validateNewAppParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_App{}

	_jsii_.Create(
		"cdk8s.App",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Defines an app.
func NewApp_Override(a App, props *AppProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s.App",
		[]interface{}{props},
		a,
	)
}

func (a *jsiiProxy_App) OnPrepare() {
	_jsii_.InvokeVoid(
		a,
		"onPrepare",
		nil, // no parameters
	)
}

func (a *jsiiProxy_App) OnSynthesize(session constructs.ISynthesisSession) {
	if err := a.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (a *jsiiProxy_App) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_App) Synth() {
	_jsii_.InvokeVoid(
		a,
		"synth",
		nil, // no parameters
	)
}

func (a *jsiiProxy_App) SynthYaml() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"synthYaml",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_App) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

