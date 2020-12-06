// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsCognitoUserPoolInvalidEmailVerificationMessageRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_cognito_user_pool" "foo" {
	email_verification_message = "Verification code"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsCognitoUserPoolInvalidEmailVerificationMessageRule(),
					Message: `"Verification code" does not match valid pattern ^[\p{L}\p{M}\p{S}\p{N}\p{P}\s*]*\{####\}[\p{L}\p{M}\p{S}\p{N}\p{P}\s*]*$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_cognito_user_pool" "foo" {
	email_verification_message = "Verification code is {####}"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAwsCognitoUserPoolInvalidEmailVerificationMessageRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
