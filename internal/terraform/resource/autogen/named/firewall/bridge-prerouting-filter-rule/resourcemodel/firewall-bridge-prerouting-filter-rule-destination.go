// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallBrIDgePreroutingFilterRuleDestination{}

// FirewallBrIDgePreroutingFilterRuleDestination describes the resource data model.
type FirewallBrIDgePreroutingFilterRuleDestination struct {
	// LeafNodes
	LeafFirewallBrIDgePreroutingFilterRuleDestinationMacAddress  types.String `tfsdk:"mac_address" vyos:"mac-address,omitempty"`
	LeafFirewallBrIDgePreroutingFilterRuleDestinationAddress     types.String `tfsdk:"address" vyos:"address,omitempty"`
	LeafFirewallBrIDgePreroutingFilterRuleDestinationAddressMask types.String `tfsdk:"address_mask" vyos:"address-mask,omitempty"`
	LeafFirewallBrIDgePreroutingFilterRuleDestinationPort        types.String `tfsdk:"port" vyos:"port,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeFirewallBrIDgePreroutingFilterRuleDestinationGroup *FirewallBrIDgePreroutingFilterRuleDestinationGroup `tfsdk:"group" vyos:"group,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallBrIDgePreroutingFilterRuleDestination) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"mac_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `MAC address

    |  Format    |  Description                                        |
    |------------|-----------------------------------------------------|
    |  macaddr   |  MAC address to match                               |
    |  !macaddr  |  Match everything except the specified MAC address  |
`,
			Description: `MAC address

    |  Format    |  Description                                        |
    |------------|-----------------------------------------------------|
    |  macaddr   |  MAC address to match                               |
    |  !macaddr  |  Match everything except the specified MAC address  |
`,
		},

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address, subnet, or range

    |  Format      |  Description                                    |
    |--------------|-------------------------------------------------|
    |  ipv4        |  IPv4 address to match                          |
    |  ipv4net     |  IPv4 prefix to match                           |
    |  ipv4range   |  IPv4 address range to match                    |
    |  !ipv4       |  Match everything except the specified address  |
    |  !ipv4net    |  Match everything except the specified prefix   |
    |  !ipv4range  |  Match everything except the specified range    |
    |  ipv6net     |  Subnet to match                                |
    |  ipv6range   |  IP range to match                              |
    |  !ipv6       |  Match everything except the specified address  |
    |  !ipv6net    |  Match everything except the specified prefix   |
    |  !ipv6range  |  Match everything except the specified range    |
`,
			Description: `IP address, subnet, or range

    |  Format      |  Description                                    |
    |--------------|-------------------------------------------------|
    |  ipv4        |  IPv4 address to match                          |
    |  ipv4net     |  IPv4 prefix to match                           |
    |  ipv4range   |  IPv4 address range to match                    |
    |  !ipv4       |  Match everything except the specified address  |
    |  !ipv4net    |  Match everything except the specified prefix   |
    |  !ipv4range  |  Match everything except the specified range    |
    |  ipv6net     |  Subnet to match                                |
    |  ipv6range   |  IP range to match                              |
    |  !ipv6       |  Match everything except the specified address  |
    |  !ipv6net    |  Match everything except the specified prefix   |
    |  !ipv6range  |  Match everything except the specified range    |
`,
		},

		"address_mask": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP mask

    |  Format  |  Description         |
    |----------|----------------------|
    |  ipv4    |  IPv4 mask to apply  |
    |  ipv6    |  IP mask to apply    |
`,
			Description: `IP mask

    |  Format  |  Description         |
    |----------|----------------------|
    |  ipv4    |  IPv4 mask to apply  |
    |  ipv6    |  IP mask to apply    |
`,
		},

		"port": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Port

    |  Format       |  Description                                                                                                                        |
    |---------------|-------------------------------------------------------------------------------------------------------------------------------------|
    |  txt          |  Named port (any name in /etc/services, e.g., http)                                                                                 |
    |  1-65535      |  Numbered port                                                                                                                      |
    |  <start-end>  |  Numbered port range (e.g. 1001-1005)                                                                                               |
    |               |  </br></br>  Multiple destination ports can be specified as a comma-separated list.</br>  For example: 'telnet,http,123,1001-1005'  |
`,
			Description: `Port

    |  Format       |  Description                                                                                                                        |
    |---------------|-------------------------------------------------------------------------------------------------------------------------------------|
    |  txt          |  Named port (any name in /etc/services, e.g., http)                                                                                 |
    |  1-65535      |  Numbered port                                                                                                                      |
    |  <start-end>  |  Numbered port range (e.g. 1001-1005)                                                                                               |
    |               |  </br></br>  Multiple destination ports can be specified as a comma-separated list.</br>  For example: 'telnet,http,123,1001-1005'  |
`,
		},

		// Nodes

		"group": schema.SingleNestedAttribute{
			Attributes: FirewallBrIDgePreroutingFilterRuleDestinationGroup{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Group

`,
			Description: `Group

`,
		},
	}
}
