// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsEc2CapacityReservationInvalidTenancyRule checks the pattern is valid
type AwsEc2CapacityReservationInvalidTenancyRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsEc2CapacityReservationInvalidTenancyRule returns new rule with default attributes
func NewAwsEc2CapacityReservationInvalidTenancyRule() *AwsEc2CapacityReservationInvalidTenancyRule {
	return &AwsEc2CapacityReservationInvalidTenancyRule{
		resourceType:  "aws_ec2_capacity_reservation",
		attributeName: "tenancy",
		enum: []string{
			"default",
			"dedicated",
		},
	}
}

// Name returns the rule name
func (r *AwsEc2CapacityReservationInvalidTenancyRule) Name() string {
	return "aws_ec2_capacity_reservation_invalid_tenancy"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEc2CapacityReservationInvalidTenancyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEc2CapacityReservationInvalidTenancyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEc2CapacityReservationInvalidTenancyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEc2CapacityReservationInvalidTenancyRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

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
					fmt.Sprintf(`"%s" is an invalid value as tenancy`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
