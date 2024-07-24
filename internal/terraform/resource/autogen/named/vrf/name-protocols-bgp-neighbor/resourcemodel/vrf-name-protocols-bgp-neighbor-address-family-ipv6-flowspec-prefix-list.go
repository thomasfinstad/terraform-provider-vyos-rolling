// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecPrefixList{}

// VrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecPrefixList describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecPrefixList struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecPrefixListExport types.String `tfsdk:"export" vyos:"export,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecPrefixListImport types.String `tfsdk:"import" vyos:"import,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvsixFlowspecPrefixList) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"export": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Prefix-list to filter outgoing route updates to this peer

    |  Format  |  Description               |
    |----------|----------------------------|
    |  txt     |  Name of IPv6 prefix-list  |
`,
			Description: `Prefix-list to filter outgoing route updates to this peer

    |  Format  |  Description               |
    |----------|----------------------------|
    |  txt     |  Name of IPv6 prefix-list  |
`,
		},

		"import": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Prefix-list to filter incoming route updates from this peer

    |  Format  |  Description               |
    |----------|----------------------------|
    |  txt     |  Name of IPv6 prefix-list  |
`,
			Description: `Prefix-list to filter incoming route updates from this peer

    |  Format  |  Description               |
    |----------|----------------------------|
    |  txt     |  Name of IPv6 prefix-list  |
`,
		},

		// Nodes

	}
}
