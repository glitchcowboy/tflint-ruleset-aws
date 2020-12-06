// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSagemakerNotebookInstanceInvalidLifecycleConfigNameRule checks the pattern is valid
type AwsSagemakerNotebookInstanceInvalidLifecycleConfigNameRule struct {
	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsSagemakerNotebookInstanceInvalidLifecycleConfigNameRule returns new rule with default attributes
func NewAwsSagemakerNotebookInstanceInvalidLifecycleConfigNameRule() *AwsSagemakerNotebookInstanceInvalidLifecycleConfigNameRule {
	return &AwsSagemakerNotebookInstanceInvalidLifecycleConfigNameRule{
		resourceType:  "aws_sagemaker_notebook_instance",
		attributeName: "lifecycle_config_name",
		max:           63,
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9](-*[a-zA-Z0-9])*`),
	}
}

// Name returns the rule name
func (r *AwsSagemakerNotebookInstanceInvalidLifecycleConfigNameRule) Name() string {
	return "aws_sagemaker_notebook_instance_invalid_lifecycle_config_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSagemakerNotebookInstanceInvalidLifecycleConfigNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSagemakerNotebookInstanceInvalidLifecycleConfigNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSagemakerNotebookInstanceInvalidLifecycleConfigNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSagemakerNotebookInstanceInvalidLifecycleConfigNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"lifecycle_config_name must be 63 characters or less",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9](-*[a-zA-Z0-9])*`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
