// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsIAMRoleInvalidAssumeRolePolicyRule checks the pattern is valid
type AwsIAMRoleInvalidAssumeRolePolicyRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsIAMRoleInvalidAssumeRolePolicyRule returns new rule with default attributes
func NewAwsIAMRoleInvalidAssumeRolePolicyRule() *AwsIAMRoleInvalidAssumeRolePolicyRule {
	return &AwsIAMRoleInvalidAssumeRolePolicyRule{
		resourceType:  "aws_iam_role",
		attributeName: "assume_role_policy",
		max:           131072,
		min:           1,
		pattern:       regexp.MustCompile(`^[\x{0009}\x{000A}\x{000D}\x{0020}-\x{00FF}]+$`),
	}
}

// Name returns the rule name
func (r *AwsIAMRoleInvalidAssumeRolePolicyRule) Name() string {
	return "aws_iam_role_invalid_assume_role_policy"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIAMRoleInvalidAssumeRolePolicyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIAMRoleInvalidAssumeRolePolicyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsIAMRoleInvalidAssumeRolePolicyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsIAMRoleInvalidAssumeRolePolicyRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"assume_role_policy must be 131072 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"assume_role_policy must be 1 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[\x{0009}\x{000A}\x{000D}\x{0020}-\x{00FF}]+$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
