// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsBudgetsBudgetInvalidNameRule(t *testing.T) {
	cases := []struct {
		Content  string
		Expected helper.Issues
	}{
		{
			Content: `
resource "aws_budgets_budget" "foo" {
	name = budget-ec2-monthly
	}`,
			Expected: helper.Issues{},
		},
		{
			Content: `
resource "aws_budgets_budget" "foo" {
	name = budget:ec2:monthly
	}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsBudgetsBudgetInvalidNameRule(),
					Message: `budget:ec2:monthly does not match valid pattern ^[^:\\]+$`,
				},
			},
		},
	}

	rule := NewAwsBudgetsBudgetInvalidNameRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
