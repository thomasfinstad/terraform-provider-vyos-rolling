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
var _ resource.Resource = &system_console_device{}

// var _ resource.ResourceWithImportState = &system_console_device{}

// system_console_device defines the resource implementation.
type system_console_device struct {
	client   *http.Client
	vyosPath []string
}

// system_console_deviceModel describes the resource data model.
type system_console_deviceModel struct {
	ID types.String `tfsdk:"identifier"`

	// LeafNodes
	Speed types.String `tfsdk:"speed"`

	// TagNodes

	// Nodes

}

// Metadata method to define the resource type name.
func (r *system_console_device) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_console_device"
}

// system_console_deviceResource method to return the example resource reference
func system_console_deviceResource() resource.Resource {
	return &system_console_device{
		vyosPath: []string{
			"system",
			"console",
			"device",
		},
	}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r *system_console_device) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Serial console configuration

`,

		Attributes: map[string]schema.Attribute{
			"identifier": schema.StringAttribute{
				Required: true,
				MarkdownDescription: `Serial console device name

|  Format  |  Description  |
|----------|---------------|
|  ttySN  |  TTY device name, regular serial port  |
|  usbNbXpY  |  TTY device name, USB based  |
|  hvcN  |  Xen console  |
`,
			},

			"speed": schema.StringAttribute{

				Optional: true,
				MarkdownDescription: `Console baud rate

|  Format  |  Description  |
|----------|---------------|
|  1200  |  1200 bps  |
|  2400  |  2400 bps  |
|  4800  |  4800 bps  |
|  9600  |  9600 bps  |
|  19200  |  19200 bps  |
|  38400  |  38400 bps  |
|  57600  |  57600 bps  |
|  115200  |  115200 bps  |
`,

				Default:  stringdefault.StaticString(`115200`),
				Computed: true,
			},
		},
	}
}

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func (r *system_console_device) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *system_console_deviceModel

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
func (r *system_console_device) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *system_console_deviceModel

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
func (r *system_console_device) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *system_console_deviceModel

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
func (r *system_console_device) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *system_console_deviceModel

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
