// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsEcrLifecyclePolicyInvalidRepositoryRule(t *testing.T) {
	cases := []struct {
		Content  string
		Expected helper.Issues
	}{
		{
			Content: `
resource "aws_ecr_lifecycle_policy" "foo" {
	repository = example
	}`,
			Expected: helper.Issues{},
		},
		{
			Content: `
resource "aws_ecr_lifecycle_policy" "foo" {
	repository = example@com
	}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsEcrLifecyclePolicyInvalidRepositoryRule(),
					Message: `example@com does not match valid pattern ^(?:[a-z0-9]+(?:[._-][a-z0-9]+)*/)*[a-z0-9]+(?:[._-][a-z0-9]+)*$`,
				},
			},
		},
	}

	rule := NewAwsEcrLifecyclePolicyInvalidRepositoryRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
