// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsCodepipelineWebhookInvalidTargetActionRule(t *testing.T) {
	cases := []struct {
		Content  string
		Expected helper.Issues
	}{
		{
			Content: `
resource "aws_codepipeline_webhook" "foo" {
	target_action = Source
	}`,
			Expected: helper.Issues{},
		},
		{
			Content: `
resource "aws_codepipeline_webhook" "foo" {
	target_action = Source/Example
	}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsCodepipelineWebhookInvalidTargetActionRule(),
					Message: `Source/Example does not match valid pattern ^[A-Za-z0-9.@\-_]+$`,
				},
			},
		},
	}

	rule := NewAwsCodepipelineWebhookInvalidTargetActionRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
