package cdk8s


// The method to divide YAML output into files.
type YamlOutputType string

const (
	// All resources are output into a single YAML file.
	YamlOutputType_FILE_PER_APP YamlOutputType = "FILE_PER_APP"
	// Resources are split into seperate files by chart.
	YamlOutputType_FILE_PER_CHART YamlOutputType = "FILE_PER_CHART"
	// Each resource is output to its own file.
	YamlOutputType_FILE_PER_RESOURCE YamlOutputType = "FILE_PER_RESOURCE"
	// Each chart in its own folder and each resource in its own file.
	YamlOutputType_FOLDER_PER_CHART_FILE_PER_RESOURCE YamlOutputType = "FOLDER_PER_CHART_FILE_PER_RESOURCE"
)

