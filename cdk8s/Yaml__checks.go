//go:build !no_runtime_type_checking

package cdk8s

import (
	"fmt"
)

func validateYaml_FormatObjectsParameters(docs *[]interface{}) error {
	if docs == nil {
		return fmt.Errorf("parameter docs is required, but nil was provided")
	}

	return nil
}

func validateYaml_LoadParameters(urlOrFile *string) error {
	if urlOrFile == nil {
		return fmt.Errorf("parameter urlOrFile is required, but nil was provided")
	}

	return nil
}

func validateYaml_SaveParameters(filePath *string, docs *[]interface{}) error {
	if filePath == nil {
		return fmt.Errorf("parameter filePath is required, but nil was provided")
	}

	if docs == nil {
		return fmt.Errorf("parameter docs is required, but nil was provided")
	}

	return nil
}

func validateYaml_TmpParameters(docs *[]interface{}) error {
	if docs == nil {
		return fmt.Errorf("parameter docs is required, but nil was provided")
	}

	return nil
}

