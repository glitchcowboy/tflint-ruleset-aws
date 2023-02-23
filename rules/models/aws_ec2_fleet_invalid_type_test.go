// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsEc2FleetInvalidTypeRule(t *testing.T) {
	cases := []struct {
		Content  string
		Expected helper.Issues
	}{
		{
			Content: `
resource "aws_ec2_fleet" "foo" {
	type = maintain
	}`,
			Expected: helper.Issues{},
		},
		{
			Content: `
resource "aws_ec2_fleet" "foo" {
	type = remain
	}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsEc2FleetInvalidTypeRule(),
					Message: `remain is an invalid value as type`,
				},
			},
		},
	}

	rule := NewAwsEc2FleetInvalidTypeRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
