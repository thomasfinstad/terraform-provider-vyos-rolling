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

var _ helpers.VyosResourceDataModel = &FirewallIPvsixNameRuleTCP{}

// FirewallIPvsixNameRuleTCP describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type FirewallIPvsixNameRuleTCP struct {
	// LeafNodes
	LeafFirewallIPvsixNameRuleTCPMss types.String `tfsdk:"mss" vyos:"mss,omitempty"`

	// TagNodes

	// Nodes

	NodeFirewallIPvsixNameRuleTCPFlags *FirewallIPvsixNameRuleTCPFlags `tfsdk:"flags" vyos:"flags,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixNameRuleTCP) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"mss":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum segment size (MSS)

    |  Format       |  Description                           |
    |---------------|----------------------------------------|
    |  1-16384      |  Maximum segment size                  |
    |  <min>-<max>  |  TCP MSS range (use '-' as delimiter)  |
`,
			Description: `Maximum segment size (MSS)

    |  Format       |  Description                           |
    |---------------|----------------------------------------|
    |  1-16384      |  Maximum segment size                  |
    |  <min>-<max>  |  TCP MSS range (use '-' as delimiter)  |
`,
		},

		// TagNodes

		// Nodes

		"flags": schema.SingleNestedAttribute{
			Attributes: FirewallIPvsixNameRuleTCPFlags{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `TCP flags to match

`,
			Description: `TCP flags to match

`,
		},
	}
}
