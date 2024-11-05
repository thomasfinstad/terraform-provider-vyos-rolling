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

var _ helpers.VyosResourceDataModel = &ProtocolsBgpNeighborInterface{}

// ProtocolsBgpNeighborInterface describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type ProtocolsBgpNeighborInterface struct {
	// LeafNodes
	LeafProtocolsBgpNeighborInterfacePeerGroup       types.String `tfsdk:"peer_group" vyos:"peer-group,omitempty"`
	LeafProtocolsBgpNeighborInterfaceRemoteAs        types.String `tfsdk:"remote_as" vyos:"remote-as,omitempty"`
	LeafProtocolsBgpNeighborInterfaceSourceInterface types.String `tfsdk:"source_interface" vyos:"source-interface,omitempty"`

	// TagNodes

	// Nodes

	NodeProtocolsBgpNeighborInterfaceVsixonly *ProtocolsBgpNeighborInterfaceVsixonly `tfsdk:"v6only" vyos:"v6only,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborInterface) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"peer_group":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Peer group for this peer

    |  Format  |  Description      |
    |----------|-------------------|
    |  txt     |  Peer-group name  |
`,
			Description: `Peer group for this peer

    |  Format  |  Description      |
    |----------|-------------------|
    |  txt     |  Peer-group name  |
`,
		},

		"remote_as":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Neighbor BGP AS number

    |  Format        |  Description                         |
    |----------------|--------------------------------------|
    |  1-4294967294  |  Neighbor AS number                  |
    |  external      |  Any AS different from the local AS  |
    |  internal      |  Neighbor AS number                  |
`,
			Description: `Neighbor BGP AS number

    |  Format        |  Description                         |
    |----------------|--------------------------------------|
    |  1-4294967294  |  Neighbor AS number                  |
    |  external      |  Any AS different from the local AS  |
    |  internal      |  Neighbor AS number                  |
`,
		},

		"source_interface":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interface used to establish connection

    |  Format     |  Description     |
    |-------------|------------------|
    |  interface  |  Interface name  |
`,
			Description: `Interface used to establish connection

    |  Format     |  Description     |
    |-------------|------------------|
    |  interface  |  Interface name  |
`,
		},

		// TagNodes

		// Nodes

		"v6only": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborInterfaceVsixonly{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Enable BGP with v6 link-local only

`,
			Description: `Enable BGP with v6 link-local only

`,
		},
	}
}
