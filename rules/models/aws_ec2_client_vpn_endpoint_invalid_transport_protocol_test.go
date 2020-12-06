// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsEc2ClientVpnEndpointInvalidTransportProtocolRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_ec2_client_vpn_endpoint" "foo" {
	transport_protocol = "http"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsEc2ClientVpnEndpointInvalidTransportProtocolRule(),
					Message: `"http" is an invalid value as transport_protocol`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_ec2_client_vpn_endpoint" "foo" {
	transport_protocol = "udp"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAwsEc2ClientVpnEndpointInvalidTransportProtocolRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
