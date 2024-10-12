// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &ProtocolsBgpNeighborCapability{}

// ProtocolsBgpNeighborCapability describes the resource data model.
type ProtocolsBgpNeighborCapability struct {
	// LeafNodes
	LeafProtocolsBgpNeighborCapabilityDynamic         types.Bool `tfsdk:"dynamic" vyos:"dynamic,omitempty"`
	LeafProtocolsBgpNeighborCapabilityExtendedNexthop types.Bool `tfsdk:"extended_nexthop" vyos:"extended-nexthop,omitempty"`
	LeafProtocolsBgpNeighborCapabilitySoftwareVersion types.Bool `tfsdk:"software_version" vyos:"software-version,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighborCapability) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"dynamic": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Advertise dynamic capability to this neighbor

`,
			Description: `Advertise dynamic capability to this neighbor

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"extended_nexthop": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Advertise extended-nexthop capability to this neighbor

`,
			Description: `Advertise extended-nexthop capability to this neighbor

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"software_version": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Advertise Software Version capability to the peer

`,
			Description: `Advertise Software Version capability to the peer

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
