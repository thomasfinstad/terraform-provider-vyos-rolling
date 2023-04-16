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
var _ resource.Resource = &vpn_sstp_client_ipvsix_pool_prefix{}

// var _ resource.ResourceWithImportState = &vpn_sstp_client_ipvsix_pool_prefix{}

// vpn_sstp_client_ipvsix_pool_prefix defines the resource implementation.
type vpn_sstp_client_ipvsix_pool_prefix struct {
	client   *http.Client
	vyosPath []string
}

// vpn_sstp_client_ipvsix_pool_prefixModel describes the resource data model.
type vpn_sstp_client_ipvsix_pool_prefixModel struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	Mask types.String `tfsdk:"mask"`

	// TagNodes

	// Nodes

}

// Metadata method to define the resource type name.
func (r *vpn_sstp_client_ipvsix_pool_prefix) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_sstp_client_ipv6_pool_prefix"
}

// vpn_sstp_client_ipvsix_pool_prefixResource method to return the example resource reference
func vpn_sstp_client_ipvsix_pool_prefixResource() resource.Resource {
	return &vpn_sstp_client_ipvsix_pool_prefix{
		vyosPath: []string{
			"vpn",
			"sstp",
			"client-ipv6-pool",
			"prefix",
		},
	}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r *vpn_sstp_client_ipvsix_pool_prefix) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Secure Socket Tunneling Protocol (SSTP) server

Pool of client IPv6 addresses

`,

		Attributes: map[string]schema.Attribute{
			"identifier": schema.StringAttribute{
				Required: true,
				MarkdownDescription: `Pool of addresses used to assign to clients

|  Format  |  Description  |
|----------|---------------|
|  ipv6net  |  IPv6 address and prefix length  |
`,
			},

			"mask": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Prefix length used for individual client

|  Format  |  Description  |
|----------|---------------|
|  u32:48-128  |  Client prefix length  |
`,

				Default:  stringdefault.StaticString(`64`),
				Computed: true,
			},
		},
	}
}

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func (r *vpn_sstp_client_ipvsix_pool_prefix) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *vpn_sstp_client_ipvsix_pool_prefixModel

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
func (r *vpn_sstp_client_ipvsix_pool_prefix) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *vpn_sstp_client_ipvsix_pool_prefixModel

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
func (r *vpn_sstp_client_ipvsix_pool_prefix) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *vpn_sstp_client_ipvsix_pool_prefixModel

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
func (r *vpn_sstp_client_ipvsix_pool_prefix) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *vpn_sstp_client_ipvsix_pool_prefixModel

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
