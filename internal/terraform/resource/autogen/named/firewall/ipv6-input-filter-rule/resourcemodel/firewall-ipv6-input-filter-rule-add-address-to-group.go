/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvsixInputFilterRuleAddAddressToGroup{}

// FirewallIPvsixInputFilterRuleAddAddressToGroup describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type FirewallIPvsixInputFilterRuleAddAddressToGroup struct {
	// LeafNodes

	// TagNodes

	// Nodes

	NodeFirewallIPvsixInputFilterRuleAddAddressToGroupSourceAddress *FirewallIPvsixInputFilterRuleAddAddressToGroupSourceAddress `tfsdk:"source_address" vyos:"source-address,omitempty"`

	NodeFirewallIPvsixInputFilterRuleAddAddressToGroupDestinationAddress *FirewallIPvsixInputFilterRuleAddAddressToGroupDestinationAddress `tfsdk:"destination_address" vyos:"destination-address,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixInputFilterRuleAddAddressToGroup) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"source_address": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixInputFilterRuleAddAddressToGroupSourceAddress{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Add source ipv6 addresses to dynamic ipv6-address-group

`,
			Description: `Add source ipv6 addresses to dynamic ipv6-address-group

`,
		},

		"destination_address": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixInputFilterRuleAddAddressToGroupDestinationAddress{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Add destination ipv6 addresses to dynamic ipv6-address-group

`,
			Description: `Add destination ipv6 addresses to dynamic ipv6-address-group

`,
		},
	}
}
