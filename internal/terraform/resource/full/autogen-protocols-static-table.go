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
var _ resource.Resource = &protocols_static_table{}

// var _ resource.ResourceWithImportState = &protocols_static_table{}

// protocols_static_table defines the resource implementation.
type protocols_static_table struct {
	client *http.Client
}

// protocols static tableModel describes the resource data model.
type protocols_static_tableModel struct {
	ConfigurableAttribute types.String `tfsdk:"configurable_attribute"`
	Defaulted             types.String `tfsdk:"defaulted"`
	ID                    types.String `tfsdk:"id"`

	Route types.Map `tfsdk:route`

	Route_ types.Map `tfsdk:route6`

	Description types.String `tfsdk:description`
}

// Metadata method to define the resource type name.
func (r *protocols_static_table) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protocols static table"
}

// protocols_static_tableResource method to return the example resource reference
func protocols_static_tableResource() resource.Resource {
	return &protocols_static_table{}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r *protocols_static_table) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example resource",

		Attributes: map[string]schema.Attribute{
			"route": schema.MapNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface": schema.MapNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"disable": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Disable instance

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"distance": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Distance for this route

Format: u32:1-255
Distance for this route
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"vrf": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `VRF to leak route

Format: txt
Name of VRF to leak to
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
							MarkdownDescription: `Next-hop IPv4 router interface

Format: txt
Gateway interface name
`,
							// DeprecationMessage:  "",
							// Validators:          []validator.Map(nil),
							// PlanModifiers:       []planmodifier.Map(nil),
							// TODO investigate if tagnode defaults can be handled
							// Default:             defaults.Map(nil),
						},

						"next_hop": schema.MapNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"disable": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Disable instance

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"distance": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Distance for this route

Format: u32:1-255
Distance for this route
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"interface": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Gateway interface name

Format: txt
Gateway interface name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"vrf": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `VRF to leak route

Format: txt
Name of VRF to leak to
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
							MarkdownDescription: `Next-hop IPv4 router address

Format: ipv4
Next-hop router address
`,
							// DeprecationMessage:  "",
							// Validators:          []validator.Map(nil),
							// PlanModifiers:       []planmodifier.Map(nil),
							// TODO investigate if tagnode defaults can be handled
							// Default:             defaults.Map(nil),
						},

						// TODO handle non-string types
						"dhcp_interface": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `DHCP interface supplying next-hop IP address

Format: txt
DHCP interface name
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

						"blackhole": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								// TODO handle non-string types
								"distance": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Distance for this route

Format: u32:1-255
Distance for this route
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"tag": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Tag value for this route

Format: u32:1-4294967295
Tag value for this route
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
							MarkdownDescription: `Silently discard pkts when matched

`,
							// DeprecationMessage:  "",
							// Validators:          []validator.Map(nil),
							// PlanModifiers:       []planmodifier.Map(nil),
							// TODO investigate if node defaults can be handled
							// Default:             defaults.Map(nil),
						},

						"reject": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								// TODO handle non-string types
								"distance": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Distance for this route

Format: u32:1-255
Distance for this route
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"tag": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Tag value for this route

Format: u32:1-4294967295
Tag value for this route
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
							MarkdownDescription: `Emit an ICMP unreachable when matched

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
				MarkdownDescription: `Static IPv4 route

Format: ipv4net
IPv4 static route
`,
				// DeprecationMessage:  "",
				// Validators:          []validator.Map(nil),
				// PlanModifiers:       []planmodifier.Map(nil),
				// TODO investigate if tagnode defaults can be handled
				// Default:             defaults.Map(nil),
			},

			"route6": schema.MapNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface": schema.MapNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"disable": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Disable instance

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"distance": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Distance for this route

Format: u32:1-255
Distance for this route
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"vrf": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `VRF to leak route

Format: txt
Name of VRF to leak to
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
							MarkdownDescription: `IPv6 gateway interface name

Format: txt
Gateway interface name
`,
							// DeprecationMessage:  "",
							// Validators:          []validator.Map(nil),
							// PlanModifiers:       []planmodifier.Map(nil),
							// TODO investigate if tagnode defaults can be handled
							// Default:             defaults.Map(nil),
						},

						"next_hop": schema.MapNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									// TODO handle non-string types
									"disable": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Disable instance

`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"distance": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Distance for this route

Format: u32:1-255
Distance for this route
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"interface": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `Gateway interface name

Format: txt
Gateway interface name
`,
										// DeprecationMessage:  "",
										// TODO Recreate some of vyos validators for use in leafnodes
										// Validators:          []validator.String(nil),
										// PlanModifiers:       []planmodifier.String(nil),

									},

									// TODO handle non-string types
									"vrf": schema.StringAttribute{
										// CustomType:          basetypes.StringTypable(nil),
										// Required:            false,
										Optional: true,
										// Sensitive:           false,
										// Description:         "",
										MarkdownDescription: `VRF to leak route

Format: txt
Name of VRF to leak to
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
							MarkdownDescription: `IPv6 gateway address

Format: ipv6
Next-hop IPv6 router
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

						"blackhole": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								// TODO handle non-string types
								"distance": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Distance for this route

Format: u32:1-255
Distance for this route
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"tag": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Tag value for this route

Format: u32:1-4294967295
Tag value for this route
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
							MarkdownDescription: `Silently discard pkts when matched

`,
							// DeprecationMessage:  "",
							// Validators:          []validator.Map(nil),
							// PlanModifiers:       []planmodifier.Map(nil),
							// TODO investigate if node defaults can be handled
							// Default:             defaults.Map(nil),
						},

						"reject": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								// TODO handle non-string types
								"distance": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Distance for this route

Format: u32:1-255
Distance for this route
`,
									// DeprecationMessage:  "",
									// TODO Recreate some of vyos validators for use in leafnodes
									// Validators:          []validator.String(nil),
									// PlanModifiers:       []planmodifier.String(nil),

								},

								// TODO handle non-string types
								"tag": schema.StringAttribute{
									// CustomType:          basetypes.StringTypable(nil),
									// Required:            false,
									Optional: true,
									// Sensitive:           false,
									// Description:         "",
									MarkdownDescription: `Tag value for this route

Format: u32:1-4294967295
Tag value for this route
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
							MarkdownDescription: `Emit an ICMP unreachable when matched

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
				MarkdownDescription: `Static IPv6 route

Format: ipv6net
IPv6 static route
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
func (r *protocols_static_table) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *protocols_static_tableModel

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
func (r *protocols_static_table) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *protocols_static_tableModel

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
func (r *protocols_static_table) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *protocols_static_tableModel

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
func (r *protocols_static_table) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *protocols_static_tableModel

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
