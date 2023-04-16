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
var _ resource.Resource = &service_snmp_vthree_trap_target{}

// var _ resource.ResourceWithImportState = &service_snmp_vthree_trap_target{}

// service_snmp_vthree_trap_target defines the resource implementation.
type service_snmp_vthree_trap_target struct {
	client   *http.Client
	vyosPath []string
}

// service_snmp_vthree_trap_targetModel describes the resource data model.
type service_snmp_vthree_trap_targetModel struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	Port     types.String `tfsdk:"port"`
	Protocol types.String `tfsdk:"protocol"`
	Type     types.String `tfsdk:"type"`
	User     types.String `tfsdk:"user"`

	// TagNodes

	// Nodes
	Auth    types.List `tfsdk:"auth"`
	Privacy types.List `tfsdk:"privacy"`
}

// Metadata method to define the resource type name.
func (r *service_snmp_vthree_trap_target) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_snmp_v3_trap_target"
}

// service_snmp_vthree_trap_targetResource method to return the example resource reference
func service_snmp_vthree_trap_targetResource() resource.Resource {
	return &service_snmp_vthree_trap_target{
		vyosPath: []string{
			"service",
			"snmp",
			"v3",
			"trap-target",
		},
	}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r *service_snmp_vthree_trap_target) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Simple Network Management Protocol (SNMP)

Simple Network Management Protocol (SNMP) v3

`,

		Attributes: map[string]schema.Attribute{
			"identifier": schema.StringAttribute{
				Required: true,
				MarkdownDescription: `Defines SNMP target for inform or traps for IP

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IP address of trap target  |
|  ipv6  |  IPv6 address of trap target  |
`,
			},

			"port": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Port number used by connection

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Numeric IP port  |
`,

				Default:  stringdefault.StaticString(`162`),
				Computed: true,
			},

			"protocol": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Protocol to be used (TCP/UDP)

|  Format  |  Description  |
|----------|---------------|
|  udp  |  Listen protocol UDP  |
|  tcp  |  Listen protocol TCP  |
`,

				Default:  stringdefault.StaticString(`udp`),
				Computed: true,
			},

			"type": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Specifies the type of notification between inform and trap

|  Format  |  Description  |
|----------|---------------|
|  inform  |  Use INFORM  |
|  trap  |  Use TRAP  |
`,

				Default:  stringdefault.StaticString(`inform`),
				Computed: true,
			},

			"user": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Defines username for authentication

`,
			},

			"auth": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{

					"encrypted_password": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Defines the encrypted key for authentication

`,
					},

					"plaintext_password": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Defines the clear text key for authentication

`,
					},

					"type": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Define used protocol

|  Format  |  Description  |
|----------|---------------|
|  md5  |  Message Digest 5  |
|  sha  |  Secure Hash Algorithm  |
`,

						Default:  stringdefault.StaticString(`md5`),
						Computed: true,
					},
				},
				Optional: true,
				MarkdownDescription: `Defines the privacy

`,
			},

			"privacy": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{

					"encrypted_password": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Defines the encrypted key for privacy protocol

`,
					},

					"plaintext_password": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Defines the clear text key for privacy protocol

`,
					},

					"type": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Defines the protocol for privacy

|  Format  |  Description  |
|----------|---------------|
|  des  |  Data Encryption Standard  |
|  aes  |  Advanced Encryption Standard  |
`,

						Default:  stringdefault.StaticString(`des`),
						Computed: true,
					},
				},
				Optional: true,
				MarkdownDescription: `Defines the privacy

`,
			},
		},
	}
}

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func (r *service_snmp_vthree_trap_target) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *service_snmp_vthree_trap_targetModel

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
func (r *service_snmp_vthree_trap_target) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *service_snmp_vthree_trap_targetModel

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
func (r *service_snmp_vthree_trap_target) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *service_snmp_vthree_trap_targetModel

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
func (r *service_snmp_vthree_trap_target) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *service_snmp_vthree_trap_targetModel

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
