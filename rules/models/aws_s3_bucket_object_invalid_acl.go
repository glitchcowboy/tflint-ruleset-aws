// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsS3BucketObjectInvalidACLRule checks the pattern is valid
type AwsS3BucketObjectInvalidACLRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsS3BucketObjectInvalidACLRule returns new rule with default attributes
func NewAwsS3BucketObjectInvalidACLRule() *AwsS3BucketObjectInvalidACLRule {
	return &AwsS3BucketObjectInvalidACLRule{
		resourceType:  "aws_s3_bucket_object",
		attributeName: "acl",
		enum: []string{
			"private",
			"public-read",
			"public-read-write",
			"authenticated-read",
			"aws-exec-read",
			"bucket-owner-read",
			"bucket-owner-full-control",
		},
	}
}

// Name returns the rule name
func (r *AwsS3BucketObjectInvalidACLRule) Name() string {
	return "aws_s3_bucket_object_invalid_acl"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsS3BucketObjectInvalidACLRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsS3BucketObjectInvalidACLRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsS3BucketObjectInvalidACLRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsS3BucketObjectInvalidACLRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as acl`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
