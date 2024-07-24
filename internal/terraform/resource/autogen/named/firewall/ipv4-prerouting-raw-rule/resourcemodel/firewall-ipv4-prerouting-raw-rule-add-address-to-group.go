// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvfourPreroutingRawRuleAddAddressToGroup{}

// FirewallIPvfourPreroutingRawRuleAddAddressToGroup describes the resource data model.
type FirewallIPvfourPreroutingRawRuleAddAddressToGroup struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeFirewallIPvfourPreroutingRawRuleAddAddressToGroupSourceAddress      *FirewallIPvfourPreroutingRawRuleAddAddressToGroupSourceAddress      `tfsdk:"source_address" vyos:"source-address,omitempty"`
	NodeFirewallIPvfourPreroutingRawRuleAddAddressToGroupDestinationAddress *FirewallIPvfourPreroutingRawRuleAddAddressToGroupDestinationAddress `tfsdk:"destination_address" vyos:"destination-address,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourPreroutingRawRuleAddAddressToGroup) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"source_address": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourPreroutingRawRuleAddAddressToGroupSourceAddress{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Add source ip addresses to dynamic address-group

`,
			Description: `Add source ip addresses to dynamic address-group

`,
		},

		"destination_address": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourPreroutingRawRuleAddAddressToGroupDestinationAddress{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Add destination ip addresses to dynamic address-group

`,
			Description: `Add destination ip addresses to dynamic address-group

`,
		},
	}
}
