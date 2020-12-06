// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAppmeshVirtualRouterInvalidMeshNameRule checks the pattern is valid
type AwsAppmeshVirtualRouterInvalidMeshNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsAppmeshVirtualRouterInvalidMeshNameRule returns new rule with default attributes
func NewAwsAppmeshVirtualRouterInvalidMeshNameRule() *AwsAppmeshVirtualRouterInvalidMeshNameRule {
	return &AwsAppmeshVirtualRouterInvalidMeshNameRule{
		resourceType:  "aws_appmesh_virtual_router",
		attributeName: "mesh_name",
		max:           255,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsAppmeshVirtualRouterInvalidMeshNameRule) Name() string {
	return "aws_appmesh_virtual_router_invalid_mesh_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppmeshVirtualRouterInvalidMeshNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAppmeshVirtualRouterInvalidMeshNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAppmeshVirtualRouterInvalidMeshNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppmeshVirtualRouterInvalidMeshNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"mesh_name must be 255 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"mesh_name must be 1 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
