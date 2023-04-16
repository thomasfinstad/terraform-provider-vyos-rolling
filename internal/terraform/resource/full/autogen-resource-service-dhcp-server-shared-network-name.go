// Code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.

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
	client   *http.Client
	vyosPath []string
}

// service_dhcp_server_shared_network_nameModel describes the resource data model.
type service_dhcp_server_shared_network_nameModel struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	Authoritative             types.String `tfsdk:"authoritative"`
	Domain_name               types.String `tfsdk:"domain_name"`
	Domain_search             types.String `tfsdk:"domain_search"`
	Ntp_server                types.String `tfsdk:"ntp_server"`
	Ping_check                types.String `tfsdk:"ping_check"`
	Description               types.String `tfsdk:"description"`
	Disable                   types.String `tfsdk:"disable"`
	Name_server               types.String `tfsdk:"name_server"`
	Shared_network_parameters types.String `tfsdk:"shared_network_parameters"`

	// TagNodes
	Subnet types.Map `tfsdk:"subnet"`

	// Nodes

}

// Metadata method to define the resource type name.
func (r *service_dhcp_server_shared_network_name) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_dhcp_server_shared_network_name"
}

// service_dhcp_server_shared_network_nameResource method to return the example resource reference
func service_dhcp_server_shared_network_nameResource() resource.Resource {
	return &service_dhcp_server_shared_network_name{
		vyosPath: []string{
			"service",
			"dhcp-server",
			"shared-network-name",
		},
	}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r *service_dhcp_server_shared_network_name) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Dynamic Host Configuration Protocol (DHCP) for DHCP server

`,

		Attributes: map[string]schema.Attribute{
			"identifier": schema.StringAttribute{
				Required: true,
				MarkdownDescription: `Name of DHCP shared network

`,
			},

			"subnet": schema.MapNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"range": schema.MapNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"start": schema.StringAttribute{

										Optional: true,
										MarkdownDescription: `First IP address for DHCP lease range

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 start address of pool  |
`,
									},

									"stop": schema.StringAttribute{

										Optional: true,
										MarkdownDescription: `Last IP address for DHCP lease range

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 end address of pool  |
`,
									},
								},
							},
							Optional: true,
							MarkdownDescription: `DHCP lease range

`,
						},

						"static_mapping": schema.MapNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"disable": schema.StringAttribute{

										Optional: true,
										MarkdownDescription: `Disable instance

`,
									},

									"ip_address": schema.StringAttribute{

										Optional: true,
										MarkdownDescription: `Fixed IP address of static mapping

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address used in static mapping  |
`,
									},

									"mac_address": schema.StringAttribute{

										Optional: true,
										MarkdownDescription: `Media Access Control (MAC) address

|  Format  |  Description  |
|----------|---------------|
|  macaddr  |  Hardware (MAC) address  |
`,
									},

									"static_mapping_parameters": schema.StringAttribute{

										Optional: true,
										MarkdownDescription: `Additional static-mapping parameters for DHCP server. Will be placed inside the "host" block of the mapping. You must use the syntax of dhcpd.conf in this text-field. Using this without proper knowledge may result in a crashed DHCP server. Check system log to look for errors.

`,
									},
								},
							},
							Optional: true,
							MarkdownDescription: `Name of static mapping

`,
						},

						"static_route": schema.MapNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"next_hop": schema.StringAttribute{

										Optional: true,
										MarkdownDescription: `IP address of router to be used to reach the destination subnet

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address of router  |
`,
									},
								},
							},
							Optional: true,
							MarkdownDescription: `Classless static route destination subnet

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  IPv4 address and prefix length  |
`,
						},

						"bootfile_name": schema.StringAttribute{

							Optional: true,
							MarkdownDescription: `Bootstrap file name

`,
						},

						"bootfile_server": schema.StringAttribute{

							Optional: true,
							MarkdownDescription: `Server from which the initial boot file is to be loaded

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Bootfile server IPv4 address  |
|  hostname  |  Bootfile server FQDN  |
`,
						},

						"bootfile_size": schema.StringAttribute{

							Optional: true,
							MarkdownDescription: `Bootstrap file size

|  Format  |  Description  |
|----------|---------------|
|  u32:1-16  |  Bootstrap file size in 512 byte blocks  |
`,
						},

						"client_prefix_length": schema.StringAttribute{

							Optional: true,
							MarkdownDescription: `Specifies the clients subnet mask as per RFC 950. If unset, subnet declaration is used.

|  Format  |  Description  |
|----------|---------------|
|  u32:0-32  |  DHCP client prefix length must be 0 to 32  |
`,
						},

						"default_router": schema.StringAttribute{

							Optional: true,
							MarkdownDescription: `IP address of default router

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Default router IPv4 address  |
`,
						},

						"domain_name": schema.StringAttribute{

							Optional: true,
							MarkdownDescription: `Client Domain Name

`,
						},

						"domain_search": schema.StringAttribute{

							Optional: true,
							MarkdownDescription: `Client Domain Name search list

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

						"name_server": schema.StringAttribute{

							Optional: true,
							MarkdownDescription: `Domain Name Servers (DNS) addresses

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Domain Name Server (DNS) IPv4 address  |
`,
						},

						"enable_failover": schema.StringAttribute{

							Optional: true,
							MarkdownDescription: `Enable DHCP failover support for this subnet

`,
						},

						"exclude": schema.StringAttribute{

							Optional: true,
							MarkdownDescription: `IP address to exclude from DHCP lease range

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address to exclude from lease range  |
`,
						},

						"ip_forwarding": schema.StringAttribute{

							Optional: true,
							MarkdownDescription: `Enable IP forwarding on client

`,
						},

						"lease": schema.StringAttribute{

							Optional: true,
							MarkdownDescription: `Lease timeout in seconds

|  Format  |  Description  |
|----------|---------------|
|  u32  |  DHCP lease time in seconds  |
`,

							Default:  stringdefault.StaticString(`86400`),
							Computed: true,
						},

						"ntp_server": schema.StringAttribute{

							Optional: true,
							MarkdownDescription: `IP address of NTP server

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  NTP server IPv4 address  |
`,
						},

						"ping_check": schema.StringAttribute{

							Optional: true,
							MarkdownDescription: `Sends ICMP Echo request to the address being assigned

`,
						},

						"pop_server": schema.StringAttribute{

							Optional: true,
							MarkdownDescription: `IP address of POP3 server

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  POP3 server IPv4 address  |
`,
						},

						"server_identifier": schema.StringAttribute{

							Optional: true,
							MarkdownDescription: `Address for DHCP server identifier

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  DHCP server identifier IPv4 address  |
`,
						},

						"smtp_server": schema.StringAttribute{

							Optional: true,
							MarkdownDescription: `IP address of SMTP server

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  SMTP server IPv4 address  |
`,
						},

						"ipv6_only_preferred": schema.StringAttribute{

							Optional: true,
							MarkdownDescription: `Disable IPv4 on IPv6 only hosts (RFC 8925)

|  Format  |  Description  |
|----------|---------------|
|  u32  |  Seconds  |
`,
						},

						"subnet_parameters": schema.StringAttribute{

							Optional: true,
							MarkdownDescription: `Additional subnet parameters for DHCP server. You must use the syntax of dhcpd.conf in this text-field. Using this without proper knowledge may result in a crashed DHCP server. Check system log to look for errors.

`,
						},

						"tftp_server_name": schema.StringAttribute{

							Optional: true,
							MarkdownDescription: `TFTP server name

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  TFTP server IPv4 address  |
|  hostname  |  TFTP server FQDN  |
`,
						},

						"time_offset": schema.StringAttribute{

							Optional: true,
							MarkdownDescription: `Client subnet offset in seconds from Coordinated Universal Time (UTC)

|  Format  |  Description  |
|----------|---------------|
|  [-]N  |  Time offset (number, may be negative)  |
`,
						},

						"time_server": schema.StringAttribute{

							Optional: true,
							MarkdownDescription: `IP address of time server

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Time server IPv4 address  |
`,
						},

						"wins_server": schema.StringAttribute{

							Optional: true,
							MarkdownDescription: `IP address for Windows Internet Name Service (WINS) server

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  WINS server IPv4 address  |
`,
						},

						"wpad_url": schema.StringAttribute{

							Optional: true,
							MarkdownDescription: `Web Proxy Autodiscovery (WPAD) URL

`,
						},

						"vendor_option": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{

								"ubiquiti": schema.SingleNestedAttribute{
									Attributes: map[string]schema.Attribute{

										"unifi_controller": schema.StringAttribute{

											Optional: true,
											MarkdownDescription: `Address of UniFi controller

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IP address of UniFi controller  |
`,
										},
									},
									Optional: true,
									MarkdownDescription: `Ubiquiti specific parameters

`,
								},
							},
							Optional: true,
							MarkdownDescription: `Vendor Specific Options

`,
						},
					},
				},
				Optional: true,
				MarkdownDescription: `DHCP subnet for shared network

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  IPv4 address and prefix length  |
`,
			},

			"authoritative": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Option to make DHCP server authoritative for this physical network

`,
			},

			"domain_name": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Client Domain Name

`,
			},

			"domain_search": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Client Domain Name search list

`,
			},

			"ntp_server": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `IP address of NTP server

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  NTP server IPv4 address  |
`,
			},

			"ping_check": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Sends ICMP Echo request to the address being assigned

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

			"disable": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Disable instance

`,
			},

			"name_server": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Domain Name Servers (DNS) addresses

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Domain Name Server (DNS) IPv4 address  |
`,
			},

			"shared_network_parameters": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Additional shared-network parameters for DHCP server. You must use the syntax of dhcpd.conf in this text-field. Using this without proper knowledge may result in a crashed DHCP server. Check system log to look for errors.

`,
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
