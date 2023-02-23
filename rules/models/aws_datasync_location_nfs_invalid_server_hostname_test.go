// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsDatasyncLocationNfsInvalidServerHostnameRule(t *testing.T) {
	cases := []struct {
		Content  string
		Expected helper.Issues
	}{
		{
			Content: `
resource "aws_datasync_location_nfs" "foo" {
	server_hostname = nfs.example.com
	}`,
			Expected: helper.Issues{},
		},
		{
			Content: `
resource "aws_datasync_location_nfs" "foo" {
	server_hostname = nfs^example^com
	}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsDatasyncLocationNfsInvalidServerHostnameRule(),
					Message: `nfs^example^com does not match valid pattern ^(([a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9\-]*[A-Za-z0-9])$`,
				},
			},
		},
	}

	rule := NewAwsDatasyncLocationNfsInvalidServerHostnameRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
