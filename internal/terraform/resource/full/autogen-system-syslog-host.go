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
var _ resource.Resource = &system_syslog_host{}

// var _ resource.ResourceWithImportState = &system_syslog_host{}

// system_syslog_host defines the resource implementation.
type system_syslog_host struct {
	client *http.Client
}

// system syslog hostModel describes the resource data model.
type system_syslog_hostModel struct {
	ConfigurableAttribute types.String `tfsdk:"configurable_attribute"`
	Defaulted             types.String `tfsdk:"defaulted"`
	ID                    types.String `tfsdk:"id"`

	Facility types.Map `tfsdk:facility`

	Format types.String `tfsdk:format`

	Port types.String `tfsdk:port`
}

// Metadata method to define the resource type name.
func (r *system_syslog_host) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system syslog host"
}

// system_syslog_hostResource method to return the example resource reference
func system_syslog_hostResource() resource.Resource {
	return &system_syslog_host{}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r *system_syslog_host) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example resource",

		Attributes: map[string]schema.Attribute{
			"facility": schema.MapNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						// TODO handle non-string types
						"protocol": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `syslog communication protocol

Format: udp
send log messages to remote syslog server over udp
Format: tcp
send log messages to remote syslog server over tcp
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"level": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Logging level

Format: emerg
Emergency messages
Format: alert
Urgent messages
Format: crit
Critical messages
Format: err
Error messages
Format: warning
Warning messages
Format: notice
Messages for further investigation
Format: info
Informational messages
Format: debug
Debug messages
Format: all
Log everything
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// CustomType:    basetypes.ObjectTypable(nil),
						// Validators:    []validator.Object(nil),
						// PlanModifiers: []planmodifier.Object(nil),
					},
				},
				// CustomType:          basetypes.MapTypable(nil),
				// Required:            false,
				Optional: true,
				// Computed:            false,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Facility for logging

Format: all
All facilities excluding "mark"
Format: auth
Authentication and authorization
Format: authpriv
Non-system authorization
Format: cron
Cron daemon
Format: daemon
System daemons
Format: kern
Kernel
Format: lpr
Line printer spooler
Format: mail
Mail subsystem
Format: mark
Timestamp
Format: news
USENET subsystem
Format: protocols
depricated will be set to local7
Format: security
depricated will be set to auth
Format: syslog
Authentication and authorization
Format: user
Application processes
Format: uucp
UUCP subsystem
Format: local0
Local facility 0
Format: local1
Local facility 1
Format: local2
Local facility 2
Format: local3
Local facility 3
Format: local4
Local facility 4
Format: local5
Local facility 5
Format: local6
Local facility 6
Format: local7
Local facility 7
`,
				// DeprecationMessage:  "",
				// Validators:          []validator.Map(nil),
				// PlanModifiers:       []planmodifier.Map(nil),
				// TODO investigate if tagnode defaults can be handled
				// Default:             defaults.Map(nil),
			},

			// TODO handle non-string types
			"port": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Port number used by connection

Format: u32:1-65535
Numeric IP port
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			"format": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					// TODO handle non-string types
					"octet_counted": schema.StringAttribute{
						// CustomType:          basetypes.StringTypable(nil),
						// Required:            false,
						Optional: true,
						// Sensitive:           false,
						// Description:         "",
						MarkdownDescription: `Allows for the transmission of all characters inside a syslog message

`,
						// DeprecationMessage:  "",
						// TODO Recreate some of vyos validators for use in leafnodes
						// Validators:          []validator.String(nil),
						// PlanModifiers:       []planmodifier.String(nil),

					},
				},
				// CustomType:          basetypes.MapTypable(nil),
				// Required:            false,
				Optional: true,
				// Computed:            false,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Logging format

`,
				// DeprecationMessage:  "",
				// Validators:          []validator.Map(nil),
				// PlanModifiers:       []planmodifier.Map(nil),
				// TODO investigate if node defaults can be handled
				// Default:             defaults.Map(nil),
			},
		},
	}
}

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func (r *system_syslog_host) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *system_syslog_hostModel

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
func (r *system_syslog_host) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *system_syslog_hostModel

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
func (r *system_syslog_host) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *system_syslog_hostModel

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
func (r *system_syslog_host) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *system_syslog_hostModel

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
