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

var _ helpers.VyosResourceDataModel = &FirewallIPvfourNameRuleAddAddressToGroupSourceAddress{}

// FirewallIPvfourNameRuleAddAddressToGroupSourceAddress describes the resource data model.
type FirewallIPvfourNameRuleAddAddressToGroupSourceAddress struct {
	// LeafNodes
	LeafFirewallIPvfourNameRuleAddAddressToGroupSourceAddressAddressGroup types.String `tfsdk:"address_group" vyos:"address-group,omitempty"`
	LeafFirewallIPvfourNameRuleAddAddressToGroupSourceAddressTimeout      types.String `tfsdk:"timeout" vyos:"timeout,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourNameRuleAddAddressToGroupSourceAddress) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address_group":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Dynamic address-group

`,
			Description: `Dynamic address-group

`,
		},

		"timeout":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
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

		// Nodes

	}
}
