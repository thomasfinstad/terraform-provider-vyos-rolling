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
var _ resource.Resource = &system_flow_accounting_sflow_server{}

// var _ resource.ResourceWithImportState = &system_flow_accounting_sflow_server{}

// system_flow_accounting_sflow_server defines the resource implementation.
type system_flow_accounting_sflow_server struct {
	client   *http.Client
	vyosPath []string
}

// system_flow_accounting_sflow_serverModel describes the resource data model.
type system_flow_accounting_sflow_serverModel struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	Port types.String `tfsdk:"port"`

	// TagNodes

	// Nodes

}

// Metadata method to define the resource type name.
func (r *system_flow_accounting_sflow_server) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_flow_accounting_sflow_server"
}

// system_flow_accounting_sflow_serverResource method to return the example resource reference
func system_flow_accounting_sflow_serverResource() resource.Resource {
	return &system_flow_accounting_sflow_server{
		vyosPath: []string{
			"system",
			"flow-accounting",
			"sflow",
			"server",
		},
	}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r *system_flow_accounting_sflow_server) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Flow accounting settings

sFlow settings

`,

		Attributes: map[string]schema.Attribute{
			"identifier": schema.StringAttribute{
				Required: true,
				MarkdownDescription: `sFlow destination server

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 server to export sFlow  |
|  ipv6  |  IPv6 server to export sFlow  |
`,
			},

			"port": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `sFlow port number

|  Format  |  Description  |
|----------|---------------|
|  u32:1025-65535  |  sFlow port number  |
`,

				Default:  stringdefault.StaticString(`6343`),
				Computed: true,
			},
		},
	}
}

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func (r *system_flow_accounting_sflow_server) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *system_flow_accounting_sflow_serverModel

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
func (r *system_flow_accounting_sflow_server) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *system_flow_accounting_sflow_serverModel

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
func (r *system_flow_accounting_sflow_server) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *system_flow_accounting_sflow_serverModel

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
func (r *system_flow_accounting_sflow_server) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *system_flow_accounting_sflow_serverModel

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
