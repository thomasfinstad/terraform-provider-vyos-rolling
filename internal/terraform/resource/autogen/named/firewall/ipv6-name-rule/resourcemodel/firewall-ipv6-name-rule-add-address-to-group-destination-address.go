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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvsixNameRuleAddAddressToGroupDestinationAddress{}

// FirewallIPvsixNameRuleAddAddressToGroupDestinationAddress describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type FirewallIPvsixNameRuleAddAddressToGroupDestinationAddress struct {
	// LeafNodes
	LeafFirewallIPvsixNameRuleAddAddressToGroupDestinationAddressAddressGroup types.String `tfsdk:"address_group" vyos:"address-group,omitempty"`
	LeafFirewallIPvsixNameRuleAddAddressToGroupDestinationAddressTimeout      types.String `tfsdk:"timeout" vyos:"timeout,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixNameRuleAddAddressToGroupDestinationAddress) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address_group":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Dynamic ipv6-address-group

`,
			Description: `Dynamic ipv6-address-group

`,
		},

		"timeout":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set timeout

    |  Format     |  Description               |
    |-------------|----------------------------|
    |  <number>s  |  Timeout value in seconds  |
    |  <number>m  |  Timeout value in minutes  |
    |  <number>h  |  Timeout value in hours    |
    |  <number>d  |  Timeout value in days     |
`,
			Description: `Set timeout

    |  Format     |  Description               |
    |-------------|----------------------------|
    |  <number>s  |  Timeout value in seconds  |
    |  <number>m  |  Timeout value in minutes  |
    |  <number>h  |  Timeout value in hours    |
    |  <number>d  |  Timeout value in days     |
`,
		},

		// TagNodes

		// Nodes

	}
}
