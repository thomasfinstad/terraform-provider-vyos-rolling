// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsBgpNeighbor describes the resource data model.
type VrfNameProtocolsBgpNeighbor struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

	ParentIDVrfName types.String `tfsdk:"name" vyos:"name_identifier,parent-id"`

	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAdvertisementInterval        types.Number `tfsdk:"advertisement_interval" vyos:"advertisement-interval,omitempty"`
	LeafVrfNameProtocolsBgpNeighborDescrIPtion                  types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafVrfNameProtocolsBgpNeighborDisableCapabilityNegotiation types.Bool   `tfsdk:"disable_capability_negotiation" vyos:"disable-capability-negotiation,omitempty"`
	LeafVrfNameProtocolsBgpNeighborDisableConnectedCheck        types.Bool   `tfsdk:"disable_connected_check" vyos:"disable-connected-check,omitempty"`
	LeafVrfNameProtocolsBgpNeighborEbgpMultihop                 types.Number `tfsdk:"ebgp_multihop" vyos:"ebgp-multihop,omitempty"`
	LeafVrfNameProtocolsBgpNeighborGracefulRestart              types.String `tfsdk:"graceful_restart" vyos:"graceful-restart,omitempty"`
	LeafVrfNameProtocolsBgpNeighborOverrIDeCapability           types.Bool   `tfsdk:"override_capability" vyos:"override-capability,omitempty"`
	LeafVrfNameProtocolsBgpNeighborPassive                      types.Bool   `tfsdk:"passive" vyos:"passive,omitempty"`
	LeafVrfNameProtocolsBgpNeighborPassword                     types.String `tfsdk:"password" vyos:"password,omitempty"`
	LeafVrfNameProtocolsBgpNeighborPeerGroup                    types.String `tfsdk:"peer_group" vyos:"peer-group,omitempty"`
	LeafVrfNameProtocolsBgpNeighborPort                         types.Number `tfsdk:"port" vyos:"port,omitempty"`
	LeafVrfNameProtocolsBgpNeighborRemoteAs                     types.String `tfsdk:"remote_as" vyos:"remote-as,omitempty"`
	LeafVrfNameProtocolsBgpNeighborShutdown                     types.Bool   `tfsdk:"shutdown" vyos:"shutdown,omitempty"`
	LeafVrfNameProtocolsBgpNeighborSolo                         types.Bool   `tfsdk:"solo" vyos:"solo,omitempty"`
	LeafVrfNameProtocolsBgpNeighborStrictCapabilityMatch        types.Bool   `tfsdk:"strict_capability_match" vyos:"strict-capability-match,omitempty"`
	LeafVrfNameProtocolsBgpNeighborUpdateSource                 types.String `tfsdk:"update_source" vyos:"update-source,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagVrfNameProtocolsBgpNeighborLocalAs   bool `tfsdk:"local_as" vyos:"local-as,child"`
	ExistsTagVrfNameProtocolsBgpNeighborLocalRole bool `tfsdk:"local_role" vyos:"local-role,child"`

	// Nodes
	NodeVrfNameProtocolsBgpNeighborAddressFamily *VrfNameProtocolsBgpNeighborAddressFamily `tfsdk:"address_family" vyos:"address-family,omitempty"`
	NodeVrfNameProtocolsBgpNeighborBfd           *VrfNameProtocolsBgpNeighborBfd           `tfsdk:"bfd" vyos:"bfd,omitempty"`
	NodeVrfNameProtocolsBgpNeighborCapability    *VrfNameProtocolsBgpNeighborCapability    `tfsdk:"capability" vyos:"capability,omitempty"`
	NodeVrfNameProtocolsBgpNeighborInterface     *VrfNameProtocolsBgpNeighborInterface     `tfsdk:"interface" vyos:"interface,omitempty"`
	NodeVrfNameProtocolsBgpNeighborTimers        *VrfNameProtocolsBgpNeighborTimers        `tfsdk:"timers" vyos:"timers,omitempty"`
	NodeVrfNameProtocolsBgpNeighborTTLSecURIty   *VrfNameProtocolsBgpNeighborTTLSecURIty   `tfsdk:"ttl_security" vyos:"ttl-security,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VrfNameProtocolsBgpNeighbor) GetVyosPath() []string {
	return []string{
		"vrf",

		"name",
		o.ParentIDVrfName.ValueString(),

		"protocols",

		"bgp",

		"neighbor",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighbor) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `BGP neighbor

    |  Format  |  Description  |
    |----------|---------------|
    |  ipv4  |  BGP neighbor IP address  |
    |  ipv6  |  BGP neighbor IPv6 address  |
    |  txt  |  Interface name  |

`,
		},

		"name_identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Virtual Routing and Forwarding instance

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  VRF instance name  |

`,
		},

		// LeafNodes

		"advertisement_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Minimum interval for sending routing updates

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:0-600  |  Advertisement interval in seconds  |

`,
		},

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  Description  |

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

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-255  |  Number of hops  |

`,
		},

		"graceful_restart": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `BGP graceful restart functionality

    |  Format  |  Description  |
    |----------|---------------|
    |  enable  |  Enable BGP graceful restart at peer level  |
    |  disable  |  Disable BGP graceful restart at peer level  |
    |  restart-helper  |  Enable BGP graceful restart helper only functionality  |

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

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  Peer-group name  |

`,
		},

		"port": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Neighbor BGP port

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-65535  |  Neighbor BGP port number  |

`,
		},

		"remote_as": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Neighbor BGP AS number

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-4294967294  |  Neighbor AS number  |
    |  external  |  Any AS different from the local AS  |
    |  internal  |  Neighbor AS number  |

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

    |  Format  |  Description  |
    |----------|---------------|
    |  ipv4  |  IPv4 address of route source  |
    |  ipv6  |  IPv6 address of route source  |
    |  txt  |  Interface as route source  |

`,
		},

		// Nodes

		"address_family": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborAddressFamily{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Address-family parameters

`,
		},

		"bfd": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborBfd{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Enable Bidirectional Forwarding Detection (BFD) support

`,
		},

		"capability": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborCapability{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Advertise capabilities to this peer-group

`,
		},

		"interface": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborInterface{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Interface parameters

`,
		},

		"timers": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborTimers{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Neighbor timers

`,
		},

		"ttl_security": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpNeighborTTLSecURIty{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Ttl security mechanism

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsBgpNeighbor) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsBgpNeighbor) UnmarshalJSON(_ []byte) error {
	return nil
}
