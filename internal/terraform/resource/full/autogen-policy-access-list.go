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
var _ resource.Resource = &policy_access_list{}

// var _ resource.ResourceWithImportState = &policy_access_list{}

// policy_access_list defines the resource implementation.
type policy_access_list struct {
	client *http.Client
}

// policy access-listModel describes the resource data model.
type policy_access_listModel struct {
	ConfigurableAttribute types.String `tfsdk:"configurable_attribute"`
	Defaulted             types.String `tfsdk:"defaulted"`
	ID                    types.String `tfsdk:"id"`

	Rule types.Map `tfsdk:rule`

	Description types.String `tfsdk:description`
}

// Metadata method to define the resource type name.
func (r *policy_access_list) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policy access-list"
}

// policy_access_listResource method to return the example resource reference
func policy_access_listResource() resource.Resource {
	return &policy_access_list{}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r *policy_access_list) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example resource",

		Attributes: map[string]schema.Attribute{
			"rule": schema.MapNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						// TODO handle non-string types
						"action": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Action to take on entries matching this rule

Format: permit
Permit matching entries
Format: deny
Deny matching entries
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						// TODO handle non-string types
						"description": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Description

Format: txt
Description
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

						},

						"destination": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								// TODO handle non-string types
								"any": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Any IP address to match

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"host": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Single host IP address to match

Format: ipv4
Host address to match
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"inverse_mask": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Network/netmask to match (requires network be defined)

Format: ipv4
Inverse-mask to match
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"network": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Network/netmask to match (requires inverse-mask be defined)

Format: ipv4net
Inverse-mask to match
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
							MarkdownDescription: `Destination network or address

`,
							// DeprecationMessage:  "",
							// Validators:          []validator.Map(nil),
							// PlanModifiers:       []planmodifier.Map(nil),
							// TODO investigate if node defaults can be handled
							// Default:             defaults.Map(nil),
						},

						"source": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								// TODO handle non-string types
								"any": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Any IP address to match

`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"host": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Single host IP address to match

Format: ipv4
Host address to match
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"inverse_mask": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Network/netmask to match (requires network be defined)

Format: ipv4
Inverse-mask to match
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"network": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Network/netmask to match (requires inverse-mask be defined)

Format: ipv4net
Inverse-mask to match
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
							MarkdownDescription: `Source network or address to match

`,
							// DeprecationMessage:  "",
							// Validators:          []validator.Map(nil),
							// PlanModifiers:       []planmodifier.Map(nil),
							// TODO investigate if node defaults can be handled
							// Default:             defaults.Map(nil),
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
				MarkdownDescription: `Rule for this access-list

Format: u32:1-65535
Access-list rule number
`,
				// DeprecationMessage:  "",
				// Validators:          []validator.Map(nil),
				// PlanModifiers:       []planmodifier.Map(nil),
				// TODO investigate if tagnode defaults can be handled
				// Default:             defaults.Map(nil),
			},

			// TODO handle non-string types
			"description": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Description

Format: txt
Description
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},
		},
	}
}

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func (r *policy_access_list) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *policy_access_listModel

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
func (r *policy_access_list) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *policy_access_listModel

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
func (r *policy_access_list) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *policy_access_listModel

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
func (r *policy_access_list) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *policy_access_listModel

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
