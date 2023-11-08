// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsBgpPeerGroupCapability describes the resource data model.
type VrfNameProtocolsBgpPeerGroupCapability struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpPeerGroupCapabilityDynamic         types.Bool `tfsdk:"dynamic" vyos:"dynamic,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupCapabilityExtendedNexthop types.Bool `tfsdk:"extended_nexthop" vyos:"extended-nexthop,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpPeerGroupCapability) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"dynamic": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Advertise dynamic capability to this neighbor

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"extended_nexthop": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Advertise extended-nexthop capability to this neighbor

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
