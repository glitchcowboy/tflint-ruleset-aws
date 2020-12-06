// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsElasticBeanstalkEnvironmentInvalidApplicationRule checks the pattern is valid
type AwsElasticBeanstalkEnvironmentInvalidApplicationRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsElasticBeanstalkEnvironmentInvalidApplicationRule returns new rule with default attributes
func NewAwsElasticBeanstalkEnvironmentInvalidApplicationRule() *AwsElasticBeanstalkEnvironmentInvalidApplicationRule {
	return &AwsElasticBeanstalkEnvironmentInvalidApplicationRule{
		resourceType:  "aws_elastic_beanstalk_environment",
		attributeName: "application",
		max:           100,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsElasticBeanstalkEnvironmentInvalidApplicationRule) Name() string {
	return "aws_elastic_beanstalk_environment_invalid_application"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsElasticBeanstalkEnvironmentInvalidApplicationRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsElasticBeanstalkEnvironmentInvalidApplicationRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsElasticBeanstalkEnvironmentInvalidApplicationRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsElasticBeanstalkEnvironmentInvalidApplicationRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"application must be 100 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"application must be 1 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
