//go:build !no_runtime_type_checking

package cdk8s

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v3"
)

func validateNewDependencyGraphParameters(node constructs.Node) error {
	if node == nil {
		return fmt.Errorf("parameter node is required, but nil was provided")
	}

	return nil
}

