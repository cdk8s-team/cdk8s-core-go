//go:build no_runtime_type_checking

package cdk8s

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Chart) validateGenerateObjectNameParameters(apiObject ApiObject) error {
	return nil
}

func (c *jsiiProxy_Chart) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func validateChart_IsChartParameters(x interface{}) error {
	return nil
}

func validateChart_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewChartParameters(scope constructs.Construct, id *string, props *ChartProps) error {
	return nil
}

