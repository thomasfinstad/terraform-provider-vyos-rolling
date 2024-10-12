// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixVpnDistributeList{}

// VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixVpnDistributeList describes the resource data model.
type VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixVpnDistributeList struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixVpnDistributeListExport types.Number `tfsdk:"export" vyos:"export,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixVpnDistributeListImport types.Number `tfsdk:"import" vyos:"import,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixVpnDistributeList) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"export": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Access-list to filter outgoing route updates to this peer-group

    |  Format   |  Description                                                      |
    |-----------|-------------------------------------------------------------------|
    |  1-65535  |  Access-list to filter outgoing route updates to this peer-group  |
`,
			Description: `Access-list to filter outgoing route updates to this peer-group

    |  Format   |  Description                                                      |
    |-----------|-------------------------------------------------------------------|
    |  1-65535  |  Access-list to filter outgoing route updates to this peer-group  |
`,
		},

		"import": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Access-list to filter incoming route updates from this peer-group

    |  Format   |  Description                                                        |
    |-----------|---------------------------------------------------------------------|
    |  1-65535  |  Access-list to filter incoming route updates from this peer-group  |
`,
			Description: `Access-list to filter incoming route updates from this peer-group

    |  Format   |  Description                                                        |
    |-----------|---------------------------------------------------------------------|
    |  1-65535  |  Access-list to filter incoming route updates from this peer-group  |
`,
		},

		// Nodes

	}
}
