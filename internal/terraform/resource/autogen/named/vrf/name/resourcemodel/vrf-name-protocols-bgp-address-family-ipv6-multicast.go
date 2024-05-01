// Package resourcemodel code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpAddressFamilyIPvsixMulticast{}

// VrfNameProtocolsBgpAddressFamilyIPvsixMulticast describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvsixMulticast struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagVrfNameProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress bool `tfsdk:"aggregate_address" vyos:"aggregate-address,child"`
	ExistsTagVrfNameProtocolsBgpAddressFamilyIPvsixMulticastNetwork          bool `tfsdk:"network" vyos:"network,child"`

	// Nodes
	NodeVrfNameProtocolsBgpAddressFamilyIPvsixMulticastDistance *VrfNameProtocolsBgpAddressFamilyIPvsixMulticastDistance `tfsdk:"distance" vyos:"distance,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvsixMulticast) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"distance": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyIPvsixMulticastDistance{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Administrative distances for BGP routes

`,
			Description: `Administrative distances for BGP routes

`,
		},
	}
}
