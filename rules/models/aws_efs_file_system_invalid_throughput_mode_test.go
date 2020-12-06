// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsEfsFileSystemInvalidThroughputModeRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_efs_file_system" "foo" {
	throughput_mode = "generalPurpose"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsEfsFileSystemInvalidThroughputModeRule(),
					Message: `"generalPurpose" is an invalid value as throughput_mode`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_efs_file_system" "foo" {
	throughput_mode = "bursting"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAwsEfsFileSystemInvalidThroughputModeRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
