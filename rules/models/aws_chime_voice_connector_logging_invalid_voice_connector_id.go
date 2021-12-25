// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsChimeVoiceConnectorLoggingInvalidVoiceConnectorIDRule checks the pattern is valid
type AwsChimeVoiceConnectorLoggingInvalidVoiceConnectorIDRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsChimeVoiceConnectorLoggingInvalidVoiceConnectorIDRule returns new rule with default attributes
func NewAwsChimeVoiceConnectorLoggingInvalidVoiceConnectorIDRule() *AwsChimeVoiceConnectorLoggingInvalidVoiceConnectorIDRule {
	return &AwsChimeVoiceConnectorLoggingInvalidVoiceConnectorIDRule{
		resourceType:  "aws_chime_voice_connector_logging",
		attributeName: "voice_connector_id",
		pattern:       regexp.MustCompile(`^.*\S.*$`),
	}
}

// Name returns the rule name
func (r *AwsChimeVoiceConnectorLoggingInvalidVoiceConnectorIDRule) Name() string {
	return "aws_chime_voice_connector_logging_invalid_voice_connector_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsChimeVoiceConnectorLoggingInvalidVoiceConnectorIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsChimeVoiceConnectorLoggingInvalidVoiceConnectorIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsChimeVoiceConnectorLoggingInvalidVoiceConnectorIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsChimeVoiceConnectorLoggingInvalidVoiceConnectorIDRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^.*\S.*$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
