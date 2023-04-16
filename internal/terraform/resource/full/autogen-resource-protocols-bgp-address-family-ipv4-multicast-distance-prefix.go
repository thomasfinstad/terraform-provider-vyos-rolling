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
var _ resource.Resource = &protocols_bgp_address_family_ipvfour_multicast_distance_prefix{}

// var _ resource.ResourceWithImportState = &protocols_bgp_address_family_ipvfour_multicast_distance_prefix{}

// protocols_bgp_address_family_ipvfour_multicast_distance_prefix defines the resource implementation.
type protocols_bgp_address_family_ipvfour_multicast_distance_prefix struct {
	client   *http.Client
	vyosPath []string
}

// protocols_bgp_address_family_ipvfour_multicast_distance_prefixModel describes the resource data model.
type protocols_bgp_address_family_ipvfour_multicast_distance_prefixModel struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	Distance types.String `tfsdk:"distance"`

	// TagNodes

	// Nodes

}

// Metadata method to define the resource type name.
func (r *protocols_bgp_address_family_ipvfour_multicast_distance_prefix) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protocols_bgp_address_family_ipv4_multicast_distance_prefix"
}

// protocols_bgp_address_family_ipvfour_multicast_distance_prefixResource method to return the example resource reference
func protocols_bgp_address_family_ipvfour_multicast_distance_prefixResource() resource.Resource {
	return &protocols_bgp_address_family_ipvfour_multicast_distance_prefix{
		vyosPath: []string{
			"protocols",
			"bgp",
			"address-family",
			"ipv4-multicast",
			"distance",
			"prefix",
		},
	}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r *protocols_bgp_address_family_ipvfour_multicast_distance_prefix) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Border Gateway Protocol (BGP)

BGP address-family parameters

Multicast IPv4 BGP settings

Administrative distances for BGP routes

`,

		Attributes: map[string]schema.Attribute{
			"identifier": schema.StringAttribute{
				Required: true,
				MarkdownDescription: `Administrative distance for a specific BGP prefix

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  Administrative distance for a specific BGP prefix  |
`,
			},

			"distance": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Administrative distance for prefix

|  Format  |  Description  |
|----------|---------------|
|  u32:1-255  |  Administrative distance for external BGP routes  |
`,
			},
		},
	}
}

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func (r *protocols_bgp_address_family_ipvfour_multicast_distance_prefix) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *protocols_bgp_address_family_ipvfour_multicast_distance_prefixModel

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
func (r *protocols_bgp_address_family_ipvfour_multicast_distance_prefix) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *protocols_bgp_address_family_ipvfour_multicast_distance_prefixModel

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
func (r *protocols_bgp_address_family_ipvfour_multicast_distance_prefix) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *protocols_bgp_address_family_ipvfour_multicast_distance_prefixModel

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
func (r *protocols_bgp_address_family_ipvfour_multicast_distance_prefix) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *protocols_bgp_address_family_ipvfour_multicast_distance_prefixModel

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
