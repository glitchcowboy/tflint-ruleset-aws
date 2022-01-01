// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsShieldProtectionGroupInvalidAggregationRule checks the pattern is valid
type AwsShieldProtectionGroupInvalidAggregationRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsShieldProtectionGroupInvalidAggregationRule returns new rule with default attributes
func NewAwsShieldProtectionGroupInvalidAggregationRule() *AwsShieldProtectionGroupInvalidAggregationRule {
	return &AwsShieldProtectionGroupInvalidAggregationRule{
		resourceType:  "aws_shield_protection_group",
		attributeName: "aggregation",
		enum: []string{
			"SUM",
			"MEAN",
			"MAX",
		},
	}
}

// Name returns the rule name
func (r *AwsShieldProtectionGroupInvalidAggregationRule) Name() string {
	return "aws_shield_protection_group_invalid_aggregation"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsShieldProtectionGroupInvalidAggregationRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsShieldProtectionGroupInvalidAggregationRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsShieldProtectionGroupInvalidAggregationRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsShieldProtectionGroupInvalidAggregationRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is an invalid value as aggregation`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
