/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallBrIDgePreroutingFilterRuleDestinationGroup{}

// FirewallBrIDgePreroutingFilterRuleDestinationGroup describes the resource data model.
type FirewallBrIDgePreroutingFilterRuleDestinationGroup struct {
	// LeafNodes
	LeafFirewallBrIDgePreroutingFilterRuleDestinationGroupIPvfourAddressGroup types.String `tfsdk:"ipv4_address_group" vyos:"ipv4-address-group,omitempty"`
	LeafFirewallBrIDgePreroutingFilterRuleDestinationGroupIPvsixAddressGroup  types.String `tfsdk:"ipv6_address_group" vyos:"ipv6-address-group,omitempty"`
	LeafFirewallBrIDgePreroutingFilterRuleDestinationGroupMacGroup            types.String `tfsdk:"mac_group" vyos:"mac-group,omitempty"`
	LeafFirewallBrIDgePreroutingFilterRuleDestinationGroupIPvfourNetworkGroup types.String `tfsdk:"ipv4_network_group" vyos:"ipv4-network-group,omitempty"`
	LeafFirewallBrIDgePreroutingFilterRuleDestinationGroupIPvsixNetworkGroup  types.String `tfsdk:"ipv6_network_group" vyos:"ipv6-network-group,omitempty"`
	LeafFirewallBrIDgePreroutingFilterRuleDestinationGroupPortGroup           types.String `tfsdk:"port_group" vyos:"port-group,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallBrIDgePreroutingFilterRuleDestinationGroup) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"ipv4_address_group":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of IPv4 addresses

`,
			Description: `Group of IPv4 addresses

`,
		},

		"ipv6_address_group":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of IPv6 addresses

`,
			Description: `Group of IPv6 addresses

`,
		},

		"mac_group":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of MAC addresses

`,
			Description: `Group of MAC addresses

`,
		},

		"ipv4_network_group":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of IPv4 networks

`,
			Description: `Group of IPv4 networks

`,
		},

		"ipv6_network_group":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of IPv6 networks

`,
			Description: `Group of IPv6 networks

`,
		},

		"port_group":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of ports

`,
			Description: `Group of ports

`,
		},

		// Nodes

	}
}
