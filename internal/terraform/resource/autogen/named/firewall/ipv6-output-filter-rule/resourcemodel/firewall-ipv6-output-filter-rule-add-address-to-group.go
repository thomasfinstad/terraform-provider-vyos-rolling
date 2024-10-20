/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvsixOutputFilterRuleAddAddressToGroup{}

// FirewallIPvsixOutputFilterRuleAddAddressToGroup describes the resource data model.
type FirewallIPvsixOutputFilterRuleAddAddressToGroup struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeFirewallIPvsixOutputFilterRuleAddAddressToGroupSourceAddress      *FirewallIPvsixOutputFilterRuleAddAddressToGroupSourceAddress      `tfsdk:"source_address" vyos:"source-address,omitempty"`
	NodeFirewallIPvsixOutputFilterRuleAddAddressToGroupDestinationAddress *FirewallIPvsixOutputFilterRuleAddAddressToGroupDestinationAddress `tfsdk:"destination_address" vyos:"destination-address,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixOutputFilterRuleAddAddressToGroup) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"source_address": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixOutputFilterRuleAddAddressToGroupSourceAddress{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Add source ipv6 addresses to dynamic ipv6-address-group

`,
			Description: `Add source ipv6 addresses to dynamic ipv6-address-group

`,
		},

		"destination_address": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixOutputFilterRuleAddAddressToGroupDestinationAddress{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Add destination ipv6 addresses to dynamic ipv6-address-group

`,
			Description: `Add destination ipv6 addresses to dynamic ipv6-address-group

`,
		},
	}
}
