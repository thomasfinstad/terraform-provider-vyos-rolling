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
var _ resource.Resource = &system_syslog_global_facility{}

// var _ resource.ResourceWithImportState = &system_syslog_global_facility{}

// system_syslog_global_facility defines the resource implementation.
type system_syslog_global_facility struct {
	client   *http.Client
	vyosPath []string
}

// system_syslog_global_facilityModel describes the resource data model.
type system_syslog_global_facilityModel struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	Level types.String `tfsdk:"level"`

	// TagNodes

	// Nodes

}

// Metadata method to define the resource type name.
func (r *system_syslog_global_facility) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_syslog_global_facility"
}

// system_syslog_global_facilityResource method to return the example resource reference
func system_syslog_global_facilityResource() resource.Resource {
	return &system_syslog_global_facility{
		vyosPath: []string{
			"system",
			"syslog",
			"global",
			"facility",
		},
	}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r *system_syslog_global_facility) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `System logging

Logging to system standard location

`,

		Attributes: map[string]schema.Attribute{
			"identifier": schema.StringAttribute{
				Required: true,
				MarkdownDescription: `Facility for logging

|  Format  |  Description  |
|----------|---------------|
|  all  |  All facilities excluding "mark"  |
|  auth  |  Authentication and authorization  |
|  authpriv  |  Non-system authorization  |
|  cron  |  Cron daemon  |
|  daemon  |  System daemons  |
|  kern  |  Kernel  |
|  lpr  |  Line printer spooler  |
|  mail  |  Mail subsystem  |
|  mark  |  Timestamp  |
|  news  |  USENET subsystem  |
|  protocols  |  depricated will be set to local7  |
|  security  |  depricated will be set to auth  |
|  syslog  |  Authentication and authorization  |
|  user  |  Application processes  |
|  uucp  |  UUCP subsystem  |
|  local0  |  Local facility 0  |
|  local1  |  Local facility 1  |
|  local2  |  Local facility 2  |
|  local3  |  Local facility 3  |
|  local4  |  Local facility 4  |
|  local5  |  Local facility 5  |
|  local6  |  Local facility 6  |
|  local7  |  Local facility 7  |
`,
			},

			"level": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Logging level

|  Format  |  Description  |
|----------|---------------|
|  emerg  |  Emergency messages  |
|  alert  |  Urgent messages  |
|  crit  |  Critical messages  |
|  err  |  Error messages  |
|  warning  |  Warning messages  |
|  notice  |  Messages for further investigation  |
|  info  |  Informational messages  |
|  debug  |  Debug messages  |
|  all  |  Log everything  |
`,
			},
		},
	}
}

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func (r *system_syslog_global_facility) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *system_syslog_global_facilityModel

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
func (r *system_syslog_global_facility) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *system_syslog_global_facilityModel

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
func (r *system_syslog_global_facility) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *system_syslog_global_facilityModel

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
func (r *system_syslog_global_facility) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *system_syslog_global_facilityModel

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
