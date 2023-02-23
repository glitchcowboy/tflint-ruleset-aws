// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsBatchJobDefinitionInvalidTypeRule(t *testing.T) {
	cases := []struct {
		Content  string
		Expected helper.Issues
	}{
		{
			Content: `
resource "aws_batch_job_definition" "foo" {
	type = container
	}`,
			Expected: helper.Issues{},
		},
		{
			Content: `
resource "aws_batch_job_definition" "foo" {
	type = docker
	}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsBatchJobDefinitionInvalidTypeRule(),
					Message: `docker is an invalid value as type`,
				},
			},
		},
	}

	rule := NewAwsBatchJobDefinitionInvalidTypeRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
