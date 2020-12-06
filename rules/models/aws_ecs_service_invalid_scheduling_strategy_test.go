// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsEcsServiceInvalidSchedulingStrategyRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_ecs_service" "foo" {
	scheduling_strategy = "SERVER"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsEcsServiceInvalidSchedulingStrategyRule(),
					Message: `"SERVER" is an invalid value as scheduling_strategy`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_ecs_service" "foo" {
	scheduling_strategy = "REPLICA"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAwsEcsServiceInvalidSchedulingStrategyRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
