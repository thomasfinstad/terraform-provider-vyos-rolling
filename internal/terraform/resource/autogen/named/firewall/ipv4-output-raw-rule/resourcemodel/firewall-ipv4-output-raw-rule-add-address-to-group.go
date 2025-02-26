/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (add-address-to-group) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvfourOutputRawRuleAddAddressToGroup{}

// FirewallIPvfourOutputRawRuleAddAddressToGroup describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type FirewallIPvfourOutputRawRuleAddAddressToGroup struct {
	// LeafNodes

	// TagNodes

	// Nodes

	NodeFirewallIPvfourOutputRawRuleAddAddressToGroupSourceAddress *FirewallIPvfourOutputRawRuleAddAddressToGroupSourceAddress `tfsdk:"source_address" vyos:"source-address,omitempty"`

	NodeFirewallIPvfourOutputRawRuleAddAddressToGroupDestinationAddress *FirewallIPvfourOutputRawRuleAddAddressToGroupDestinationAddress `tfsdk:"destination_address" vyos:"destination-address,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourOutputRawRuleAddAddressToGroup) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"source_address": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourOutputRawRuleAddAddressToGroupSourceAddress{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Add source ip addresses to dynamic address-group

`,
			Description: `Add source ip addresses to dynamic address-group

`,
		},

		"destination_address": schema.SingleNestedAttribute{
			Attributes: FirewallIPvfourOutputRawRuleAddAddressToGroupDestinationAddress{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Add destination ip addresses to dynamic address-group

`,
			Description: `Add destination ip addresses to dynamic address-group

`,
		},
	}
}
