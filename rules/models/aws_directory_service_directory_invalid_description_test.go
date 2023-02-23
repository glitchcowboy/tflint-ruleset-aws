// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsDirectoryServiceDirectoryInvalidDescriptionRule(t *testing.T) {
	cases := []struct {
		Content  string
		Expected helper.Issues
	}{
		{
			Content: `
resource "aws_directory_service_directory" "foo" {
	description = example
	}`,
			Expected: helper.Issues{},
		},
		{
			Content: `
resource "aws_directory_service_directory" "foo" {
	description = @example
	}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsDirectoryServiceDirectoryInvalidDescriptionRule(),
					Message: `@example does not match valid pattern ^([a-zA-Z0-9_])[\\a-zA-Z0-9_@#%*+=:?./!\s-]*$`,
				},
			},
		},
	}

	rule := NewAwsDirectoryServiceDirectoryInvalidDescriptionRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
