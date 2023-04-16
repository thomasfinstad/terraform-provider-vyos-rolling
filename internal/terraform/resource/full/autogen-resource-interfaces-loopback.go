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
var _ resource.Resource = &interfaces_loopback{}

// var _ resource.ResourceWithImportState = &interfaces_loopback{}

// interfaces_loopback defines the resource implementation.
type interfaces_loopback struct {
	client   *http.Client
	vyosPath []string
}

// interfaces_loopbackModel describes the resource data model.
type interfaces_loopbackModel struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	Address     types.String `tfsdk:"address"`
	Description types.String `tfsdk:"description"`
	Redirect    types.String `tfsdk:"redirect"`

	// TagNodes

	// Nodes
	Ip     types.List `tfsdk:"ip"`
	Mirror types.List `tfsdk:"mirror"`
}

// Metadata method to define the resource type name.
func (r *interfaces_loopback) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interfaces_loopback"
}

// interfaces_loopbackResource method to return the example resource reference
func interfaces_loopbackResource() resource.Resource {
	return &interfaces_loopback{
		vyosPath: []string{
			"interfaces",
			"loopback",
		},
	}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r *interfaces_loopback) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: ``,

		Attributes: map[string]schema.Attribute{
			"identifier": schema.StringAttribute{
				Required: true,
				MarkdownDescription: `Loopback Interface

|  Format  |  Description  |
|----------|---------------|
|  lo  |  Loopback interface  |
`,
			},

			"address": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `IP address

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  IPv4 address and prefix length  |
|  ipv6net  |  IPv6 address and prefix length  |
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

			"redirect": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Redirect incoming packet to destination

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Destination interface name  |
`,
			},

			"ip": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{

					"source_validation": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Source validation by reversed path (RFC3704)

|  Format  |  Description  |
|----------|---------------|
|  strict  |  Enable Strict Reverse Path Forwarding as defined in RFC3704  |
|  loose  |  Enable Loose Reverse Path Forwarding as defined in RFC3704  |
|  disable  |  No source validation  |
`,
					},
				},
				Optional: true,
				MarkdownDescription: `IPv4 routing parameters

`,
			},

			"mirror": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{

					"ingress": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Mirror ingress traffic to destination interface

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Destination interface name  |
`,
					},

					"egress": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Mirror egress traffic to destination interface

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Destination interface name  |
`,
					},
				},
				Optional: true,
				MarkdownDescription: `Mirror ingress/egress packets

`,
			},
		},
	}
}

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func (r *interfaces_loopback) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *interfaces_loopbackModel

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
func (r *interfaces_loopback) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *interfaces_loopbackModel

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
func (r *interfaces_loopback) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *interfaces_loopbackModel

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
func (r *interfaces_loopback) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *interfaces_loopbackModel

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
