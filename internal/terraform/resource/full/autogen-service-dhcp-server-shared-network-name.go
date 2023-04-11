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
var _ resource.Resource = &service_dhcp_server_shared_network_name{}

// var _ resource.ResourceWithImportState = &service_dhcp_server_shared_network_name{}

// service_dhcp_server_shared_network_name defines the resource implementation.
type service_dhcp_server_shared_network_name struct {
	client *http.Client
}

// service dhcp-server shared-network-nameModel describes the resource data model.
type service_dhcp_server_shared_network_nameModel struct {
	ConfigurableAttribute types.String `tfsdk:"configurable_attribute"`
	Defaulted             types.String `tfsdk:"defaulted"`
	ID                    types.String `tfsdk:"id"`

	Subnet types.Map `tfsdk:subnet`

	Authoritative types.String `tfsdk:authoritative`

	Domain_name types.String `tfsdk:domain_name`

	Domain_search types.String `tfsdk:domain_search`

	Ntp_server types.String `tfsdk:ntp_server`

	Ping_check types.String `tfsdk:ping_check`

	Description types.String `tfsdk:description`

	Disable types.String `tfsdk:disable`

	Name_server types.String `tfsdk:name_server`

	Shared_network_parameters types.String `tfsdk:shared_network_parameters`
}

// Metadata method to define the resource type name.
func (r *service_dhcp_server_shared_network_name) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service dhcp-server shared-network-name"
}

// service_dhcp_server_shared_network_nameResource method to return the example resource reference
func service_dhcp_server_shared_network_nameResource() resource.Resource {
	return &service_dhcp_server_shared_network_name{}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r *service_dhcp_server_shared_network_name) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example resource",

		Attributes: map[string]schema.Attribute{
			"subnet": schema.MapNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"range": schema.MapNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"start": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `First IP address for DHCP lease range

Format: ipv4
IPv4 start address of pool
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"stop": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Last IP address for DHCP lease range

Format: ipv4
IPv4 end address of pool
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
							MarkdownDescription: `DHCP lease range

`,
							// DeprecationMessage:  "",
							// Validators:          []validator.Map(nil),
							// PlanModifiers:       []planmodifier.Map(nil),
							// TODO investigate if tagnode defaults can be handled
							// Default:             defaults.Map(nil),
						},

						"static_mapping": schema.MapNestedAttribute{
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
									"ip_address": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Fixed IP address of static mapping

Format: ipv4
IPv4 address used in static mapping
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
									"static_mapping_parameters": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Additional static-mapping parameters for DHCP server. Will be placed inside the "host" block of the mapping. You must use the syntax of dhcpd.conf in this text-field. Using this without proper knowledge may result in a crashed DHCP server. Check system log to look for errors.

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
							MarkdownDescription: `Name of static mapping

`,
							// DeprecationMessage:  "",
							// Validators:          []validator.Map(nil),
							// PlanModifiers:       []planmodifier.Map(nil),
							// TODO investigate if tagnode defaults can be handled
							// Default:             defaults.Map(nil),
						},

						"static_route": schema.MapNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"next_hop": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `IP address of router to be used to reach the destination subnet

Format: ipv4
IPv4 address of router
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
							MarkdownDescription: `Classless static route destination subnet

Format: ipv4net
IPv4 address and prefix length
`,
							// DeprecationMessage:  "",
							// Validators:          []validator.Map(nil),
							// PlanModifiers:       []planmodifier.Map(nil),
							// TODO investigate if tagnode defaults can be handled
							// Default:             defaults.Map(nil),
						},

						// TODO handle non-string types
						"bootfile_name": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Bootstrap file name

`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"bootfile_server": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Server from which the initial boot file is to be loaded

Format: ipv4
Bootfile server IPv4 address
Format: hostname
Bootfile server FQDN
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"bootfile_size": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Bootstrap file size

Format: u32:1-16
Bootstrap file size in 512 byte blocks
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"client_prefix_length": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Specifies the clients subnet mask as per RFC 950. If unset, subnet declaration is used.

Format: u32:0-32
DHCP client prefix length must be 0 to 32
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"default_router": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `IP address of default router

Format: ipv4
Default router IPv4 address
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"domain_name": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Client Domain Name

`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"domain_search": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Client Domain Name search list

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
						"name_server": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Domain Name Servers (DNS) addresses

Format: ipv4
Domain Name Server (DNS) IPv4 address
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"enable_failover": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Enable DHCP failover support for this subnet

`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"exclude": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `IP address to exclude from DHCP lease range

Format: ipv4
IPv4 address to exclude from lease range
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"ip_forwarding": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Enable IP forwarding on client

`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"lease": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Lease timeout in seconds

Format: u32
DHCP lease time in seconds
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

							Default:  stringdefault.StaticString(`86400`),
							Computed: true,
						},

						// TODO handle non-string types
						"ntp_server": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `IP address of NTP server

Format: ipv4
NTP server IPv4 address
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"ping_check": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Sends ICMP Echo request to the address being assigned

`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"pop_server": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `IP address of POP3 server

Format: ipv4
POP3 server IPv4 address
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"server_identifier": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Address for DHCP server identifier

Format: ipv4
DHCP server identifier IPv4 address
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"smtp_server": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `IP address of SMTP server

Format: ipv4
SMTP server IPv4 address
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"ipv6_only_preferred": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Disable IPv4 on IPv6 only hosts (RFC 8925)

Format: u32
Seconds
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"subnet_parameters": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Additional subnet parameters for DHCP server. You must use the syntax of dhcpd.conf in this text-field. Using this without proper knowledge may result in a crashed DHCP server. Check system log to look for errors.

`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"tftp_server_name": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `TFTP server name

Format: ipv4
TFTP server IPv4 address
Format: hostname
TFTP server FQDN
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"time_offset": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Client subnet offset in seconds from Coordinated Universal Time (UTC)

Format: [-]N
Time offset (number, may be negative)
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"time_server": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `IP address of time server

Format: ipv4
Time server IPv4 address
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"wins_server": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `IP address for Windows Internet Name Service (WINS) server

Format: ipv4
WINS server IPv4 address
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"wpad_url": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Web Proxy Autodiscovery (WPAD) URL

`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						"vendor_option": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"ubiquiti": schema.SingleNestedAttribute{
									Attributes: map[string]schema.Attribute{
										// TODO handle non-string types
										"unifi_controller": schema.StringAttribute{
											// CustomType:          basetypes.StringTypable(nil),
											// Required:            false,
											Optional: true,
											// Sensitive:           false,
											// Description:         "",
											MarkdownDescription: `Address of UniFi controller

Format: ipv4
IP address of UniFi controller
`,
											// DeprecationMessage:  "",
											// TODO Recreate some of vyos validators for use in leafnodes
											// Validators:          []validator.String(nil),
											// PlanModifiers:       []planmodifier.String(nil),

										},
									},
									// CustomType:          basetypes.MapTypable(nil),
									// Required:            false,
									Optional: true,
									// Computed:            false,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Ubiquiti specific parameters

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
							MarkdownDescription: `Vendor Specific Options

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
				MarkdownDescription: `DHCP subnet for shared network

Format: ipv4net
IPv4 address and prefix length
`,
				// DeprecationMessage:  "",
				// Validators:          []validator.Map(nil),
				// PlanModifiers:       []planmodifier.Map(nil),
				// TODO investigate if tagnode defaults can be handled
				// Default:             defaults.Map(nil),
			},

			// TODO handle non-string types
			"authoritative": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Option to make DHCP server authoritative for this physical network

`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"domain_name": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Client Domain Name

`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"domain_search": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Client Domain Name search list

`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"ntp_server": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `IP address of NTP server

Format: ipv4
NTP server IPv4 address
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"ping_check": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Sends ICMP Echo request to the address being assigned

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
				MarkdownDescription: `Disable instance

`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"name_server": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Domain Name Servers (DNS) addresses

Format: ipv4
Domain Name Server (DNS) IPv4 address
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"shared_network_parameters": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Additional shared-network parameters for DHCP server. You must use the syntax of dhcpd.conf in this text-field. Using this without proper knowledge may result in a crashed DHCP server. Check system log to look for errors.

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
func (r *service_dhcp_server_shared_network_name) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *service_dhcp_server_shared_network_nameModel

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
func (r *service_dhcp_server_shared_network_name) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *service_dhcp_server_shared_network_nameModel

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
func (r *service_dhcp_server_shared_network_name) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *service_dhcp_server_shared_network_nameModel

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
func (r *service_dhcp_server_shared_network_name) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *service_dhcp_server_shared_network_nameModel

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
