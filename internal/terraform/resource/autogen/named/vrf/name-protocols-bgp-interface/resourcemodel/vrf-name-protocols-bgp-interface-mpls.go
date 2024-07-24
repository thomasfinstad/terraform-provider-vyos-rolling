// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpInterfaceMpls{}

// VrfNameProtocolsBgpInterfaceMpls describes the resource data model.
type VrfNameProtocolsBgpInterfaceMpls struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpInterfaceMplsForwarding types.Bool `tfsdk:"forwarding" vyos:"forwarding,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpInterfaceMpls) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"forwarding": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable MPLS forwarding for eBGP directly connected peers

`,
			Description: `Enable MPLS forwarding for eBGP directly connected peers

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
