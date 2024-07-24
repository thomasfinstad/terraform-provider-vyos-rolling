// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &VrfNameProtocolsBgpPeerGroup{}

// VrfNameProtocolsBgpPeerGroup describes the resource data model.
type VrfNameProtocolsBgpPeerGroup struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"peer_group_id" vyos:"-,self-id"`

	ParentIDVrfName types.String `tfsdk:"name_id" vyos:"name,parent-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafVrfNameProtocolsBgpPeerGroupDescrIPtion                  types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupDisableCapabilityNegotiation types.Bool   `tfsdk:"disable_capability_negotiation" vyos:"disable-capability-negotiation,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupDisableConnectedCheck        types.Bool   `tfsdk:"disable_connected_check" vyos:"disable-connected-check,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupEbgpMultihop                 types.Number `tfsdk:"ebgp_multihop" vyos:"ebgp-multihop,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupGracefulRestart              types.String `tfsdk:"graceful_restart" vyos:"graceful-restart,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupOverrIDeCapability           types.Bool   `tfsdk:"override_capability" vyos:"override-capability,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupPassive                      types.Bool   `tfsdk:"passive" vyos:"passive,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupPassword                     types.String `tfsdk:"password" vyos:"password,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupShutdown                     types.Bool   `tfsdk:"shutdown" vyos:"shutdown,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupUpdateSource                 types.String `tfsdk:"update_source" vyos:"update-source,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupRemoteAs                     types.String `tfsdk:"remote_as" vyos:"remote-as,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupPort                         types.Number `tfsdk:"port" vyos:"port,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagVrfNameProtocolsBgpPeerGroupLocalAs   bool `tfsdk:"-" vyos:"local-as,child"`
	ExistsTagVrfNameProtocolsBgpPeerGroupLocalRole bool `tfsdk:"-" vyos:"local-role,child"`

	// Nodes
	NodeVrfNameProtocolsBgpPeerGroupAddressFamily *VrfNameProtocolsBgpPeerGroupAddressFamily `tfsdk:"address_family" vyos:"address-family,omitempty"`
	NodeVrfNameProtocolsBgpPeerGroupBfd           *VrfNameProtocolsBgpPeerGroupBfd           `tfsdk:"bfd" vyos:"bfd,omitempty"`
	NodeVrfNameProtocolsBgpPeerGroupCapability    *VrfNameProtocolsBgpPeerGroupCapability    `tfsdk:"capability" vyos:"capability,omitempty"`
	NodeVrfNameProtocolsBgpPeerGroupPathAttribute *VrfNameProtocolsBgpPeerGroupPathAttribute `tfsdk:"path_attribute" vyos:"path-attribute,omitempty"`
	NodeVrfNameProtocolsBgpPeerGroupTTLSecURIty   *VrfNameProtocolsBgpPeerGroupTTLSecURIty   `tfsdk:"ttl_security" vyos:"ttl-security,omitempty"`
}

// SetID configures the resource ID
func (o *VrfNameProtocolsBgpPeerGroup) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *VrfNameProtocolsBgpPeerGroup) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *VrfNameProtocolsBgpPeerGroup) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VrfNameProtocolsBgpPeerGroup) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"peer-group",
		o.SelfIdentifier.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *VrfNameProtocolsBgpPeerGroup) GetVyosParentPath() []string {
	return []string{
		"vrf",

		"name",
		o.ParentIDVrfName.ValueString(),

		"protocols",

		"bgp",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *VrfNameProtocolsBgpPeerGroup) GetVyosNamedParentPath() []string {
	return []string{
		"vrf",

		"name",
		o.ParentIDVrfName.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpPeerGroup) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"peer_group_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Name of peer-group

`,
			Description: `Name of peer-group

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in peer_group_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illegal character in  peer_group_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		"name_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Virtual Routing and Forwarding instance

    |  Format  |  Description        |
    |----------|---------------------|
    |  txt     |  VRF instance name  |
`,
			Description: `Virtual Routing and Forwarding instance

    |  Format  |  Description        |
    |----------|---------------------|
    |  txt     |  VRF instance name  |
`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in name_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illegal character in  name_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Description  |
`,
			Description: `Description

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Description  |
`,
		},

		"disable_capability_negotiation": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable capability negotiation with this neighbor

`,
			Description: `Disable capability negotiation with this neighbor

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"disable_connected_check": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Allow peerings between eBGP peer using loopback/dummy address

`,
			Description: `Allow peerings between eBGP peer using loopback/dummy address

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"ebgp_multihop": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Allow this EBGP neighbor to not be on a directly connected network

    |  Format  |  Description     |
    |----------|------------------|
    |  1-255   |  Number of hops  |
`,
			Description: `Allow this EBGP neighbor to not be on a directly connected network

    |  Format  |  Description     |
    |----------|------------------|
    |  1-255   |  Number of hops  |
`,
		},

		"graceful_restart": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `BGP graceful restart functionality

    |  Format          |  Description                                            |
    |------------------|---------------------------------------------------------|
    |  enable          |  Enable BGP graceful restart at peer level              |
    |  disable         |  Disable BGP graceful restart at peer level             |
    |  restart-helper  |  Enable BGP graceful restart helper only functionality  |
`,
			Description: `BGP graceful restart functionality

    |  Format          |  Description                                            |
    |------------------|---------------------------------------------------------|
    |  enable          |  Enable BGP graceful restart at peer level              |
    |  disable         |  Disable BGP graceful restart at peer level             |
    |  restart-helper  |  Enable BGP graceful restart helper only functionality  |
`,
		},

		"override_capability": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Ignore capability negotiation with specified neighbor

`,
			Description: `Ignore capability negotiation with specified neighbor

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"passive": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not initiate a session with this neighbor

`,
			Description: `Do not initiate a session with this neighbor

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `BGP MD5 password

`,
			Description: `BGP MD5 password

`,
		},

		"shutdown": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Administratively shutdown this neighbor

`,
			Description: `Administratively shutdown this neighbor

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"update_source": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Source IP of routing updates

    |  Format  |  Description                   |
    |----------|--------------------------------|
    |  ipv4    |  IPv4 address of route source  |
    |  ipv6    |  IPv6 address of route source  |
    |  txt     |  Interface as route source     |
`,
			Description: `Source IP of routing updates

    |  Format  |  Description                   |
    |----------|--------------------------------|
    |  ipv4    |  IPv4 address of route source  |
    |  ipv6    |  IPv6 address of route source  |
    |  txt     |  Interface as route source     |
`,
		},

		"remote_as": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Neighbor BGP AS number

    |  Format        |  Description                         |
    |----------------|--------------------------------------|
    |  1-4294967294  |  Neighbor AS number                  |
    |  external      |  Any AS different from the local AS  |
    |  internal      |  Neighbor AS number                  |
`,
			Description: `Neighbor BGP AS number

    |  Format        |  Description                         |
    |----------------|--------------------------------------|
    |  1-4294967294  |  Neighbor AS number                  |
    |  external      |  Any AS different from the local AS  |
    |  internal      |  Neighbor AS number                  |
`,
		},

		"port": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Port number used by connection

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Numeric IP port  |
`,
			Description: `Port number used by connection

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Numeric IP port  |
`,
		},

		// Nodes

		"address_family": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamily{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Address-family parameters

`,
			Description: `Address-family parameters

`,
		},

		"bfd": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupBfd{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Enable Bidirectional Forwarding Detection (BFD) support

`,
			Description: `Enable Bidirectional Forwarding Detection (BFD) support

`,
		},

		"capability": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupCapability{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Advertise capabilities to this peer-group

`,
			Description: `Advertise capabilities to this peer-group

`,
		},

		"path_attribute": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupPathAttribute{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Manipulate path attributes from incoming UPDATE messages

`,
			Description: `Manipulate path attributes from incoming UPDATE messages

`,
		},

		"ttl_security": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupTTLSecURIty{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Ttl security mechanism

`,
			Description: `Ttl security mechanism

`,
		},
	}
}
