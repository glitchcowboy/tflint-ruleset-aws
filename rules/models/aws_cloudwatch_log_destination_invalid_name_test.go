// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsCloudwatchLogDestinationInvalidNameRule(t *testing.T) {
	cases := []struct {
		Content  string
		Expected helper.Issues
	}{
		{
			Content: `
resource "aws_cloudwatch_log_destination" "foo" {
	name = test_destination
	}`,
			Expected: helper.Issues{},
		},
		{
			Content: `
resource "aws_cloudwatch_log_destination" "foo" {
	name = test:destination
	}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsCloudwatchLogDestinationInvalidNameRule(),
					Message: `test:destination does not match valid pattern ^[^:*]*$`,
				},
			},
		},
	}

	rule := NewAwsCloudwatchLogDestinationInvalidNameRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
