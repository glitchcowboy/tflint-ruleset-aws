// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsOrganizationsAccountInvalidIAMUserAccessToBillingRule checks the pattern is valid
type AwsOrganizationsAccountInvalidIAMUserAccessToBillingRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsOrganizationsAccountInvalidIAMUserAccessToBillingRule returns new rule with default attributes
func NewAwsOrganizationsAccountInvalidIAMUserAccessToBillingRule() *AwsOrganizationsAccountInvalidIAMUserAccessToBillingRule {
	return &AwsOrganizationsAccountInvalidIAMUserAccessToBillingRule{
		resourceType:  "aws_organizations_account",
		attributeName: "iam_user_access_to_billing",
		enum: []string{
			"ALLOW",
			"DENY",
		},
	}
}

// Name returns the rule name
func (r *AwsOrganizationsAccountInvalidIAMUserAccessToBillingRule) Name() string {
	return "aws_organizations_account_invalid_iam_user_access_to_billing"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsOrganizationsAccountInvalidIAMUserAccessToBillingRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsOrganizationsAccountInvalidIAMUserAccessToBillingRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsOrganizationsAccountInvalidIAMUserAccessToBillingRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsOrganizationsAccountInvalidIAMUserAccessToBillingRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as iam_user_access_to_billing`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
