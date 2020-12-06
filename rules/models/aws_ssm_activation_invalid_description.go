// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSsmActivationInvalidDescriptionRule checks the pattern is valid
type AwsSsmActivationInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsSsmActivationInvalidDescriptionRule returns new rule with default attributes
func NewAwsSsmActivationInvalidDescriptionRule() *AwsSsmActivationInvalidDescriptionRule {
	return &AwsSsmActivationInvalidDescriptionRule{
		resourceType:  "aws_ssm_activation",
		attributeName: "description",
		max:           256,
	}
}

// Name returns the rule name
func (r *AwsSsmActivationInvalidDescriptionRule) Name() string {
	return "aws_ssm_activation_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmActivationInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsmActivationInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmActivationInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmActivationInvalidDescriptionRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"description must be 256 characters or less",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
