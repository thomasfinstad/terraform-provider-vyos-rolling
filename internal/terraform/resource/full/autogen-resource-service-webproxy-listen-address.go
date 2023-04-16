// Code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.

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
var _ resource.Resource = &service_webproxy_listen_address{}

// var _ resource.ResourceWithImportState = &service_webproxy_listen_address{}

// service_webproxy_listen_address defines the resource implementation.
type service_webproxy_listen_address struct {
	client   *http.Client
	vyosPath []string
}

// service_webproxy_listen_addressModel describes the resource data model.
type service_webproxy_listen_addressModel struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	Port                types.String `tfsdk:"port"`
	Disable_transparent types.String `tfsdk:"disable_transparent"`

	// TagNodes

	// Nodes

}

// Metadata method to define the resource type name.
func (r *service_webproxy_listen_address) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_webproxy_listen_address"
}

// service_webproxy_listen_addressResource method to return the example resource reference
func service_webproxy_listen_addressResource() resource.Resource {
	return &service_webproxy_listen_address{
		vyosPath: []string{
			"service",
			"webproxy",
			"listen-address",
		},
	}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r *service_webproxy_listen_address) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Webproxy service settings

`,

		Attributes: map[string]schema.Attribute{
			"identifier": schema.StringAttribute{
				Required: true,
				MarkdownDescription: `IPv4 listen-address for WebProxy

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address listen on  |
`,
			},

			"port": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Default Proxy Port

|  Format  |  Description  |
|----------|---------------|
|  u32:1025-65535  |  Default port number  |
`,
			},

			"disable_transparent": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Disable transparent mode

`,
			},
		},
	}
}

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func (r *service_webproxy_listen_address) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *service_webproxy_listen_addressModel

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
func (r *service_webproxy_listen_address) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *service_webproxy_listen_addressModel

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
func (r *service_webproxy_listen_address) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *service_webproxy_listen_addressModel

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
func (r *service_webproxy_listen_address) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *service_webproxy_listen_addressModel

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
