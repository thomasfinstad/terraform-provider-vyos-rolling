// Package resourcemodel code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpAddressFamilyIPvsixUnicastLabelVpn{}

// VrfNameProtocolsBgpAddressFamilyIPvsixUnicastLabelVpn describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvsixUnicastLabelVpn struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyIPvsixUnicastLabelVpnExport types.String `tfsdk:"export" vyos:"export,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsBgpAddressFamilyIPvsixUnicastLabelVpnAllocationMode *VrfNameProtocolsBgpAddressFamilyIPvsixUnicastLabelVpnAllocationMode `tfsdk:"allocation_mode" vyos:"allocation-mode,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvsixUnicastLabelVpn) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"export": schema.StringAttribute{
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
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvsixUnicastLabelVpnAllocationMode{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Label allocation mode

`,
			Description: `Label allocation mode

`,
		},
	}
}
