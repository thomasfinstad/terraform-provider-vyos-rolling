// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// FirewallIPvfourOutputFilterRuleDestination describes the resource data model.
type FirewallIPvfourOutputFilterRuleDestination struct {
	// LeafNodes
	LeafFirewallIPvfourOutputFilterRuleDestinationAddress     types.String `tfsdk:"address" vyos:"address,omitempty"`
	LeafFirewallIPvfourOutputFilterRuleDestinationAddressMask types.String `tfsdk:"address_mask" vyos:"address-mask,omitempty"`
	LeafFirewallIPvfourOutputFilterRuleDestinationFqdn        types.String `tfsdk:"fqdn" vyos:"fqdn,omitempty"`
	LeafFirewallIPvfourOutputFilterRuleDestinationMacAddress  types.String `tfsdk:"mac_address" vyos:"mac-address,omitempty"`
	LeafFirewallIPvfourOutputFilterRuleDestinationPort        types.String `tfsdk:"port" vyos:"port,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeFirewallIPvfourOutputFilterRuleDestinationGeoIP *FirewallIPvfourOutputFilterRuleDestinationGeoIP `tfsdk:"geoip" vyos:"geoip,omitempty"`
	NodeFirewallIPvfourOutputFilterRuleDestinationGroup *FirewallIPvfourOutputFilterRuleDestinationGroup `tfsdk:"group" vyos:"group,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourOutputFilterRuleDestination) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address, subnet, or range

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IPv4 address to match  |
    |  ipv4net  &emsp; |  IPv4 prefix to match  |
    |  ipv4range  &emsp; |  IPv4 address range to match  |
    |  !ipv4  &emsp; |  Match everything except the specified address  |
    |  !ipv4net  &emsp; |  Match everything except the specified prefix  |
    |  !ipv4range  &emsp; |  Match everything except the specified range  |

`,
		},

		"address_mask": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP mask

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IPv4 mask to apply  |

`,
		},

		"fqdn": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Fully qualified domain name

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  <fqdn>  &emsp; |  Fully qualified domain name  |

`,
		},

		"mac_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `MAC address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  macaddr  &emsp; |  MAC address to match  |
    |  !macaddr  &emsp; |  Match everything except the specified MAC address  |

`,
		},

		"port": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Port

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Named port (any name in /etc/services, e.g., http)  |
    |  number: 1-65535  &emsp; |  Numbered port  |
    |  <start-end>  &emsp; |  Numbered port range (e.g. 1001-1005)  |
    |     &emsp; |  \n\n  Multiple destination ports can be specified as a comma-separated list.\n  For example: 'telnet,http,123,1001-1005'  |

`,
		},

		// Nodes

		"geoip": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourOutputFilterRuleDestinationGeoIP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `GeoIP options - Data provided by DB-IP.com

`,
		},

		"group": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourOutputFilterRuleDestinationGroup{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Group

`,
		},
	}
}
