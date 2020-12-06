// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsLambdaFunctionInvalidKmsKeyArnRule checks the pattern is valid
type AwsLambdaFunctionInvalidKmsKeyArnRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsLambdaFunctionInvalidKmsKeyArnRule returns new rule with default attributes
func NewAwsLambdaFunctionInvalidKmsKeyArnRule() *AwsLambdaFunctionInvalidKmsKeyArnRule {
	return &AwsLambdaFunctionInvalidKmsKeyArnRule{
		resourceType:  "aws_lambda_function",
		attributeName: "kms_key_arn",
		pattern:       regexp.MustCompile(`^(arn:(aws[a-zA-Z-]*)?:[a-z0-9-.]+:.*)|()$`),
	}
}

// Name returns the rule name
func (r *AwsLambdaFunctionInvalidKmsKeyArnRule) Name() string {
	return "aws_lambda_function_invalid_kms_key_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLambdaFunctionInvalidKmsKeyArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLambdaFunctionInvalidKmsKeyArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLambdaFunctionInvalidKmsKeyArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLambdaFunctionInvalidKmsKeyArnRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^(arn:(aws[a-zA-Z-]*)?:[a-z0-9-.]+:.*)|()$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
