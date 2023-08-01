// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsBgpNeighbor describes the resource data model.
type ProtocolsBgpNeighbor struct {
	SelfIdentifier types.String `tfsdk:"neighbor_id" vyos:",self-id"`

	// LeafNodes
	LeafProtocolsBgpNeighborAdvertisementInterval        types.Number `tfsdk:"advertisement_interval" vyos:"advertisement-interval,omitempty"`
	LeafProtocolsBgpNeighborDescrIPtion                  types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafProtocolsBgpNeighborDisableCapabilityNegotiation types.Bool   `tfsdk:"disable_capability_negotiation" vyos:"disable-capability-negotiation,omitempty"`
	LeafProtocolsBgpNeighborDisableConnectedCheck        types.Bool   `tfsdk:"disable_connected_check" vyos:"disable-connected-check,omitempty"`
	LeafProtocolsBgpNeighborEbgpMultihop                 types.Number `tfsdk:"ebgp_multihop" vyos:"ebgp-multihop,omitempty"`
	LeafProtocolsBgpNeighborGracefulRestart              types.String `tfsdk:"graceful_restart" vyos:"graceful-restart,omitempty"`
	LeafProtocolsBgpNeighborOverrIDeCapability           types.Bool   `tfsdk:"override_capability" vyos:"override-capability,omitempty"`
	LeafProtocolsBgpNeighborPassive                      types.Bool   `tfsdk:"passive" vyos:"passive,omitempty"`
	LeafProtocolsBgpNeighborPassword                     types.String `tfsdk:"password" vyos:"password,omitempty"`
	LeafProtocolsBgpNeighborPeerGroup                    types.String `tfsdk:"peer_group" vyos:"peer-group,omitempty"`
	LeafProtocolsBgpNeighborPort                         types.Number `tfsdk:"port" vyos:"port,omitempty"`
	LeafProtocolsBgpNeighborRemoteAs                     types.String `tfsdk:"remote_as" vyos:"remote-as,omitempty"`
	LeafProtocolsBgpNeighborShutdown                     types.Bool   `tfsdk:"shutdown" vyos:"shutdown,omitempty"`
	LeafProtocolsBgpNeighborSolo                         types.Bool   `tfsdk:"solo" vyos:"solo,omitempty"`
	LeafProtocolsBgpNeighborStrictCapabilityMatch        types.Bool   `tfsdk:"strict_capability_match" vyos:"strict-capability-match,omitempty"`
	LeafProtocolsBgpNeighborUpdateSource                 types.String `tfsdk:"update_source" vyos:"update-source,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagProtocolsBgpNeighborLocalAs   bool `tfsdk:"local_as" vyos:"local-as,child"`
	ExistsTagProtocolsBgpNeighborLocalRole bool `tfsdk:"local_role" vyos:"local-role,child"`

	// Nodes
	NodeProtocolsBgpNeighborAddressFamily *ProtocolsBgpNeighborAddressFamily `tfsdk:"address_family" vyos:"address-family,omitempty"`
	NodeProtocolsBgpNeighborBfd           *ProtocolsBgpNeighborBfd           `tfsdk:"bfd" vyos:"bfd,omitempty"`
	NodeProtocolsBgpNeighborCapability    *ProtocolsBgpNeighborCapability    `tfsdk:"capability" vyos:"capability,omitempty"`
	NodeProtocolsBgpNeighborInterface     *ProtocolsBgpNeighborInterface     `tfsdk:"interface" vyos:"interface,omitempty"`
	NodeProtocolsBgpNeighborTimers        *ProtocolsBgpNeighborTimers        `tfsdk:"timers" vyos:"timers,omitempty"`
	NodeProtocolsBgpNeighborTTLSecURIty   *ProtocolsBgpNeighborTTLSecURIty   `tfsdk:"ttl_security" vyos:"ttl-security,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBgpNeighbor) GetVyosPath() []string {
	return []string{
		"protocols",

		"bgp",

		"neighbor",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighbor) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"neighbor_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `BGP neighbor

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  BGP neighbor IP address  |
    |  ipv6  &emsp; |  BGP neighbor IPv6 address  |
    |  txt  &emsp; |  Interface name  |

`,
		},

		// LeafNodes

		"advertisement_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Minimum interval for sending routing updates

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-600  &emsp; |  Advertisement interval in seconds  |

`,
		},

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Description  |

`,
		},

		"disable_capability_negotiation": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable capability negotiation with this neighbor

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"disable_connected_check": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable check to see if eBGP peer address is a connected route

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"ebgp_multihop": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Allow this EBGP neighbor to not be on a directly connected network

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-255  &emsp; |  Number of hops  |

`,
		},

		"graceful_restart": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `BGP graceful restart functionality

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  enable  &emsp; |  Enable BGP graceful restart at peer level  |
    |  disable  &emsp; |  Disable BGP graceful restart at peer level  |
    |  restart-helper  &emsp; |  Enable BGP graceful restart helper only functionality  |

`,
		},

		"override_capability": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Ignore capability negotiation with specified neighbor

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"passive": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not initiate a session with this neighbor

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `BGP MD5 password

`,
		},

		"peer_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Peer group for this peer

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Peer-group name  |

`,
		},

		"port": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Neighbor BGP port

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Neighbor BGP port number  |

`,
		},

		"remote_as": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Neighbor BGP AS number

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-4294967294  &emsp; |  Neighbor AS number  |
    |  external  &emsp; |  Any AS different from the local AS  |
    |  internal  &emsp; |  Neighbor AS number  |

`,
		},

		"shutdown": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Administratively shutdown this neighbor

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"solo": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not send back prefixes learned from the neighbor

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"strict_capability_match": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable strict capability negotiation

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"update_source": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Source IP of routing updates

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IPv4 address of route source  |
    |  ipv6  &emsp; |  IPv6 address of route source  |
    |  txt  &emsp; |  Interface as route source  |

`,
		},

		// Nodes

		"address_family": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborAddressFamily{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Address-family parameters

`,
		},

		"bfd": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborBfd{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Enable Bidirectional Forwarding Detection (BFD) support

`,
		},

		"capability": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborCapability{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Advertise capabilities to this peer-group

`,
		},

		"interface": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborInterface{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Interface parameters

`,
		},

		"timers": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborTimers{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Neighbor timers

`,
		},

		"ttl_security": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborTTLSecURIty{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Ttl security mechanism

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBgpNeighbor) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsBgpNeighbor) UnmarshalJSON(_ []byte) error {
	return nil
}