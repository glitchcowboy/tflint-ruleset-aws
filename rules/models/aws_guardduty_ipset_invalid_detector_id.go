// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsGuarddutyIpsetInvalidDetectorIDRule checks the pattern is valid
type AwsGuarddutyIpsetInvalidDetectorIDRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsGuarddutyIpsetInvalidDetectorIDRule returns new rule with default attributes
func NewAwsGuarddutyIpsetInvalidDetectorIDRule() *AwsGuarddutyIpsetInvalidDetectorIDRule {
	return &AwsGuarddutyIpsetInvalidDetectorIDRule{
		resourceType:  "aws_guardduty_ipset",
		attributeName: "detector_id",
		max:           300,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsGuarddutyIpsetInvalidDetectorIDRule) Name() string {
	return "aws_guardduty_ipset_invalid_detector_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsGuarddutyIpsetInvalidDetectorIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsGuarddutyIpsetInvalidDetectorIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsGuarddutyIpsetInvalidDetectorIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsGuarddutyIpsetInvalidDetectorIDRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"detector_id must be 300 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"detector_id must be 1 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
