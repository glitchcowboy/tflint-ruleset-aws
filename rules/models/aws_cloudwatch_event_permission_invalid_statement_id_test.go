// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsCloudwatchEventPermissionInvalidStatementIDRule(t *testing.T) {
	cases := []struct {
		Content  string
		Expected helper.Issues
	}{
		{
			Content: `
resource "aws_cloudwatch_event_permission" "foo" {
	statement_id = OrganizationAccess
	}`,
			Expected: helper.Issues{},
		},
		{
			Content: `
resource "aws_cloudwatch_event_permission" "foo" {
	statement_id = Organization Access
	}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsCloudwatchEventPermissionInvalidStatementIDRule(),
					Message: `Organization Access does not match valid pattern ^[a-zA-Z0-9-_]+$`,
				},
			},
		},
	}

	rule := NewAwsCloudwatchEventPermissionInvalidStatementIDRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
