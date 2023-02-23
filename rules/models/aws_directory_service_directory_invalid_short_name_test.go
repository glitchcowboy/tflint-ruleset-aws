// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsDirectoryServiceDirectoryInvalidShortNameRule(t *testing.T) {
	cases := []struct {
		Content  string
		Expected helper.Issues
	}{
		{
			Content: `
resource "aws_directory_service_directory" "foo" {
	short_name = CORP
	}`,
			Expected: helper.Issues{},
		},
		{
			Content: `
resource "aws_directory_service_directory" "foo" {
	short_name = CORP:EXAMPLE
	}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsDirectoryServiceDirectoryInvalidShortNameRule(),
					Message: `CORP:EXAMPLE does not match valid pattern ^[^\\/:*?"<>|.]+[^\\/:*?"<>|]*$`,
				},
			},
		},
	}

	rule := NewAwsDirectoryServiceDirectoryInvalidShortNameRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
