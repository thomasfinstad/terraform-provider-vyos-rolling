/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (bridge) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &InterfacesOpenvpnServerBrIDge{}

// InterfacesOpenvpnServerBrIDge describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type InterfacesOpenvpnServerBrIDge struct {
	// LeafNodes
	LeafInterfacesOpenvpnServerBrIDgeDisable    types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafInterfacesOpenvpnServerBrIDgeStart      types.String `tfsdk:"start" vyos:"start,omitempty"`
	LeafInterfacesOpenvpnServerBrIDgeStop       types.String `tfsdk:"stop" vyos:"stop,omitempty"`
	LeafInterfacesOpenvpnServerBrIDgeSubnetMask types.String `tfsdk:"subnet_mask" vyos:"subnet-mask,omitempty"`
	LeafInterfacesOpenvpnServerBrIDgeGateway    types.String `tfsdk:"gateway" vyos:"gateway,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesOpenvpnServerBrIDge) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"disable":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (disable) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
			Description: `Disable instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"start":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (start) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `First IP address in the pool

    |  Format  |  Description   |
    |----------|----------------|
    |  ipv4    |  IPv4 address  |
`,
			Description: `First IP address in the pool

    |  Format  |  Description   |
    |----------|----------------|
    |  ipv4    |  IPv4 address  |
`,
		},

		"stop":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (stop) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Last IP address in the pool

    |  Format  |  Description   |
    |----------|----------------|
    |  ipv4    |  IPv4 address  |
`,
			Description: `Last IP address in the pool

    |  Format  |  Description   |
    |----------|----------------|
    |  ipv4    |  IPv4 address  |
`,
		},

		"subnet_mask":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (subnet-mask) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Subnet mask pushed to dynamic clients.

    |  Format  |  Description       |
    |----------|--------------------|
    |  ipv4    |  IPv4 subnet mask  |
`,
			Description: `Subnet mask pushed to dynamic clients.

    |  Format  |  Description       |
    |----------|--------------------|
    |  ipv4    |  IPv4 subnet mask  |
`,
		},

		"gateway":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (gateway) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Gateway IP address

    |  Format  |  Description   |
    |----------|----------------|
    |  ipv4    |  IPv4 address  |
`,
			Description: `Gateway IP address

    |  Format  |  Description   |
    |----------|----------------|
    |  ipv4    |  IPv4 address  |
`,
		},

		// TagNodes

		// Nodes

	}
}
