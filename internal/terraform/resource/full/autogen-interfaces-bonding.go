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
var _ resource.Resource = &interfaces_bonding{}

// var _ resource.ResourceWithImportState = &interfaces_bonding{}

// interfaces_bonding defines the resource implementation.
type interfaces_bonding struct {
	client *http.Client
}

// interfaces bondingModel describes the resource data model.
type interfaces_bondingModel struct {
	ConfigurableAttribute types.String `tfsdk:"configurable_attribute"`
	Defaulted             types.String `tfsdk:"defaulted"`
	ID                    types.String `tfsdk:"id"`

	Vif_s types.Map `tfsdk:vif_s`

	Vif types.Map `tfsdk:vif`

	Arp_monitor types.String `tfsdk:arp_monitor`

	Dhcp_options types.String `tfsdk:dhcp_options`

	Dhcpv__options types.String `tfsdk:dhcpv6_options`

	Mirror types.String `tfsdk:mirror`

	Ip types.String `tfsdk:ip`

	Ipv_ types.String `tfsdk:ipv6`

	Member types.String `tfsdk:member`

	Address types.String `tfsdk:address`

	Description types.String `tfsdk:description`

	Disable_link_detect types.String `tfsdk:disable_link_detect`

	Disable types.String `tfsdk:disable`

	Vrf types.String `tfsdk:vrf`

	Hash_policy types.String `tfsdk:hash_policy`

	Mac types.String `tfsdk:mac`

	Mii_mon_interval types.String `tfsdk:mii_mon_interval`

	Min_links types.String `tfsdk:min_links`

	Lacp_rate types.String `tfsdk:lacp_rate`

	Mode types.String `tfsdk:mode`

	Mtu types.String `tfsdk:mtu`

	Primary types.String `tfsdk:primary`

	Redirect types.String `tfsdk:redirect`

	Xdp types.String `tfsdk:xdp`
}

// Metadata method to define the resource type name.
func (r *interfaces_bonding) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interfaces bonding"
}

// interfaces_bondingResource method to return the example resource reference
func interfaces_bondingResource() resource.Resource {
	return &interfaces_bonding{}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r *interfaces_bonding) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example resource",

		Attributes: map[string]schema.Attribute{
			"vif_s": schema.MapNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"vif_c": schema.MapNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
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
									"address": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `IP address

Format: ipv4net
IPv4 address and prefix length
Format: ipv6net
IPv6 address and prefix length
Format: dhcp
Dynamic Host Configuration Protocol
Format: dhcpv6
Dynamic Host Configuration Protocol for IPv6
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"disable_link_detect": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Ignore link state changes

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
									"mac": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Media Access Control (MAC) address

Format: macaddr
Hardware (MAC) address
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"mtu": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Maximum Transmission Unit (MTU)

Format: u32:68-16000
Maximum Transmission Unit in byte
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

										Default:  stringdefault.StaticString(`1500`),
										Computed: true,
									},

									// TODO handle non-string types
									"redirect": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Redirect incoming packet to destination

Format: txt
Destination interface name
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
										MarkdownDescription: `VRF instance name

Format: txt
VRF instance name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									"dhcp_options": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											// TODO handle non-string types
											"client_id": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Identifier used by client to identify itself to the DHCP server

`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"host_name": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Override system host-name sent to DHCP server

`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"mtu": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Use MTU value from DHCP server - ignore interface setting

`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"vendor_class_id": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Identify the vendor client type to the DHCP server

`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"no_default_route": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Do not install default route to system

`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"default_route_distance": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Distance for installed default route

Format: u32:1-255
Distance for the default route from DHCP server
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

												Default:  stringdefault.StaticString(`210`),
												Computed: true,
											},

											// TODO handle non-string types
											"reject": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `IP addresses or subnets from which to reject DHCP leases

Format: ipv4
IPv4 address to match
Format: ipv4net
IPv4 prefix to match
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},
										},
										// CustomType:          basetypes.MapTypable(nil),
										// Required:            false,
										Optional: true,
										// Computed:            false,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `DHCP client settings/options

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},

									"dhcpv6_options": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											"pd": schema.MapNestedAttribute{
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"interface": schema.MapNestedAttribute{
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	// TODO handle non-string types
																	"address": schema.StringAttribute{
																		// CustomType:          basetypes.StringTypable(nil),
																		// Required:            false,
																		Optional: true,
																		// Sensitive:           false,
																		// Description:         "",
																		MarkdownDescription: `Local interface address assigned to interface (default: EUI-64)

Format: >0
Used to form IPv6 interface address
`,
																		// DeprecationMessage:  "",
																		// TODO Recreate some of vyos validators for use in leafnodes
																		// Validators:          []validator.String(nil),
																		// PlanModifiers:       []planmodifier.String(nil),

																	},

																	// TODO handle non-string types
																	"sla_id": schema.StringAttribute{
																		// CustomType:          basetypes.StringTypable(nil),
																		// Required:            false,
																		Optional: true,
																		// Sensitive:           false,
																		// Description:         "",
																		MarkdownDescription: `Interface site-Level aggregator (SLA)

Format: u32:0-65535
Decimal integer which fits in the length of SLA IDs
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
															MarkdownDescription: `Delegate IPv6 prefix from provider to this interface

`,
															// DeprecationMessage:  "",
															// Validators:          []validator.Map(nil),
															// PlanModifiers:       []planmodifier.Map(nil),
															// TODO investigate if tagnode defaults can be handled
															// Default:             defaults.Map(nil),
														},

														// TODO handle non-string types
														"length": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Request IPv6 prefix length from peer

Format: u32:32-64
Length of delegated prefix
`,
															// DeprecationMessage:  "",
															// TODO Recreate some of vyos validators for use in leafnodes
															// Validators:          []validator.String(nil),
															// PlanModifiers:       []planmodifier.String(nil),

															Default:  stringdefault.StaticString(`64`),
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
												MarkdownDescription: `DHCPv6 prefix delegation interface statement

Format: instance number
Prefix delegation instance (>= 0)
`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if tagnode defaults can be handled
												// Default:             defaults.Map(nil),
											},

											// TODO handle non-string types
											"duid": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `DHCP unique identifier (DUID) to be sent by dhcpv6 client

Format: duid
DHCP unique identifier (DUID)
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"parameters_only": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Acquire only config parameters, no address

`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"rapid_commit": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Wait for immediate reply instead of advertisements

`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"temporary": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `IPv6 temporary address

`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},
										},
										// CustomType:          basetypes.MapTypable(nil),
										// Required:            false,
										Optional: true,
										// Computed:            false,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `DHCPv6 client settings/options

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},

									"ip": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											// TODO handle non-string types
											"adjust_mss": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Adjust TCP MSS value

Format: clamp-mss-to-pmtu
Automatically sets the MSS to the proper value
Format: u32:536-65535
TCP Maximum segment size in bytes
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"arp_cache_timeout": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `ARP cache entry timeout in seconds

Format: u32:1-86400
ARP cache entry timout in seconds
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

												Default:  stringdefault.StaticString(`30`),
												Computed: true,
											},

											// TODO handle non-string types
											"disable_arp_filter": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Disable ARP filter on this interface

`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

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

											// TODO handle non-string types
											"enable_directed_broadcast": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Enable directed broadcast forwarding on this interface

`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"enable_arp_accept": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Enable ARP accept on this interface

`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"enable_arp_announce": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Enable ARP announce on this interface

`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"enable_arp_ignore": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Enable ARP ignore on this interface

`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"enable_proxy_arp": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Enable proxy-arp on this interface

`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"proxy_arp_pvlan": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Enable private VLAN proxy ARP on this interface

`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"source_validation": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Source validation by reversed path (RFC3704)

Format: strict
Enable Strict Reverse Path Forwarding as defined in RFC3704
Format: loose
Enable Loose Reverse Path Forwarding as defined in RFC3704
Format: disable
No source validation
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
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
											"adjust_mss": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Adjust TCP MSS value

Format: clamp-mss-to-pmtu
Automatically sets the MSS to the proper value
Format: u32:536-65535
TCP Maximum segment size in bytes
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

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

											// TODO handle non-string types
											"dup_addr_detect_transmits": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Number of NS messages to send while performing DAD (default: 1)

Format: u32:0
Disable Duplicate Address Dectection (DAD)
Format: u32:1-n
Number of NS messages to send while performing DAD
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											"address": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													// TODO handle non-string types
													"autoconf": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Enable acquisition of IPv6 address using stateless autoconfig (SLAAC)

`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"eui64": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Prefix for IPv6 address with MAC-based EUI-64

Format: <h:h:h:h:h:h:h:h/64>
IPv6 /64 network
`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},

													// TODO handle non-string types
													"no_default_link_local": schema.StringAttribute{
														// CustomType:          basetypes.StringTypable(nil),
														// Required:            false,
														Optional: true,
														// Sensitive:           false,
														// Description:         "",
														MarkdownDescription: `Remove the default link-local address from the interface

`,
														// DeprecationMessage:  "",
														// TODO Recreate some of vyos validators for use in leafnodes
														// Validators:          []validator.String(nil),
														// PlanModifiers:       []planmodifier.String(nil),

													},
												},
												// CustomType:          basetypes.MapTypable(nil),
												// Required:            false,
												Optional: true,
												// Computed:            false,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `IPv6 address configuration modes

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
										MarkdownDescription: `IPv6 routing parameters

`,
										// DeprecationMessage:  "",
										// Validators:          []validator.Map(nil),
										// PlanModifiers:       []planmodifier.Map(nil),
										// TODO investigate if node defaults can be handled
										// Default:             defaults.Map(nil),
									},

									"mirror": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											// TODO handle non-string types
											"ingress": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Mirror ingress traffic to destination interface

Format: txt
Destination interface name
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"egress": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Mirror egress traffic to destination interface

Format: txt
Destination interface name
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},
										},
										// CustomType:          basetypes.MapTypable(nil),
										// Required:            false,
										Optional: true,
										// Computed:            false,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Mirror ingress/egress packets

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
							MarkdownDescription: `QinQ TAG-C Virtual Local Area Network (VLAN) ID

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
						"address": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `IP address

Format: ipv4net
IPv4 address and prefix length
Format: ipv6net
IPv6 address and prefix length
Format: dhcp
Dynamic Host Configuration Protocol
Format: dhcpv6
Dynamic Host Configuration Protocol for IPv6
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"disable_link_detect": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Ignore link state changes

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
						"protocol": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Protocol used for service VLAN (default: 802.1ad)

Format: 802.1ad
Provider Bridging (IEEE 802.1ad, Q-inQ), ethertype 0x88a8
Format: 802.1q
VLAN-tagged frame (IEEE 802.1q), ethertype 0x8100
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

							Default:  stringdefault.StaticString(`802.1ad`),
							Computed: true,
						},

						// TODO handle non-string types
						"mac": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Media Access Control (MAC) address

Format: macaddr
Hardware (MAC) address
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"mtu": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Maximum Transmission Unit (MTU)

Format: u32:68-16000
Maximum Transmission Unit in byte
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

							Default:  stringdefault.StaticString(`1500`),
							Computed: true,
						},

						// TODO handle non-string types
						"redirect": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Redirect incoming packet to destination

Format: txt
Destination interface name
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
							MarkdownDescription: `VRF instance name

Format: txt
VRF instance name
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						"dhcp_options": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								// TODO handle non-string types
								"client_id": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Identifier used by client to identify itself to the DHCP server

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"host_name": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Override system host-name sent to DHCP server

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"mtu": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Use MTU value from DHCP server - ignore interface setting

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"vendor_class_id": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Identify the vendor client type to the DHCP server

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"no_default_route": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Do not install default route to system

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"default_route_distance": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Distance for installed default route

Format: u32:1-255
Distance for the default route from DHCP server
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

									Default:  stringdefault.StaticString(`210`),
									Computed: true,
								},

								// TODO handle non-string types
								"reject": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `IP addresses or subnets from which to reject DHCP leases

Format: ipv4
IPv4 address to match
Format: ipv4net
IPv4 prefix to match
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},
							},
							// CustomType:          basetypes.MapTypable(nil),
							// Required:            false,
							Optional: true,
							// Computed:            false,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `DHCP client settings/options

`,
							// DeprecationMessage:  "",
							// Validators:          []validator.Map(nil),
							// PlanModifiers:       []planmodifier.Map(nil),
							// TODO investigate if node defaults can be handled
							// Default:             defaults.Map(nil),
						},

						"dhcpv6_options": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"pd": schema.MapNestedAttribute{
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"interface": schema.MapNestedAttribute{
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														// TODO handle non-string types
														"address": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Local interface address assigned to interface (default: EUI-64)

Format: >0
Used to form IPv6 interface address
`,
															// DeprecationMessage:  "",
															// TODO Recreate some of vyos validators for use in leafnodes
															// Validators:          []validator.String(nil),
															// PlanModifiers:       []planmodifier.String(nil),

														},

														// TODO handle non-string types
														"sla_id": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Interface site-Level aggregator (SLA)

Format: u32:0-65535
Decimal integer which fits in the length of SLA IDs
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
												MarkdownDescription: `Delegate IPv6 prefix from provider to this interface

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if tagnode defaults can be handled
												// Default:             defaults.Map(nil),
											},

											// TODO handle non-string types
											"length": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Request IPv6 prefix length from peer

Format: u32:32-64
Length of delegated prefix
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

												Default:  stringdefault.StaticString(`64`),
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
									MarkdownDescription: `DHCPv6 prefix delegation interface statement

Format: instance number
Prefix delegation instance (>= 0)
`,
									// DeprecationMessage:  "",
									// Validators:          []validator.Map(nil),
									// PlanModifiers:       []planmodifier.Map(nil),
									// TODO investigate if tagnode defaults can be handled
									// Default:             defaults.Map(nil),
								},

								// TODO handle non-string types
								"duid": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `DHCP unique identifier (DUID) to be sent by dhcpv6 client

Format: duid
DHCP unique identifier (DUID)
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"parameters_only": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Acquire only config parameters, no address

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"rapid_commit": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Wait for immediate reply instead of advertisements

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"temporary": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `IPv6 temporary address

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},
							},
							// CustomType:          basetypes.MapTypable(nil),
							// Required:            false,
							Optional: true,
							// Computed:            false,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `DHCPv6 client settings/options

`,
							// DeprecationMessage:  "",
							// Validators:          []validator.Map(nil),
							// PlanModifiers:       []planmodifier.Map(nil),
							// TODO investigate if node defaults can be handled
							// Default:             defaults.Map(nil),
						},

						"ip": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								// TODO handle non-string types
								"adjust_mss": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Adjust TCP MSS value

Format: clamp-mss-to-pmtu
Automatically sets the MSS to the proper value
Format: u32:536-65535
TCP Maximum segment size in bytes
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"arp_cache_timeout": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `ARP cache entry timeout in seconds

Format: u32:1-86400
ARP cache entry timout in seconds
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

									Default:  stringdefault.StaticString(`30`),
									Computed: true,
								},

								// TODO handle non-string types
								"disable_arp_filter": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Disable ARP filter on this interface

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

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

								// TODO handle non-string types
								"enable_directed_broadcast": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Enable directed broadcast forwarding on this interface

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"enable_arp_accept": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Enable ARP accept on this interface

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"enable_arp_announce": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Enable ARP announce on this interface

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"enable_arp_ignore": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Enable ARP ignore on this interface

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"enable_proxy_arp": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Enable proxy-arp on this interface

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"proxy_arp_pvlan": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Enable private VLAN proxy ARP on this interface

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"source_validation": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Source validation by reversed path (RFC3704)

Format: strict
Enable Strict Reverse Path Forwarding as defined in RFC3704
Format: loose
Enable Loose Reverse Path Forwarding as defined in RFC3704
Format: disable
No source validation
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
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
								"adjust_mss": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Adjust TCP MSS value

Format: clamp-mss-to-pmtu
Automatically sets the MSS to the proper value
Format: u32:536-65535
TCP Maximum segment size in bytes
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

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

								// TODO handle non-string types
								"dup_addr_detect_transmits": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Number of NS messages to send while performing DAD (default: 1)

Format: u32:0
Disable Duplicate Address Dectection (DAD)
Format: u32:1-n
Number of NS messages to send while performing DAD
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								"address": schema.SingleNestedAttribute{
									Attributes: map[string]schema.Attribute{
										// TODO handle non-string types
										"autoconf": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Enable acquisition of IPv6 address using stateless autoconfig (SLAAC)

`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"eui64": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Prefix for IPv6 address with MAC-based EUI-64

Format: <h:h:h:h:h:h:h:h/64>
IPv6 /64 network
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"no_default_link_local": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Remove the default link-local address from the interface

`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},
									},
									// CustomType:          basetypes.MapTypable(nil),
									// Required:            false,
									Optional: true,
									// Computed:            false,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `IPv6 address configuration modes

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
							MarkdownDescription: `IPv6 routing parameters

`,
							// DeprecationMessage:  "",
							// Validators:          []validator.Map(nil),
							// PlanModifiers:       []planmodifier.Map(nil),
							// TODO investigate if node defaults can be handled
							// Default:             defaults.Map(nil),
						},

						"mirror": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								// TODO handle non-string types
								"ingress": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Mirror ingress traffic to destination interface

Format: txt
Destination interface name
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"egress": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Mirror egress traffic to destination interface

Format: txt
Destination interface name
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},
							},
							// CustomType:          basetypes.MapTypable(nil),
							// Required:            false,
							Optional: true,
							// Computed:            false,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Mirror ingress/egress packets

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
				MarkdownDescription: `QinQ TAG-S Virtual Local Area Network (VLAN) ID

Format: u32:0-4094
QinQ Virtual Local Area Network (VLAN) ID
`,
				// DeprecationMessage:  "",
				// Validators:          []validator.Map(nil),
				// PlanModifiers:       []planmodifier.Map(nil),
				// TODO investigate if tagnode defaults can be handled
				// Default:             defaults.Map(nil),
			},

			"vif": schema.MapNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
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
						"address": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `IP address

Format: ipv4net
IPv4 address and prefix length
Format: ipv6net
IPv6 address and prefix length
Format: dhcp
Dynamic Host Configuration Protocol
Format: dhcpv6
Dynamic Host Configuration Protocol for IPv6
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"disable_link_detect": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Ignore link state changes

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
						"egress_qos": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `VLAN egress QoS

Format: txt
Format for qos mapping, e.g.: '0:1 1:6 7:6'
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"ingress_qos": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `VLAN ingress QoS

Format: txt
Format for qos mapping, e.g.: '0:1 1:6 7:6'
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"mac": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Media Access Control (MAC) address

Format: macaddr
Hardware (MAC) address
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"mtu": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Maximum Transmission Unit (MTU)

Format: u32:68-16000
Maximum Transmission Unit in byte
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

							Default:  stringdefault.StaticString(`1500`),
							Computed: true,
						},

						// TODO handle non-string types
						"redirect": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Redirect incoming packet to destination

Format: txt
Destination interface name
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
							MarkdownDescription: `VRF instance name

Format: txt
VRF instance name
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						"dhcp_options": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								// TODO handle non-string types
								"client_id": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Identifier used by client to identify itself to the DHCP server

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"host_name": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Override system host-name sent to DHCP server

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"mtu": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Use MTU value from DHCP server - ignore interface setting

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"vendor_class_id": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Identify the vendor client type to the DHCP server

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"no_default_route": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Do not install default route to system

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"default_route_distance": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Distance for installed default route

Format: u32:1-255
Distance for the default route from DHCP server
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

									Default:  stringdefault.StaticString(`210`),
									Computed: true,
								},

								// TODO handle non-string types
								"reject": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `IP addresses or subnets from which to reject DHCP leases

Format: ipv4
IPv4 address to match
Format: ipv4net
IPv4 prefix to match
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},
							},
							// CustomType:          basetypes.MapTypable(nil),
							// Required:            false,
							Optional: true,
							// Computed:            false,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `DHCP client settings/options

`,
							// DeprecationMessage:  "",
							// Validators:          []validator.Map(nil),
							// PlanModifiers:       []planmodifier.Map(nil),
							// TODO investigate if node defaults can be handled
							// Default:             defaults.Map(nil),
						},

						"dhcpv6_options": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"pd": schema.MapNestedAttribute{
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"interface": schema.MapNestedAttribute{
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														// TODO handle non-string types
														"address": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Local interface address assigned to interface (default: EUI-64)

Format: >0
Used to form IPv6 interface address
`,
															// DeprecationMessage:  "",
															// TODO Recreate some of vyos validators for use in leafnodes
															// Validators:          []validator.String(nil),
															// PlanModifiers:       []planmodifier.String(nil),

														},

														// TODO handle non-string types
														"sla_id": schema.StringAttribute{
															// CustomType:          basetypes.StringTypable(nil),
															// Required:            false,
															Optional: true,
															// Sensitive:           false,
															// Description:         "",
															MarkdownDescription: `Interface site-Level aggregator (SLA)

Format: u32:0-65535
Decimal integer which fits in the length of SLA IDs
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
												MarkdownDescription: `Delegate IPv6 prefix from provider to this interface

`,
												// DeprecationMessage:  "",
												// Validators:          []validator.Map(nil),
												// PlanModifiers:       []planmodifier.Map(nil),
												// TODO investigate if tagnode defaults can be handled
												// Default:             defaults.Map(nil),
											},

											// TODO handle non-string types
											"length": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Request IPv6 prefix length from peer

Format: u32:32-64
Length of delegated prefix
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

												Default:  stringdefault.StaticString(`64`),
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
									MarkdownDescription: `DHCPv6 prefix delegation interface statement

Format: instance number
Prefix delegation instance (>= 0)
`,
									// DeprecationMessage:  "",
									// Validators:          []validator.Map(nil),
									// PlanModifiers:       []planmodifier.Map(nil),
									// TODO investigate if tagnode defaults can be handled
									// Default:             defaults.Map(nil),
								},

								// TODO handle non-string types
								"duid": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `DHCP unique identifier (DUID) to be sent by dhcpv6 client

Format: duid
DHCP unique identifier (DUID)
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"parameters_only": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Acquire only config parameters, no address

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"rapid_commit": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Wait for immediate reply instead of advertisements

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"temporary": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `IPv6 temporary address

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},
							},
							// CustomType:          basetypes.MapTypable(nil),
							// Required:            false,
							Optional: true,
							// Computed:            false,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `DHCPv6 client settings/options

`,
							// DeprecationMessage:  "",
							// Validators:          []validator.Map(nil),
							// PlanModifiers:       []planmodifier.Map(nil),
							// TODO investigate if node defaults can be handled
							// Default:             defaults.Map(nil),
						},

						"ip": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								// TODO handle non-string types
								"adjust_mss": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Adjust TCP MSS value

Format: clamp-mss-to-pmtu
Automatically sets the MSS to the proper value
Format: u32:536-65535
TCP Maximum segment size in bytes
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"arp_cache_timeout": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `ARP cache entry timeout in seconds

Format: u32:1-86400
ARP cache entry timout in seconds
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

									Default:  stringdefault.StaticString(`30`),
									Computed: true,
								},

								// TODO handle non-string types
								"disable_arp_filter": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Disable ARP filter on this interface

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

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

								// TODO handle non-string types
								"enable_directed_broadcast": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Enable directed broadcast forwarding on this interface

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"enable_arp_accept": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Enable ARP accept on this interface

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"enable_arp_announce": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Enable ARP announce on this interface

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"enable_arp_ignore": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Enable ARP ignore on this interface

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"enable_proxy_arp": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Enable proxy-arp on this interface

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"proxy_arp_pvlan": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Enable private VLAN proxy ARP on this interface

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"source_validation": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Source validation by reversed path (RFC3704)

Format: strict
Enable Strict Reverse Path Forwarding as defined in RFC3704
Format: loose
Enable Loose Reverse Path Forwarding as defined in RFC3704
Format: disable
No source validation
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
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
								"adjust_mss": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Adjust TCP MSS value

Format: clamp-mss-to-pmtu
Automatically sets the MSS to the proper value
Format: u32:536-65535
TCP Maximum segment size in bytes
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

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

								// TODO handle non-string types
								"dup_addr_detect_transmits": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Number of NS messages to send while performing DAD (default: 1)

Format: u32:0
Disable Duplicate Address Dectection (DAD)
Format: u32:1-n
Number of NS messages to send while performing DAD
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								"address": schema.SingleNestedAttribute{
									Attributes: map[string]schema.Attribute{
										// TODO handle non-string types
										"autoconf": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Enable acquisition of IPv6 address using stateless autoconfig (SLAAC)

`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"eui64": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Prefix for IPv6 address with MAC-based EUI-64

Format: <h:h:h:h:h:h:h:h/64>
IPv6 /64 network
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"no_default_link_local": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Remove the default link-local address from the interface

`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},
									},
									// CustomType:          basetypes.MapTypable(nil),
									// Required:            false,
									Optional: true,
									// Computed:            false,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `IPv6 address configuration modes

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
							MarkdownDescription: `IPv6 routing parameters

`,
							// DeprecationMessage:  "",
							// Validators:          []validator.Map(nil),
							// PlanModifiers:       []planmodifier.Map(nil),
							// TODO investigate if node defaults can be handled
							// Default:             defaults.Map(nil),
						},

						"mirror": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								// TODO handle non-string types
								"ingress": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Mirror ingress traffic to destination interface

Format: txt
Destination interface name
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"egress": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Mirror egress traffic to destination interface

Format: txt
Destination interface name
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},
							},
							// CustomType:          basetypes.MapTypable(nil),
							// Required:            false,
							Optional: true,
							// Computed:            false,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Mirror ingress/egress packets

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
				MarkdownDescription: `Virtual Local Area Network (VLAN) ID

Format: u32:0-4094
Virtual Local Area Network (VLAN) ID
`,
				// DeprecationMessage:  "",
				// Validators:          []validator.Map(nil),
				// PlanModifiers:       []planmodifier.Map(nil),
				// TODO investigate if tagnode defaults can be handled
				// Default:             defaults.Map(nil),
			},

			// TODO handle non-string types
			"address": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `IP address

Format: ipv4net
IPv4 address and prefix length
Format: ipv6net
IPv6 address and prefix length
Format: dhcp
Dynamic Host Configuration Protocol
Format: dhcpv6
Dynamic Host Configuration Protocol for IPv6
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
			"disable_link_detect": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Ignore link state changes

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
			"vrf": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `VRF instance name

Format: txt
VRF instance name
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"hash_policy": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Bonding transmit hash policy

Format: layer2
use MAC addresses to generate the hash
Format: layer2+3
combine MAC address and IP address to make hash
Format: layer3+4
combine IP address and port to make hash
Format: encap2+3
combine encapsulated MAC address and IP address to make hash
Format: encap3+4
combine encapsulated IP address and port to make hash
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

				Default:  stringdefault.StaticString(`layer2`),
				Computed: true,
			},

			// TODO handle non-string types
			"mac": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Media Access Control (MAC) address

Format: macaddr
Hardware (MAC) address
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"mii_mon_interval": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Specifies the MII link monitoring frequency in milliseconds

Format: u32:0
Disable MII link monitoring
Format: u32:50-1000
MII link monitoring frequency in milliseconds
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

				Default:  stringdefault.StaticString(`100`),
				Computed: true,
			},

			// TODO handle non-string types
			"min_links": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Minimum number of member interfaces required up before enabling bond

Format: u32:0-16
Minimum number of member interfaces required up before enabling bond
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

				Default:  stringdefault.StaticString(`0`),
				Computed: true,
			},

			// TODO handle non-string types
			"lacp_rate": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Rate in which we will ask our link partner to transmit LACPDU packets

Format: slow
Request partner to transmit LACPDUs every 30 seconds
Format: fast
Request partner to transmit LACPDUs every 1 second
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

				Default:  stringdefault.StaticString(`slow`),
				Computed: true,
			},

			// TODO handle non-string types
			"mode": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Bonding mode

Format: 802.3ad
IEEE 802.3ad Dynamic link aggregation
Format: active-backup
Fault tolerant: only one slave in the bond is active
Format: broadcast
Fault tolerant: transmits everything on all slave interfaces
Format: round-robin
Load balance: transmit packets in sequential order
Format: transmit-load-balance
Load balance: adapts based on transmit load and speed
Format: adaptive-load-balance
Load balance: adapts based on transmit and receive plus ARP
Format: xor-hash
Distribute based on MAC address
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

				Default:  stringdefault.StaticString(`802.3ad`),
				Computed: true,
			},

			// TODO handle non-string types
			"mtu": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Maximum Transmission Unit (MTU)

Format: u32:68-16000
Maximum Transmission Unit in byte
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

				Default:  stringdefault.StaticString(`1500`),
				Computed: true,
			},

			// TODO handle non-string types
			"primary": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Primary device interface

Format: txt
Interface name
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"redirect": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Redirect incoming packet to destination

Format: txt
Destination interface name
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"xdp": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Enable eXpress Data Path

`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			"arp_monitor": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					// TODO handle non-string types
					"interval": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `ARP link monitoring interval

Format: u32
Specifies the ARP link monitoring frequency in milliseconds
`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},

					// TODO handle non-string types
					"target": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `IP address used for ARP monitoring

Format: ipv4
Specify IPv4 address of ARP requests when interval is enabled
`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},
				},
				// CustomType:          basetypes.MapTypable(nil),
				// Required:            false,
				Optional: true,
				// Computed:            false,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `ARP link monitoring parameters

`,
				// DeprecationMessage:  "",
				// Validators:          []validator.Map(nil),
				// PlanModifiers:       []planmodifier.Map(nil),
				// TODO investigate if node defaults can be handled
				// Default:             defaults.Map(nil),
			},

			"dhcp_options": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					// TODO handle non-string types
					"client_id": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Identifier used by client to identify itself to the DHCP server

`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},

					// TODO handle non-string types
					"host_name": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Override system host-name sent to DHCP server

`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},

					// TODO handle non-string types
					"mtu": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Use MTU value from DHCP server - ignore interface setting

`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},

					// TODO handle non-string types
					"vendor_class_id": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Identify the vendor client type to the DHCP server

`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},

					// TODO handle non-string types
					"no_default_route": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Do not install default route to system

`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},

					// TODO handle non-string types
					"default_route_distance": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Distance for installed default route

Format: u32:1-255
Distance for the default route from DHCP server
`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

						Default:  stringdefault.StaticString(`210`),
						Computed: true,
					},

					// TODO handle non-string types
					"reject": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `IP addresses or subnets from which to reject DHCP leases

Format: ipv4
IPv4 address to match
Format: ipv4net
IPv4 prefix to match
`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},
				},
				// CustomType:          basetypes.MapTypable(nil),
				// Required:            false,
				Optional: true,
				// Computed:            false,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `DHCP client settings/options

`,
				// DeprecationMessage:  "",
				// Validators:          []validator.Map(nil),
				// PlanModifiers:       []planmodifier.Map(nil),
				// TODO investigate if node defaults can be handled
				// Default:             defaults.Map(nil),
			},

			"dhcpv6_options": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"pd": schema.MapNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"interface": schema.MapNestedAttribute{
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											// TODO handle non-string types
											"address": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Local interface address assigned to interface (default: EUI-64)

Format: >0
Used to form IPv6 interface address
`,
												// DeprecationMessage:  "",
												// TODO Recreate some of vyos validators for use in leafnodes
												// Validators:          []validator.String(nil),
												// PlanModifiers:       []planmodifier.String(nil),

											},

											// TODO handle non-string types
											"sla_id": schema.StringAttribute{
												// CustomType:          basetypes.StringTypable(nil),
												// Required:            false,
												Optional: true,
												// Sensitive:           false,
												// Description:         "",
												MarkdownDescription: `Interface site-Level aggregator (SLA)

Format: u32:0-65535
Decimal integer which fits in the length of SLA IDs
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
									MarkdownDescription: `Delegate IPv6 prefix from provider to this interface

`,
									// DeprecationMessage:  "",
									// Validators:          []validator.Map(nil),
									// PlanModifiers:       []planmodifier.Map(nil),
									// TODO investigate if tagnode defaults can be handled
									// Default:             defaults.Map(nil),
								},

								// TODO handle non-string types
								"length": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Request IPv6 prefix length from peer

Format: u32:32-64
Length of delegated prefix
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

									Default:  stringdefault.StaticString(`64`),
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
						MarkdownDescription: `DHCPv6 prefix delegation interface statement

Format: instance number
Prefix delegation instance (>= 0)
`,
						// DeprecationMessage:  "",
						// Validators:          []validator.Map(nil),
						// PlanModifiers:       []planmodifier.Map(nil),
						// TODO investigate if tagnode defaults can be handled
						// Default:             defaults.Map(nil),
					},

					// TODO handle non-string types
					"duid": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `DHCP unique identifier (DUID) to be sent by dhcpv6 client

Format: duid
DHCP unique identifier (DUID)
`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},

					// TODO handle non-string types
					"parameters_only": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Acquire only config parameters, no address

`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},

					// TODO handle non-string types
					"rapid_commit": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Wait for immediate reply instead of advertisements

`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},

					// TODO handle non-string types
					"temporary": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `IPv6 temporary address

`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},
				},
				// CustomType:          basetypes.MapTypable(nil),
				// Required:            false,
				Optional: true,
				// Computed:            false,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `DHCPv6 client settings/options

`,
				// DeprecationMessage:  "",
				// Validators:          []validator.Map(nil),
				// PlanModifiers:       []planmodifier.Map(nil),
				// TODO investigate if node defaults can be handled
				// Default:             defaults.Map(nil),
			},

			"mirror": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					// TODO handle non-string types
					"ingress": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Mirror ingress traffic to destination interface

Format: txt
Destination interface name
`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},

					// TODO handle non-string types
					"egress": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Mirror egress traffic to destination interface

Format: txt
Destination interface name
`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},
				},
				// CustomType:          basetypes.MapTypable(nil),
				// Required:            false,
				Optional: true,
				// Computed:            false,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Mirror ingress/egress packets

`,
				// DeprecationMessage:  "",
				// Validators:          []validator.Map(nil),
				// PlanModifiers:       []planmodifier.Map(nil),
				// TODO investigate if node defaults can be handled
				// Default:             defaults.Map(nil),
			},

			"ip": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					// TODO handle non-string types
					"adjust_mss": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Adjust TCP MSS value

Format: clamp-mss-to-pmtu
Automatically sets the MSS to the proper value
Format: u32:536-65535
TCP Maximum segment size in bytes
`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},

					// TODO handle non-string types
					"arp_cache_timeout": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `ARP cache entry timeout in seconds

Format: u32:1-86400
ARP cache entry timout in seconds
`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

						Default:  stringdefault.StaticString(`30`),
						Computed: true,
					},

					// TODO handle non-string types
					"disable_arp_filter": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Disable ARP filter on this interface

`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},

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

					// TODO handle non-string types
					"enable_directed_broadcast": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Enable directed broadcast forwarding on this interface

`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},

					// TODO handle non-string types
					"enable_arp_accept": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Enable ARP accept on this interface

`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},

					// TODO handle non-string types
					"enable_arp_announce": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Enable ARP announce on this interface

`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},

					// TODO handle non-string types
					"enable_arp_ignore": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Enable ARP ignore on this interface

`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},

					// TODO handle non-string types
					"enable_proxy_arp": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Enable proxy-arp on this interface

`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},

					// TODO handle non-string types
					"proxy_arp_pvlan": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Enable private VLAN proxy ARP on this interface

`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},

					// TODO handle non-string types
					"source_validation": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Source validation by reversed path (RFC3704)

Format: strict
Enable Strict Reverse Path Forwarding as defined in RFC3704
Format: loose
Enable Loose Reverse Path Forwarding as defined in RFC3704
Format: disable
No source validation
`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
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
					"adjust_mss": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Adjust TCP MSS value

Format: clamp-mss-to-pmtu
Automatically sets the MSS to the proper value
Format: u32:536-65535
TCP Maximum segment size in bytes
`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},

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

					// TODO handle non-string types
					"dup_addr_detect_transmits": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Number of NS messages to send while performing DAD (default: 1)

Format: u32:0
Disable Duplicate Address Dectection (DAD)
Format: u32:1-n
Number of NS messages to send while performing DAD
`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},

					"address": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							// TODO handle non-string types
							"autoconf": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Enable acquisition of IPv6 address using stateless autoconfig (SLAAC)

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"eui64": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Prefix for IPv6 address with MAC-based EUI-64

Format: <h:h:h:h:h:h:h:h/64>
IPv6 /64 network
`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},

							// TODO handle non-string types
							"no_default_link_local": schema.StringAttribute{
								// CustomType:          basetypes.StringTypable(nil),
								// Required:            false,
								Optional: true,
								// Sensitive:           false,
								// Description:         "",
								MarkdownDescription: `Remove the default link-local address from the interface

`,
								// DeprecationMessage:  "",
								// TODO Recreate some of vyos validators for use in leafnodes
								// Validators:          []validator.String(nil),
								// PlanModifiers:       []planmodifier.String(nil),

							},
						},
						// CustomType:          basetypes.MapTypable(nil),
						// Required:            false,
						Optional: true,
						// Computed:            false,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `IPv6 address configuration modes

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
				MarkdownDescription: `IPv6 routing parameters

`,
				// DeprecationMessage:  "",
				// Validators:          []validator.Map(nil),
				// PlanModifiers:       []planmodifier.Map(nil),
				// TODO investigate if node defaults can be handled
				// Default:             defaults.Map(nil),
			},

			"member": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					// TODO handle non-string types
					"interface": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Member interface name

Format: txt
Interface name
`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},
				},
				// CustomType:          basetypes.MapTypable(nil),
				// Required:            false,
				Optional: true,
				// Computed:            false,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Bridge member interfaces

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
func (r *interfaces_bonding) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *interfaces_bondingModel

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
func (r *interfaces_bonding) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *interfaces_bondingModel

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
func (r *interfaces_bonding) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *interfaces_bondingModel

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
func (r *interfaces_bonding) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *interfaces_bondingModel

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
