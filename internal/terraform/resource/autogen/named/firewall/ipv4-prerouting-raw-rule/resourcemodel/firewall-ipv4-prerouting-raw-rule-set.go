// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvfourPreroutingRawRuleSet{}

// FirewallIPvfourPreroutingRawRuleSet describes the resource data model.
type FirewallIPvfourPreroutingRawRuleSet struct {
	// LeafNodes
	LeafFirewallIPvfourPreroutingRawRuleSetDscp   types.Number `tfsdk:"dscp" vyos:"dscp,omitempty"`
	LeafFirewallIPvfourPreroutingRawRuleSetMark   types.Number `tfsdk:"mark" vyos:"mark,omitempty"`
	LeafFirewallIPvfourPreroutingRawRuleSetTCPMss types.Number `tfsdk:"tcp_mss" vyos:"tcp-mss,omitempty"`
	LeafFirewallIPvfourPreroutingRawRuleSetTTL    types.Number `tfsdk:"ttl" vyos:"ttl,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourPreroutingRawRuleSet) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"dscp": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set DSCP (Packet Differentiated Services Codepoint) bits

    |  Format  |  Description  |
    |----------|---------------|
    |  0-63    |  DSCP number  |
`,
			Description: `Set DSCP (Packet Differentiated Services Codepoint) bits

    |  Format  |  Description  |
    |----------|---------------|
    |  0-63    |  DSCP number  |
`,
		},

		"mark": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set packet mark

    |  Format        |  Description  |
    |----------------|---------------|
    |  1-2147483647  |  Packet mark  |
`,
			Description: `Set packet mark

    |  Format        |  Description  |
    |----------------|---------------|
    |  1-2147483647  |  Packet mark  |
`,
		},

		"tcp_mss": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set TCP Maximum Segment Size

    |  Format    |  Description                   |
    |------------|--------------------------------|
    |  500-1460  |  Explicitly set TCP MSS value  |
`,
			Description: `Set TCP Maximum Segment Size

    |  Format    |  Description                   |
    |------------|--------------------------------|
    |  500-1460  |  Explicitly set TCP MSS value  |
`,
		},

		"ttl": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set TTL (time to live)

    |  Format  |  Description  |
    |----------|---------------|
    |  0-255   |  TTL number   |
`,
			Description: `Set TTL (time to live)

    |  Format  |  Description  |
    |----------|---------------|
    |  0-255   |  TTL number   |
`,
		},

		// Nodes

	}
}