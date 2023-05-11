//go:build no_runtime_type_checking

package cdk8s

// Building without runtime type checking enabled, so all the below just return nil

func validateYaml_FormatObjectsParameters(docs *[]interface{}) error {
	return nil
}

func validateYaml_LoadParameters(urlOrFile *string) error {
	return nil
}

func validateYaml_SaveParameters(filePath *string, docs *[]interface{}) error {
	return nil
}

func validateYaml_TmpParameters(docs *[]interface{}) error {
	return nil
}

