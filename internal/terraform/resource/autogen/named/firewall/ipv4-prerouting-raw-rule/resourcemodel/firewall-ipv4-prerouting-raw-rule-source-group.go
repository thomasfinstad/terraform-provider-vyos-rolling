/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (group) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvfourPreroutingRawRuleSourceGroup{}

// FirewallIPvfourPreroutingRawRuleSourceGroup describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type FirewallIPvfourPreroutingRawRuleSourceGroup struct {
	// LeafNodes
	LeafFirewallIPvfourPreroutingRawRuleSourceGroupAddressGroup types.String `tfsdk:"address_group" vyos:"address-group,omitempty"`
	LeafFirewallIPvfourPreroutingRawRuleSourceGroupDomainGroup  types.String `tfsdk:"domain_group" vyos:"domain-group,omitempty"`
	LeafFirewallIPvfourPreroutingRawRuleSourceGroupMacGroup     types.String `tfsdk:"mac_group" vyos:"mac-group,omitempty"`
	LeafFirewallIPvfourPreroutingRawRuleSourceGroupNetworkGroup types.String `tfsdk:"network_group" vyos:"network-group,omitempty"`
	LeafFirewallIPvfourPreroutingRawRuleSourceGroupPortGroup    types.String `tfsdk:"port_group" vyos:"port-group,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourPreroutingRawRuleSourceGroup) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address_group":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (address-group) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of addresses

`,
			Description: `Group of addresses

`,
		},

		"domain_group":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (domain-group) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of domains

`,
			Description: `Group of domains

`,
		},

		"mac_group":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (mac-group) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of MAC addresses

`,
			Description: `Group of MAC addresses

`,
		},

		"network_group":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (network-group) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of networks

`,
			Description: `Group of networks

`,
		},

		"port_group":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (port-group) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Group of ports

`,
			Description: `Group of ports

`,
		},

		// TagNodes

		// Nodes

	}
}
