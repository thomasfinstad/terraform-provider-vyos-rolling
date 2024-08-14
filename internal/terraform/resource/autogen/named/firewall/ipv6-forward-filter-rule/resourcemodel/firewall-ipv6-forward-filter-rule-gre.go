// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvsixForwardFilterRuleGre{}

// FirewallIPvsixForwardFilterRuleGre describes the resource data model.
type FirewallIPvsixForwardFilterRuleGre struct {
	// LeafNodes
	LeafFirewallIPvsixForwardFilterRuleGreInnerProto types.String `tfsdk:"inner_proto" vyos:"inner-proto,omitempty"`
	LeafFirewallIPvsixForwardFilterRuleGreKey        types.Number `tfsdk:"key" vyos:"key,omitempty"`
	LeafFirewallIPvsixForwardFilterRuleGreVersion    types.String `tfsdk:"version" vyos:"version,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeFirewallIPvsixForwardFilterRuleGreFlags *FirewallIPvsixForwardFilterRuleGreFlags `tfsdk:"flags" vyos:"flags,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixForwardFilterRuleGre) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"inner_proto": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `EtherType of encapsulated packet

    |  Format      |  Description                                                   |
    |--------------|----------------------------------------------------------------|
    |  0-65535     |  Ethernet protocol number                                      |
    |  0x0-0xffff  |  Ethernet protocol number (hex)                                |
    |  ip          |  IPv4                                                          |
    |  ip6         |  IPv6                                                          |
    |  arp         |  Address Resolution Protocol                                   |
    |  802.1q      |  VLAN-tagged frames (IEEE 802.1q)                              |
    |  802.1ad     |  Provider Bridging (IEEE 802.1ad, Q-in-Q)                      |
    |  gretap      |  Transparent Ethernet Bridging (L2 Ethernet over GRE, gretap)  |
`,
			Description: `EtherType of encapsulated packet

    |  Format      |  Description                                                   |
    |--------------|----------------------------------------------------------------|
    |  0-65535     |  Ethernet protocol number                                      |
    |  0x0-0xffff  |  Ethernet protocol number (hex)                                |
    |  ip          |  IPv4                                                          |
    |  ip6         |  IPv6                                                          |
    |  arp         |  Address Resolution Protocol                                   |
    |  802.1q      |  VLAN-tagged frames (IEEE 802.1q)                              |
    |  802.1ad     |  Provider Bridging (IEEE 802.1ad, Q-in-Q)                      |
    |  gretap      |  Transparent Ethernet Bridging (L2 Ethernet over GRE, gretap)  |
`,
		},

		"key": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Tunnel key (only GRE tunnels)

    |  Format  |  Description  |
    |----------|---------------|
    |  u32     |  Tunnel key   |
`,
			Description: `Tunnel key (only GRE tunnels)

    |  Format  |  Description  |
    |----------|---------------|
    |  u32     |  Tunnel key   |
`,
		},

		"version": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `GRE Version

    |  Format  |  Description                         |
    |----------|--------------------------------------|
    |  gre     |  Standard GRE                        |
    |  pptp    |  Point to Point Tunnelling Protocol  |
`,
			Description: `GRE Version

    |  Format  |  Description                         |
    |----------|--------------------------------------|
    |  gre     |  Standard GRE                        |
    |  pptp    |  Point to Point Tunnelling Protocol  |
`,
		},

		// Nodes

		"flags": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixForwardFilterRuleGreFlags{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `GRE flag bits to match

`,
			Description: `GRE flag bits to match

`,
		},
	}
}
