// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VrfNameProtocolsBgpPeerGroup describes the resource data model.
type VrfNameProtocolsBgpPeerGroup struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"peer_group_id" vyos:",self-id"`

	ParentIDVrfName types.String `tfsdk:"name" vyos:"name,parent-id"`

	// LeafNodes
	LeafVrfNameProtocolsBgpPeerGroupDescrIPtion                  types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupDisableCapabilityNegotiation types.Bool   `tfsdk:"disable_capability_negotiation" vyos:"disable-capability-negotiation,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupDisableConnectedCheck        types.Bool   `tfsdk:"disable_connected_check" vyos:"disable-connected-check,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupEbgpMultihop                 types.Number `tfsdk:"ebgp_multihop" vyos:"ebgp-multihop,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupGracefulRestart              types.String `tfsdk:"graceful_restart" vyos:"graceful-restart,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupOverrIDeCapability           types.Bool   `tfsdk:"override_capability" vyos:"override-capability,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupPassive                      types.Bool   `tfsdk:"passive" vyos:"passive,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupPassword                     types.String `tfsdk:"password" vyos:"password,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupRemoteAs                     types.String `tfsdk:"remote_as" vyos:"remote-as,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupShutdown                     types.Bool   `tfsdk:"shutdown" vyos:"shutdown,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupUpdateSource                 types.String `tfsdk:"update_source" vyos:"update-source,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagVrfNameProtocolsBgpPeerGroupLocalAs   bool `tfsdk:"-" vyos:"local-as,ignore,child"`
	ExistsTagVrfNameProtocolsBgpPeerGroupLocalRole bool `tfsdk:"-" vyos:"local-role,ignore,child"`

	// Nodes
	NodeVrfNameProtocolsBgpPeerGroupAddressFamily *VrfNameProtocolsBgpPeerGroupAddressFamily `tfsdk:"address_family" vyos:"address-family,omitempty"`
	NodeVrfNameProtocolsBgpPeerGroupBfd           *VrfNameProtocolsBgpPeerGroupBfd           `tfsdk:"bfd" vyos:"bfd,omitempty"`
	NodeVrfNameProtocolsBgpPeerGroupCapability    *VrfNameProtocolsBgpPeerGroupCapability    `tfsdk:"capability" vyos:"capability,omitempty"`
	NodeVrfNameProtocolsBgpPeerGroupTTLSecURIty   *VrfNameProtocolsBgpPeerGroupTTLSecURIty   `tfsdk:"ttl_security" vyos:"ttl-security,omitempty"`
}

// SetID configures the resource ID
func (o *VrfNameProtocolsBgpPeerGroup) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VrfNameProtocolsBgpPeerGroup) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"vrf",

		"name",
		o.ParentIDVrfName.ValueString(),

		"protocols",

		"bgp",

		"peer-group",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpPeerGroup) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"peer_group_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Name of peer-group

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		"name_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Virtual Routing and Forwarding instance

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  VRF instance name  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

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
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamily{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Address-family parameters

`,
		},

		"bfd": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupBfd{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Enable Bidirectional Forwarding Detection (BFD) support

`,
		},

		"capability": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupCapability{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Advertise capabilities to this peer-group

`,
		},

		"ttl_security": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupTTLSecURIty{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Ttl security mechanism

`,
		},
	}
}
