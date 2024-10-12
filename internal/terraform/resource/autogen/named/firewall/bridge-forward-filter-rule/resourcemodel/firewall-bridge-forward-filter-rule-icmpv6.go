// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallBrIDgeForwardFilterRuleIcmpvsix{}

// FirewallBrIDgeForwardFilterRuleIcmpvsix describes the resource data model.
type FirewallBrIDgeForwardFilterRuleIcmpvsix struct {
	// LeafNodes
	LeafFirewallBrIDgeForwardFilterRuleIcmpvsixCode     types.Number `tfsdk:"code" vyos:"code,omitempty"`
	LeafFirewallBrIDgeForwardFilterRuleIcmpvsixType     types.Number `tfsdk:"type" vyos:"type,omitempty"`
	LeafFirewallBrIDgeForwardFilterRuleIcmpvsixTypeName types.String `tfsdk:"type_name" vyos:"type-name,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallBrIDgeForwardFilterRuleIcmpvsix) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"code": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `ICMPv6 code

    |  Format  |  Description          |
    |----------|-----------------------|
    |  0-255   |  ICMPv6 code (0-255)  |
`,
			Description: `ICMPv6 code

    |  Format  |  Description          |
    |----------|-----------------------|
    |  0-255   |  ICMPv6 code (0-255)  |
`,
		},

		"type": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `ICMPv6 type

    |  Format  |  Description          |
    |----------|-----------------------|
    |  0-255   |  ICMPv6 type (0-255)  |
`,
			Description: `ICMPv6 type

    |  Format  |  Description          |
    |----------|-----------------------|
    |  0-255   |  ICMPv6 type (0-255)  |
`,
		},

		"type_name": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `ICMPv6 type-name

    |  Format                   |  Description                              |
    |---------------------------|-------------------------------------------|
    |  destination-unreachable  |  ICMPv6 type 1: destination-unreachable   |
    |  packet-too-big           |  ICMPv6 type 2: packet-too-big            |
    |  time-exceeded            |  ICMPv6 type 3: time-exceeded             |
    |  echo-request             |  ICMPv6 type 128: echo-request            |
    |  echo-reply               |  ICMPv6 type 129: echo-reply              |
    |  mld-listener-query       |  ICMPv6 type 130: mld-listener-query      |
    |  mld-listener-report      |  ICMPv6 type 131: mld-listener-report     |
    |  mld-listener-reduction   |  ICMPv6 type 132: mld-listener-reduction  |
    |  nd-router-solicit        |  ICMPv6 type 133: nd-router-solicit       |
    |  nd-router-advert         |  ICMPv6 type 134: nd-router-advert        |
    |  nd-neighbor-solicit      |  ICMPv6 type 135: nd-neighbor-solicit     |
    |  nd-neighbor-advert       |  ICMPv6 type 136: nd-neighbor-advert      |
    |  nd-redirect              |  ICMPv6 type 137: nd-redirect             |
    |  parameter-problem        |  ICMPv6 type 4: parameter-problem         |
    |  router-renumbering       |  ICMPv6 type 138: router-renumbering      |
    |  ind-neighbor-solicit     |  ICMPv6 type 141: ind-neighbor-solicit    |
    |  ind-neighbor-advert      |  ICMPv6 type 142: ind-neighbor-advert     |
    |  mld2-listener-report     |  ICMPv6 type 143: mld2-listener-report    |
`,
			Description: `ICMPv6 type-name

    |  Format                   |  Description                              |
    |---------------------------|-------------------------------------------|
    |  destination-unreachable  |  ICMPv6 type 1: destination-unreachable   |
    |  packet-too-big           |  ICMPv6 type 2: packet-too-big            |
    |  time-exceeded            |  ICMPv6 type 3: time-exceeded             |
    |  echo-request             |  ICMPv6 type 128: echo-request            |
    |  echo-reply               |  ICMPv6 type 129: echo-reply              |
    |  mld-listener-query       |  ICMPv6 type 130: mld-listener-query      |
    |  mld-listener-report      |  ICMPv6 type 131: mld-listener-report     |
    |  mld-listener-reduction   |  ICMPv6 type 132: mld-listener-reduction  |
    |  nd-router-solicit        |  ICMPv6 type 133: nd-router-solicit       |
    |  nd-router-advert         |  ICMPv6 type 134: nd-router-advert        |
    |  nd-neighbor-solicit      |  ICMPv6 type 135: nd-neighbor-solicit     |
    |  nd-neighbor-advert       |  ICMPv6 type 136: nd-neighbor-advert      |
    |  nd-redirect              |  ICMPv6 type 137: nd-redirect             |
    |  parameter-problem        |  ICMPv6 type 4: parameter-problem         |
    |  router-renumbering       |  ICMPv6 type 138: router-renumbering      |
    |  ind-neighbor-solicit     |  ICMPv6 type 141: ind-neighbor-solicit    |
    |  ind-neighbor-advert      |  ICMPv6 type 142: ind-neighbor-advert     |
    |  mld2-listener-report     |  ICMPv6 type 143: mld2-listener-report    |
`,
		},

		// Nodes

	}
}
