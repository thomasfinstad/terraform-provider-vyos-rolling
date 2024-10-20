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

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpAddressFamilyIPvfourUnicastLabelVpn{}

// VrfNameProtocolsBgpAddressFamilyIPvfourUnicastLabelVpn describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvfourUnicastLabelVpn struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastLabelVpnExport types.String `tfsdk:"export" vyos:"export,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeVrfNameProtocolsBgpAddressFamilyIPvfourUnicastLabelVpnAllocationMode *VrfNameProtocolsBgpAddressFamilyIPvfourUnicastLabelVpnAllocationMode `tfsdk:"allocation_mode" vyos:"allocation-mode,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourUnicastLabelVpn) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"export":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `For routes leaked from current address-family to VPN

    |  Format     |  Description                   |
    |-------------|--------------------------------|
    |  auto       |  Automatically assign a label  |
    |  0-1048575  |  Label Value                   |
`,
			Description: `For routes leaked from current address-family to VPN

    |  Format     |  Description                   |
    |-------------|--------------------------------|
    |  auto       |  Automatically assign a label  |
    |  0-1048575  |  Label Value                   |
`,
		},

		// Nodes

		"allocation_mode": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvfourUnicastLabelVpnAllocationMode{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Label allocation mode

`,
			Description: `Label allocation mode

`,
		},
	}
}
