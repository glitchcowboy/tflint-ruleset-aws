// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsServicequotasServiceQuotaInvalidServiceCodeRule checks the pattern is valid
type AwsServicequotasServiceQuotaInvalidServiceCodeRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsServicequotasServiceQuotaInvalidServiceCodeRule returns new rule with default attributes
func NewAwsServicequotasServiceQuotaInvalidServiceCodeRule() *AwsServicequotasServiceQuotaInvalidServiceCodeRule {
	return &AwsServicequotasServiceQuotaInvalidServiceCodeRule{
		resourceType:  "aws_servicequotas_service_quota",
		attributeName: "service_code",
		max:           63,
		min:           1,
		pattern:       regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9-]{1,63}$`),
	}
}

// Name returns the rule name
func (r *AwsServicequotasServiceQuotaInvalidServiceCodeRule) Name() string {
	return "aws_servicequotas_service_quota_invalid_service_code"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsServicequotasServiceQuotaInvalidServiceCodeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsServicequotasServiceQuotaInvalidServiceCodeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsServicequotasServiceQuotaInvalidServiceCodeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsServicequotasServiceQuotaInvalidServiceCodeRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"service_code must be 63 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"service_code must be 1 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z][a-zA-Z0-9-]{1,63}$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
