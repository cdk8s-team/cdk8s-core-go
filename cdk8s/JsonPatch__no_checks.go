//go:build no_runtime_type_checking

package cdk8s

// Building without runtime type checking enabled, so all the below just return nil

func validateJsonPatch_AddParameters(path *string, value interface{}) error {
	return nil
}

func validateJsonPatch_ApplyParameters(document interface{}) error {
	return nil
}

func validateJsonPatch_CopyParameters(from *string, path *string) error {
	return nil
}

func validateJsonPatch_MoveParameters(from *string, path *string) error {
	return nil
}

func validateJsonPatch_RemoveParameters(path *string) error {
	return nil
}

func validateJsonPatch_ReplaceParameters(path *string, value interface{}) error {
	return nil
}

func validateJsonPatch_TestParameters(path *string, value interface{}) error {
	return nil
}

