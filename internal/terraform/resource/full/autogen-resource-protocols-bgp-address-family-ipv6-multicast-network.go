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
var _ resource.Resource = &protocols_bgp_address_family_ipvsix_multicast_network{}

// var _ resource.ResourceWithImportState = &protocols_bgp_address_family_ipvsix_multicast_network{}

// protocols_bgp_address_family_ipvsix_multicast_network defines the resource implementation.
type protocols_bgp_address_family_ipvsix_multicast_network struct {
	client   *http.Client
	vyosPath []string
}

// protocols_bgp_address_family_ipvsix_multicast_networkModel describes the resource data model.
type protocols_bgp_address_family_ipvsix_multicast_networkModel struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	Path_limit types.String `tfsdk:"path_limit"`
	Route_map  types.String `tfsdk:"route_map"`

	// TagNodes

	// Nodes

}

// Metadata method to define the resource type name.
func (r *protocols_bgp_address_family_ipvsix_multicast_network) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protocols_bgp_address_family_ipv6_multicast_network"
}

// protocols_bgp_address_family_ipvsix_multicast_networkResource method to return the example resource reference
func protocols_bgp_address_family_ipvsix_multicast_networkResource() resource.Resource {
	return &protocols_bgp_address_family_ipvsix_multicast_network{
		vyosPath: []string{
			"protocols",
			"bgp",
			"address-family",
			"ipv6-multicast",
			"network",
		},
	}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r *protocols_bgp_address_family_ipvsix_multicast_network) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Border Gateway Protocol (BGP)

BGP address-family parameters

Multicast IPv6 BGP settings

`,

		Attributes: map[string]schema.Attribute{
			"identifier": schema.StringAttribute{
				Required: true,
				MarkdownDescription: `Import BGP network/prefix into multicast IPv6 RIB

|  Format  |  Description  |
|----------|---------------|
|  ipv6net  |  Multicast IPv6 BGP network/prefix  |
`,
			},

			"path_limit": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `AS-path hopcount limit

|  Format  |  Description  |
|----------|---------------|
|  u32:0-255  |  AS path hop count limit  |
`,
			},

			"route_map": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Specify route-map name to use

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Route map name  |
`,
			},
		},
	}
}

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func (r *protocols_bgp_address_family_ipvsix_multicast_network) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *protocols_bgp_address_family_ipvsix_multicast_networkModel

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
func (r *protocols_bgp_address_family_ipvsix_multicast_network) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *protocols_bgp_address_family_ipvsix_multicast_networkModel

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
func (r *protocols_bgp_address_family_ipvsix_multicast_network) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *protocols_bgp_address_family_ipvsix_multicast_networkModel

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
func (r *protocols_bgp_address_family_ipvsix_multicast_network) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *protocols_bgp_address_family_ipvsix_multicast_networkModel

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
