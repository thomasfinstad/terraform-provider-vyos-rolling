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
var _ resource.Resource = &service_snmp_script_extensions_extension_name{}

// var _ resource.ResourceWithImportState = &service_snmp_script_extensions_extension_name{}

// service_snmp_script_extensions_extension_name defines the resource implementation.
type service_snmp_script_extensions_extension_name struct {
	client   *http.Client
	vyosPath []string
}

// service_snmp_script_extensions_extension_nameModel describes the resource data model.
type service_snmp_script_extensions_extension_nameModel struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	Script types.String `tfsdk:"script"`

	// TagNodes

	// Nodes

}

// Metadata method to define the resource type name.
func (r *service_snmp_script_extensions_extension_name) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_snmp_script_extensions_extension_name"
}

// service_snmp_script_extensions_extension_nameResource method to return the example resource reference
func service_snmp_script_extensions_extension_nameResource() resource.Resource {
	return &service_snmp_script_extensions_extension_name{
		vyosPath: []string{
			"service",
			"snmp",
			"script-extensions",
			"extension-name",
		},
	}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r *service_snmp_script_extensions_extension_name) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Simple Network Management Protocol (SNMP)

SNMP script extensions

`,

		Attributes: map[string]schema.Attribute{
			"identifier": schema.StringAttribute{
				Required: true,
				MarkdownDescription: `Extension name

`,
			},

			"script": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Script location and name

`,
			},
		},
	}
}

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func (r *service_snmp_script_extensions_extension_name) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *service_snmp_script_extensions_extension_nameModel

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
func (r *service_snmp_script_extensions_extension_name) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *service_snmp_script_extensions_extension_nameModel

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
func (r *service_snmp_script_extensions_extension_name) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *service_snmp_script_extensions_extension_nameModel

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
func (r *service_snmp_script_extensions_extension_name) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *service_snmp_script_extensions_extension_nameModel

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
