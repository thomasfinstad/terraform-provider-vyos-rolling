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
var _ resource.Resource = &firewall_interface{}

// var _ resource.ResourceWithImportState = &firewall_interface{}

// firewall_interface defines the resource implementation.
type firewall_interface struct {
	client   *http.Client
	vyosPath []string
}

// firewall_interfaceModel describes the resource data model.
type firewall_interfaceModel struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes

	// TagNodes

	// Nodes
	In    types.List `tfsdk:"in"`
	Out   types.List `tfsdk:"out"`
	Local types.List `tfsdk:"local"`
}

// Metadata method to define the resource type name.
func (r *firewall_interface) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_firewall_interface"
}

// firewall_interfaceResource method to return the example resource reference
func firewall_interfaceResource() resource.Resource {
	return &firewall_interface{
		vyosPath: []string{
			"firewall",
			"interface",
		},
	}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r *firewall_interface) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Firewall

`,

		Attributes: map[string]schema.Attribute{
			"identifier": schema.StringAttribute{
				Required: true,
				MarkdownDescription: `Interface name to apply firewall configuration

`,
			},

			"in": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{

					"name": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Local IPv4 firewall ruleset name for interface

`,
					},

					"ipv6_name": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Local IPv6 firewall ruleset name for interface

`,
					},
				},
				Optional: true,
				MarkdownDescription: `Forwarded packets on inbound interface

`,
			},

			"out": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{

					"name": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Local IPv4 firewall ruleset name for interface

`,
					},

					"ipv6_name": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Local IPv6 firewall ruleset name for interface

`,
					},
				},
				Optional: true,
				MarkdownDescription: `Forwarded packets on outbound interface

`,
			},

			"local": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{

					"name": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Local IPv4 firewall ruleset name for interface

`,
					},

					"ipv6_name": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Local IPv6 firewall ruleset name for interface

`,
					},
				},
				Optional: true,
				MarkdownDescription: `Packets destined for this router

`,
			},
		},
	}
}

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func (r *firewall_interface) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *firewall_interfaceModel

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
func (r *firewall_interface) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *firewall_interfaceModel

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
func (r *firewall_interface) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *firewall_interfaceModel

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
func (r *firewall_interface) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *firewall_interfaceModel

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
