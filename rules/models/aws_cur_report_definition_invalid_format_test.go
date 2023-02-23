// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsCurReportDefinitionInvalidFormatRule(t *testing.T) {
	cases := []struct {
		Content  string
		Expected helper.Issues
	}{
		{
			Content: `
resource "aws_cur_report_definition" "foo" {
	format = textORcsv
	}`,
			Expected: helper.Issues{},
		},
		{
			Content: `
resource "aws_cur_report_definition" "foo" {
	format = textORjson
	}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsCurReportDefinitionInvalidFormatRule(),
					Message: `textORjson is an invalid value as format`,
				},
			},
		},
	}

	rule := NewAwsCurReportDefinitionInvalidFormatRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
