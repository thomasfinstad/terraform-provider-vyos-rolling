package resourcefull

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &vrf_name{}

// var _ resource.ResourceWithImportState = &vrf_name{}

// vrf_name defines the resource implementation.
type vrf_name struct {
	client *http.Client
}

// vrf nameModel describes the resource data model.
type vrf_nameModel struct {
	ConfigurableAttribute types.String `tfsdk:"configurable_attribute"`
	Defaulted             types.String `tfsdk:"defaulted"`
	ID                    types.String `tfsdk:"id"`

	Ip types.String `tfsdk:ip`

	Ipv_ types.String `tfsdk:ipv6`

	Protocols types.String `tfsdk:protocols`

	Description types.String `tfsdk:description`

	Disable types.String `tfsdk:disable`

	Table types.String `tfsdk:table`

	Vni types.String `tfsdk:vni`
}

// Metadata method to define the resource type name.
func (r *vrf_name) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vrf name"
}

// vrf_nameResource method to return the example resource reference
func vrf_nameResource() resource.Resource {
	return &vrf_name{}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r *vrf_name) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example resource",

		Attributes: map[string]schema.Attribute{
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
			"disable": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Administratively disable interface

`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"table": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Routing table associated with this instance

Format: u32:100-65535
Routing table ID
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"vni": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Virtual Network Identifier

Format: u32:0-16777214
VXLAN virtual network identifier
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			"ip": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					// TODO handle non-string types
					"disable_forwarding": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Disable IP forwarding on this interface

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
				MarkdownDescription: `IPv4 routing parameters

`,
				// DeprecationMessage:  "",
				// Validators:          []validator.Map(nil),
				// PlanModifiers:       []planmodifier.Map(nil),
				// TODO investigate if node defaults can be handled
				// Default:             defaults.Map(nil),
			},

			"ipv6": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					// TODO handle non-string types
					"disable_forwarding": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Disable IP forwarding on this interface

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
				MarkdownDescription: `IPv6 routing parameters

`,
				// DeprecationMessage:  "",
				// Validators:          []validator.Map(nil),
				// PlanModifiers:       []planmodifier.Map(nil),
				// TODO investigate if node defaults can be handled
				// Default:             defaults.Map(nil),
			},

			"protocols": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"bgp": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"neighbor": schema.MapNestedAttribute{
								NestedObject: schema.NestedAttributeObject{
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
								MarkdownDescription: `BGP neighbor

Format: ipv4
BGP neighbor IP address
Format: ipv6
BGP neighbor IPv6 address
Format: txt
Interface name
`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if tagnode defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"peer_group": schema.MapNestedAttribute{
								NestedObject: schema.NestedAttributeObject{
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
								MarkdownDescription: `Name of peer-group

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if tagnode defaults can be handled
								// Default:             defaults.Map(nil),
							},

							// TODO handle non-string types
							"system_as": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Autonomous System Number (ASN)

Format: u32:1-4294967294
Autonomous System Number
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

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

							"address_family": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									"ipv4_unicast": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											"aggregate_address": schema.MapNestedAttribute{
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														// TODO handle non-string types
														"as_set": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Generate AS-set path information for this aggregate address

`,
															// DeprecationMessage:  "",
															// TODO Recreate some of vyos validators for use in leafnodes
															// Validators:          []validator.String(nil),
															// PlanModifiers:       []planmodifier.String(nil),

														},

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

														// TODO handle non-string types
														"summary_only": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Announce the aggregate summary network only

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
												MarkdownDescription: `BGP aggregate network

Format: ipv4net
BGP aggregate network
`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if tagnode defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"network": schema.MapNestedAttribute{
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														// TODO handle non-string types
														"backdoor": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Network as a backdoor route

`,
															// DeprecationMessage:  "",
															// TODO Recreate some of vyos validators for use in leafnodes
															// Validators:          []validator.String(nil),
															// PlanModifiers:       []planmodifier.String(nil),

														},

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
												MarkdownDescription: `BGP network

Format: ipv4net
BGP network
`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if tagnode defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"distance": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													"prefix": schema.MapNestedAttribute{
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																// TODO handle non-string types
																"distance": schema.StringAttribute{
																	// CustomType:          basetypes.StringTypable(nil),
																	// Required:            false,
																	Optional: true,
																	// Sensitive:           false,
																	// Description:         "",
																	MarkdownDescription: `Administrative distance for prefix

Format: u32:1-255
Administrative distance for external BGP routes
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
														MarkdownDescription: `Administrative distance for a specific BGP prefix

Format: ipv4net
Administrative distance for a specific BGP prefix
`,
														// DeprecationMessage:  "",
														// Validators:          []validator.Map(nil),
														// PlanModifiers:       []planmodifier.Map(nil),
														// TODO investigate if tagnode defaults can be handled
														// Default:             defaults.Map(nil),
													},

													// TODO handle non-string types
													"external": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `eBGP routes administrative distance

Format: u32:1-255
eBGP routes administrative distance
`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"internal": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `iBGP routes administrative distance

Format: u32:1-255
iBGP routes administrative distance
`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"local": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Locally originated BGP routes administrative distance

Format: u32:1-255
Locally originated BGP routes administrative distance
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
												MarkdownDescription: `Administrative distances for BGP routes

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"export": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													// TODO handle non-string types
													"vpn": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `to/from default instance VPN RIB

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
												MarkdownDescription: `Export routes from this address-family

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"import": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													// TODO handle non-string types
													"vpn": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `to/from default instance VPN RIB

`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"vrf": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `VRF to import from

Format: txt
VRF instance name
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
												MarkdownDescription: `Import routes to this address-family

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"label": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													"vpn": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"export": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `For routes leaked from current address-family to VPN

Format: auto
Automatically assign a label
Format: u32:0-1048575
Label Value
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
														MarkdownDescription: `Between current address-family and VPN

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
												MarkdownDescription: `Label value for VRF

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"maximum_paths": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													// TODO handle non-string types
													"ebgp": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `eBGP maximum paths

Format: u32:1-256
Number of paths to consider
`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"ibgp": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `iBGP maximum paths

Format: u32:1-256
Number of paths to consider
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
												MarkdownDescription: `Forward packets over multiple paths

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"rd": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													"vpn": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"export": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `For routes leaked from current address-family to VPN

Format: ASN:NN_OR_IP-ADDRESS:NN
Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)
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
														MarkdownDescription: `Between current address-family and VPN

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
												MarkdownDescription: `Specify route distinguisher

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"route_map": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													"vpn": schema.SingleNestedAttribute{
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
														MarkdownDescription: `Between current address-family and VPN

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
												MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"route_target": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													"vpn": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"both": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Route Target both import and export

Format: txt
Space separated route target list (A.B.C.D:MN|EF:OPQR|GHJK:MN)
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
																MarkdownDescription: `Route Target import

Format: txt
Space separated route target list (A.B.C.D:MN|EF:OPQR|GHJK:MN)
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

															// TODO handle non-string types
															"export": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Route Target export

Format: txt
Space separated route target list (A.B.C.D:MN|EF:OPQR|GHJK:MN)
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
														MarkdownDescription: `Between current address-family and VPN

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
												MarkdownDescription: `Specify route target list

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"redistribute": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													// TODO handle non-string types
													"table": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Redistribute non-main Kernel Routing Table

`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													"connected": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Metric for redistributed routes

Format: u32:1-4294967295
Metric for redistributed routes
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute connected routes into BGP

`,
														// DeprecationMessage:  "",
														// Validators:          []validator.Map(nil),
														// PlanModifiers:       []planmodifier.Map(nil),
														// TODO investigate if node defaults can be handled
														// Default:             defaults.Map(nil),
													},

													"isis": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Metric for redistributed routes

Format: u32:1-4294967295
Metric for redistributed routes
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute IS-IS routes into BGP

`,
														// DeprecationMessage:  "",
														// Validators:          []validator.Map(nil),
														// PlanModifiers:       []planmodifier.Map(nil),
														// TODO investigate if node defaults can be handled
														// Default:             defaults.Map(nil),
													},

													"kernel": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Metric for redistributed routes

Format: u32:1-4294967295
Metric for redistributed routes
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute kernel routes into BGP

`,
														// DeprecationMessage:  "",
														// Validators:          []validator.Map(nil),
														// PlanModifiers:       []planmodifier.Map(nil),
														// TODO investigate if node defaults can be handled
														// Default:             defaults.Map(nil),
													},

													"ospf": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Metric for redistributed routes

Format: u32:1-4294967295
Metric for redistributed routes
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute OSPF routes into BGP

`,
														// DeprecationMessage:  "",
														// Validators:          []validator.Map(nil),
														// PlanModifiers:       []planmodifier.Map(nil),
														// TODO investigate if node defaults can be handled
														// Default:             defaults.Map(nil),
													},

													"rip": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Metric for redistributed routes

Format: u32:1-4294967295
Metric for redistributed routes
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute RIP routes into BGP

`,
														// DeprecationMessage:  "",
														// Validators:          []validator.Map(nil),
														// PlanModifiers:       []planmodifier.Map(nil),
														// TODO investigate if node defaults can be handled
														// Default:             defaults.Map(nil),
													},

													"babel": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Metric for redistributed routes

Format: u32:1-4294967295
Metric for redistributed routes
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute Babel routes into BGP

`,
														// DeprecationMessage:  "",
														// Validators:          []validator.Map(nil),
														// PlanModifiers:       []planmodifier.Map(nil),
														// TODO investigate if node defaults can be handled
														// Default:             defaults.Map(nil),
													},

													"static": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Metric for redistributed routes

Format: u32:1-4294967295
Metric for redistributed routes
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute static routes into BGP

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
												MarkdownDescription: `Redistribute routes from other protocols into BGP

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
										MarkdownDescription: `IPv4 BGP settings

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},

									"ipv4_multicast": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											"aggregate_address": schema.MapNestedAttribute{
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														// TODO handle non-string types
														"as_set": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Generate AS-set path information for this aggregate address

`,
															// DeprecationMessage:  "",
															// TODO Recreate some of vyos validators for use in leafnodes
															// Validators:          []validator.String(nil),
															// PlanModifiers:       []planmodifier.String(nil),

														},

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

														// TODO handle non-string types
														"summary_only": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Announce the aggregate summary network only

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
												MarkdownDescription: `BGP aggregate network/prefix

Format: ipv4net
BGP aggregate network/prefix
`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if tagnode defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"network": schema.MapNestedAttribute{
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														// TODO handle non-string types
														"backdoor": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Use BGP network/prefix as a backdoor route

`,
															// DeprecationMessage:  "",
															// TODO Recreate some of vyos validators for use in leafnodes
															// Validators:          []validator.String(nil),
															// PlanModifiers:       []planmodifier.String(nil),

														},

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
												MarkdownDescription: `Import BGP network/prefix into multicast IPv4 RIB

Format: ipv4net
Multicast IPv4 BGP network/prefix
`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if tagnode defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"distance": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													"prefix": schema.MapNestedAttribute{
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																// TODO handle non-string types
																"distance": schema.StringAttribute{
																	// CustomType:          basetypes.StringTypable(nil),
																	// Required:            false,
																	Optional: true,
																	// Sensitive:           false,
																	// Description:         "",
																	MarkdownDescription: `Administrative distance for prefix

Format: u32:1-255
Administrative distance for external BGP routes
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
														MarkdownDescription: `Administrative distance for a specific BGP prefix

Format: ipv4net
Administrative distance for a specific BGP prefix
`,
														// DeprecationMessage:  "",
														// Validators:          []validator.Map(nil),
														// PlanModifiers:       []planmodifier.Map(nil),
														// TODO investigate if tagnode defaults can be handled
														// Default:             defaults.Map(nil),
													},

													// TODO handle non-string types
													"external": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `eBGP routes administrative distance

Format: u32:1-255
eBGP routes administrative distance
`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"internal": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `iBGP routes administrative distance

Format: u32:1-255
iBGP routes administrative distance
`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"local": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Locally originated BGP routes administrative distance

Format: u32:1-255
Locally originated BGP routes administrative distance
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
												MarkdownDescription: `Administrative distances for BGP routes

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
										MarkdownDescription: `Multicast IPv4 BGP settings

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},

									"ipv4_labeled_unicast": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											"aggregate_address": schema.MapNestedAttribute{
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														// TODO handle non-string types
														"as_set": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Generate AS-set path information for this aggregate address

`,
															// DeprecationMessage:  "",
															// TODO Recreate some of vyos validators for use in leafnodes
															// Validators:          []validator.String(nil),
															// PlanModifiers:       []planmodifier.String(nil),

														},

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

														// TODO handle non-string types
														"summary_only": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Announce the aggregate summary network only

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
												MarkdownDescription: `BGP aggregate network/prefix

Format: ipv4net
BGP aggregate network/prefix
`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if tagnode defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"network": schema.MapNestedAttribute{
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														// TODO handle non-string types
														"backdoor": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Use BGP network/prefix as a backdoor route

`,
															// DeprecationMessage:  "",
															// TODO Recreate some of vyos validators for use in leafnodes
															// Validators:          []validator.String(nil),
															// PlanModifiers:       []planmodifier.String(nil),

														},

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
												MarkdownDescription: `Import BGP network/prefix into labeled unicast IPv4 RIB

Format: ipv4net
Labeled Unicast IPv4 BGP network/prefix
`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if tagnode defaults can be handled
												// Default:             defaults.Map(nil),
											},
										},
										// CustomType:          basetypes.MapTypable(nil),
										// Required:            false,
										Optional: true,
										// Computed:            false,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Labeled Unicast IPv4 BGP settings

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},

									"ipv4_flowspec": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											"local_install": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													// TODO handle non-string types
													"interface": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Interface

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
												MarkdownDescription: `Apply local policy routing to interface

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
										MarkdownDescription: `Flowspec IPv4 BGP settings

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},

									"ipv4_vpn": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											"network": schema.MapNestedAttribute{
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														// TODO handle non-string types
														"rd": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Route Distinguisher

Format: ASN:NN_OR_IP-ADDRESS:NN
Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)
`,
															// DeprecationMessage:  "",
															// TODO Recreate some of vyos validators for use in leafnodes
															// Validators:          []validator.String(nil),
															// PlanModifiers:       []planmodifier.String(nil),

														},

														// TODO handle non-string types
														"label": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `MPLS label value assigned to route

Format: u32:0-1048575
MPLS label value
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
												MarkdownDescription: `Import BGP network/prefix into unicast VPN IPv4 RIB

Format: ipv4net
Unicast VPN IPv4 BGP network/prefix
`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if tagnode defaults can be handled
												// Default:             defaults.Map(nil),
											},
										},
										// CustomType:          basetypes.MapTypable(nil),
										// Required:            false,
										Optional: true,
										// Computed:            false,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Unicast VPN IPv4 BGP settings

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},

									"ipv6_unicast": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											"aggregate_address": schema.MapNestedAttribute{
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														// TODO handle non-string types
														"as_set": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Generate AS-set path information for this aggregate address

`,
															// DeprecationMessage:  "",
															// TODO Recreate some of vyos validators for use in leafnodes
															// Validators:          []validator.String(nil),
															// PlanModifiers:       []planmodifier.String(nil),

														},

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

														// TODO handle non-string types
														"summary_only": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Announce the aggregate summary network only

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
												MarkdownDescription: `BGP aggregate network

Format: ipv6net
Aggregate network
`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if tagnode defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"network": schema.MapNestedAttribute{
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														// TODO handle non-string types
														"path_limit": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `AS-path hopcount limit

Format: u32:0-255
AS path hop count limit
`,
															// DeprecationMessage:  "",
															// TODO Recreate some of vyos validators for use in leafnodes
															// Validators:          []validator.String(nil),
															// PlanModifiers:       []planmodifier.String(nil),

														},

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
												MarkdownDescription: `BGP network

Format: ipv6net
Aggregate network
`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if tagnode defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"distance": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													"prefix": schema.MapNestedAttribute{
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																// TODO handle non-string types
																"distance": schema.StringAttribute{
																	// CustomType:          basetypes.StringTypable(nil),
																	// Required:            false,
																	Optional: true,
																	// Sensitive:           false,
																	// Description:         "",
																	MarkdownDescription: `Administrative distance for prefix

Format: u32:1-255
Administrative distance for external BGP routes
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
														MarkdownDescription: `Administrative distance for a specific BGP prefix

Format: ipv6net
Administrative distance for a specific BGP prefix
`,
														// DeprecationMessage:  "",
														// Validators:          []validator.Map(nil),
														// PlanModifiers:       []planmodifier.Map(nil),
														// TODO investigate if tagnode defaults can be handled
														// Default:             defaults.Map(nil),
													},

													// TODO handle non-string types
													"external": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `eBGP routes administrative distance

Format: u32:1-255
eBGP routes administrative distance
`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"internal": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `iBGP routes administrative distance

Format: u32:1-255
iBGP routes administrative distance
`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"local": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Locally originated BGP routes administrative distance

Format: u32:1-255
Locally originated BGP routes administrative distance
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
												MarkdownDescription: `Administrative distances for BGP routes

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"export": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													// TODO handle non-string types
													"vpn": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `to/from default instance VPN RIB

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
												MarkdownDescription: `Export routes from this address-family

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"import": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													// TODO handle non-string types
													"vpn": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `to/from default instance VPN RIB

`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"vrf": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `VRF to import from

Format: txt
VRF instance name
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
												MarkdownDescription: `Import routes to this address-family

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"label": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													"vpn": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"export": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `For routes leaked from current address-family to VPN

Format: auto
Automatically assign a label
Format: u32:0-1048575
Label Value
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
														MarkdownDescription: `Between current address-family and VPN

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
												MarkdownDescription: `Label value for VRF

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"maximum_paths": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													// TODO handle non-string types
													"ebgp": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `eBGP maximum paths

Format: u32:1-256
Number of paths to consider
`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"ibgp": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `iBGP maximum paths

Format: u32:1-256
Number of paths to consider
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
												MarkdownDescription: `Forward packets over multiple paths

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"rd": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													"vpn": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"export": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `For routes leaked from current address-family to VPN

Format: ASN:NN_OR_IP-ADDRESS:NN
Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)
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
														MarkdownDescription: `Between current address-family and VPN

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
												MarkdownDescription: `Specify route distinguisher

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"route_map": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													"vpn": schema.SingleNestedAttribute{
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
														MarkdownDescription: `Between current address-family and VPN

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
												MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"route_target": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													"vpn": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"both": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Route Target both import and export

Format: txt
Space separated route target list (A.B.C.D:MN|EF:OPQR|GHJK:MN)
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
																MarkdownDescription: `Route Target import

Format: txt
Space separated route target list (A.B.C.D:MN|EF:OPQR|GHJK:MN)
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

															// TODO handle non-string types
															"export": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Route Target export

Format: txt
Space separated route target list (A.B.C.D:MN|EF:OPQR|GHJK:MN)
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
														MarkdownDescription: `Between current address-family and VPN

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
												MarkdownDescription: `Specify route target list

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"redistribute": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													// TODO handle non-string types
													"table": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Redistribute non-main Kernel Routing Table

`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													"connected": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Metric for redistributed routes

Format: u32:1-4294967295
Metric for redistributed routes
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute connected routes into BGP

`,
														// DeprecationMessage:  "",
														// Validators:          []validator.Map(nil),
														// PlanModifiers:       []planmodifier.Map(nil),
														// TODO investigate if node defaults can be handled
														// Default:             defaults.Map(nil),
													},

													"kernel": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Metric for redistributed routes

Format: u32:1-4294967295
Metric for redistributed routes
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute kernel routes into BGP

`,
														// DeprecationMessage:  "",
														// Validators:          []validator.Map(nil),
														// PlanModifiers:       []planmodifier.Map(nil),
														// TODO investigate if node defaults can be handled
														// Default:             defaults.Map(nil),
													},

													"ospfv3": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Metric for redistributed routes

Format: u32:1-4294967295
Metric for redistributed routes
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute OSPFv3 routes into BGP

`,
														// DeprecationMessage:  "",
														// Validators:          []validator.Map(nil),
														// PlanModifiers:       []planmodifier.Map(nil),
														// TODO investigate if node defaults can be handled
														// Default:             defaults.Map(nil),
													},

													"ripng": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Metric for redistributed routes

Format: u32:1-4294967295
Metric for redistributed routes
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute RIPng routes into BGP

`,
														// DeprecationMessage:  "",
														// Validators:          []validator.Map(nil),
														// PlanModifiers:       []planmodifier.Map(nil),
														// TODO investigate if node defaults can be handled
														// Default:             defaults.Map(nil),
													},

													"babel": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Metric for redistributed routes

Format: u32:1-4294967295
Metric for redistributed routes
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute Babel routes into BGP

`,
														// DeprecationMessage:  "",
														// Validators:          []validator.Map(nil),
														// PlanModifiers:       []planmodifier.Map(nil),
														// TODO investigate if node defaults can be handled
														// Default:             defaults.Map(nil),
													},

													"static": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Metric for redistributed routes

Format: u32:1-4294967295
Metric for redistributed routes
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute static routes into BGP

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
												MarkdownDescription: `Redistribute routes from other protocols into BGP

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
										MarkdownDescription: `IPv6 BGP settings

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},

									"ipv6_multicast": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											"aggregate_address": schema.MapNestedAttribute{
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														// TODO handle non-string types
														"as_set": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Generate AS-set path information for this aggregate address

`,
															// DeprecationMessage:  "",
															// TODO Recreate some of vyos validators for use in leafnodes
															// Validators:          []validator.String(nil),
															// PlanModifiers:       []planmodifier.String(nil),

														},

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

														// TODO handle non-string types
														"summary_only": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Announce the aggregate summary network only

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
												MarkdownDescription: `BGP aggregate network/prefix

Format: ipv6net
BGP aggregate network/prefix
`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if tagnode defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"network": schema.MapNestedAttribute{
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														// TODO handle non-string types
														"path_limit": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `AS-path hopcount limit

Format: u32:0-255
AS path hop count limit
`,
															// DeprecationMessage:  "",
															// TODO Recreate some of vyos validators for use in leafnodes
															// Validators:          []validator.String(nil),
															// PlanModifiers:       []planmodifier.String(nil),

														},

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
												MarkdownDescription: `Import BGP network/prefix into multicast IPv6 RIB

Format: ipv6net
Multicast IPv6 BGP network/prefix
`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if tagnode defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"distance": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													"prefix": schema.MapNestedAttribute{
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																// TODO handle non-string types
																"distance": schema.StringAttribute{
																	// CustomType:          basetypes.StringTypable(nil),
																	// Required:            false,
																	Optional: true,
																	// Sensitive:           false,
																	// Description:         "",
																	MarkdownDescription: `Administrative distance for prefix

Format: u32:1-255
Administrative distance for external BGP routes
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
														MarkdownDescription: `Administrative distance for a specific BGP prefix

Format: ipv6net
Administrative distance for a specific BGP prefix
`,
														// DeprecationMessage:  "",
														// Validators:          []validator.Map(nil),
														// PlanModifiers:       []planmodifier.Map(nil),
														// TODO investigate if tagnode defaults can be handled
														// Default:             defaults.Map(nil),
													},

													// TODO handle non-string types
													"external": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `eBGP routes administrative distance

Format: u32:1-255
eBGP routes administrative distance
`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"internal": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `iBGP routes administrative distance

Format: u32:1-255
iBGP routes administrative distance
`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"local": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Locally originated BGP routes administrative distance

Format: u32:1-255
Locally originated BGP routes administrative distance
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
												MarkdownDescription: `Administrative distances for BGP routes

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
										MarkdownDescription: `Multicast IPv6 BGP settings

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},

									"ipv6_labeled_unicast": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											"aggregate_address": schema.MapNestedAttribute{
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														// TODO handle non-string types
														"as_set": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Generate AS-set path information for this aggregate address

`,
															// DeprecationMessage:  "",
															// TODO Recreate some of vyos validators for use in leafnodes
															// Validators:          []validator.String(nil),
															// PlanModifiers:       []planmodifier.String(nil),

														},

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

														// TODO handle non-string types
														"summary_only": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Announce the aggregate summary network only

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
												MarkdownDescription: `BGP aggregate network/prefix

Format: ipv6net
BGP aggregate network/prefix
`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if tagnode defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"network": schema.MapNestedAttribute{
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														// TODO handle non-string types
														"backdoor": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Use BGP network/prefix as a backdoor route

`,
															// DeprecationMessage:  "",
															// TODO Recreate some of vyos validators for use in leafnodes
															// Validators:          []validator.String(nil),
															// PlanModifiers:       []planmodifier.String(nil),

														},

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
												MarkdownDescription: `Import BGP network/prefix into labeled unicast IPv6 RIB

Format: ipv6net
Labeled Unicast IPv6 BGP network/prefix
`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if tagnode defaults can be handled
												// Default:             defaults.Map(nil),
											},
										},
										// CustomType:          basetypes.MapTypable(nil),
										// Required:            false,
										Optional: true,
										// Computed:            false,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Labeled Unicast IPv6 BGP settings

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},

									"ipv6_flowspec": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											"local_install": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													// TODO handle non-string types
													"interface": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Interface

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
												MarkdownDescription: `Apply local policy routing to interface

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
										MarkdownDescription: `Flowspec IPv6 BGP settings

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},

									"ipv6_vpn": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											"network": schema.MapNestedAttribute{
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														// TODO handle non-string types
														"rd": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Route Distinguisher

Format: ASN:NN_OR_IP-ADDRESS:NN
Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)
`,
															// DeprecationMessage:  "",
															// TODO Recreate some of vyos validators for use in leafnodes
															// Validators:          []validator.String(nil),
															// PlanModifiers:       []planmodifier.String(nil),

														},

														// TODO handle non-string types
														"label": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `MPLS label value assigned to route

Format: u32:0-1048575
MPLS label value
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
												MarkdownDescription: `Import BGP network/prefix into unicast VPN IPv6 RIB

Format: ipv6net
Unicast VPN IPv6 BGP network/prefix
`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if tagnode defaults can be handled
												// Default:             defaults.Map(nil),
											},
										},
										// CustomType:          basetypes.MapTypable(nil),
										// Required:            false,
										Optional: true,
										// Computed:            false,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Unicast VPN IPv6 BGP settings

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},

									"l2vpn_evpn": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											"vni": schema.MapNestedAttribute{
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														// TODO handle non-string types
														"advertise_default_gw": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Advertise All default g/w mac-ip routes in EVPN

`,
															// DeprecationMessage:  "",
															// TODO Recreate some of vyos validators for use in leafnodes
															// Validators:          []validator.String(nil),
															// PlanModifiers:       []planmodifier.String(nil),

														},

														// TODO handle non-string types
														"advertise_svi_ip": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Advertise svi mac-ip routes in EVPN

`,
															// DeprecationMessage:  "",
															// TODO Recreate some of vyos validators for use in leafnodes
															// Validators:          []validator.String(nil),
															// PlanModifiers:       []planmodifier.String(nil),

														},

														// TODO handle non-string types
														"rd": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Route Distinguisher

Format: ASN:NN_OR_IP-ADDRESS:NN
Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)
`,
															// DeprecationMessage:  "",
															// TODO Recreate some of vyos validators for use in leafnodes
															// Validators:          []validator.String(nil),
															// PlanModifiers:       []planmodifier.String(nil),

														},

														"route_target": schema.SingleNestedAttribute{
															Attributes: map[string]schema.Attribute{
																// TODO handle non-string types
																"both": schema.StringAttribute{
																	// CustomType:          basetypes.StringTypable(nil),
																	// Required:            false,
																	Optional: true,
																	// Sensitive:           false,
																	// Description:         "",
																	MarkdownDescription: `Route Target both import and export

Format: txt
Route target (A.B.C.D:MN|EF:OPQR|GHJK:MN)
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
																	MarkdownDescription: `Route Target import

Format: txt
Route target (A.B.C.D:MN|EF:OPQR|GHJK:MN)
`,
																	// DeprecationMessage:  "",
																	// TODO Recreate some of vyos validators for use in leafnodes
																	// Validators:          []validator.String(nil),
																	// PlanModifiers:       []planmodifier.String(nil),

																},

																// TODO handle non-string types
																"export": schema.StringAttribute{
																	// CustomType:          basetypes.StringTypable(nil),
																	// Required:            false,
																	Optional: true,
																	// Sensitive:           false,
																	// Description:         "",
																	MarkdownDescription: `Route Target export

Format: txt
Route target (A.B.C.D:MN|EF:OPQR|GHJK:MN)
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
															MarkdownDescription: `Route Target

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
												MarkdownDescription: `VXLAN Network Identifier

Format: u32:1-16777215
VNI number
`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if tagnode defaults can be handled
												// Default:             defaults.Map(nil),
											},

											// TODO handle non-string types
											"advertise_all_vni": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Advertise All local VNIs

`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"advertise_default_gw": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Advertise All default g/w mac-ip routes in EVPN

`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"advertise_svi_ip": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Advertise svi mac-ip routes in EVPN

`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"rd": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Route Distinguisher

Format: ASN:NN_OR_IP-ADDRESS:NN
Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"advertise_pip": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `EVPN system primary IP

Format: ipv4
IP address
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"rt_auto_derive": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Auto derivation of Route Target (RFC8365)

`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											"advertise": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													"ipv4": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															"unicast": schema.SingleNestedAttribute{
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
																MarkdownDescription: `IPv4 address family

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
														MarkdownDescription: `IPv4 address family

`,
														// DeprecationMessage:  "",
														// Validators:          []validator.Map(nil),
														// PlanModifiers:       []planmodifier.Map(nil),
														// TODO investigate if node defaults can be handled
														// Default:             defaults.Map(nil),
													},

													"ipv6": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															"unicast": schema.SingleNestedAttribute{
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
																MarkdownDescription: `IPv4 address family

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
														MarkdownDescription: `IPv6 address family

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
												MarkdownDescription: `Advertise prefix routes

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"route_target": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													// TODO handle non-string types
													"both": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Route Target both import and export

Format: txt
Route target (A.B.C.D:MN|EF:OPQR|GHJK:MN)
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
														MarkdownDescription: `Route Target import

Format: txt
Route target (A.B.C.D:MN|EF:OPQR|GHJK:MN)
`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"export": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Route Target export

Format: txt
Route target (A.B.C.D:MN|EF:OPQR|GHJK:MN)
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
												MarkdownDescription: `Route Target

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"flooding": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													// TODO handle non-string types
													"disable": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Do not flood any BUM packets

`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"head_end_replication": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Flood BUM packets using head-end replication

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
												MarkdownDescription: `Specify handling for BUM packets

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
								MarkdownDescription: `BGP address-family parameters

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"listen": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									"range": schema.MapNestedAttribute{
										NestedObject: schema.NestedAttributeObject{
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
										MarkdownDescription: `BGP dynamic neighbors listen range

Format: ipv4net
IPv4 dynamic neighbors listen range
Format: ipv6net
IPv6 dynamic neighbors listen range
`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if tagnode defaults can be handled
										// Default:             defaults.Map(nil),
									},

									// TODO handle non-string types
									"limit": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Maximum number of dynamic neighbors that can be created

Format: u32:1-5000
BGP neighbor limit
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
								MarkdownDescription: `Listen for and accept BGP dynamic neighbors from range

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"parameters": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"always_compare_med": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Always compare MEDs from different neighbors

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"cluster_id": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Route-reflector cluster-id

Format: ipv4
Route-reflector cluster-id
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"deterministic_med": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Compare MEDs between different peers in the same AS

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"ebgp_requires_policy": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Require in and out policy for eBGP peers (RFC8212)

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"fast_convergence": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Teardown sessions immediately whenever peer becomes unreachable

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"graceful_shutdown": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Graceful shutdown

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"log_neighbor_changes": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Log neighbor up/down changes and reset reason

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"minimum_holdtime": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `BGP minimum holdtime

Format: u32:1-65535
Minimum holdtime in seconds
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"network_import_check": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Enable IGP route check for network statements

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"route_reflector_allow_outbound_policy": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Route reflector client allow policy outbound

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"no_client_to_client_reflection": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Disable client to client route reflection

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"no_fast_external_failover": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Disable immediate session reset on peer link down event

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"no_suppress_duplicates": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Disable suppress duplicate updates if the route actually not changed

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"reject_as_sets": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Reject routes with AS_SET or AS_CONFED_SET flag

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
										MarkdownDescription: `Administrative shutdown of the BGP instance

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"suppress_fib_pending": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Advertise only routes that are programmed in kernel to peers

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"router_id": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Override default router identifier

Format: ipv4
Router-ID in IP address format
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									"bestpath": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											// TODO handle non-string types
											"bandwidth": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Link Bandwidth attribute

Format: default-weight-for-missing
Assign low default weight (1) to paths not having link bandwidth
Format: ignore
Ignore link bandwidth (do regular ECMP, not weighted)
Format: skip-missing
Ignore paths without link bandwidth for ECMP (if other paths have it)
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"compare_routerid": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Compare the router-id for identical EBGP paths

`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											"as_path": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													// TODO handle non-string types
													"confed": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Compare AS-path lengths including confederation sets and sequences

`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"ignore": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Ignore AS-path length in selecting a route

`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"multipath_relax": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Allow load sharing across routes that have different AS paths (but same length)

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
												MarkdownDescription: `AS-path attribute comparison parameters

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"med": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													// TODO handle non-string types
													"confed": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Compare MEDs among confederation paths

`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"missing_as_worst": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Treat missing route as a MED as the least preferred one

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
												MarkdownDescription: `MED attribute comparison parameters

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"peer_type": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													// TODO handle non-string types
													"multipath_relax": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Allow load sharing across routes learned from different peer types

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
												MarkdownDescription: `Peer type

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
										MarkdownDescription: `Default bestpath selection mechanism

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},

									"confederation": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											// TODO handle non-string types
											"identifier": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Confederation AS identifier

Format: u32:1-4294967294
Confederation AS id
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"peers": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Peer ASs in the BGP confederation

Format: u32:1-4294967294
Peer AS number
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
										MarkdownDescription: `AS confederation parameters

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},

									"conditional_advertisement": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											// TODO handle non-string types
											"timer": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Set period to rescan BGP table to check if condition is met

Format: u32:5-240
Period to rerun the conditional advertisement scanner process
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

												Default:  stringdefault.StaticString(`60`),
												Computed: true,
											},
										},
										// CustomType:          basetypes.MapTypable(nil),
										// Required:            false,
										Optional: true,
										// Computed:            false,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Conditional advertisement settings

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},

									"dampening": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											// TODO handle non-string types
											"half_life": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Half-life time for dampening

Format: u32:1-45
Half-life penalty in minutes
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"max_suppress_time": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Maximum duration to suppress a stable route

Format: u32:1-255
Maximum suppress duration in minutes
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"re_use": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Threshold to start reusing a route

Format: u32:1-20000
Re-use penalty points
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"start_suppress_time": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `When to start suppressing a route

Format: u32:1-20000
Start-suppress penalty points
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
										MarkdownDescription: `Enable route-flap dampening

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},

									"default": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											// TODO handle non-string types
											"local_pref": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Default local preference

Format: u32
Local preference
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
										MarkdownDescription: `BGP defaults

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},

									"distance": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											"prefix": schema.MapNestedAttribute{
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														// TODO handle non-string types
														"distance": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Administrative distance for prefix

Format: u32:1-255
Administrative distance for external BGP routes
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
												MarkdownDescription: `Administrative distance for a specific BGP prefix

Format: ipv4net
Administrative distance for a specific BGP prefix
`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if tagnode defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"global": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													// TODO handle non-string types
													"external": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Administrative distance for external BGP routes

Format: u32:1-255
Administrative distance for external BGP routes
`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"internal": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Administrative distance for internal BGP routes

Format: u32:1-255
Administrative distance for internal BGP routes
`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"local": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Administrative distance for local BGP routes

Format: u32:1-255
Administrative distance for internal BGP routes
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
												MarkdownDescription: `Global administratives distances for BGP routes

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
										MarkdownDescription: `Administratives distances for BGP routes

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},

									"graceful_restart": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											// TODO handle non-string types
											"stalepath_time": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Maximum time to hold onto restarting neighbors stale paths

Format: u32:1-3600
Hold time in seconds
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
										MarkdownDescription: `Graceful restart capability parameters

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
								MarkdownDescription: `BGP parameters

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
								MarkdownDescription: `BGP protocol timers

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
						MarkdownDescription: `Border Gateway Protocol (BGP)

`,
						// DeprecationMessage:  "",
						// Validators:          []validator.Map(nil),
						// PlanModifiers:       []planmodifier.Map(nil),
						// TODO investigate if node defaults can be handled
						// Default:             defaults.Map(nil),
					},

					"eigrp": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							// TODO handle non-string types
							"local_as": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Autonomous System Number (ASN)

Format: u32:1-65535
Autonomous System Number
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"maximum_paths": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Forward packets over multiple paths

Format: u32:1-32
Number of paths
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"network": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Enable routing on an IP network

Format: ipv4net
EIGRP network prefix
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"passive_interface": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Suppress routing updates on an interface

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"redistribute": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Redistribute information from another routing protocol

Format: bgp
Border Gateway Protocol (BGP)
Format: connected
Connected routes
Format: nhrp
Next Hop Resolution Protocol (NHRP)
Format: ospf
Open Shortest Path First (OSPFv2)
Format: rip
Routing Information Protocol (RIP)
Format: babel
Babel routing protocol (Babel)
Format: static
Statically configured routes
Format: vnc
Virtual Network Control (VNC)
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

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

							// TODO handle non-string types
							"router_id": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Override default router identifier

Format: ipv4
Router-ID in IP address format
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"variance": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Control load balancing variance

Format: u32:1-128
Metric variance multiplier
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							"metric": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"weights": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Modify metric coefficients

Format: u32:0-255
K1
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
								MarkdownDescription: `Modify metrics and parameters for advertisement

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
						MarkdownDescription: `Enhanced Interior Gateway Routing Protocol (EIGRP)

`,
						// DeprecationMessage:  "",
						// Validators:          []validator.Map(nil),
						// PlanModifiers:       []planmodifier.Map(nil),
						// TODO investigate if node defaults can be handled
						// Default:             defaults.Map(nil),
					},

					"isis": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"interface": schema.MapNestedAttribute{
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										// TODO handle non-string types
										"circuit_type": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Configure circuit type for interface

Format: level-1
Level-1 only adjacencies are formed
Format: level-1-2
Level-1-2 adjacencies are formed
Format: level-2-only
Level-2 only adjacencies are formed
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"hello_padding": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Add padding to IS-IS hello packets

`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"hello_interval": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Set Hello interval

Format: u32:1-600
Set Hello interval
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"hello_multiplier": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Set Hello interval

Format: u32:2-100
Set multiplier for Hello holding time
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"metric": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Set default metric for circuit

Format: u32:0-16777215
Default metric value
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
											MarkdownDescription: `Configure passive mode for interface

`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"priority": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Set priority for Designated Router election

Format: u32:0-127
Priority value
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"psnp_interval": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Set PSNP interval

Format: u32:0-127
PSNP interval in seconds
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"no_three_way_handshake": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Disable three-way handshake

`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

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
											},
											// CustomType:          basetypes.MapTypable(nil),
											// Required:            false,
											Optional: true,
											// Computed:            false,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Enable Bidirectional Forwarding Detection (BFD)

`,
											// DeprecationMessage:  "",
											// Validators:          []validator.Map(nil),
											// PlanModifiers:       []planmodifier.Map(nil),
											// TODO investigate if node defaults can be handled
											// Default:             defaults.Map(nil),
										},

										"network": schema.SingleNestedAttribute{
											Attributes: map[string]schema.Attribute{
												// TODO handle non-string types
												"point_to_point": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `point-to-point network type

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
											MarkdownDescription: `Set network type

`,
											// DeprecationMessage:  "",
											// Validators:          []validator.Map(nil),
											// PlanModifiers:       []planmodifier.Map(nil),
											// TODO investigate if node defaults can be handled
											// Default:             defaults.Map(nil),
										},

										"password": schema.SingleNestedAttribute{
											Attributes: map[string]schema.Attribute{
												// TODO handle non-string types
												"plaintext_password": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `Plain-text authentication type

Format: txt
Circuit password
`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

												},

												// TODO handle non-string types
												"md5": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `MD5 authentication type

Format: txt
Level-wide password
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
											MarkdownDescription: `Configure the authentication password for a circuit

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
								MarkdownDescription: `Interface params

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if tagnode defaults can be handled
								// Default:             defaults.Map(nil),
							},

							// TODO handle non-string types
							"dynamic_hostname": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Dynamic hostname for IS-IS

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"level": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `IS-IS level number

Format: level-1
Act as a station router
Format: level-1-2
Act as both a station and an area router
Format: level-2
Act as an area router
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"log_adjacency_changes": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Log adjacency state changes

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"lsp_gen_interval": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Minimum interval between regenerating same LSP

Format: u32:1-120
Minimum interval in seconds
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"lsp_mtu": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Configure the maximum size of generated LSPs

Format: u32:128-4352
Maximum size of generated LSPs
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

								Default:  stringdefault.StaticString(`1497`),
								Computed: true,
							},

							// TODO handle non-string types
							"lsp_refresh_interval": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `LSP refresh interval

Format: u32:1-65235
LSP refresh interval in seconds
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"max_lsp_lifetime": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Maximum LSP lifetime

Format: u32:350-65535
LSP lifetime in seconds
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"metric_style": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Use old-style (ISO 10589) or new-style packet formats

Format: narrow
Use old style of TLVs with narrow metric
Format: transition
Send and accept both styles of TLVs during transition
Format: wide
Use new style of TLVs to carry wider metric
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"net": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `A Network Entity Title for this process (ISO only)

Format: XX.XXXX. ... .XXX.XX
Network entity title (NET)
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"purge_originator": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Use the RFC 6232 purge-originator

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"set_attached_bit": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Set attached bit to identify as L1/L2 router for inter-area traffic

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"set_overload_bit": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Set overload bit to avoid any transit traffic

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"spf_interval": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Minimum interval between SPF calculations

Format: u32:1-120
Interval in seconds
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

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

							"area_password": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"plaintext_password": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Plain-text authentication type

Format: txt
Circuit password
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"md5": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `MD5 authentication type

Format: txt
Level-wide password
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
								MarkdownDescription: `Configure the authentication password for an area

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"default_information": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									"originate": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											"ipv4": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													"level_1": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"always": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Always advertise default route

`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Set default metric for circuit

Format: u32:0-16777215
Default metric value
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Distribute default route into level-1

`,
														// DeprecationMessage:  "",
														// Validators:          []validator.Map(nil),
														// PlanModifiers:       []planmodifier.Map(nil),
														// TODO investigate if node defaults can be handled
														// Default:             defaults.Map(nil),
													},

													"level_2": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"always": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Always advertise default route

`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Set default metric for circuit

Format: u32:0-16777215
Default metric value
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Distribute default route into level-2

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
												MarkdownDescription: `Distribute default route for IPv4

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"ipv6": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													"level_1": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"always": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Always advertise default route

`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Set default metric for circuit

Format: u32:0-16777215
Default metric value
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Distribute default route into level-1

`,
														// DeprecationMessage:  "",
														// Validators:          []validator.Map(nil),
														// PlanModifiers:       []planmodifier.Map(nil),
														// TODO investigate if node defaults can be handled
														// Default:             defaults.Map(nil),
													},

													"level_2": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"always": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Always advertise default route

`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Set default metric for circuit

Format: u32:0-16777215
Default metric value
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Distribute default route into level-2

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
												MarkdownDescription: `Distribute default route for IPv6

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
										MarkdownDescription: `Distribute a default route

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
								MarkdownDescription: `Control distribution of default information

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"domain_password": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"plaintext_password": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Plain-text authentication type

Format: txt
Circuit password
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"md5": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `MD5 authentication type

Format: txt
Level-wide password
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
								MarkdownDescription: `Set the authentication password for a routing domain

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"traffic_engineering": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"enable": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Enable MPLS traffic engineering extensions

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"address": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `MPLS traffic engineering router ID

Format: ipv4
IPv4 address
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
								MarkdownDescription: `Show IS-IS neighbor adjacencies

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"segment_routing": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									"prefix": schema.MapNestedAttribute{
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"absolute": schema.SingleNestedAttribute{
													Attributes: map[string]schema.Attribute{
														// TODO handle non-string types
														"value": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Specify the absolute value of prefix segment/label ID

Format: u32:16-1048575
The absolute segment/label ID value
`,
															// DeprecationMessage:  "",
															// TODO Recreate some of vyos validators for use in leafnodes
															// Validators:          []validator.String(nil),
															// PlanModifiers:       []planmodifier.String(nil),

														},

														// TODO handle non-string types
														"explicit_null": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Request upstream neighbor to replace segment/label with explicit null label

`,
															// DeprecationMessage:  "",
															// TODO Recreate some of vyos validators for use in leafnodes
															// Validators:          []validator.String(nil),
															// PlanModifiers:       []planmodifier.String(nil),

														},

														// TODO handle non-string types
														"no_php_flag": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Do not request penultimate hop popping for segment/label

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
													MarkdownDescription: `Specify the absolute value of prefix segment/label ID

`,
													// DeprecationMessage:  "",
													// Validators:          []validator.Map(nil),
													// PlanModifiers:       []planmodifier.Map(nil),
													// TODO investigate if node defaults can be handled
													// Default:             defaults.Map(nil),
												},

												"index": schema.SingleNestedAttribute{
													Attributes: map[string]schema.Attribute{
														// TODO handle non-string types
														"value": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Specify the index value of prefix segment/label ID

Format: u32:0-65535
The index segment/label ID value
`,
															// DeprecationMessage:  "",
															// TODO Recreate some of vyos validators for use in leafnodes
															// Validators:          []validator.String(nil),
															// PlanModifiers:       []planmodifier.String(nil),

														},

														// TODO handle non-string types
														"explicit_null": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Request upstream neighbor to replace segment/label with explicit null label

`,
															// DeprecationMessage:  "",
															// TODO Recreate some of vyos validators for use in leafnodes
															// Validators:          []validator.String(nil),
															// PlanModifiers:       []planmodifier.String(nil),

														},

														// TODO handle non-string types
														"no_php_flag": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Do not request penultimate hop popping for segment/label

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
													MarkdownDescription: `Specify the index value of prefix segment/label ID

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
										MarkdownDescription: `Static IPv4/IPv6 prefix segment/label mapping

Format: ipv4net
IPv4 prefix segment
Format: ipv6net
IPv6 prefix segment
`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if tagnode defaults can be handled
										// Default:             defaults.Map(nil),
									},

									// TODO handle non-string types
									"maximum_label_depth": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Maximum MPLS labels allowed for this router

Format: u32:1-16
MPLS label depth
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									"global_block": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											// TODO handle non-string types
											"low_label_value": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `MPLS label lower bound

Format: u32:16-1048575
Label value (recommended minimum value: 300)
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"high_label_value": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `MPLS label upper bound

Format: u32:16-1048575
Label value
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
										MarkdownDescription: `Segment Routing Global Block label range

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},

									"local_block": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											// TODO handle non-string types
											"low_label_value": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `MPLS label lower bound

Format: u32:16-1048575
Label value (recommended minimum value: 300)
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"high_label_value": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `MPLS label upper bound

Format: u32:16-1048575
Label value
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
										MarkdownDescription: `Segment Routing Local Block label range

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
								MarkdownDescription: `Segment-Routing (SPRING) settings

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"redistribute": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									"ipv4": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											"bgp": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													"level_1": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Set default metric for circuit

Format: u32:0-16777215
Default metric value
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute into level-1

`,
														// DeprecationMessage:  "",
														// Validators:          []validator.Map(nil),
														// PlanModifiers:       []planmodifier.Map(nil),
														// TODO investigate if node defaults can be handled
														// Default:             defaults.Map(nil),
													},

													"level_2": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Set default metric for circuit

Format: u32:0-16777215
Default metric value
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute into level-2

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
												MarkdownDescription: `Border Gateway Protocol (BGP)

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"connected": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													"level_1": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Set default metric for circuit

Format: u32:0-16777215
Default metric value
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute into level-1

`,
														// DeprecationMessage:  "",
														// Validators:          []validator.Map(nil),
														// PlanModifiers:       []planmodifier.Map(nil),
														// TODO investigate if node defaults can be handled
														// Default:             defaults.Map(nil),
													},

													"level_2": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Set default metric for circuit

Format: u32:0-16777215
Default metric value
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute into level-2

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
												MarkdownDescription: `Redistribute connected routes into IS-IS

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"kernel": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													"level_1": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Set default metric for circuit

Format: u32:0-16777215
Default metric value
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute into level-1

`,
														// DeprecationMessage:  "",
														// Validators:          []validator.Map(nil),
														// PlanModifiers:       []planmodifier.Map(nil),
														// TODO investigate if node defaults can be handled
														// Default:             defaults.Map(nil),
													},

													"level_2": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Set default metric for circuit

Format: u32:0-16777215
Default metric value
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute into level-2

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
												MarkdownDescription: `Redistribute kernel routes into IS-IS

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"ospf": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													"level_1": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Set default metric for circuit

Format: u32:0-16777215
Default metric value
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute into level-1

`,
														// DeprecationMessage:  "",
														// Validators:          []validator.Map(nil),
														// PlanModifiers:       []planmodifier.Map(nil),
														// TODO investigate if node defaults can be handled
														// Default:             defaults.Map(nil),
													},

													"level_2": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Set default metric for circuit

Format: u32:0-16777215
Default metric value
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute into level-2

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
												MarkdownDescription: `Redistribute OSPF routes into IS-IS

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"rip": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													"level_1": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Set default metric for circuit

Format: u32:0-16777215
Default metric value
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute into level-1

`,
														// DeprecationMessage:  "",
														// Validators:          []validator.Map(nil),
														// PlanModifiers:       []planmodifier.Map(nil),
														// TODO investigate if node defaults can be handled
														// Default:             defaults.Map(nil),
													},

													"level_2": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Set default metric for circuit

Format: u32:0-16777215
Default metric value
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute into level-2

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
												MarkdownDescription: `Redistribute RIP routes into IS-IS

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"babel": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													"level_1": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Set default metric for circuit

Format: u32:0-16777215
Default metric value
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute into level-1

`,
														// DeprecationMessage:  "",
														// Validators:          []validator.Map(nil),
														// PlanModifiers:       []planmodifier.Map(nil),
														// TODO investigate if node defaults can be handled
														// Default:             defaults.Map(nil),
													},

													"level_2": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Set default metric for circuit

Format: u32:0-16777215
Default metric value
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute into level-2

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
												MarkdownDescription: `Redistribute Babel routes into IS-IS

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"static": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													"level_1": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Set default metric for circuit

Format: u32:0-16777215
Default metric value
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute into level-1

`,
														// DeprecationMessage:  "",
														// Validators:          []validator.Map(nil),
														// PlanModifiers:       []planmodifier.Map(nil),
														// TODO investigate if node defaults can be handled
														// Default:             defaults.Map(nil),
													},

													"level_2": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Set default metric for circuit

Format: u32:0-16777215
Default metric value
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute into level-2

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
												MarkdownDescription: `Redistribute static routes into IS-IS

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
										MarkdownDescription: `Redistribute IPv4 routes

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},

									"ipv6": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											"bgp": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													"level_1": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Set default metric for circuit

Format: u32:0-16777215
Default metric value
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute into level-1

`,
														// DeprecationMessage:  "",
														// Validators:          []validator.Map(nil),
														// PlanModifiers:       []planmodifier.Map(nil),
														// TODO investigate if node defaults can be handled
														// Default:             defaults.Map(nil),
													},

													"level_2": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Set default metric for circuit

Format: u32:0-16777215
Default metric value
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute into level-2

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
												MarkdownDescription: `Redistribute BGP routes into IS-IS

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"connected": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													"level_1": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Set default metric for circuit

Format: u32:0-16777215
Default metric value
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute into level-1

`,
														// DeprecationMessage:  "",
														// Validators:          []validator.Map(nil),
														// PlanModifiers:       []planmodifier.Map(nil),
														// TODO investigate if node defaults can be handled
														// Default:             defaults.Map(nil),
													},

													"level_2": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Set default metric for circuit

Format: u32:0-16777215
Default metric value
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute into level-2

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
												MarkdownDescription: `Redistribute connected routes into IS-IS

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"kernel": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													"level_1": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Set default metric for circuit

Format: u32:0-16777215
Default metric value
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute into level-1

`,
														// DeprecationMessage:  "",
														// Validators:          []validator.Map(nil),
														// PlanModifiers:       []planmodifier.Map(nil),
														// TODO investigate if node defaults can be handled
														// Default:             defaults.Map(nil),
													},

													"level_2": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Set default metric for circuit

Format: u32:0-16777215
Default metric value
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute into level-2

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
												MarkdownDescription: `Redistribute kernel routes into IS-IS

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"ospf6": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													"level_1": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Set default metric for circuit

Format: u32:0-16777215
Default metric value
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute into level-1

`,
														// DeprecationMessage:  "",
														// Validators:          []validator.Map(nil),
														// PlanModifiers:       []planmodifier.Map(nil),
														// TODO investigate if node defaults can be handled
														// Default:             defaults.Map(nil),
													},

													"level_2": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Set default metric for circuit

Format: u32:0-16777215
Default metric value
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute into level-2

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
												MarkdownDescription: `Redistribute OSPFv3 routes into IS-IS

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"ripng": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													"level_1": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Set default metric for circuit

Format: u32:0-16777215
Default metric value
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute into level-1

`,
														// DeprecationMessage:  "",
														// Validators:          []validator.Map(nil),
														// PlanModifiers:       []planmodifier.Map(nil),
														// TODO investigate if node defaults can be handled
														// Default:             defaults.Map(nil),
													},

													"level_2": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Set default metric for circuit

Format: u32:0-16777215
Default metric value
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute into level-2

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
												MarkdownDescription: `Redistribute RIPng routes into IS-IS

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"babel": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													"level_1": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Set default metric for circuit

Format: u32:0-16777215
Default metric value
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute into level-1

`,
														// DeprecationMessage:  "",
														// Validators:          []validator.Map(nil),
														// PlanModifiers:       []planmodifier.Map(nil),
														// TODO investigate if node defaults can be handled
														// Default:             defaults.Map(nil),
													},

													"level_2": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Set default metric for circuit

Format: u32:0-16777215
Default metric value
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute into level-2

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
												MarkdownDescription: `Redistribute Babel routes into IS-IS

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if node defaults can be handled
												// Default:             defaults.Map(nil),
											},

											"static": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													"level_1": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Set default metric for circuit

Format: u32:0-16777215
Default metric value
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute into level-1

`,
														// DeprecationMessage:  "",
														// Validators:          []validator.Map(nil),
														// PlanModifiers:       []planmodifier.Map(nil),
														// TODO investigate if node defaults can be handled
														// Default:             defaults.Map(nil),
													},

													"level_2": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"metric": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Set default metric for circuit

Format: u32:0-16777215
Default metric value
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

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
														MarkdownDescription: `Redistribute into level-2

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
												MarkdownDescription: `Redistribute static routes into IS-IS

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
										MarkdownDescription: `Redistribute IPv6 routes

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
								MarkdownDescription: `Redistribute information from another routing protocol

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"spf_delay_ietf": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"init_delay": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Delay used while in QUIET state

Format: u32:0-60000
Delay used while in QUIET state (in ms)
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"short_delay": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Delay used while in SHORT_WAIT state

Format: u32:0-60000
Delay used while in SHORT_WAIT state (in ms)
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"long_delay": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Delay used while in LONG_WAIT

Format: u32:0-60000
Delay used while in LONG_WAIT state in ms
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"holddown": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Time with no received IGP events before considering IGP stable

Format: u32:0-60000
Time with no received IGP events before considering IGP stable in ms
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"time_to_learn": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Maximum duration needed to learn all the events related to a single failure

Format: u32:0-60000
Maximum duration needed to learn all the events related to a single failure in ms
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
								MarkdownDescription: `IETF SPF delay algorithm

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
						MarkdownDescription: `Intermediate System to Intermediate System (IS-IS)

`,
						// DeprecationMessage:  "",
						// Validators:          []validator.Map(nil),
						// PlanModifiers:       []planmodifier.Map(nil),
						// TODO investigate if node defaults can be handled
						// Default:             defaults.Map(nil),
					},

					"ospf": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"access_list": schema.MapNestedAttribute{
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										// TODO handle non-string types
										"export": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Filter for outgoing routing update

Format: bgp
Filter BGP routes
Format: connected
Filter connected routes
Format: isis
Filter IS-IS routes
Format: kernel
Filter Kernel routes
Format: rip
Filter RIP routes
Format: static
Filter static routes
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
								MarkdownDescription: `Access list to filter networks in routing updates

Format: u32
Access-list number
`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if tagnode defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"area": schema.MapNestedAttribute{
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"range": schema.MapNestedAttribute{
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													// TODO handle non-string types
													"cost": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Metric for this range

Format: u32:0-16777215
Metric for this range
`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"not_advertise": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Do not advertise this range

`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"substitute": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Advertise area range as another prefix

Format: ipv4net
Advertise area range as another prefix
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
											MarkdownDescription: `Summarize routes matching a prefix (border routers only)

Format: ipv4net
Area range prefix
`,
											// DeprecationMessage:  "",
											// Validators:          []validator.Map(nil),
											// PlanModifiers:       []planmodifier.Map(nil),
											// TODO investigate if tagnode defaults can be handled
											// Default:             defaults.Map(nil),
										},

										"virtual_link": schema.MapNestedAttribute{
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													// TODO handle non-string types
													"dead_interval": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Interval after which a neighbor is declared dead

Format: u32:1-65535
Neighbor dead interval (seconds)
`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

														Default:  stringdefault.StaticString(`40`),
														Computed: true,
													},

													// TODO handle non-string types
													"hello_interval": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Interval between hello packets

Format: u32:1-65535
Hello interval (seconds)
`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

														Default:  stringdefault.StaticString(`10`),
														Computed: true,
													},

													// TODO handle non-string types
													"retransmit_interval": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Interval between retransmitting lost link state advertisements

Format: u32:1-65535
Retransmit interval (seconds)
`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

														Default:  stringdefault.StaticString(`5`),
														Computed: true,
													},

													// TODO handle non-string types
													"transmit_delay": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Link state transmit delay

Format: u32:1-65535
Link state transmit delay (seconds)
`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

														Default:  stringdefault.StaticString(`1`),
														Computed: true,
													},

													"authentication": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															// TODO handle non-string types
															"plaintext_password": schema.StringAttribute{
																// CustomType:          basetypes.StringTypable(nil),
																// Required:            false,
																Optional: true,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `Plain text password

Format: txt
Plain text password (8 characters or less)
`,
																// DeprecationMessage:  "",
																// TODO Recreate some of vyos validators for use in leafnodes
																// Validators:          []validator.String(nil),
																// PlanModifiers:       []planmodifier.String(nil),

															},

															"md5": schema.SingleNestedAttribute{
																Attributes: map[string]schema.Attribute{
																	"key_id": schema.MapNestedAttribute{
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				// TODO handle non-string types
																				"md5_key": schema.StringAttribute{
																					// CustomType:          basetypes.StringTypable(nil),
																					// Required:            false,
																					Optional: true,
																					// Sensitive:           false,
																					// Description:         "",
																					MarkdownDescription: `MD5 authentication type

Format: txt
MD5 Key (16 characters or less)
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
																		MarkdownDescription: `MD5 key id

Format: u32:1-255
MD5 key id
`,
																		// DeprecationMessage:  "",
																		// Validators:          []validator.Map(nil),
																		// PlanModifiers:       []planmodifier.Map(nil),
																		// TODO investigate if tagnode defaults can be handled
																		// Default:             defaults.Map(nil),
																	},
																},
																// CustomType:          basetypes.MapTypable(nil),
																// Required:            false,
																Optional: true,
																// Computed:            false,
																// Sensitive:           false,
																// Description:         "",
																MarkdownDescription: `MD5 key id

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
														MarkdownDescription: `Authentication

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
											MarkdownDescription: `Virtual link

Format: ipv4
OSPF area in dotted decimal notation
`,
											// DeprecationMessage:  "",
											// Validators:          []validator.Map(nil),
											// PlanModifiers:       []planmodifier.Map(nil),
											// TODO investigate if tagnode defaults can be handled
											// Default:             defaults.Map(nil),
										},

										// TODO handle non-string types
										"authentication": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `OSPF area authentication type

Format: plaintext-password
Use plain-text authentication
Format: md5
Use MD5 authentication
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"network": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `OSPF network

Format: ipv4net
OSPF network
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"shortcut": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Area shortcut mode

Format: default
Set default
Format: disable
Disable shortcutting mode
Format: enable
Enable shortcutting mode
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"export_list": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Set the filter for networks announced to other areas

Format: u32
Access-list number
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"import_list": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Set the filter for networks from other areas announced

Format: u32
Access-list number
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										"area_type": schema.SingleNestedAttribute{
											Attributes: map[string]schema.Attribute{
												// TODO handle non-string types
												"normal": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `Normal OSPF area

`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

												},

												"nssa": schema.SingleNestedAttribute{
													Attributes: map[string]schema.Attribute{
														// TODO handle non-string types
														"default_cost": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Summary-default cost of an NSSA area

Format: u32:0-16777215
Summary default cost
`,
															// DeprecationMessage:  "",
															// TODO Recreate some of vyos validators for use in leafnodes
															// Validators:          []validator.String(nil),
															// PlanModifiers:       []planmodifier.String(nil),

														},

														// TODO handle non-string types
														"no_summary": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Do not inject inter-area routes into stub

`,
															// DeprecationMessage:  "",
															// TODO Recreate some of vyos validators for use in leafnodes
															// Validators:          []validator.String(nil),
															// PlanModifiers:       []planmodifier.String(nil),

														},

														// TODO handle non-string types
														"translate": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Configure NSSA-ABR

Format: always
Always translate LSA types
Format: candidate
Translate for election
Format: never
Never translate LSA types
`,
															// DeprecationMessage:  "",
															// TODO Recreate some of vyos validators for use in leafnodes
															// Validators:          []validator.String(nil),
															// PlanModifiers:       []planmodifier.String(nil),

															Default:  stringdefault.StaticString(`candidate`),
															Computed: true,
														},
													},
													// CustomType:          basetypes.MapTypable(nil),
													// Required:            false,
													Optional: true,
													// Computed:            false,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `Not-So-Stubby OSPF area

`,
													// DeprecationMessage:  "",
													// Validators:          []validator.Map(nil),
													// PlanModifiers:       []planmodifier.Map(nil),
													// TODO investigate if node defaults can be handled
													// Default:             defaults.Map(nil),
												},

												"stub": schema.SingleNestedAttribute{
													Attributes: map[string]schema.Attribute{
														// TODO handle non-string types
														"default_cost": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Summary-default cost

Format: u32:0-16777215
Summary default cost
`,
															// DeprecationMessage:  "",
															// TODO Recreate some of vyos validators for use in leafnodes
															// Validators:          []validator.String(nil),
															// PlanModifiers:       []planmodifier.String(nil),

														},

														// TODO handle non-string types
														"no_summary": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Do not inject inter-area routes into the stub

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
													MarkdownDescription: `Stub OSPF area

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
											MarkdownDescription: `Area type

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
								MarkdownDescription: `OSPF area settings

Format: u32
OSPF area number in decimal notation
Format: ipv4
OSPF area number in dotted decimal notation
`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if tagnode defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"interface": schema.MapNestedAttribute{
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										// TODO handle non-string types
										"area": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Enable OSPF on this interface

Format: u32
OSPF area ID as decimal notation
Format: ipv4
OSPF area ID in IP address notation
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"dead_interval": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Interval after which a neighbor is declared dead

Format: u32:1-65535
Neighbor dead interval (seconds)
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

											Default:  stringdefault.StaticString(`40`),
											Computed: true,
										},

										// TODO handle non-string types
										"hello_interval": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Interval between hello packets

Format: u32:1-65535
Hello interval (seconds)
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

											Default:  stringdefault.StaticString(`10`),
											Computed: true,
										},

										// TODO handle non-string types
										"retransmit_interval": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Interval between retransmitting lost link state advertisements

Format: u32:1-65535
Retransmit interval (seconds)
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

											Default:  stringdefault.StaticString(`5`),
											Computed: true,
										},

										// TODO handle non-string types
										"transmit_delay": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Link state transmit delay

Format: u32:1-65535
Link state transmit delay (seconds)
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

											Default:  stringdefault.StaticString(`1`),
											Computed: true,
										},

										// TODO handle non-string types
										"cost": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Interface cost

Format: u32:1-65535
OSPF interface cost
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"mtu_ignore": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Disable Maximum Transmission Unit (MTU) mismatch detection

`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"priority": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Router priority

Format: u32:0-255
OSPF router priority cost
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

											Default:  stringdefault.StaticString(`1`),
											Computed: true,
										},

										// TODO handle non-string types
										"bandwidth": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Interface bandwidth (Mbit/s)

Format: u32:1-100000
Bandwidth in Megabit/sec (for calculating OSPF cost)
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"hello_multiplier": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Hello multiplier factor

Format: u32:1-10
Number of Hellos to send each second
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"network": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Network type

Format: broadcast
Broadcast network type
Format: non-broadcast
Non-broadcast network type
Format: point-to-multipoint
Point-to-multipoint network type
Format: point-to-point
Point-to-point network type
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										"authentication": schema.SingleNestedAttribute{
											Attributes: map[string]schema.Attribute{
												// TODO handle non-string types
												"plaintext_password": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `Plain text password

Format: txt
Plain text password (8 characters or less)
`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

												},

												"md5": schema.SingleNestedAttribute{
													Attributes: map[string]schema.Attribute{
														"key_id": schema.MapNestedAttribute{
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	// TODO handle non-string types
																	"md5_key": schema.StringAttribute{
																		// CustomType:          basetypes.StringTypable(nil),
																		// Required:            false,
																		Optional: true,
																		// Sensitive:           false,
																		// Description:         "",
																		MarkdownDescription: `MD5 authentication type

Format: txt
MD5 Key (16 characters or less)
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
															MarkdownDescription: `MD5 key id

Format: u32:1-255
MD5 key id
`,
															// DeprecationMessage:  "",
															// Validators:          []validator.Map(nil),
															// PlanModifiers:       []planmodifier.Map(nil),
															// TODO investigate if tagnode defaults can be handled
															// Default:             defaults.Map(nil),
														},
													},
													// CustomType:          basetypes.MapTypable(nil),
													// Required:            false,
													Optional: true,
													// Computed:            false,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `MD5 key id

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
											MarkdownDescription: `Authentication

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
											},
											// CustomType:          basetypes.MapTypable(nil),
											// Required:            false,
											Optional: true,
											// Computed:            false,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Enable Bidirectional Forwarding Detection (BFD)

`,
											// DeprecationMessage:  "",
											// Validators:          []validator.Map(nil),
											// PlanModifiers:       []planmodifier.Map(nil),
											// TODO investigate if node defaults can be handled
											// Default:             defaults.Map(nil),
										},

										"passive": schema.SingleNestedAttribute{
											Attributes: map[string]schema.Attribute{
												// TODO handle non-string types
												"disable": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `Disable instance

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
											MarkdownDescription: `Suppress routing updates on an interface

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
								MarkdownDescription: `Interface configuration

Format: txt
Interface name
`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if tagnode defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"neighbor": schema.MapNestedAttribute{
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										// TODO handle non-string types
										"poll_interval": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Dead neighbor polling interval

Format: u32:1-65535
Seconds between dead neighbor polling interval
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

											Default:  stringdefault.StaticString(`60`),
											Computed: true,
										},

										// TODO handle non-string types
										"priority": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Neighbor priority in seconds

Format: u32:0-255
Neighbor priority
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

											Default:  stringdefault.StaticString(`0`),
											Computed: true,
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
								MarkdownDescription: `Specify neighbor router

Format: ipv4
Neighbor IP address
`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if tagnode defaults can be handled
								// Default:             defaults.Map(nil),
							},

							// TODO handle non-string types
							"default_metric": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Metric of redistributed routes

Format: u32:0-16777214
Metric of redistributed routes
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"maximum_paths": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Maximum multiple paths (ECMP)

Format: u32:1-64
Maximum multiple paths (ECMP)
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"passive_interface": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Suppress routing updates on an interface

Format: default
Default to suppress routing updates on all interfaces
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

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

							"auto_cost": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"reference_bandwidth": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Reference bandwidth method to assign cost

Format: u32:1-4294967
Reference bandwidth cost in Mbits/sec
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

										Default:  stringdefault.StaticString(`100`),
										Computed: true,
									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Calculate interface cost according to bandwidth

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"default_information": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									"originate": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											// TODO handle non-string types
											"always": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Always advertise a default route

`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"metric": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `OSPF default metric

Format: u32:0-16777214
Default metric
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"metric_type": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `OSPF metric type for default routes

Format: u32:1-2
Set OSPF External Type 1/2 metrics
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

												Default:  stringdefault.StaticString(`2`),
												Computed: true,
											},

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
										MarkdownDescription: `Distribute a default route

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
								MarkdownDescription: `Default route advertisment settings

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"distance": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"global": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Administrative distance

Format: u32:1-255
Administrative distance
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									"ospf": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											// TODO handle non-string types
											"external": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Distance for external routes

Format: u32:1-255
Distance for external routes
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"inter_area": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Distance for inter-area routes

Format: u32:1-255
Distance for inter-area routes
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"intra_area": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Distance for intra-area routes

Format: u32:1-255
Distance for intra-area routes
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
										MarkdownDescription: `OSPF administrative distance

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
								MarkdownDescription: `Administrative distance

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"log_adjacency_changes": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"detail": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Log all state changes

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
								MarkdownDescription: `Log adjacency state changes

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"max_metric": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									"router_lsa": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											// TODO handle non-string types
											"administrative": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Administratively apply, for an indefinite period

`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"on_shutdown": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Advertise stub-router prior to full shutdown of OSPF

Format: u32:5-100
Time (seconds) to advertise self as stub-router
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"on_startup": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Automatically advertise stub Router-LSA on startup of OSPF

Format: u32:5-86400
Time (seconds) to advertise self as stub-router
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
										MarkdownDescription: `Advertise own Router-LSA with infinite distance (stub router)

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
								MarkdownDescription: `OSPF maximum and infinite-distance metric

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"mpls_te": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"enable": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Enable MPLS-TE functionality

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"router_address": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Stable IP address of the advertising router

Format: ipv4
Stable IP address of the advertising router
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

										Default:  stringdefault.StaticString(`0.0.0.0`),
										Computed: true,
									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `MultiProtocol Label Switching-Traffic Engineering (MPLS-TE) parameters

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"parameters": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"abr_type": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `OSPF ABR type

Format: cisco
Cisco ABR type
Format: ibm
IBM ABR type
Format: shortcut
Shortcut ABR type
Format: standard
Standard ABR type
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

										Default:  stringdefault.StaticString(`cisco`),
										Computed: true,
									},

									// TODO handle non-string types
									"opaque_lsa": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Enable the Opaque-LSA capability (rfc2370)

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"rfc1583_compatibility": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Enable RFC1583 criteria for handling AS external routes

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"router_id": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Override default router identifier

Format: ipv4
Router-ID in IP address format
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
								MarkdownDescription: `OSPF specific parameters

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"segment_routing": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									"prefix": schema.MapNestedAttribute{
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"index": schema.SingleNestedAttribute{
													Attributes: map[string]schema.Attribute{
														// TODO handle non-string types
														"value": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Specify the index value of prefix segment/label ID

Format: u32:0-65535
The index segment/label ID value
`,
															// DeprecationMessage:  "",
															// TODO Recreate some of vyos validators for use in leafnodes
															// Validators:          []validator.String(nil),
															// PlanModifiers:       []planmodifier.String(nil),

														},

														// TODO handle non-string types
														"explicit_null": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Request upstream neighbor to replace segment/label with explicit null label

`,
															// DeprecationMessage:  "",
															// TODO Recreate some of vyos validators for use in leafnodes
															// Validators:          []validator.String(nil),
															// PlanModifiers:       []planmodifier.String(nil),

														},

														// TODO handle non-string types
														"no_php_flag": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Do not request penultimate hop popping for segment/label

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
													MarkdownDescription: `Specify the index value of prefix segment/label ID

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
										MarkdownDescription: `Static IPv4 prefix segment/label mapping

Format: ipv4net
IPv4 prefix segment
`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if tagnode defaults can be handled
										// Default:             defaults.Map(nil),
									},

									// TODO handle non-string types
									"maximum_label_depth": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Maximum MPLS labels allowed for this router

Format: u32:1-16
MPLS label depth
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									"global_block": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											// TODO handle non-string types
											"low_label_value": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `MPLS label lower bound

Format: u32:16-1048575
Label value (recommended minimum value: 300)
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"high_label_value": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `MPLS label upper bound

Format: u32:16-1048575
Label value
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
										MarkdownDescription: `Segment Routing Global Block label range

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},

									"local_block": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											// TODO handle non-string types
											"low_label_value": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `MPLS label lower bound

Format: u32:16-1048575
Label value (recommended minimum value: 300)
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"high_label_value": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `MPLS label upper bound

Format: u32:16-1048575
Label value
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
										MarkdownDescription: `Segment Routing Local Block label range

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
								MarkdownDescription: `Segment-Routing (SPRING) settings

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"redistribute": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									"table": schema.MapNestedAttribute{
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												// TODO handle non-string types
												"metric": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `OSPF default metric

Format: u32:0-16777214
Default metric
`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

												},

												// TODO handle non-string types
												"metric_type": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `OSPF metric type for default routes

Format: u32:1-2
Set OSPF External Type 1/2 metrics
`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

													Default:  stringdefault.StaticString(`2`),
													Computed: true,
												},

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
										MarkdownDescription: `Redistribute non-main Kernel Routing Table

Format: u32:1-200
Policy route table number
`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if tagnode defaults can be handled
										// Default:             defaults.Map(nil),
									},

									"bgp": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											// TODO handle non-string types
											"metric": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `OSPF default metric

Format: u32:0-16777214
Default metric
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"metric_type": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `OSPF metric type for default routes

Format: u32:1-2
Set OSPF External Type 1/2 metrics
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

												Default:  stringdefault.StaticString(`2`),
												Computed: true,
											},

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
										MarkdownDescription: `Redistribute BGP routes

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},

									"connected": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											// TODO handle non-string types
											"metric": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `OSPF default metric

Format: u32:0-16777214
Default metric
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"metric_type": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `OSPF metric type for default routes

Format: u32:1-2
Set OSPF External Type 1/2 metrics
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

												Default:  stringdefault.StaticString(`2`),
												Computed: true,
											},

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
										MarkdownDescription: `Redistribute connected routes

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},

									"isis": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											// TODO handle non-string types
											"metric": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `OSPF default metric

Format: u32:0-16777214
Default metric
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"metric_type": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `OSPF metric type for default routes

Format: u32:1-2
Set OSPF External Type 1/2 metrics
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

												Default:  stringdefault.StaticString(`2`),
												Computed: true,
											},

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
										MarkdownDescription: `Redistribute IS-IS routes

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},

									"kernel": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											// TODO handle non-string types
											"metric": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `OSPF default metric

Format: u32:0-16777214
Default metric
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"metric_type": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `OSPF metric type for default routes

Format: u32:1-2
Set OSPF External Type 1/2 metrics
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

												Default:  stringdefault.StaticString(`2`),
												Computed: true,
											},

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
										MarkdownDescription: `Redistribute Kernel routes

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},

									"rip": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											// TODO handle non-string types
											"metric": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `OSPF default metric

Format: u32:0-16777214
Default metric
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"metric_type": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `OSPF metric type for default routes

Format: u32:1-2
Set OSPF External Type 1/2 metrics
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

												Default:  stringdefault.StaticString(`2`),
												Computed: true,
											},

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
										MarkdownDescription: `Redistribute RIP routes

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},

									"babel": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											// TODO handle non-string types
											"metric": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `OSPF default metric

Format: u32:0-16777214
Default metric
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"metric_type": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `OSPF metric type for default routes

Format: u32:1-2
Set OSPF External Type 1/2 metrics
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

												Default:  stringdefault.StaticString(`2`),
												Computed: true,
											},

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
										MarkdownDescription: `Redistribute Babel routes

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},

									"static": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											// TODO handle non-string types
											"metric": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `OSPF default metric

Format: u32:0-16777214
Default metric
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"metric_type": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `OSPF metric type for default routes

Format: u32:1-2
Set OSPF External Type 1/2 metrics
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

												Default:  stringdefault.StaticString(`2`),
												Computed: true,
											},

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
										MarkdownDescription: `Redistribute statically configured routes

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
								MarkdownDescription: `Redistribute information from another routing protocol

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"refresh": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"timers": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Refresh timer

Format: u32:10-1800
Timer value in seconds
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
								MarkdownDescription: `Adjust refresh parameters

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"timers": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									"throttle": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											"spf": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													// TODO handle non-string types
													"delay": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Delay from the first change received to SPF calculation

Format: u32:0-600000
Delay in milliseconds
`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

														Default:  stringdefault.StaticString(`200`),
														Computed: true,
													},

													// TODO handle non-string types
													"initial_holdtime": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Initial hold time between consecutive SPF calculations

Format: u32:0-600000
Initial hold time in milliseconds
`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

														Default:  stringdefault.StaticString(`1000`),
														Computed: true,
													},

													// TODO handle non-string types
													"max_holdtime": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Maximum hold time

Format: u32:0-600000
Max hold time in milliseconds
`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

														Default:  stringdefault.StaticString(`10000`),
														Computed: true,
													},
												},
												// CustomType:          basetypes.MapTypable(nil),
												// Required:            false,
												Optional: true,
												// Computed:            false,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `OSPF SPF timers

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
										MarkdownDescription: `Throttling adaptive timers

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
								MarkdownDescription: `Adjust routing timers

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
						MarkdownDescription: `Open Shortest Path First (OSPF)

`,
						// DeprecationMessage:  "",
						// Validators:          []validator.Map(nil),
						// PlanModifiers:       []planmodifier.Map(nil),
						// TODO investigate if node defaults can be handled
						// Default:             defaults.Map(nil),
					},

					"ospfv3": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"area": schema.MapNestedAttribute{
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"range": schema.MapNestedAttribute{
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													// TODO handle non-string types
													"advertise": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Advertise this range

`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"not_advertise": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Do not advertise this range

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
											MarkdownDescription: `Specify IPv6 prefix (border routers only)

Format: ipv6net
Specify IPv6 prefix (border routers only)
`,
											// DeprecationMessage:  "",
											// Validators:          []validator.Map(nil),
											// PlanModifiers:       []planmodifier.Map(nil),
											// TODO investigate if tagnode defaults can be handled
											// Default:             defaults.Map(nil),
										},

										// TODO handle non-string types
										"export_list": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Name of export-list

`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"import_list": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Name of import-list

`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										"area_type": schema.SingleNestedAttribute{
											Attributes: map[string]schema.Attribute{
												"nssa": schema.SingleNestedAttribute{
													Attributes: map[string]schema.Attribute{
														// TODO handle non-string types
														"default_information_originate": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Originate Type 7 default into NSSA area

`,
															// DeprecationMessage:  "",
															// TODO Recreate some of vyos validators for use in leafnodes
															// Validators:          []validator.String(nil),
															// PlanModifiers:       []planmodifier.String(nil),

														},

														// TODO handle non-string types
														"no_summary": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Do not inject inter-area routes into the stub

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
													MarkdownDescription: `NSSA OSPFv3 area

`,
													// DeprecationMessage:  "",
													// Validators:          []validator.Map(nil),
													// PlanModifiers:       []planmodifier.Map(nil),
													// TODO investigate if node defaults can be handled
													// Default:             defaults.Map(nil),
												},

												"stub": schema.SingleNestedAttribute{
													Attributes: map[string]schema.Attribute{
														// TODO handle non-string types
														"no_summary": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Do not inject inter-area routes into the stub

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
													MarkdownDescription: `Stub OSPFv3 area

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
											MarkdownDescription: `OSPFv3 Area type

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
								MarkdownDescription: `OSPFv3 Area

Format: u32
Area ID as a decimal value
Format: ipv4
Area ID in IP address forma
`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if tagnode defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"interface": schema.MapNestedAttribute{
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										// TODO handle non-string types
										"area": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Enable OSPF on this interface

Format: u32
OSPF area ID as decimal notation
Format: ipv4
OSPF area ID in IP address notation
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"dead_interval": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Interval after which a neighbor is declared dead

Format: u32:1-65535
Neighbor dead interval (seconds)
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

											Default:  stringdefault.StaticString(`40`),
											Computed: true,
										},

										// TODO handle non-string types
										"hello_interval": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Interval between hello packets

Format: u32:1-65535
Hello interval (seconds)
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

											Default:  stringdefault.StaticString(`10`),
											Computed: true,
										},

										// TODO handle non-string types
										"retransmit_interval": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Interval between retransmitting lost link state advertisements

Format: u32:1-65535
Retransmit interval (seconds)
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

											Default:  stringdefault.StaticString(`5`),
											Computed: true,
										},

										// TODO handle non-string types
										"transmit_delay": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Link state transmit delay

Format: u32:1-65535
Link state transmit delay (seconds)
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

											Default:  stringdefault.StaticString(`1`),
											Computed: true,
										},

										// TODO handle non-string types
										"cost": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Interface cost

Format: u32:1-65535
OSPF interface cost
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"mtu_ignore": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Disable Maximum Transmission Unit (MTU) mismatch detection

`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"priority": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Router priority

Format: u32:0-255
OSPF router priority cost
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

											Default:  stringdefault.StaticString(`1`),
											Computed: true,
										},

										// TODO handle non-string types
										"ifmtu": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Interface MTU

Format: u32:1-65535
Interface MTU
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"instance_id": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Instance ID

Format: u32:0-255
Instance Id
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

											Default:  stringdefault.StaticString(`0`),
											Computed: true,
										},

										// TODO handle non-string types
										"network": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Network type

Format: broadcast
Broadcast network type
Format: point-to-point
Point-to-point network type
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
											MarkdownDescription: `Configure passive mode for interface

`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

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
											},
											// CustomType:          basetypes.MapTypable(nil),
											// Required:            false,
											Optional: true,
											// Computed:            false,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Enable Bidirectional Forwarding Detection (BFD)

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
								MarkdownDescription: `Enable routing on an IPv6 interface

Format: txt
Interface used for routing information exchange
`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if tagnode defaults can be handled
								// Default:             defaults.Map(nil),
							},

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

							"auto_cost": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"reference_bandwidth": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Reference bandwidth method to assign cost

Format: u32:1-4294967
Reference bandwidth cost in Mbits/sec
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

										Default:  stringdefault.StaticString(`100`),
										Computed: true,
									},
								},
								// CustomType:          basetypes.MapTypable(nil),
								// Required:            false,
								Optional: true,
								// Computed:            false,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Calculate interface cost according to bandwidth

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"default_information": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									"originate": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											// TODO handle non-string types
											"always": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Always advertise a default route

`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"metric": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `OSPF default metric

Format: u32:0-16777214
Default metric
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"metric_type": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `OSPF metric type for default routes

Format: u32:1-2
Set OSPF External Type 1/2 metrics
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

												Default:  stringdefault.StaticString(`2`),
												Computed: true,
											},

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
										MarkdownDescription: `Distribute a default route

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
								MarkdownDescription: `Default route advertisment settings

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"distance": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"global": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Administrative distance

Format: u32:1-255
Administrative distance
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									"ospfv3": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											// TODO handle non-string types
											"external": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Distance for external routes

Format: u32:1-255
Distance for external routes
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"inter_area": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Distance for inter-area routes

Format: u32:1-255
Distance for inter-area routes
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"intra_area": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Distance for intra-area routes

Format: u32:1-255
Distance for intra-area routes
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
										MarkdownDescription: `OSPFv3 administrative distance

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
								MarkdownDescription: `Administrative distance

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"log_adjacency_changes": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"detail": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Log all state changes

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
								MarkdownDescription: `Log adjacency state changes

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"parameters": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"router_id": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Override default router identifier

Format: ipv4
Router-ID in IP address format
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
								MarkdownDescription: `OSPFv3 specific parameters

`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if node defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"redistribute": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									"bgp": schema.SingleNestedAttribute{
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
										MarkdownDescription: `Redistribute BGP routes

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},

									"connected": schema.SingleNestedAttribute{
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
										MarkdownDescription: `Redistribute connected routes

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},

									"kernel": schema.SingleNestedAttribute{
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
										MarkdownDescription: `Redistribute kernel routes

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},

									"ripng": schema.SingleNestedAttribute{
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
										MarkdownDescription: `Redistribute RIPNG routes

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},

									"babel": schema.SingleNestedAttribute{
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
										MarkdownDescription: `Redistribute Babel routes

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},

									"static": schema.SingleNestedAttribute{
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
										MarkdownDescription: `Redistribute static routes

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
								MarkdownDescription: `Redistribute information from another routing protocol

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
						MarkdownDescription: `Open Shortest Path First (OSPF) for IPv6

`,
						// DeprecationMessage:  "",
						// Validators:          []validator.Map(nil),
						// PlanModifiers:       []planmodifier.Map(nil),
						// TODO investigate if node defaults can be handled
						// Default:             defaults.Map(nil),
					},

					"static": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"route": schema.MapNestedAttribute{
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"interface": schema.MapNestedAttribute{
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													// TODO handle non-string types
													"disable": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Disable instance

`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"distance": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Distance for this route

Format: u32:1-255
Distance for this route
`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"vrf": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `VRF to leak route

Format: txt
Name of VRF to leak to
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
											MarkdownDescription: `Next-hop IPv4 router interface

Format: txt
Gateway interface name
`,
											// DeprecationMessage:  "",
											// Validators:          []validator.Map(nil),
											// PlanModifiers:       []planmodifier.Map(nil),
											// TODO investigate if tagnode defaults can be handled
											// Default:             defaults.Map(nil),
										},

										"next_hop": schema.MapNestedAttribute{
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													// TODO handle non-string types
													"disable": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Disable instance

`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"distance": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Distance for this route

Format: u32:1-255
Distance for this route
`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"interface": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Gateway interface name

Format: txt
Gateway interface name
`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"vrf": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `VRF to leak route

Format: txt
Name of VRF to leak to
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
											MarkdownDescription: `Next-hop IPv4 router address

Format: ipv4
Next-hop router address
`,
											// DeprecationMessage:  "",
											// Validators:          []validator.Map(nil),
											// PlanModifiers:       []planmodifier.Map(nil),
											// TODO investigate if tagnode defaults can be handled
											// Default:             defaults.Map(nil),
										},

										// TODO handle non-string types
										"dhcp_interface": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `DHCP interface supplying next-hop IP address

Format: txt
DHCP interface name
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

										"blackhole": schema.SingleNestedAttribute{
											Attributes: map[string]schema.Attribute{
												// TODO handle non-string types
												"distance": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `Distance for this route

Format: u32:1-255
Distance for this route
`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

												},

												// TODO handle non-string types
												"tag": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `Tag value for this route

Format: u32:1-4294967295
Tag value for this route
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
											MarkdownDescription: `Silently discard pkts when matched

`,
											// DeprecationMessage:  "",
											// Validators:          []validator.Map(nil),
											// PlanModifiers:       []planmodifier.Map(nil),
											// TODO investigate if node defaults can be handled
											// Default:             defaults.Map(nil),
										},

										"reject": schema.SingleNestedAttribute{
											Attributes: map[string]schema.Attribute{
												// TODO handle non-string types
												"distance": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `Distance for this route

Format: u32:1-255
Distance for this route
`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

												},

												// TODO handle non-string types
												"tag": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `Tag value for this route

Format: u32:1-4294967295
Tag value for this route
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
											MarkdownDescription: `Emit an ICMP unreachable when matched

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
								MarkdownDescription: `Static IPv4 route

Format: ipv4net
IPv4 static route
`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if tagnode defaults can be handled
								// Default:             defaults.Map(nil),
							},

							"route6": schema.MapNestedAttribute{
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"interface": schema.MapNestedAttribute{
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													// TODO handle non-string types
													"disable": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Disable instance

`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"distance": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Distance for this route

Format: u32:1-255
Distance for this route
`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"vrf": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `VRF to leak route

Format: txt
Name of VRF to leak to
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
											MarkdownDescription: `IPv6 gateway interface name

Format: txt
Gateway interface name
`,
											// DeprecationMessage:  "",
											// Validators:          []validator.Map(nil),
											// PlanModifiers:       []planmodifier.Map(nil),
											// TODO investigate if tagnode defaults can be handled
											// Default:             defaults.Map(nil),
										},

										"next_hop": schema.MapNestedAttribute{
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													// TODO handle non-string types
													"disable": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Disable instance

`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"distance": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Distance for this route

Format: u32:1-255
Distance for this route
`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"interface": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Gateway interface name

Format: txt
Gateway interface name
`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"vrf": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `VRF to leak route

Format: txt
Name of VRF to leak to
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
											MarkdownDescription: `IPv6 gateway address

Format: ipv6
Next-hop IPv6 router
`,
											// DeprecationMessage:  "",
											// Validators:          []validator.Map(nil),
											// PlanModifiers:       []planmodifier.Map(nil),
											// TODO investigate if tagnode defaults can be handled
											// Default:             defaults.Map(nil),
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

										"blackhole": schema.SingleNestedAttribute{
											Attributes: map[string]schema.Attribute{
												// TODO handle non-string types
												"distance": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `Distance for this route

Format: u32:1-255
Distance for this route
`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

												},

												// TODO handle non-string types
												"tag": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `Tag value for this route

Format: u32:1-4294967295
Tag value for this route
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
											MarkdownDescription: `Silently discard pkts when matched

`,
											// DeprecationMessage:  "",
											// Validators:          []validator.Map(nil),
											// PlanModifiers:       []planmodifier.Map(nil),
											// TODO investigate if node defaults can be handled
											// Default:             defaults.Map(nil),
										},

										"reject": schema.SingleNestedAttribute{
											Attributes: map[string]schema.Attribute{
												// TODO handle non-string types
												"distance": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `Distance for this route

Format: u32:1-255
Distance for this route
`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

												},

												// TODO handle non-string types
												"tag": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `Tag value for this route

Format: u32:1-4294967295
Tag value for this route
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
											MarkdownDescription: `Emit an ICMP unreachable when matched

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
								MarkdownDescription: `Static IPv6 route

Format: ipv6net
IPv6 static route
`,
								// DeprecationMessage:  "",
								// Validators:          []validator.Map(nil),
								// PlanModifiers:       []planmodifier.Map(nil),
								// TODO investigate if tagnode defaults can be handled
								// Default:             defaults.Map(nil),
							},
						},
						// CustomType:          basetypes.MapTypable(nil),
						// Required:            false,
						Optional: true,
						// Computed:            false,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Static Routing

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
				MarkdownDescription: `Routing protocol parameters

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
func (r *vrf_name) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *vrf_nameModel

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
func (r *vrf_name) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *vrf_nameModel

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
func (r *vrf_name) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *vrf_nameModel

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
func (r *vrf_name) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *vrf_nameModel

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
