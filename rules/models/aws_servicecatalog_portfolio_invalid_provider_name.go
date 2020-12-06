// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsServicecatalogPortfolioInvalidProviderNameRule checks the pattern is valid
type AwsServicecatalogPortfolioInvalidProviderNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsServicecatalogPortfolioInvalidProviderNameRule returns new rule with default attributes
func NewAwsServicecatalogPortfolioInvalidProviderNameRule() *AwsServicecatalogPortfolioInvalidProviderNameRule {
	return &AwsServicecatalogPortfolioInvalidProviderNameRule{
		resourceType:  "aws_servicecatalog_portfolio",
		attributeName: "provider_name",
		max:           50,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsServicecatalogPortfolioInvalidProviderNameRule) Name() string {
	return "aws_servicecatalog_portfolio_invalid_provider_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsServicecatalogPortfolioInvalidProviderNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsServicecatalogPortfolioInvalidProviderNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsServicecatalogPortfolioInvalidProviderNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsServicecatalogPortfolioInvalidProviderNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"provider_name must be 50 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"provider_name must be 1 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
