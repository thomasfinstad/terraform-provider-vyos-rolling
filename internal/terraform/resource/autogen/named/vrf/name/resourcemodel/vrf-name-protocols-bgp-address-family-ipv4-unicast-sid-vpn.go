// Package resourcemodel code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpAddressFamilyIPvfourUnicastSIDVpn{}

// VrfNameProtocolsBgpAddressFamilyIPvfourUnicastSIDVpn describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvfourUnicastSIDVpn struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastSIDVpnExport types.String `tfsdk:"export" vyos:"export,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourUnicastSIDVpn) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"export": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `For routes leaked from current VRF to VPN

    |  Format     |  Description                   |
    |-------------|--------------------------------|
    |  1-1048575  |  SID allocation index          |
    |  auto       |  Automatically assign a label  |
`,
			Description: `For routes leaked from current VRF to VPN

    |  Format     |  Description                   |
    |-------------|--------------------------------|
    |  1-1048575  |  SID allocation index          |
    |  auto       |  Automatically assign a label  |
`,
		},

		// Nodes

	}
}
