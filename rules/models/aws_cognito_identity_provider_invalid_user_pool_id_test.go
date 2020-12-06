// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsCognitoIdentityProviderInvalidUserPoolIDRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_cognito_identity_provider" "foo" {
	user_pool_id = "foobar"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsCognitoIdentityProviderInvalidUserPoolIDRule(),
					Message: `"foobar" does not match valid pattern ^[\w-]+_[0-9a-zA-Z]+$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_cognito_identity_provider" "foo" {
	user_pool_id = "foo_bar"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAwsCognitoIdentityProviderInvalidUserPoolIDRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
