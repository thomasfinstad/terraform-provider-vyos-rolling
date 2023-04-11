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
var _ resource.Resource = &firewall_ipv__name{}

// var _ resource.ResourceWithImportState = &firewall_ipv__name{}

// firewall_ipv__name defines the resource implementation.
type firewall_ipv__name struct {
	client *http.Client
}

// firewall ipv6-nameModel describes the resource data model.
type firewall_ipv__nameModel struct {
	ConfigurableAttribute types.String `tfsdk:"configurable_attribute"`
	Defaulted             types.String `tfsdk:"defaulted"`
	ID                    types.String `tfsdk:"id"`

	Rule types.Map `tfsdk:rule`

	Default_action types.String `tfsdk:default_action`

	Enable_default_log types.String `tfsdk:enable_default_log`

	Description types.String `tfsdk:description`

	Default_jump_target types.String `tfsdk:default_jump_target`
}

// Metadata method to define the resource type name.
func (r *firewall_ipv__name) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_firewall ipv6-name"
}

// firewall_ipv__nameResource method to return the example resource reference
func firewall_ipv__nameResource() resource.Resource {
	return &firewall_ipv__name{}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r *firewall_ipv__name) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example resource",

		Attributes: map[string]schema.Attribute{
			"rule": schema.MapNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						// TODO handle non-string types
						"action": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Rule action

Format: accept
Accept matching entries
Format: jump
Jump to another chain
Format: reject
Reject matching entries
Format: return
Return from the current chain and continue at the next rule of the
                      last chain
Format: drop
Drop matching entries
Format: queue
Enqueue packet to userspace
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
						"disable": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Option to disable firewall rule

`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"log": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Option to log packets matching rule

Format: enable
Enable log
Format: disable
Disable log
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"log_level": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Set log-level. Log must be enable.

Format: emerg
Emerg log level
Format: alert
Alert log level
Format: crit
Critical log level
Format: err
Error log level
Format: warn
Warning log level
Format: notice
Notice log level
Format: info
Info log level
Format: debug
Debug log level
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
							MarkdownDescription: `Protocol to match (protocol name, number, or "all")

Format: all
All IP protocols
Format: tcp_udp
Both TCP and UDP
Format: u32:0-255
IP protocol number
Format: <protocol>
IP protocol name
Format: !<protocol>
IP protocol name
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"dscp": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `DSCP value

Format: u32:0-63
DSCP value to match
Format: <start-end>
DSCP range to match
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"dscp_exclude": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `DSCP value not to match

Format: u32:0-63
DSCP value not to match
Format: <start-end>
DSCP range not to match
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"packet_length": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Payload size in bytes, including header and data to match

Format: u32:1-65535
Packet length to match
Format: <start-end>
Packet length range to match
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"packet_length_exclude": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Payload size in bytes, including header and data not to match

Format: u32:1-65535
Packet length not to match
Format: <start-end>
Packet length range not to match
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"packet_type": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Packet type

Format: broadcast
Match broadcast packet type
Format: host
Match host packet type, addressed to local host
Format: multicast
Match multicast packet type
Format: other
Match packet addressed to another host
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"connection_mark": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Connection mark

Format: u32:0-2147483647
Connection-mark to match
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"jump_target": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Set jump target. Action jump must be defined to use this setting

`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"queue": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Queue target to use. Action queue must be defined to use this setting

Format: u32:0-65535
Queue target
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"queue_options": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Options used for queue target. Action queue must be defined to use this
                    setting

Format: bypass
Let packets go through if userspace application cannot back off
Format: fanout
Distribute packets between several queues
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						"destination": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								// TODO handle non-string types
								"mac_address": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `MAC address

Format: macaddr
MAC address to match
Format: !macaddr
Match everything except the specified MAC address
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},
							},
							// CustomType:          basetypes.MapTypable(nil),
							// Required:            false,
							Optional: true,
							// Computed:            false,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Destination parameters

`,
							// DeprecationMessage:  "",
							// Validators:          []validator.Map(nil),
							// PlanModifiers:       []planmodifier.Map(nil),
							// TODO investigate if node defaults can be handled
							// Default:             defaults.Map(nil),
						},

						"fragment": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								// TODO handle non-string types
								"match_frag": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Second and further fragments of fragmented packets

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"match_non_frag": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Head fragments or unfragmented packets

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},
							},
							// CustomType:          basetypes.MapTypable(nil),
							// Required:            false,
							Optional: true,
							// Computed:            false,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `IP fragment match

`,
							// DeprecationMessage:  "",
							// Validators:          []validator.Map(nil),
							// PlanModifiers:       []planmodifier.Map(nil),
							// TODO investigate if node defaults can be handled
							// Default:             defaults.Map(nil),
						},

						"inbound_interface": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								// TODO handle non-string types
								"interface_name": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Match interface

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"interface_group": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Match interface-group

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},
							},
							// CustomType:          basetypes.MapTypable(nil),
							// Required:            false,
							Optional: true,
							// Computed:            false,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Match inbound-interface

`,
							// DeprecationMessage:  "",
							// Validators:          []validator.Map(nil),
							// PlanModifiers:       []planmodifier.Map(nil),
							// TODO investigate if node defaults can be handled
							// Default:             defaults.Map(nil),
						},

						"outbound_interface": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								// TODO handle non-string types
								"interface_name": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Match interface

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"interface_group": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Match interface-group

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},
							},
							// CustomType:          basetypes.MapTypable(nil),
							// Required:            false,
							Optional: true,
							// Computed:            false,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Match outbound-interface

`,
							// DeprecationMessage:  "",
							// Validators:          []validator.Map(nil),
							// PlanModifiers:       []planmodifier.Map(nil),
							// TODO investigate if node defaults can be handled
							// Default:             defaults.Map(nil),
						},

						"ipsec": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								// TODO handle non-string types
								"match_ipsec": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Inbound IPsec packets

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"match_none": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Inbound non-IPsec packets

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},
							},
							// CustomType:          basetypes.MapTypable(nil),
							// Required:            false,
							Optional: true,
							// Computed:            false,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Inbound IPsec packets

`,
							// DeprecationMessage:  "",
							// Validators:          []validator.Map(nil),
							// PlanModifiers:       []planmodifier.Map(nil),
							// TODO investigate if node defaults can be handled
							// Default:             defaults.Map(nil),
						},

						"limit": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								// TODO handle non-string types
								"burst": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Maximum number of packets to allow in excess of rate

Format: u32:0-4294967295
Maximum number of packets to allow in excess of rate
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"rate": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Maximum average matching rate

Format: txt
integer/unit (Example: 5/minute)
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},
							},
							// CustomType:          basetypes.MapTypable(nil),
							// Required:            false,
							Optional: true,
							// Computed:            false,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Rate limit using a token bucket filter

`,
							// DeprecationMessage:  "",
							// Validators:          []validator.Map(nil),
							// PlanModifiers:       []planmodifier.Map(nil),
							// TODO investigate if node defaults can be handled
							// Default:             defaults.Map(nil),
						},

						"connection_status": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								// TODO handle non-string types
								"nat": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `NAT connection status

Format: destination
Match connections that are subject to destination NAT
Format: source
Match connections that are subject to source NAT
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},
							},
							// CustomType:          basetypes.MapTypable(nil),
							// Required:            false,
							Optional: true,
							// Computed:            false,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Connection status

`,
							// DeprecationMessage:  "",
							// Validators:          []validator.Map(nil),
							// PlanModifiers:       []planmodifier.Map(nil),
							// TODO investigate if node defaults can be handled
							// Default:             defaults.Map(nil),
						},

						"recent": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								// TODO handle non-string types
								"count": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Source addresses seen more than N times

Format: u32:1-255
Source addresses seen more than N times
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"time": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Source addresses seen in the last second/minute/hour

Format: second
Source addresses seen COUNT times in the last second
Format: minute
Source addresses seen COUNT times in the last minute
Format: hour
Source addresses seen COUNT times in the last hour
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},
							},
							// CustomType:          basetypes.MapTypable(nil),
							// Required:            false,
							Optional: true,
							// Computed:            false,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Parameters for matching recently seen sources

`,
							// DeprecationMessage:  "",
							// Validators:          []validator.Map(nil),
							// PlanModifiers:       []planmodifier.Map(nil),
							// TODO investigate if node defaults can be handled
							// Default:             defaults.Map(nil),
						},

						"source": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								// TODO handle non-string types
								"address": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `IP address, subnet, or range

Format: ipv4
IPv4 address to match
Format: ipv4net
IPv4 prefix to match
Format: ipv4range
IPv4 address range to match
Format: !ipv4
Match everything except the specified address
Format: !ipv4net
Match everything except the specified prefix
Format: !ipv4range
Match everything except the specified range
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"mac_address": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `MAC address

Format: macaddr
MAC address to match
Format: !macaddr
Match everything except the specified MAC address
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
									MarkdownDescription: `Port

Format: txt
Named port (any name in /etc/services, e.g., http)
Format: u32:1-65535
Numbered port
Format: <start-end>
Numbered port range (e.g. 1001-1005)
Format:
\n\n Multiple destination ports can be specified as a
                          comma-separated list.\n For example: 'telnet,http,123,1001-1005'
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								"group": schema.SingleNestedAttribute{
									Attributes: map[string]schema.Attribute{
										// TODO handle non-string types
										"address_group": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Group of addresses

`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"domain_group": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Group of domains

`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"mac_group": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Group of MAC addresses

`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"network_group": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Group of networks

`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"port_group": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Group of ports

`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},
									},
									// CustomType:          basetypes.MapTypable(nil),
									// Required:            false,
									Optional: true,
									// Computed:            false,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Group

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
							MarkdownDescription: `Source parameters

`,
							// DeprecationMessage:  "",
							// Validators:          []validator.Map(nil),
							// PlanModifiers:       []planmodifier.Map(nil),
							// TODO investigate if node defaults can be handled
							// Default:             defaults.Map(nil),
						},

						"state": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								// TODO handle non-string types
								"established": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Established state

Format: enable
Enable
Format: disable
Disable
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"invalid": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Invalid state

Format: enable
Enable
Format: disable
Disable
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"new": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `New state

Format: enable
Enable
Format: disable
Disable
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"related": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Related state

Format: enable
Enable
Format: disable
Disable
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},
							},
							// CustomType:          basetypes.MapTypable(nil),
							// Required:            false,
							Optional: true,
							// Computed:            false,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Session state

`,
							// DeprecationMessage:  "",
							// Validators:          []validator.Map(nil),
							// PlanModifiers:       []planmodifier.Map(nil),
							// TODO investigate if node defaults can be handled
							// Default:             defaults.Map(nil),
						},

						"tcp": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								// TODO handle non-string types
								"mss": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Maximum segment size (MSS)

Format: u32:1-16384
Maximum segment size
Format: <min>-<max>
TCP MSS range (use '-' as delimiter)
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								"flags": schema.SingleNestedAttribute{
									Attributes: map[string]schema.Attribute{
										// TODO handle non-string types
										"syn": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Synchronise flag

`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"ack": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Acknowledge flag

`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"fin": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Finish flag

`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"rst": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Reset flag

`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"urg": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Urgent flag

`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"psh": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Push flag

`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"ecn": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Explicit Congestion Notification flag

`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										// TODO handle non-string types
										"cwr": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Congestion Window Reduced flag

`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},

										"not": schema.SingleNestedAttribute{
											Attributes: map[string]schema.Attribute{
												// TODO handle non-string types
												"syn": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `Synchronise flag

`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

												},

												// TODO handle non-string types
												"ack": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `Acknowledge flag

`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

												},

												// TODO handle non-string types
												"fin": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `Finish flag

`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

												},

												// TODO handle non-string types
												"rst": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `Reset flag

`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

												},

												// TODO handle non-string types
												"urg": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `Urgent flag

`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

												},

												// TODO handle non-string types
												"psh": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `Push flag

`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

												},

												// TODO handle non-string types
												"ecn": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `Explicit Congestion Notification flag

`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

												},

												// TODO handle non-string types
												"cwr": schema.StringAttribute{
													// CustomType:          basetypes.StringTypable(nil),
													// Required:            false,
													Optional: true,
													// Sensitive:           false,
													// Description:         "",
													MarkdownDescription: `Congestion Window Reduced flag

`,
													// DeprecationMessage:  "",
													// TODO Recreate some of vyos validators for use in leafnodes
													// Validators:          []validator.String(nil),
													// PlanModifiers:       []planmodifier.String(nil),

												},
											},
											// CustomType:          basetypes.MapTypable(nil),
											// Required:            false,
											Optional: true,
											// Computed:            false,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Match flags not set

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
									MarkdownDescription: `TCP flags to match

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
							MarkdownDescription: `TCP flags to match

`,
							// DeprecationMessage:  "",
							// Validators:          []validator.Map(nil),
							// PlanModifiers:       []planmodifier.Map(nil),
							// TODO investigate if node defaults can be handled
							// Default:             defaults.Map(nil),
						},

						"time": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								// TODO handle non-string types
								"startdate": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Date to start matching rule

Format: txt
Enter date using following notation - YYYY-MM-DD
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"starttime": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Time of day to start matching rule

Format: txt
Enter time using using 24 hour notation - hh:mm:ss
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"stopdate": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Date to stop matching rule

Format: txt
Enter date using following notation - YYYY-MM-DD
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"stoptime": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Time of day to stop matching rule

Format: txt
Enter time using using 24 hour notation - hh:mm:ss
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"weekdays": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Comma separated weekdays to match rule on

Format: txt
Name of day (Monday, Tuesday, Wednesday, Thursdays, Friday,
                          Saturday, Sunday)
Format: u32:0-6
Day number (0 = Sunday ... 6 = Saturday)
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},
							},
							// CustomType:          basetypes.MapTypable(nil),
							// Required:            false,
							Optional: true,
							// Computed:            false,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Time to match rule

`,
							// DeprecationMessage:  "",
							// Validators:          []validator.Map(nil),
							// PlanModifiers:       []planmodifier.Map(nil),
							// TODO investigate if node defaults can be handled
							// Default:             defaults.Map(nil),
						},

						"hop_limit": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								// TODO handle non-string types
								"eq": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Match on equal value

Format: u32:0-255
Equal to value
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"gt": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Match on greater then value

Format: u32:0-255
Greater then value
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"lt": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Match on less then value

Format: u32:0-255
Less then value
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},
							},
							// CustomType:          basetypes.MapTypable(nil),
							// Required:            false,
							Optional: true,
							// Computed:            false,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Hop limit

`,
							// DeprecationMessage:  "",
							// Validators:          []validator.Map(nil),
							// PlanModifiers:       []planmodifier.Map(nil),
							// TODO investigate if node defaults can be handled
							// Default:             defaults.Map(nil),
						},

						"icmpv6": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								// TODO handle non-string types
								"code": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `ICMPv6 code

Format: u32:0-255
ICMPv6 code (0-255)
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"type": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `ICMPv6 type

Format: u32:0-255
ICMPv6 type (0-255)
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"type_name": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `ICMPv6 type-name

Format: destination-unreachable
ICMPv6 type 1: destination-unreachable
Format: packet-too-big
ICMPv6 type 2: packet-too-big
Format: time-exceeded
ICMPv6 type 3: time-exceeded
Format: echo-request
ICMPv6 type 128: echo-request
Format: echo-reply
ICMPv6 type 129: echo-reply
Format: mld-listener-query
ICMPv6 type 130: mld-listener-query
Format: mld-listener-report
ICMPv6 type 131: mld-listener-report
Format: mld-listener-reduction
ICMPv6 type 132: mld-listener-reduction
Format: nd-router-solicit
ICMPv6 type 133: nd-router-solicit
Format: nd-router-advert
ICMPv6 type 134: nd-router-advert
Format: nd-neighbor-solicit
ICMPv6 type 135: nd-neighbor-solicit
Format: nd-neighbor-advert
ICMPv6 type 136: nd-neighbor-advert
Format: nd-redirect
ICMPv6 type 137: nd-redirect
Format: parameter-problem
ICMPv6 type 4: parameter-problem
Format: router-renumbering
ICMPv6 type 138: router-renumbering
Format: ind-neighbor-solicit
ICMPv6 type 141: ind-neighbor-solicit
Format: ind-neighbor-advert
ICMPv6 type 142: ind-neighbor-advert
Format: mld2-listener-report
ICMPv6 type 143: mld2-listener-report
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},
							},
							// CustomType:          basetypes.MapTypable(nil),
							// Required:            false,
							Optional: true,
							// Computed:            false,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `ICMPv6 type and code information

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
				MarkdownDescription: `Firewall rule number (IPv6)

Format: u32:1-999999
Number for this Firewall rule
`,
				// DeprecationMessage:  "",
				// Validators:          []validator.Map(nil),
				// PlanModifiers:       []planmodifier.Map(nil),
				// TODO investigate if tagnode defaults can be handled
				// Default:             defaults.Map(nil),
			},

			// TODO handle non-string types
			"default_action": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Default-action for rule-set

Format: drop
Drop if no prior rules are hit
Format: jump
Jump to another chain if no prior rules are hit
Format: reject
Drop and notify source if no prior rules are hit
Format: return
Return from the current chain and continue at the next rule of the last
                  chain
Format: accept
Accept if no prior rules are hit
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

				Default:  stringdefault.StaticString(`drop`),
				Computed: true,
			},

			// TODO handle non-string types
			"enable_default_log": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Log packets hitting default-action

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
			"default_jump_target": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Set jump target. Action jump must be defined in default-action to use this
                setting

`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},
		},
	}
}

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func (r *firewall_ipv__name) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *firewall_ipv__nameModel

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
func (r *firewall_ipv__name) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *firewall_ipv__nameModel

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
func (r *firewall_ipv__name) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *firewall_ipv__nameModel

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
func (r *firewall_ipv__name) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *firewall_ipv__nameModel

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
