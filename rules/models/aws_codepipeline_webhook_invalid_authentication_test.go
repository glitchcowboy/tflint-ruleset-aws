// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsCodepipelineWebhookInvalidAuthenticationRule(t *testing.T) {
	cases := []struct {
		Content  string
		Expected helper.Issues
	}{
		{
			Content: `
resource "aws_codepipeline_webhook" "foo" {
	authentication = GITHUB_HMAC
	}`,
			Expected: helper.Issues{},
		},
		{
			Content: `
resource "aws_codepipeline_webhook" "foo" {
	authentication = GITLAB_HMAC
	}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsCodepipelineWebhookInvalidAuthenticationRule(),
					Message: `GITLAB_HMAC is an invalid value as authentication`,
				},
			},
		},
	}

	rule := NewAwsCodepipelineWebhookInvalidAuthenticationRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
