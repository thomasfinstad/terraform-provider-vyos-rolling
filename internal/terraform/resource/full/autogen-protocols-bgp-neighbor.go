package resourcefull

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &protocols_bgp_neighbor{}

// var _ resource.ResourceWithImportState = &protocols_bgp_neighbor{}

// protocols_bgp_neighbor defines the resource implementation.
type protocols_bgp_neighbor struct {
	client *http.Client
}

// protocols bgp neighborModel describes the resource data model.
type protocols_bgp_neighborModel struct {
	ConfigurableAttribute types.String `tfsdk:"configurable_attribute"`
	Defaulted             types.String `tfsdk:"defaulted"`
	ID                    types.String `tfsdk:"id"`

	Local_as types.Map `tfsdk:local_as`

	Local_role types.Map `tfsdk:local_role`

	Address_family types.String `tfsdk:address_family`

	Bfd types.String `tfsdk:bfd`

	Capability types.String `tfsdk:capability`

	Iface types.String `tfsdk:interface`

	Timers types.String `tfsdk:timers`

	Ttl_security types.String `tfsdk:ttl_security`

	Advertisement_interval types.String `tfsdk:advertisement_interval`

	Description types.String `tfsdk:description`

	Disable_capability_negotiation types.String `tfsdk:disable_capability_negotiation`

	Disable_connected_check types.String `tfsdk:disable_connected_check`

	Ebgp_multihop types.String `tfsdk:ebgp_multihop`

	Graceful_restart types.String `tfsdk:graceful_restart`

	Override_capability types.String `tfsdk:override_capability`

	Passive types.String `tfsdk:passive`

	Password types.String `tfsdk:password`

	Peer_group types.String `tfsdk:peer_group`

	Port types.String `tfsdk:port`

	Remote_as types.String `tfsdk:remote_as`

	Shutdown types.String `tfsdk:shutdown`

	Solo types.String `tfsdk:solo`

	Strict_capability_match types.String `tfsdk:strict_capability_match`

	Update_source types.String `tfsdk:update_source`
}

// Metadata method to define the resource type name.
func (r *protocols_bgp_neighbor) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protocols bgp neighbor"
}

// protocols_bgp_neighborResource method to return the example resource reference
func protocols_bgp_neighborResource() resource.Resource {
	return &protocols_bgp_neighbor{}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r *protocols_bgp_neighbor) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example resource",

		Attributes: map[string]schema.Attribute{
			"local_as": schema.MapNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"no_prepend": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								// TODO handle non-string types
								"replace_as": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Prepend only local-as from/to updates for eBGP peers

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},
							},
							// CustomType:          basetypes.MapTypable(nil),
							// Required:            false,
							Optional: true,
							// Computed:            false,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Disable prepending local-as from/to updates for eBGP peers

`,
							// DeprecationMessage:  "",
							// Validators:          []validator.Map(nil),
							// PlanModifiers:       []planmodifier.Map(nil),
							// TODO investigate if node defaults can be handled
							// Default:             defaults.Map(nil),
						},

						// CustomType:    basetypes.ObjectTypable(nil),
						// Validators:    []validator.Object(nil),
						// PlanModifiers: []planmodifier.Object(nil),
					},
				},
				// CustomType:          basetypes.MapTypable(nil),
				// Required:            false,
				Optional: true,
				// Computed:            false,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Specify alternate ASN for this BGP process

Format: u32:1-4294967294
Autonomous System Number (ASN)
`,
				// DeprecationMessage:  "",
				// Validators:          []validator.Map(nil),
				// PlanModifiers:       []planmodifier.Map(nil),
				// TODO investigate if tagnode defaults can be handled
				// Default:             defaults.Map(nil),
			},

			"local_role": schema.MapNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						// TODO handle non-string types
						"strict": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Neighbor must send this exact capability, otherwise a role missmatch notification will be sent

`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// CustomType:    basetypes.ObjectTypable(nil),
						// Validators:    []validator.Object(nil),
						// PlanModifiers: []planmodifier.Object(nil),
					},
				},
				// CustomType:          basetypes.MapTypable(nil),
				// Required:            false,
				Optional: true,
				// Computed:            false,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Local role for BGP neighbor (RFC9234)

Format: customer
Using Transit
Format: peer
Public/Private Peering
Format: provider
Providing Transit
Format: rs-client
RS Client
Format: rs-server
Route Server
`,
				// DeprecationMessage:  "",
				// Validators:          []validator.Map(nil),
				// PlanModifiers:       []planmodifier.Map(nil),
				// TODO investigate if tagnode defaults can be handled
				// Default:             defaults.Map(nil),
			},

			// TODO handle non-string types
			"advertisement_interval": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Minimum interval for sending routing updates

Format: u32:0-600
Advertisement interval in seconds
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"description": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Description

Format: txt
Description
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"disable_capability_negotiation": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Disable capability negotiation with this neighbor

`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"disable_connected_check": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Disable check to see if eBGP peer address is a connected route

`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"ebgp_multihop": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Allow this EBGP neighbor to not be on a directly connected network

Format: u32:1-255
Number of hops
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"graceful_restart": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `BGP graceful restart functionality

Format: enable
Enable BGP graceful restart at peer level
Format: disable
Disable BGP graceful restart at peer level
Format: restart-helper
Enable BGP graceful restart helper only functionality
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"override_capability": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Ignore capability negotiation with specified neighbor

`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"passive": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Do not initiate a session with this neighbor

`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"password": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `BGP MD5 password

`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"peer_group": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Peer group for this peer

Format: txt
Peer-group name
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"port": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Neighbor BGP port

Format: u32:1-65535
Neighbor BGP port number
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"remote_as": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Neighbor BGP AS number

Format: u32:1-4294967294
Neighbor AS number
Format: external
Any AS different from the local AS
Format: internal
Neighbor AS number
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"shutdown": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Administratively shutdown this neighbor

`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"solo": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Do not send back prefixes learned from the neighbor

`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"strict_capability_match": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Enable strict capability negotiation

`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"update_source": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Source IP of routing updates

Format: ipv4
IPv4 address of route source
Format: ipv6
IPv6 address of route source
Format: txt
Interface as route source
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			"address_family": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"ipv4_unicast": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							// TODO handle non-string types
							"addpath_tx_all": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Use addpath to advertise all paths to a neighbor

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"addpath_tx_per_as": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Use addpath to advertise the bestpath per each neighboring AS

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"as_override": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Override ASN in outbound updates to configured neighbor local-as

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"maximum_prefix": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Maximum number of prefixes to accept from this peer

Format: u32:1-4294967295
Prefix limit
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"maximum_prefix_out": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Maximum number of prefixes to be sent to this peer

Format: u32:1-4294967295
Prefix limit
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"remove_private_as": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Remove private AS numbers from AS path in outbound route updates

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"route_reflector_client": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Peer is a route reflector client

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"route_server_client": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Peer is a route server client

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"unsuppress_map": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Route-map to selectively unsuppress suppressed routes

Format: txt
Route map name
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"weight": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Default weight for routes from this peer

Format: u32:1-65535
Default weight
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							"capability": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									"orf": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											"prefix_list": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													// TODO handle non-string types
													"receive": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Capability to receive the ORF

`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"send": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Capability to send the ORF

`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},
												},
												// CustomType:          basetypes.MapTypable(nil),
												// Required:            false,
												Optional: true,
												// Computed:            false,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Advertise prefix-list ORF capability to this peer

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},
										},
										// CustomType:          basetypes.MapTypable(nil),
										// Required:            false,
										Optional: true,
										// Computed:            false,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Advertise ORF capability to this peer

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Advertise capabilities to this neighbor (IPv4)

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"prefix_list": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `IPv4-Prefix-list to filter outgoing route updates to this peer

Format: txt
Name of IPv4 prefix-list
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `IPv4-Prefix-list to filter incoming route updates from this peer

Format: txt
Name of IPv4 prefix-list
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `IPv4-Prefix-list to filter route updates to/from this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"conditionally_advertise": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"advertise_map": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Route-map to conditionally advertise routes

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"exist_map": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Advertise routes only if prefixes in exist-map are installed in BGP table

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"non_exist_map": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Advertise routes only if prefixes in non-exist-map are not installed in BGP table

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Use route-map to conditionally advertise routes

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"allowas_in": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"number": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Number of occurrences of AS number

Format: u32:1-10
Number of times AS is allowed in path
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Accept route that contains the local-as in the as-path

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"attribute_unchanged": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"as_path": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Send AS path unchanged

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"med": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Send multi-exit discriminator unchanged

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"next_hop": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Send nexthop unchanged

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `BGP attributes are sent unchanged

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"disable_send_community": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"extended": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Disable sending extended community attributes to this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"standard": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Disable sending standard community attributes to this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Disable sending community attributes to this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"distribute_list": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Access-list to filter outgoing route updates to this peer-group

Format: u32:1-65535
Access-list to filter outgoing route updates to this peer-group
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Access-list to filter incoming route updates from this peer-group

Format: u32:1-65535
Access-list to filter incoming route updates from this peer-group
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Access-list to filter route updates to/from this peer-group

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"filter_list": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `As-path-list to filter outgoing route updates to this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `As-path-list to filter incoming route updates from this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `as-path-list to filter route updates to/from this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"nexthop_self": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"force": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Set the next hop to self for reflected routes

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Disable the next hop calculation for this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"route_map": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Route-map to filter outgoing route updates

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Route-map to filter incoming route updates

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"soft_reconfiguration": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"inbound": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Enable inbound soft reconfiguration

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Soft reconfiguration for peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"default_originate": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"route_map": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Specify route-map name to use

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Originate default route to this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},
						},
						// CustomType:          basetypes.MapTypable(nil),
						// Required:            false,
						Optional: true,
						// Computed:            false,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `IPv4 BGP neighbor parameters

`,
						// DeprecationMessage:  "",
						// Validators:          []validator.Map(nil),
						// PlanModifiers:       []planmodifier.Map(nil),
						// TODO investigate if node defaults can be handled
						// Default:             defaults.Map(nil),
					},

					"ipv6_unicast": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							// TODO handle non-string types
							"addpath_tx_all": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Use addpath to advertise all paths to a neighbor

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"addpath_tx_per_as": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Use addpath to advertise the bestpath per each neighboring AS

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"as_override": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Override ASN in outbound updates to configured neighbor local-as

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"maximum_prefix": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Maximum number of prefixes to accept from this peer

Format: u32:1-4294967295
Prefix limit
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"maximum_prefix_out": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Maximum number of prefixes to be sent to this peer

Format: u32:1-4294967295
Prefix limit
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"remove_private_as": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Remove private AS numbers from AS path in outbound route updates

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"route_reflector_client": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Peer is a route reflector client

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"route_server_client": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Peer is a route server client

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"unsuppress_map": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Route-map to selectively unsuppress suppressed routes

Format: txt
Route map name
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"weight": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Default weight for routes from this peer

Format: u32:1-65535
Default weight
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							"capability": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									"orf": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											"prefix_list": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													// TODO handle non-string types
													"receive": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Capability to receive the ORF

`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"send": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Capability to send the ORF

`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},
												},
												// CustomType:          basetypes.MapTypable(nil),
												// Required:            false,
												Optional: true,
												// Computed:            false,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Advertise prefix-list ORF capability to this peer

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},
										},
										// CustomType:          basetypes.MapTypable(nil),
										// Required:            false,
										Optional: true,
										// Computed:            false,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Advertise ORF capability to this peer

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Advertise capabilities to this neighbor (IPv6)

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"nexthop_local": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"unchanged": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Leave link-local nexthop unchanged for this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Nexthop attributes

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"prefix_list": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Prefix-list to filter outgoing route updates to this peer

Format: txt
Name of IPv6 prefix-list
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Prefix-list to filter incoming route updates from this peer

Format: txt
Name of IPv6 prefix-list
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Prefix-list to filter route updates to/from this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"conditionally_advertise": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"advertise_map": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Route-map to conditionally advertise routes

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"exist_map": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Advertise routes only if prefixes in exist-map are installed in BGP table

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"non_exist_map": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Advertise routes only if prefixes in non-exist-map are not installed in BGP table

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Use route-map to conditionally advertise routes

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"allowas_in": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"number": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Number of occurrences of AS number

Format: u32:1-10
Number of times AS is allowed in path
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Accept route that contains the local-as in the as-path

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"attribute_unchanged": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"as_path": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Send AS path unchanged

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"med": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Send multi-exit discriminator unchanged

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"next_hop": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Send nexthop unchanged

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `BGP attributes are sent unchanged

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"disable_send_community": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"extended": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Disable sending extended community attributes to this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"standard": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Disable sending standard community attributes to this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Disable sending community attributes to this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"distribute_list": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Access-list to filter outgoing route updates to this peer-group

Format: u32:1-65535
Access-list to filter outgoing route updates to this peer-group
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Access-list to filter incoming route updates from this peer-group

Format: u32:1-65535
Access-list to filter incoming route updates from this peer-group
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Access-list to filter route updates to/from this peer-group

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"filter_list": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `As-path-list to filter outgoing route updates to this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `As-path-list to filter incoming route updates from this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `as-path-list to filter route updates to/from this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"nexthop_self": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"force": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Set the next hop to self for reflected routes

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Disable the next hop calculation for this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"route_map": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Route-map to filter outgoing route updates

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Route-map to filter incoming route updates

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"soft_reconfiguration": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"inbound": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Enable inbound soft reconfiguration

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Soft reconfiguration for peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"default_originate": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"route_map": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Specify route-map name to use

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Originate default route to this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},
						},
						// CustomType:          basetypes.MapTypable(nil),
						// Required:            false,
						Optional: true,
						// Computed:            false,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `IPv6 BGP neighbor parameters

`,
						// DeprecationMessage:  "",
						// Validators:          []validator.Map(nil),
						// PlanModifiers:       []planmodifier.Map(nil),
						// TODO investigate if node defaults can be handled
						// Default:             defaults.Map(nil),
					},

					"ipv4_labeled_unicast": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							// TODO handle non-string types
							"addpath_tx_all": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Use addpath to advertise all paths to a neighbor

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"addpath_tx_per_as": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Use addpath to advertise the bestpath per each neighboring AS

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"as_override": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Override ASN in outbound updates to configured neighbor local-as

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"maximum_prefix": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Maximum number of prefixes to accept from this peer

Format: u32:1-4294967295
Prefix limit
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"maximum_prefix_out": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Maximum number of prefixes to be sent to this peer

Format: u32:1-4294967295
Prefix limit
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"remove_private_as": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Remove private AS numbers from AS path in outbound route updates

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"route_reflector_client": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Peer is a route reflector client

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"route_server_client": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Peer is a route server client

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"unsuppress_map": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Route-map to selectively unsuppress suppressed routes

Format: txt
Route map name
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"weight": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Default weight for routes from this peer

Format: u32:1-65535
Default weight
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							"capability": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									"orf": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											"prefix_list": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													// TODO handle non-string types
													"receive": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Capability to receive the ORF

`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"send": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Capability to send the ORF

`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},
												},
												// CustomType:          basetypes.MapTypable(nil),
												// Required:            false,
												Optional: true,
												// Computed:            false,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Advertise prefix-list ORF capability to this peer

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},
										},
										// CustomType:          basetypes.MapTypable(nil),
										// Required:            false,
										Optional: true,
										// Computed:            false,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Advertise ORF capability to this peer

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Advertise capabilities to this neighbor (IPv4)

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"prefix_list": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `IPv4-Prefix-list to filter outgoing route updates to this peer

Format: txt
Name of IPv4 prefix-list
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `IPv4-Prefix-list to filter incoming route updates from this peer

Format: txt
Name of IPv4 prefix-list
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `IPv4-Prefix-list to filter route updates to/from this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"conditionally_advertise": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"advertise_map": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Route-map to conditionally advertise routes

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"exist_map": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Advertise routes only if prefixes in exist-map are installed in BGP table

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"non_exist_map": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Advertise routes only if prefixes in non-exist-map are not installed in BGP table

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Use route-map to conditionally advertise routes

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"allowas_in": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"number": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Number of occurrences of AS number

Format: u32:1-10
Number of times AS is allowed in path
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Accept route that contains the local-as in the as-path

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"attribute_unchanged": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"as_path": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Send AS path unchanged

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"med": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Send multi-exit discriminator unchanged

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"next_hop": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Send nexthop unchanged

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `BGP attributes are sent unchanged

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"disable_send_community": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"extended": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Disable sending extended community attributes to this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"standard": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Disable sending standard community attributes to this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Disable sending community attributes to this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"distribute_list": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Access-list to filter outgoing route updates to this peer-group

Format: u32:1-65535
Access-list to filter outgoing route updates to this peer-group
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Access-list to filter incoming route updates from this peer-group

Format: u32:1-65535
Access-list to filter incoming route updates from this peer-group
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Access-list to filter route updates to/from this peer-group

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"filter_list": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `As-path-list to filter outgoing route updates to this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `As-path-list to filter incoming route updates from this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `as-path-list to filter route updates to/from this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"nexthop_self": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"force": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Set the next hop to self for reflected routes

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Disable the next hop calculation for this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"route_map": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Route-map to filter outgoing route updates

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Route-map to filter incoming route updates

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"soft_reconfiguration": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"inbound": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Enable inbound soft reconfiguration

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Soft reconfiguration for peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"default_originate": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"route_map": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Specify route-map name to use

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Originate default route to this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},
						},
						// CustomType:          basetypes.MapTypable(nil),
						// Required:            false,
						Optional: true,
						// Computed:            false,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `IPv4 Labeled Unicast BGP neighbor parameters

`,
						// DeprecationMessage:  "",
						// Validators:          []validator.Map(nil),
						// PlanModifiers:       []planmodifier.Map(nil),
						// TODO investigate if node defaults can be handled
						// Default:             defaults.Map(nil),
					},

					"ipv6_labeled_unicast": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							// TODO handle non-string types
							"addpath_tx_all": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Use addpath to advertise all paths to a neighbor

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"addpath_tx_per_as": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Use addpath to advertise the bestpath per each neighboring AS

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"as_override": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Override ASN in outbound updates to configured neighbor local-as

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"maximum_prefix": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Maximum number of prefixes to accept from this peer

Format: u32:1-4294967295
Prefix limit
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"maximum_prefix_out": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Maximum number of prefixes to be sent to this peer

Format: u32:1-4294967295
Prefix limit
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"remove_private_as": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Remove private AS numbers from AS path in outbound route updates

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"route_reflector_client": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Peer is a route reflector client

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"route_server_client": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Peer is a route server client

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"unsuppress_map": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Route-map to selectively unsuppress suppressed routes

Format: txt
Route map name
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"weight": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Default weight for routes from this peer

Format: u32:1-65535
Default weight
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							"capability": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									"orf": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											"prefix_list": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													// TODO handle non-string types
													"receive": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Capability to receive the ORF

`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"send": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Capability to send the ORF

`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},
												},
												// CustomType:          basetypes.MapTypable(nil),
												// Required:            false,
												Optional: true,
												// Computed:            false,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Advertise prefix-list ORF capability to this peer

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},
										},
										// CustomType:          basetypes.MapTypable(nil),
										// Required:            false,
										Optional: true,
										// Computed:            false,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Advertise ORF capability to this peer

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Advertise capabilities to this neighbor (IPv6)

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"nexthop_local": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"unchanged": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Leave link-local nexthop unchanged for this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Nexthop attributes

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"prefix_list": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Prefix-list to filter outgoing route updates to this peer

Format: txt
Name of IPv6 prefix-list
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Prefix-list to filter incoming route updates from this peer

Format: txt
Name of IPv6 prefix-list
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Prefix-list to filter route updates to/from this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"conditionally_advertise": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"advertise_map": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Route-map to conditionally advertise routes

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"exist_map": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Advertise routes only if prefixes in exist-map are installed in BGP table

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"non_exist_map": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Advertise routes only if prefixes in non-exist-map are not installed in BGP table

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Use route-map to conditionally advertise routes

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"allowas_in": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"number": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Number of occurrences of AS number

Format: u32:1-10
Number of times AS is allowed in path
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Accept route that contains the local-as in the as-path

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"attribute_unchanged": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"as_path": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Send AS path unchanged

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"med": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Send multi-exit discriminator unchanged

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"next_hop": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Send nexthop unchanged

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `BGP attributes are sent unchanged

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"disable_send_community": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"extended": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Disable sending extended community attributes to this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"standard": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Disable sending standard community attributes to this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Disable sending community attributes to this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"distribute_list": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Access-list to filter outgoing route updates to this peer-group

Format: u32:1-65535
Access-list to filter outgoing route updates to this peer-group
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Access-list to filter incoming route updates from this peer-group

Format: u32:1-65535
Access-list to filter incoming route updates from this peer-group
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Access-list to filter route updates to/from this peer-group

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"filter_list": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `As-path-list to filter outgoing route updates to this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `As-path-list to filter incoming route updates from this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `as-path-list to filter route updates to/from this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"nexthop_self": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"force": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Set the next hop to self for reflected routes

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Disable the next hop calculation for this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"route_map": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Route-map to filter outgoing route updates

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Route-map to filter incoming route updates

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"soft_reconfiguration": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"inbound": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Enable inbound soft reconfiguration

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Soft reconfiguration for peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"default_originate": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"route_map": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Specify route-map name to use

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Originate default route to this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},
						},
						// CustomType:          basetypes.MapTypable(nil),
						// Required:            false,
						Optional: true,
						// Computed:            false,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `IPv6 Labeled Unicast BGP neighbor parameters

`,
						// DeprecationMessage:  "",
						// Validators:          []validator.Map(nil),
						// PlanModifiers:       []planmodifier.Map(nil),
						// TODO investigate if node defaults can be handled
						// Default:             defaults.Map(nil),
					},

					"ipv4_vpn": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							// TODO handle non-string types
							"addpath_tx_all": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Use addpath to advertise all paths to a neighbor

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"addpath_tx_per_as": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Use addpath to advertise the bestpath per each neighboring AS

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"as_override": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Override ASN in outbound updates to configured neighbor local-as

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"maximum_prefix": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Maximum number of prefixes to accept from this peer

Format: u32:1-4294967295
Prefix limit
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"maximum_prefix_out": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Maximum number of prefixes to be sent to this peer

Format: u32:1-4294967295
Prefix limit
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"remove_private_as": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Remove private AS numbers from AS path in outbound route updates

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"route_reflector_client": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Peer is a route reflector client

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"route_server_client": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Peer is a route server client

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"unsuppress_map": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Route-map to selectively unsuppress suppressed routes

Format: txt
Route map name
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"weight": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Default weight for routes from this peer

Format: u32:1-65535
Default weight
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							"prefix_list": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `IPv4-Prefix-list to filter outgoing route updates to this peer

Format: txt
Name of IPv4 prefix-list
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `IPv4-Prefix-list to filter incoming route updates from this peer

Format: txt
Name of IPv4 prefix-list
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `IPv4-Prefix-list to filter route updates to/from this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"conditionally_advertise": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"advertise_map": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Route-map to conditionally advertise routes

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"exist_map": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Advertise routes only if prefixes in exist-map are installed in BGP table

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"non_exist_map": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Advertise routes only if prefixes in non-exist-map are not installed in BGP table

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Use route-map to conditionally advertise routes

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"allowas_in": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"number": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Number of occurrences of AS number

Format: u32:1-10
Number of times AS is allowed in path
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Accept route that contains the local-as in the as-path

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"attribute_unchanged": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"as_path": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Send AS path unchanged

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"med": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Send multi-exit discriminator unchanged

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"next_hop": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Send nexthop unchanged

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `BGP attributes are sent unchanged

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"disable_send_community": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"extended": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Disable sending extended community attributes to this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"standard": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Disable sending standard community attributes to this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Disable sending community attributes to this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"distribute_list": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Access-list to filter outgoing route updates to this peer-group

Format: u32:1-65535
Access-list to filter outgoing route updates to this peer-group
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Access-list to filter incoming route updates from this peer-group

Format: u32:1-65535
Access-list to filter incoming route updates from this peer-group
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Access-list to filter route updates to/from this peer-group

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"filter_list": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `As-path-list to filter outgoing route updates to this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `As-path-list to filter incoming route updates from this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `as-path-list to filter route updates to/from this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"nexthop_self": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"force": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Set the next hop to self for reflected routes

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Disable the next hop calculation for this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"route_map": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Route-map to filter outgoing route updates

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Route-map to filter incoming route updates

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"soft_reconfiguration": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"inbound": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Enable inbound soft reconfiguration

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Soft reconfiguration for peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},
						},
						// CustomType:          basetypes.MapTypable(nil),
						// Required:            false,
						Optional: true,
						// Computed:            false,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `IPv4 VPN BGP neighbor parameters

`,
						// DeprecationMessage:  "",
						// Validators:          []validator.Map(nil),
						// PlanModifiers:       []planmodifier.Map(nil),
						// TODO investigate if node defaults can be handled
						// Default:             defaults.Map(nil),
					},

					"ipv6_vpn": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							// TODO handle non-string types
							"addpath_tx_all": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Use addpath to advertise all paths to a neighbor

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"addpath_tx_per_as": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Use addpath to advertise the bestpath per each neighboring AS

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"as_override": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Override ASN in outbound updates to configured neighbor local-as

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"maximum_prefix": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Maximum number of prefixes to accept from this peer

Format: u32:1-4294967295
Prefix limit
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"maximum_prefix_out": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Maximum number of prefixes to be sent to this peer

Format: u32:1-4294967295
Prefix limit
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"remove_private_as": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Remove private AS numbers from AS path in outbound route updates

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"route_reflector_client": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Peer is a route reflector client

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"route_server_client": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Peer is a route server client

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"unsuppress_map": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Route-map to selectively unsuppress suppressed routes

Format: txt
Route map name
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"weight": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Default weight for routes from this peer

Format: u32:1-65535
Default weight
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							"nexthop_local": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"unchanged": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Leave link-local nexthop unchanged for this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Nexthop attributes

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"prefix_list": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Prefix-list to filter outgoing route updates to this peer

Format: txt
Name of IPv6 prefix-list
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Prefix-list to filter incoming route updates from this peer

Format: txt
Name of IPv6 prefix-list
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Prefix-list to filter route updates to/from this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"conditionally_advertise": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"advertise_map": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Route-map to conditionally advertise routes

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"exist_map": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Advertise routes only if prefixes in exist-map are installed in BGP table

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"non_exist_map": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Advertise routes only if prefixes in non-exist-map are not installed in BGP table

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Use route-map to conditionally advertise routes

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"allowas_in": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"number": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Number of occurrences of AS number

Format: u32:1-10
Number of times AS is allowed in path
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Accept route that contains the local-as in the as-path

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"attribute_unchanged": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"as_path": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Send AS path unchanged

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"med": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Send multi-exit discriminator unchanged

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"next_hop": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Send nexthop unchanged

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `BGP attributes are sent unchanged

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"disable_send_community": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"extended": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Disable sending extended community attributes to this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"standard": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Disable sending standard community attributes to this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Disable sending community attributes to this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"distribute_list": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Access-list to filter outgoing route updates to this peer-group

Format: u32:1-65535
Access-list to filter outgoing route updates to this peer-group
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Access-list to filter incoming route updates from this peer-group

Format: u32:1-65535
Access-list to filter incoming route updates from this peer-group
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Access-list to filter route updates to/from this peer-group

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"filter_list": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `As-path-list to filter outgoing route updates to this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `As-path-list to filter incoming route updates from this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `as-path-list to filter route updates to/from this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"nexthop_self": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"force": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Set the next hop to self for reflected routes

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Disable the next hop calculation for this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"route_map": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Route-map to filter outgoing route updates

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Route-map to filter incoming route updates

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"soft_reconfiguration": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"inbound": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Enable inbound soft reconfiguration

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Soft reconfiguration for peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},
						},
						// CustomType:          basetypes.MapTypable(nil),
						// Required:            false,
						Optional: true,
						// Computed:            false,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `IPv6 VPN BGP neighbor parameters

`,
						// DeprecationMessage:  "",
						// Validators:          []validator.Map(nil),
						// PlanModifiers:       []planmodifier.Map(nil),
						// TODO investigate if node defaults can be handled
						// Default:             defaults.Map(nil),
					},

					"ipv4_flowspec": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							// TODO handle non-string types
							"route_reflector_client": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Peer is a route reflector client

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"route_server_client": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Peer is a route server client

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							"prefix_list": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `IPv4-Prefix-list to filter outgoing route updates to this peer

Format: txt
Name of IPv4 prefix-list
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `IPv4-Prefix-list to filter incoming route updates from this peer

Format: txt
Name of IPv4 prefix-list
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `IPv4-Prefix-list to filter route updates to/from this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"filter_list": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `As-path-list to filter outgoing route updates to this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `As-path-list to filter incoming route updates from this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `as-path-list to filter route updates to/from this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"route_map": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Route-map to filter outgoing route updates

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Route-map to filter incoming route updates

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"soft_reconfiguration": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"inbound": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Enable inbound soft reconfiguration

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Soft reconfiguration for peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},
						},
						// CustomType:          basetypes.MapTypable(nil),
						// Required:            false,
						Optional: true,
						// Computed:            false,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `IPv4 Flow Specification BGP neighbor parameters

`,
						// DeprecationMessage:  "",
						// Validators:          []validator.Map(nil),
						// PlanModifiers:       []planmodifier.Map(nil),
						// TODO investigate if node defaults can be handled
						// Default:             defaults.Map(nil),
					},

					"ipv6_flowspec": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							// TODO handle non-string types
							"route_reflector_client": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Peer is a route reflector client

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"route_server_client": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Peer is a route server client

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							"prefix_list": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Prefix-list to filter outgoing route updates to this peer

Format: txt
Name of IPv6 prefix-list
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Prefix-list to filter incoming route updates from this peer

Format: txt
Name of IPv6 prefix-list
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Prefix-list to filter route updates to/from this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"filter_list": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `As-path-list to filter outgoing route updates to this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `As-path-list to filter incoming route updates from this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `as-path-list to filter route updates to/from this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"route_map": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Route-map to filter outgoing route updates

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Route-map to filter incoming route updates

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"soft_reconfiguration": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"inbound": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Enable inbound soft reconfiguration

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Soft reconfiguration for peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},
						},
						// CustomType:          basetypes.MapTypable(nil),
						// Required:            false,
						Optional: true,
						// Computed:            false,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `IPv6 Flow Specification BGP neighbor parameters

`,
						// DeprecationMessage:  "",
						// Validators:          []validator.Map(nil),
						// PlanModifiers:       []planmodifier.Map(nil),
						// TODO investigate if node defaults can be handled
						// Default:             defaults.Map(nil),
					},

					"ipv4_multicast": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							// TODO handle non-string types
							"addpath_tx_all": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Use addpath to advertise all paths to a neighbor

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"addpath_tx_per_as": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Use addpath to advertise the bestpath per each neighboring AS

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"as_override": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Override ASN in outbound updates to configured neighbor local-as

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"maximum_prefix": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Maximum number of prefixes to accept from this peer

Format: u32:1-4294967295
Prefix limit
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"maximum_prefix_out": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Maximum number of prefixes to be sent to this peer

Format: u32:1-4294967295
Prefix limit
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"remove_private_as": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Remove private AS numbers from AS path in outbound route updates

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"route_reflector_client": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Peer is a route reflector client

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"route_server_client": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Peer is a route server client

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"unsuppress_map": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Route-map to selectively unsuppress suppressed routes

Format: txt
Route map name
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"weight": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Default weight for routes from this peer

Format: u32:1-65535
Default weight
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							"capability": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									"orf": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											"prefix_list": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													// TODO handle non-string types
													"receive": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Capability to receive the ORF

`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"send": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Capability to send the ORF

`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},
												},
												// CustomType:          basetypes.MapTypable(nil),
												// Required:            false,
												Optional: true,
												// Computed:            false,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Advertise prefix-list ORF capability to this peer

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},
										},
										// CustomType:          basetypes.MapTypable(nil),
										// Required:            false,
										Optional: true,
										// Computed:            false,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Advertise ORF capability to this peer

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Advertise capabilities to this neighbor (IPv4)

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"prefix_list": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `IPv4-Prefix-list to filter outgoing route updates to this peer

Format: txt
Name of IPv4 prefix-list
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `IPv4-Prefix-list to filter incoming route updates from this peer

Format: txt
Name of IPv4 prefix-list
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `IPv4-Prefix-list to filter route updates to/from this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"conditionally_advertise": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"advertise_map": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Route-map to conditionally advertise routes

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"exist_map": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Advertise routes only if prefixes in exist-map are installed in BGP table

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"non_exist_map": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Advertise routes only if prefixes in non-exist-map are not installed in BGP table

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Use route-map to conditionally advertise routes

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"allowas_in": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"number": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Number of occurrences of AS number

Format: u32:1-10
Number of times AS is allowed in path
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Accept route that contains the local-as in the as-path

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"attribute_unchanged": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"as_path": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Send AS path unchanged

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"med": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Send multi-exit discriminator unchanged

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"next_hop": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Send nexthop unchanged

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `BGP attributes are sent unchanged

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"disable_send_community": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"extended": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Disable sending extended community attributes to this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"standard": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Disable sending standard community attributes to this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Disable sending community attributes to this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"distribute_list": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Access-list to filter outgoing route updates to this peer-group

Format: u32:1-65535
Access-list to filter outgoing route updates to this peer-group
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Access-list to filter incoming route updates from this peer-group

Format: u32:1-65535
Access-list to filter incoming route updates from this peer-group
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Access-list to filter route updates to/from this peer-group

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"filter_list": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `As-path-list to filter outgoing route updates to this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `As-path-list to filter incoming route updates from this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `as-path-list to filter route updates to/from this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"nexthop_self": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"force": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Set the next hop to self for reflected routes

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Disable the next hop calculation for this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"route_map": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Route-map to filter outgoing route updates

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Route-map to filter incoming route updates

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"soft_reconfiguration": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"inbound": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Enable inbound soft reconfiguration

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Soft reconfiguration for peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"default_originate": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"route_map": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Specify route-map name to use

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Originate default route to this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},
						},
						// CustomType:          basetypes.MapTypable(nil),
						// Required:            false,
						Optional: true,
						// Computed:            false,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `IPv4 Multicast BGP neighbor parameters

`,
						// DeprecationMessage:  "",
						// Validators:          []validator.Map(nil),
						// PlanModifiers:       []planmodifier.Map(nil),
						// TODO investigate if node defaults can be handled
						// Default:             defaults.Map(nil),
					},

					"ipv6_multicast": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							// TODO handle non-string types
							"addpath_tx_all": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Use addpath to advertise all paths to a neighbor

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"addpath_tx_per_as": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Use addpath to advertise the bestpath per each neighboring AS

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"as_override": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Override ASN in outbound updates to configured neighbor local-as

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"maximum_prefix": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Maximum number of prefixes to accept from this peer

Format: u32:1-4294967295
Prefix limit
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"maximum_prefix_out": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Maximum number of prefixes to be sent to this peer

Format: u32:1-4294967295
Prefix limit
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"remove_private_as": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Remove private AS numbers from AS path in outbound route updates

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"route_reflector_client": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Peer is a route reflector client

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"route_server_client": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Peer is a route server client

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"unsuppress_map": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Route-map to selectively unsuppress suppressed routes

Format: txt
Route map name
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"weight": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Default weight for routes from this peer

Format: u32:1-65535
Default weight
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							"nexthop_local": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"unchanged": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Leave link-local nexthop unchanged for this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Nexthop attributes

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"prefix_list": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Prefix-list to filter outgoing route updates to this peer

Format: txt
Name of IPv6 prefix-list
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Prefix-list to filter incoming route updates from this peer

Format: txt
Name of IPv6 prefix-list
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Prefix-list to filter route updates to/from this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"conditionally_advertise": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"advertise_map": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Route-map to conditionally advertise routes

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"exist_map": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Advertise routes only if prefixes in exist-map are installed in BGP table

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"non_exist_map": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Advertise routes only if prefixes in non-exist-map are not installed in BGP table

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Use route-map to conditionally advertise routes

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"allowas_in": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"number": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Number of occurrences of AS number

Format: u32:1-10
Number of times AS is allowed in path
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Accept route that contains the local-as in the as-path

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"attribute_unchanged": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"as_path": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Send AS path unchanged

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"med": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Send multi-exit discriminator unchanged

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"next_hop": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Send nexthop unchanged

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `BGP attributes are sent unchanged

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"disable_send_community": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"extended": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Disable sending extended community attributes to this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"standard": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Disable sending standard community attributes to this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Disable sending community attributes to this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"distribute_list": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Access-list to filter outgoing route updates to this peer-group

Format: u32:1-65535
Access-list to filter outgoing route updates to this peer-group
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Access-list to filter incoming route updates from this peer-group

Format: u32:1-65535
Access-list to filter incoming route updates from this peer-group
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Access-list to filter route updates to/from this peer-group

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"filter_list": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `As-path-list to filter outgoing route updates to this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `As-path-list to filter incoming route updates from this peer

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `as-path-list to filter route updates to/from this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"nexthop_self": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"force": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Set the next hop to self for reflected routes

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Disable the next hop calculation for this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"route_map": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Route-map to filter outgoing route updates

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Route-map to filter incoming route updates

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"soft_reconfiguration": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"inbound": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Enable inbound soft reconfiguration

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Soft reconfiguration for peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"default_originate": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"route_map": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Specify route-map name to use

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Originate default route to this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},
						},
						// CustomType:          basetypes.MapTypable(nil),
						// Required:            false,
						Optional: true,
						// Computed:            false,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `IPv6 Multicast BGP neighbor parameters

`,
						// DeprecationMessage:  "",
						// Validators:          []validator.Map(nil),
						// PlanModifiers:       []planmodifier.Map(nil),
						// TODO investigate if node defaults can be handled
						// Default:             defaults.Map(nil),
					},

					"l2vpn_evpn": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							// TODO handle non-string types
							"route_reflector_client": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Peer is a route reflector client

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"route_server_client": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Peer is a route server client

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							"allowas_in": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"number": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Number of occurrences of AS number

Format: u32:1-10
Number of times AS is allowed in path
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Accept route that contains the local-as in the as-path

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"attribute_unchanged": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"as_path": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Send AS path unchanged

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"med": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Send multi-exit discriminator unchanged

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"next_hop": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Send nexthop unchanged

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `BGP attributes are sent unchanged

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"nexthop_self": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"force": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Set the next hop to self for reflected routes

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Disable the next hop calculation for this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"route_map": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"export": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Route-map to filter outgoing route updates

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"import": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Route-map to filter incoming route updates

Format: txt
Route map name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"soft_reconfiguration": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"inbound": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Enable inbound soft reconfiguration

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Soft reconfiguration for peer

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},
						},
						// CustomType:          basetypes.MapTypable(nil),
						// Required:            false,
						Optional: true,
						// Computed:            false,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `L2VPN EVPN BGP settings

`,
						// DeprecationMessage:  "",
						// Validators:          []validator.Map(nil),
						// PlanModifiers:       []planmodifier.Map(nil),
						// TODO investigate if node defaults can be handled
						// Default:             defaults.Map(nil),
					},
				},
				// CustomType:          basetypes.MapTypable(nil),
				// Required:            false,
				Optional: true,
				// Computed:            false,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Address-family parameters

`,
				// DeprecationMessage:  "",
				// Validators:          []validator.Map(nil),
				// PlanModifiers:       []planmodifier.Map(nil),
				// TODO investigate if node defaults can be handled
				// Default:             defaults.Map(nil),
			},

			"bfd": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					// TODO handle non-string types
					"profile": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Use settings from BFD profile

Format: txt
BFD profile name
`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},

					// TODO handle non-string types
					"check_control_plane_failure": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Allow to write CBIT independence in BFD outgoing packets and read both C-BIT value of BFD and lookup BGP peer status

`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},
				},
				// CustomType:          basetypes.MapTypable(nil),
				// Required:            false,
				Optional: true,
				// Computed:            false,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Enable Bidirectional Forwarding Detection (BFD) support

`,
				// DeprecationMessage:  "",
				// Validators:          []validator.Map(nil),
				// PlanModifiers:       []planmodifier.Map(nil),
				// TODO investigate if node defaults can be handled
				// Default:             defaults.Map(nil),
			},

			"capability": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					// TODO handle non-string types
					"dynamic": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Advertise dynamic capability to this neighbor

`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},

					// TODO handle non-string types
					"extended_nexthop": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Advertise extended-nexthop capability to this neighbor

`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},
				},
				// CustomType:          basetypes.MapTypable(nil),
				// Required:            false,
				Optional: true,
				// Computed:            false,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Advertise capabilities to this peer-group

`,
				// DeprecationMessage:  "",
				// Validators:          []validator.Map(nil),
				// PlanModifiers:       []planmodifier.Map(nil),
				// TODO investigate if node defaults can be handled
				// Default:             defaults.Map(nil),
			},

			"interface": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					// TODO handle non-string types
					"peer_group": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Peer group for this peer

Format: txt
Peer-group name
`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},

					// TODO handle non-string types
					"remote_as": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Neighbor BGP AS number

Format: u32:1-4294967294
Neighbor AS number
Format: external
Any AS different from the local AS
Format: internal
Neighbor AS number
`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},

					// TODO handle non-string types
					"source_interface": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Interface used to establish connection

Format: interface
Interface name
`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},

					"v6only": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							// TODO handle non-string types
							"peer_group": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Peer group for this peer

Format: txt
Peer-group name
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"remote_as": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Neighbor BGP AS number

Format: u32:1-4294967294
Neighbor AS number
Format: external
Any AS different from the local AS
Format: internal
Neighbor AS number
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},
						},
						// CustomType:          basetypes.MapTypable(nil),
						// Required:            false,
						Optional: true,
						// Computed:            false,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Enable BGP with v6 link-local only

`,
						// DeprecationMessage:  "",
						// Validators:          []validator.Map(nil),
						// PlanModifiers:       []planmodifier.Map(nil),
						// TODO investigate if node defaults can be handled
						// Default:             defaults.Map(nil),
					},
				},
				// CustomType:          basetypes.MapTypable(nil),
				// Required:            false,
				Optional: true,
				// Computed:            false,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Interface parameters

`,
				// DeprecationMessage:  "",
				// Validators:          []validator.Map(nil),
				// PlanModifiers:       []planmodifier.Map(nil),
				// TODO investigate if node defaults can be handled
				// Default:             defaults.Map(nil),
			},

			"timers": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					// TODO handle non-string types
					"connect": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `BGP connect timer for this neighbor

Format: u32:1-65535
Connect timer in seconds
Format: 0
Disable connect timer
`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},

					// TODO handle non-string types
					"holdtime": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `BGP hold timer for this neighbor

Format: u32:1-65535
Hold timer in seconds
Format: 0
Hold timer disabled
`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},

					// TODO handle non-string types
					"keepalive": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `BGP keepalive interval for this neighbor

Format: u32:1-65535
Keepalive interval in seconds
`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},
				},
				// CustomType:          basetypes.MapTypable(nil),
				// Required:            false,
				Optional: true,
				// Computed:            false,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Neighbor timers

`,
				// DeprecationMessage:  "",
				// Validators:          []validator.Map(nil),
				// PlanModifiers:       []planmodifier.Map(nil),
				// TODO investigate if node defaults can be handled
				// Default:             defaults.Map(nil),
			},

			"ttl_security": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					// TODO handle non-string types
					"hops": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Number of the maximum number of hops to the BGP peer

Format: u32:1-254
Number of hops
`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},
				},
				// CustomType:          basetypes.MapTypable(nil),
				// Required:            false,
				Optional: true,
				// Computed:            false,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Ttl security mechanism

`,
				// DeprecationMessage:  "",
				// Validators:          []validator.Map(nil),
				// PlanModifiers:       []planmodifier.Map(nil),
				// TODO investigate if node defaults can be handled
				// Default:             defaults.Map(nil),
			},
		},
	}
}

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func (r *protocols_bgp_neighbor) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *protocols_bgp_neighborModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create example, got error: %s", err))
	//     return
	// }

	// For the purposes of this example code, hardcoding a response value to
	// save into the Terraform state.
	data.ID = types.StringValue("example-id")

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "created a resource")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read method to define the logic which refreshes the Terraform state for the resource.
func (r *protocols_bgp_neighbor) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *protocols_bgp_neighborModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read example, got error: %s", err))
	//     return
	// }

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update method to define the logic which updates the resource and sets the updated Terraform state on success.
func (r *protocols_bgp_neighbor) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *protocols_bgp_neighborModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update example, got error: %s", err))
	//     return
	// }

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete method to define the logic which deletes the resource and removes the Terraform state on success.
func (r *protocols_bgp_neighbor) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *protocols_bgp_neighborModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete example, got error: %s", err))
	//     return
	// }
}
