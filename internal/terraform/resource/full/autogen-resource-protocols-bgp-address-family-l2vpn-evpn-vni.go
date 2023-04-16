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
var _ resource.Resource = &protocols_bgp_address_family_ltwovpn_evpn_vni{}

// var _ resource.ResourceWithImportState = &protocols_bgp_address_family_ltwovpn_evpn_vni{}

// protocols_bgp_address_family_ltwovpn_evpn_vni defines the resource implementation.
type protocols_bgp_address_family_ltwovpn_evpn_vni struct {
	client   *http.Client
	vyosPath []string
}

// protocols_bgp_address_family_ltwovpn_evpn_vniModel describes the resource data model.
type protocols_bgp_address_family_ltwovpn_evpn_vniModel struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	Advertise_default_gw types.String `tfsdk:"advertise_default_gw"`
	Advertise_svi_ip     types.String `tfsdk:"advertise_svi_ip"`
	Rd                   types.String `tfsdk:"rd"`

	// TagNodes

	// Nodes
	Route_target types.List `tfsdk:"route_target"`
}

// Metadata method to define the resource type name.
func (r *protocols_bgp_address_family_ltwovpn_evpn_vni) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protocols_bgp_address_family_l2vpn_evpn_vni"
}

// protocols_bgp_address_family_ltwovpn_evpn_vniResource method to return the example resource reference
func protocols_bgp_address_family_ltwovpn_evpn_vniResource() resource.Resource {
	return &protocols_bgp_address_family_ltwovpn_evpn_vni{
		vyosPath: []string{
			"protocols",
			"bgp",
			"address-family",
			"l2vpn-evpn",
			"vni",
		},
	}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r *protocols_bgp_address_family_ltwovpn_evpn_vni) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Border Gateway Protocol (BGP)

BGP address-family parameters

L2VPN EVPN BGP settings

`,

		Attributes: map[string]schema.Attribute{
			"identifier": schema.StringAttribute{
				Required: true,
				MarkdownDescription: `VXLAN Network Identifier

|  Format  |  Description  |
|----------|---------------|
|  u32:1-16777215  |  VNI number  |
`,
			},

			"advertise_default_gw": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Advertise All default g/w mac-ip routes in EVPN

`,
			},

			"advertise_svi_ip": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Advertise svi mac-ip routes in EVPN

`,
			},

			"rd": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Route Distinguisher

|  Format  |  Description  |
|----------|---------------|
|  ASN:NN_OR_IP-ADDRESS:NN  |  Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)  |
`,
			},

			"route_target": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{

					"both": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Route Target both import and export

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Route target (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |
`,
					},

					"import": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Route Target import

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Route target (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |
`,
					},

					"export": schema.StringAttribute{

						Optional: true,
						MarkdownDescription: `Route Target export

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Route target (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |
`,
					},
				},
				Optional: true,
				MarkdownDescription: `Route Target

`,
			},
		},
	}
}

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func (r *protocols_bgp_address_family_ltwovpn_evpn_vni) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *protocols_bgp_address_family_ltwovpn_evpn_vniModel

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
func (r *protocols_bgp_address_family_ltwovpn_evpn_vni) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *protocols_bgp_address_family_ltwovpn_evpn_vniModel

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
func (r *protocols_bgp_address_family_ltwovpn_evpn_vni) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *protocols_bgp_address_family_ltwovpn_evpn_vniModel

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
func (r *protocols_bgp_address_family_ltwovpn_evpn_vni) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *protocols_bgp_address_family_ltwovpn_evpn_vniModel

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
