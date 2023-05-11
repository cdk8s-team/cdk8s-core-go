//go:build no_runtime_type_checking

package cdk8s

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApiObjectMetadataDefinition) validateAddParameters(key *string, value interface{}) error {
	return nil
}

func (a *jsiiProxy_ApiObjectMetadataDefinition) validateAddAnnotationParameters(key *string, value *string) error {
	return nil
}

func (a *jsiiProxy_ApiObjectMetadataDefinition) validateAddLabelParameters(key *string, value *string) error {
	return nil
}

func (a *jsiiProxy_ApiObjectMetadataDefinition) validateAddOwnerReferenceParameters(owner *OwnerReference) error {
	return nil
}

func (a *jsiiProxy_ApiObjectMetadataDefinition) validateGetLabelParameters(key *string) error {
	return nil
}

func validateNewApiObjectMetadataDefinitionParameters(options *ApiObjectMetadata) error {
	return nil
}

