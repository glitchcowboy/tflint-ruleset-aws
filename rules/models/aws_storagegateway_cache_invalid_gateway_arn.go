// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsStoragegatewayCacheInvalidGatewayArnRule checks the pattern is valid
type AwsStoragegatewayCacheInvalidGatewayArnRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsStoragegatewayCacheInvalidGatewayArnRule returns new rule with default attributes
func NewAwsStoragegatewayCacheInvalidGatewayArnRule() *AwsStoragegatewayCacheInvalidGatewayArnRule {
	return &AwsStoragegatewayCacheInvalidGatewayArnRule{
		resourceType:  "aws_storagegateway_cache",
		attributeName: "gateway_arn",
		max:           500,
		min:           50,
	}
}

// Name returns the rule name
func (r *AwsStoragegatewayCacheInvalidGatewayArnRule) Name() string {
	return "aws_storagegateway_cache_invalid_gateway_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsStoragegatewayCacheInvalidGatewayArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsStoragegatewayCacheInvalidGatewayArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsStoragegatewayCacheInvalidGatewayArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsStoragegatewayCacheInvalidGatewayArnRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"gateway_arn must be 500 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"gateway_arn must be 50 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
