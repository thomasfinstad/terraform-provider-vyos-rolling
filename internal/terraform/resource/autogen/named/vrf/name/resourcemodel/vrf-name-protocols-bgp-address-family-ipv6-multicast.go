// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// VrfNameProtocolsBgpAddressFamilyIPvsixMulticast describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvsixMulticast struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagVrfNameProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress bool `tfsdk:"aggregate_address" vyos:"aggregate-address,ignore,child"`
	ExistsTagVrfNameProtocolsBgpAddressFamilyIPvsixMulticastNetwork          bool `tfsdk:"network" vyos:"network,ignore,child"`

	// Nodes
	NodeVrfNameProtocolsBgpAddressFamilyIPvsixMulticastDistance *VrfNameProtocolsBgpAddressFamilyIPvsixMulticastDistance `tfsdk:"distance" vyos:"distance,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvsixMulticast) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"distance": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvsixMulticastDistance{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Administrative distances for BGP routes

`,
		},
	}
}
