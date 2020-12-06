// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCognitoUserGroupInvalidRoleArnRule checks the pattern is valid
type AwsCognitoUserGroupInvalidRoleArnRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCognitoUserGroupInvalidRoleArnRule returns new rule with default attributes
func NewAwsCognitoUserGroupInvalidRoleArnRule() *AwsCognitoUserGroupInvalidRoleArnRule {
	return &AwsCognitoUserGroupInvalidRoleArnRule{
		resourceType:  "aws_cognito_user_group",
		attributeName: "role_arn",
		max:           2048,
		min:           20,
		pattern:       regexp.MustCompile(`^arn:[\w+=/,.@-]+:[\w+=/,.@-]+:([\w+=/,.@-]*)?:[0-9]+:[\w+=/,.@-]+(:[\w+=/,.@-]+)?(:[\w+=/,.@-]+)?$`),
	}
}

// Name returns the rule name
func (r *AwsCognitoUserGroupInvalidRoleArnRule) Name() string {
	return "aws_cognito_user_group_invalid_role_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCognitoUserGroupInvalidRoleArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCognitoUserGroupInvalidRoleArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCognitoUserGroupInvalidRoleArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCognitoUserGroupInvalidRoleArnRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"role_arn must be 2048 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"role_arn must be 20 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^arn:[\w+=/,.@-]+:[\w+=/,.@-]+:([\w+=/,.@-]*)?:[0-9]+:[\w+=/,.@-]+(:[\w+=/,.@-]+)?(:[\w+=/,.@-]+)?$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
