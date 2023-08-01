// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspec describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspec struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspecRouteReflectorClient types.Bool `tfsdk:"route_reflector_client" vyos:"route-reflector-client,omitempty"`
	LeafVrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspecRouteServerClient    types.Bool `tfsdk:"route_server_client" vyos:"route-server-client,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspecPrefixList          *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspecPrefixList          `tfsdk:"prefix_list" vyos:"prefix-list,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspecFilterList          *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspecFilterList          `tfsdk:"filter_list" vyos:"filter-list,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspecRouteMap            *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspecRouteMap            `tfsdk:"route_map" vyos:"route-map,omitempty"`
	NodeVrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspecSoftReconfiguration *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspecSoftReconfiguration `tfsdk:"soft_reconfiguration" vyos:"soft-reconfiguration,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspec) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"route_reflector_client": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Peer is a route reflector client

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"route_server_client": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Peer is a route server client

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

		"prefix_list": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspecPrefixList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4-Prefix-list to filter route updates to/from this peer

`,
		},

		"filter_list": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspecFilterList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `as-path-list to filter route updates to/from this peer

`,
		},

		"route_map": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspecRouteMap{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
		},

		"soft_reconfiguration": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspecSoftReconfiguration{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Soft reconfiguration for peer

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspec) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsBgpNeighborAddressFamilyIPvfourFlowspec) UnmarshalJSON(_ []byte) error {
	return nil
}