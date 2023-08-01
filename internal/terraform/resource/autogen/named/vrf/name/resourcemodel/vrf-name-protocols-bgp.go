// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsBgp describes the resource data model.
type VrfNameProtocolsBgp struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpSystemAs types.Number `tfsdk:"system_as" vyos:"system-as,omitempty"`
	LeafVrfNameProtocolsBgpRouteMap types.String `tfsdk:"route_map" vyos:"route-map,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagVrfNameProtocolsBgpNeighbor  bool `tfsdk:"neighbor" vyos:"neighbor,child"`
	ExistsTagVrfNameProtocolsBgpPeerGroup bool `tfsdk:"peer_group" vyos:"peer-group,child"`

	// Nodes
	NodeVrfNameProtocolsBgpAddressFamily *VrfNameProtocolsBgpAddressFamily `tfsdk:"address_family" vyos:"address-family,omitempty"`
	NodeVrfNameProtocolsBgpListen        *VrfNameProtocolsBgpListen        `tfsdk:"listen" vyos:"listen,omitempty"`
	NodeVrfNameProtocolsBgpParameters    *VrfNameProtocolsBgpParameters    `tfsdk:"parameters" vyos:"parameters,omitempty"`
	NodeVrfNameProtocolsBgpTimers        *VrfNameProtocolsBgpTimers        `tfsdk:"timers" vyos:"timers,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgp) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"system_as": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Autonomous System Number (ASN)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-4294967294  &emsp; |  Autonomous System Number  |

`,
		},

		"route_map": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specify route-map name to use

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Route map name  |

`,
		},

		// Nodes

		"address_family": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamily{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `BGP address-family parameters

`,
		},

		"listen": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpListen{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Listen for and accept BGP dynamic neighbors from range

`,
		},

		"parameters": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpParameters{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `BGP parameters

`,
		},

		"timers": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpTimers{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `BGP protocol timers

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsBgp) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsBgp) UnmarshalJSON(_ []byte) error {
	return nil
}