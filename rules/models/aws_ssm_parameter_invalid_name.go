// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSsmParameterInvalidNameRule checks the pattern is valid
type AwsSsmParameterInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsSsmParameterInvalidNameRule returns new rule with default attributes
func NewAwsSsmParameterInvalidNameRule() *AwsSsmParameterInvalidNameRule {
	return &AwsSsmParameterInvalidNameRule{
		resourceType:  "aws_ssm_parameter",
		attributeName: "name",
		max:           2048,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsSsmParameterInvalidNameRule) Name() string {
	return "aws_ssm_parameter_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmParameterInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsmParameterInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmParameterInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmParameterInvalidNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"name must be 2048 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"name must be 1 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
