// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsDlmLifecyclePolicyInvalidStateRule(t *testing.T) {
	cases := []struct {
		Content  string
		Expected helper.Issues
	}{
		{
			Content: `
resource "aws_dlm_lifecycle_policy" "foo" {
	state = ENABLED
	}`,
			Expected: helper.Issues{},
		},
		{
			Content: `
resource "aws_dlm_lifecycle_policy" "foo" {
	state = ERROR
	}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsDlmLifecyclePolicyInvalidStateRule(),
					Message: `ERROR is an invalid value as state`,
				},
			},
		},
	}

	rule := NewAwsDlmLifecyclePolicyInvalidStateRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
