// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
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

var _ helpers.VyosTopResourceDataModel = &VrfNameProtocolsOspfvthreeArea{}

// VrfNameProtocolsOspfvthreeArea describes the resource data model.
type VrfNameProtocolsOspfvthreeArea struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafVrfNameProtocolsOspfvthreeAreaExportList types.String `tfsdk:"export_list" vyos:"export-list,omitempty"`
	LeafVrfNameProtocolsOspfvthreeAreaImportList types.String `tfsdk:"import_list" vyos:"import-list,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	ExistsTagVrfNameProtocolsOspfvthreeAreaRange bool `tfsdk:"-" vyos:"range,child"`

	// Nodes
	NodeVrfNameProtocolsOspfvthreeAreaAreaType *VrfNameProtocolsOspfvthreeAreaAreaType `tfsdk:"area_type" vyos:"area-type,omitempty"`
}

// SetID configures the resource ID
func (o *VrfNameProtocolsOspfvthreeArea) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *VrfNameProtocolsOspfvthreeArea) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *VrfNameProtocolsOspfvthreeArea) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VrfNameProtocolsOspfvthreeArea) GetVyosPath() []string {
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
func (o *VrfNameProtocolsOspfvthreeArea) GetVyosParentPath() []string {
	return []string{
		"vrf",

		"name",

		o.SelfIdentifier.Attributes()["name"].(types.String).ValueString(),

		"protocols",

		"ospfv3",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *VrfNameProtocolsOspfvthreeArea) GetVyosNamedParentPath() []string {
	return []string{
		"vrf",

		"name",

		o.SelfIdentifier.Attributes()["name"].(types.String).ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfvthreeArea) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
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
					MarkdownDescription: `OSPFv3 Area

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  u32     |  Area ID as a decimal value   |
    |  ipv4    |  Area ID in IP address forma  |
`,
					Description: `OSPFv3 Area

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  u32     |  Area ID as a decimal value   |
    |  ipv4    |  Area ID in IP address forma  |
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

		"export_list": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Name of export-list

`,
			Description: `Name of export-list

`,
		},

		"import_list": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Name of import-list

`,
			Description: `Name of import-list

`,
		},

		// Nodes

		"area_type": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeAreaAreaType{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `OSPFv3 Area type

`,
			Description: `OSPFv3 Area type

`,
		},
	}
}
