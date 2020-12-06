// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAppsyncResolverInvalidResponseTemplateRule checks the pattern is valid
type AwsAppsyncResolverInvalidResponseTemplateRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsAppsyncResolverInvalidResponseTemplateRule returns new rule with default attributes
func NewAwsAppsyncResolverInvalidResponseTemplateRule() *AwsAppsyncResolverInvalidResponseTemplateRule {
	return &AwsAppsyncResolverInvalidResponseTemplateRule{
		resourceType:  "aws_appsync_resolver",
		attributeName: "response_template",
		max:           65536,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsAppsyncResolverInvalidResponseTemplateRule) Name() string {
	return "aws_appsync_resolver_invalid_response_template"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppsyncResolverInvalidResponseTemplateRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAppsyncResolverInvalidResponseTemplateRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAppsyncResolverInvalidResponseTemplateRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppsyncResolverInvalidResponseTemplateRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"response_template must be 65536 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"response_template must be 1 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
