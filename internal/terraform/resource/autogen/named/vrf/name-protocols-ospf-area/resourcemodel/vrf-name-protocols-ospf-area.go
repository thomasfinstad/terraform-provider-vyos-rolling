// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &VrfNameProtocolsOspfArea{}

// VrfNameProtocolsOspfArea describes the resource data model.
type VrfNameProtocolsOspfArea struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafVrfNameProtocolsOspfAreaAuthentication types.String `tfsdk:"authentication" vyos:"authentication,omitempty"`
	LeafVrfNameProtocolsOspfAreaNetwork        types.List   `tfsdk:"network" vyos:"network,omitempty"`
	LeafVrfNameProtocolsOspfAreaShortcut       types.String `tfsdk:"shortcut" vyos:"shortcut,omitempty"`
	LeafVrfNameProtocolsOspfAreaExportList     types.Number `tfsdk:"export_list" vyos:"export-list,omitempty"`
	LeafVrfNameProtocolsOspfAreaImportList     types.Number `tfsdk:"import_list" vyos:"import-list,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	ExistsTagVrfNameProtocolsOspfAreaRange bool `tfsdk:"-" vyos:"range,child"`

	ExistsTagVrfNameProtocolsOspfAreaVirtualLink bool `tfsdk:"-" vyos:"virtual-link,child"`

	// Nodes
	NodeVrfNameProtocolsOspfAreaAreaType *VrfNameProtocolsOspfAreaAreaType `tfsdk:"area_type" vyos:"area-type,omitempty"`
}

// SetID configures the resource ID
func (o *VrfNameProtocolsOspfArea) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *VrfNameProtocolsOspfArea) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *VrfNameProtocolsOspfArea) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VrfNameProtocolsOspfArea) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"area",
		o.SelfIdentifier.Attributes()["area"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *VrfNameProtocolsOspfArea) GetVyosParentPath() []string {
	return []string{
		"vrf",

		"name",

		o.SelfIdentifier.Attributes()["name"].(types.String).ValueString(),

		"protocols",

		"ospf",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *VrfNameProtocolsOspfArea) GetVyosNamedParentPath() []string {
	return []string{
		"vrf",

		"name",

		o.SelfIdentifier.Attributes()["name"].(types.String).ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfArea) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"area": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `OSPF area settings

    |  Format  |  Description                                  |
    |----------|-----------------------------------------------|
    |  u32     |  OSPF area number in decimal notation         |
    |  ipv4    |  OSPF area number in dotted decimal notation  |
`,
					Description: `OSPF area settings

    |  Format  |  Description                                  |
    |----------|-----------------------------------------------|
    |  u32     |  OSPF area number in decimal notation         |
    |  ipv4    |  OSPF area number in dotted decimal notation  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in area, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
								"illegal character in  area, value must match: ^[a-zA-Z0-9-_]*$",
							),
						),
					},
				},

				"name": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Virtual Routing and Forwarding instance

    |  Format  |  Description        |
    |----------|---------------------|
    |  txt     |  VRF instance name  |
`,
					Description: `Virtual Routing and Forwarding instance

    |  Format  |  Description        |
    |----------|---------------------|
    |  txt     |  VRF instance name  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in name, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
								"illegal character in  name, value must match: ^[a-zA-Z0-9-_]*$",
							),
						),
					},
				},
			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"authentication": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `OSPF area authentication type

    |  Format              |  Description                    |
    |----------------------|---------------------------------|
    |  plaintext-password  |  Use plain-text authentication  |
    |  md5                 |  Use MD5 authentication         |
`,
			Description: `OSPF area authentication type

    |  Format              |  Description                    |
    |----------------------|---------------------------------|
    |  plaintext-password  |  Use plain-text authentication  |
    |  md5                 |  Use MD5 authentication         |
`,
		},

		"network": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `OSPF network

    |  Format   |  Description   |
    |-----------|----------------|
    |  ipv4net  |  OSPF network  |
`,
			Description: `OSPF network

    |  Format   |  Description   |
    |-----------|----------------|
    |  ipv4net  |  OSPF network  |
`,
		},

		"shortcut": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Area shortcut mode

    |  Format   |  Description                |
    |-----------|-----------------------------|
    |  default  |  Set default                |
    |  disable  |  Disable shortcutting mode  |
    |  enable   |  Enable shortcutting mode   |
`,
			Description: `Area shortcut mode

    |  Format   |  Description                |
    |-----------|-----------------------------|
    |  default  |  Set default                |
    |  disable  |  Disable shortcutting mode  |
    |  enable   |  Enable shortcutting mode   |
`,
		},

		"export_list": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set the filter for networks announced to other areas

    |  Format  |  Description         |
    |----------|----------------------|
    |  u32     |  Access-list number  |
`,
			Description: `Set the filter for networks announced to other areas

    |  Format  |  Description         |
    |----------|----------------------|
    |  u32     |  Access-list number  |
`,
		},

		"import_list": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set the filter for networks from other areas announced

    |  Format  |  Description         |
    |----------|----------------------|
    |  u32     |  Access-list number  |
`,
			Description: `Set the filter for networks from other areas announced

    |  Format  |  Description         |
    |----------|----------------------|
    |  u32     |  Access-list number  |
`,
		},

		// Nodes

		"area_type": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfAreaAreaType{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Area type

`,
			Description: `Area type

`,
		},
	}
}
