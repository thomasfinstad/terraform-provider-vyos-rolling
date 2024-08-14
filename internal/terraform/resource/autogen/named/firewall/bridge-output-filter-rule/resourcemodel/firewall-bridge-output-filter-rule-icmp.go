// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallBrIDgeOutputFilterRuleIcmp{}

// FirewallBrIDgeOutputFilterRuleIcmp describes the resource data model.
type FirewallBrIDgeOutputFilterRuleIcmp struct {
	// LeafNodes
	LeafFirewallBrIDgeOutputFilterRuleIcmpCode     types.Number `tfsdk:"code" vyos:"code,omitempty"`
	LeafFirewallBrIDgeOutputFilterRuleIcmpType     types.Number `tfsdk:"type" vyos:"type,omitempty"`
	LeafFirewallBrIDgeOutputFilterRuleIcmpTypeName types.String `tfsdk:"type_name" vyos:"type-name,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallBrIDgeOutputFilterRuleIcmp) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"code": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `ICMP code

    |  Format  |  Description        |
    |----------|---------------------|
    |  0-255   |  ICMP code (0-255)  |
`,
			Description: `ICMP code

    |  Format  |  Description        |
    |----------|---------------------|
    |  0-255   |  ICMP code (0-255)  |
`,
		},

		"type": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `ICMP type

    |  Format  |  Description        |
    |----------|---------------------|
    |  0-255   |  ICMP type (0-255)  |
`,
			Description: `ICMP type

    |  Format  |  Description        |
    |----------|---------------------|
    |  0-255   |  ICMP type (0-255)  |
`,
		},

		"type_name": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `ICMP type-name

    |  Format                   |  Description                           |
    |---------------------------|----------------------------------------|
    |  echo-reply               |  ICMP type 0: echo-reply               |
    |  destination-unreachable  |  ICMP type 3: destination-unreachable  |
    |  source-quench            |  ICMP type 4: source-quench            |
    |  redirect                 |  ICMP type 5: redirect                 |
    |  echo-request             |  ICMP type 8: echo-request             |
    |  router-advertisement     |  ICMP type 9: router-advertisement     |
    |  router-solicitation      |  ICMP type 10: router-solicitation     |
    |  time-exceeded            |  ICMP type 11: time-exceeded           |
    |  parameter-problem        |  ICMP type 12: parameter-problem       |
    |  timestamp-request        |  ICMP type 13: timestamp-request       |
    |  timestamp-reply          |  ICMP type 14: timestamp-reply         |
    |  info-request             |  ICMP type 15: info-request            |
    |  info-reply               |  ICMP type 16: info-reply              |
    |  address-mask-request     |  ICMP type 17: address-mask-request    |
    |  address-mask-reply       |  ICMP type 18: address-mask-reply      |
`,
			Description: `ICMP type-name

    |  Format                   |  Description                           |
    |---------------------------|----------------------------------------|
    |  echo-reply               |  ICMP type 0: echo-reply               |
    |  destination-unreachable  |  ICMP type 3: destination-unreachable  |
    |  source-quench            |  ICMP type 4: source-quench            |
    |  redirect                 |  ICMP type 5: redirect                 |
    |  echo-request             |  ICMP type 8: echo-request             |
    |  router-advertisement     |  ICMP type 9: router-advertisement     |
    |  router-solicitation      |  ICMP type 10: router-solicitation     |
    |  time-exceeded            |  ICMP type 11: time-exceeded           |
    |  parameter-problem        |  ICMP type 12: parameter-problem       |
    |  timestamp-request        |  ICMP type 13: timestamp-request       |
    |  timestamp-reply          |  ICMP type 14: timestamp-reply         |
    |  info-request             |  ICMP type 15: info-request            |
    |  info-reply               |  ICMP type 16: info-reply              |
    |  address-mask-request     |  ICMP type 17: address-mask-request    |
    |  address-mask-reply       |  ICMP type 18: address-mask-reply      |
`,
		},

		// Nodes

	}
}