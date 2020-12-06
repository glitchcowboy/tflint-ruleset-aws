// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsRoute53ResolverRuleInvalidDomainNameRule checks the pattern is valid
type AwsRoute53ResolverRuleInvalidDomainNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsRoute53ResolverRuleInvalidDomainNameRule returns new rule with default attributes
func NewAwsRoute53ResolverRuleInvalidDomainNameRule() *AwsRoute53ResolverRuleInvalidDomainNameRule {
	return &AwsRoute53ResolverRuleInvalidDomainNameRule{
		resourceType:  "aws_route53_resolver_rule",
		attributeName: "domain_name",
		max:           256,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsRoute53ResolverRuleInvalidDomainNameRule) Name() string {
	return "aws_route53_resolver_rule_invalid_domain_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRoute53ResolverRuleInvalidDomainNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRoute53ResolverRuleInvalidDomainNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRoute53ResolverRuleInvalidDomainNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRoute53ResolverRuleInvalidDomainNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"domain_name must be 256 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"domain_name must be 1 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
