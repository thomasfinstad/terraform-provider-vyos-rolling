// Package resourcemodel code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourLabeledUnicastNexthopSelf{}

// VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourLabeledUnicastNexthopSelf describes the resource data model.
type VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourLabeledUnicastNexthopSelf struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourLabeledUnicastNexthopSelfForce types.Bool `tfsdk:"force" vyos:"force,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourLabeledUnicastNexthopSelf) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"force": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Set the next hop to self for reflected routes

`,
			Description: `Set the next hop to self for reflected routes

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
