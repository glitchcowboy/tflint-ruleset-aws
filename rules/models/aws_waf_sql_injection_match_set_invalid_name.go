// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsWafSQLInjectionMatchSetInvalidNameRule checks the pattern is valid
type AwsWafSQLInjectionMatchSetInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsWafSQLInjectionMatchSetInvalidNameRule returns new rule with default attributes
func NewAwsWafSQLInjectionMatchSetInvalidNameRule() *AwsWafSQLInjectionMatchSetInvalidNameRule {
	return &AwsWafSQLInjectionMatchSetInvalidNameRule{
		resourceType:  "aws_waf_sql_injection_match_set",
		attributeName: "name",
		max:           128,
		min:           1,
		pattern:       regexp.MustCompile(`^.*\S.*$`),
	}
}

// Name returns the rule name
func (r *AwsWafSQLInjectionMatchSetInvalidNameRule) Name() string {
	return "aws_waf_sql_injection_match_set_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsWafSQLInjectionMatchSetInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsWafSQLInjectionMatchSetInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsWafSQLInjectionMatchSetInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsWafSQLInjectionMatchSetInvalidNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"name must be 128 characters or less",
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
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^.*\S.*$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
