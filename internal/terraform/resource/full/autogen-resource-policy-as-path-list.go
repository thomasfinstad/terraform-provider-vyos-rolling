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
var _ resource.Resource = &policy_as_path_list{}

// var _ resource.ResourceWithImportState = &policy_as_path_list{}

// policy_as_path_list defines the resource implementation.
type policy_as_path_list struct {
	client   *http.Client
	vyosPath []string
}

// policy_as_path_listModel describes the resource data model.
type policy_as_path_listModel struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	Description types.String `tfsdk:"description"`

	// TagNodes
	Rule types.Map `tfsdk:"rule"`

	// Nodes

}

// Metadata method to define the resource type name.
func (r *policy_as_path_list) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policy_as_path_list"
}

// policy_as_path_listResource method to return the example resource reference
func policy_as_path_listResource() resource.Resource {
	return &policy_as_path_list{
		vyosPath: []string{
			"policy",
			"as-path-list",
		},
	}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r *policy_as_path_list) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Routing policy

`,

		Attributes: map[string]schema.Attribute{
			"identifier": schema.StringAttribute{
				Required: true,
				MarkdownDescription: `Add a BGP autonomous system path filter

|  Format  |  Description  |
|----------|---------------|
|  txt  |  AS path list name  |
`,
			},

			"rule": schema.MapNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"action": schema.StringAttribute{

							Optional: true,
							MarkdownDescription: `Action to take on entries matching this rule

|  Format  |  Description  |
|----------|---------------|
|  permit  |  Permit matching entries  |
|  deny  |  Deny matching entries  |
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

						"regex": schema.StringAttribute{

							Optional: true,
							MarkdownDescription: `Regular expression to match against an AS path

|  Format  |  Description  |
|----------|---------------|
|  txt  |  AS path regular expression (ex: "64501 64502")  |
`,
						},
					},
				},
				Optional: true,
				MarkdownDescription: `Rule for this as-path-list

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  AS path list rule number  |
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
		},
	}
}

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func (r *policy_as_path_list) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *policy_as_path_listModel

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
func (r *policy_as_path_list) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *policy_as_path_listModel

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
func (r *policy_as_path_list) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *policy_as_path_listModel

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
func (r *policy_as_path_list) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *policy_as_path_listModel

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
