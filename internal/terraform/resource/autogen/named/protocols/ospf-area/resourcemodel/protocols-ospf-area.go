/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (area) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ProtocolsOspfArea{}

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-tag-node-identifier (area) */
// ProtocolsOspfAreaSelfIdentifier used by TagNodes to keep the id info
type ProtocolsOspfAreaSelfIdentifier struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (area) */

	ProtocolsOspfArea types.String `tfsdk:"area" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (ospf) */

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (protocols) */
}

// ProtocolsOspfArea describes the resource data model.
// This is a basenode!
// Top level basenode type: `TagNode`
type ProtocolsOspfArea struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	SelfIdentifier *ProtocolsOspfAreaSelfIdentifier `tfsdk:"identifier" vyos:"-,self-id"`

	// LeafNodes
	LeafProtocolsOspfAreaAuthentication types.String `tfsdk:"authentication" vyos:"authentication,omitempty"`
	LeafProtocolsOspfAreaNetwork        types.List   `tfsdk:"network" vyos:"network,omitempty"`
	LeafProtocolsOspfAreaShortcut       types.String `tfsdk:"shortcut" vyos:"shortcut,omitempty"`
	LeafProtocolsOspfAreaExportList     types.Number `tfsdk:"export_list" vyos:"export-list,omitempty"`
	LeafProtocolsOspfAreaImportList     types.Number `tfsdk:"import_list" vyos:"import-list,omitempty"`

	// TagNodes

	ExistsTagProtocolsOspfAreaRange bool `tfsdk:"-" vyos:"range,child"`

	ExistsTagProtocolsOspfAreaVirtualLink bool `tfsdk:"-" vyos:"virtual-link,child"`

	// Nodes

	NodeProtocolsOspfAreaAreaType *ProtocolsOspfAreaAreaType `tfsdk:"area_type" vyos:"area-type,omitempty"`
}

// SetID configures the resource ID
func (o *ProtocolsOspfArea) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsOspfArea) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsOspfArea) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsOspfArea) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"area",
		o.SelfIdentifier.ProtocolsOspfArea.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsOspfArea) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (ospf) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (protocols) */
		"protocols", // Node

		"ospf", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ProtocolsOspfArea) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (ospf) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (protocols) */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsOspfArea) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
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
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  area, value must match: ^[.:a-zA-Z0-9-_/]+$",
							),
						),
					},
				},

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (ospf) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (protocols) */

			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"authentication":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (authentication) */
		schema.StringAttribute{
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

		"network":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (network) */
		schema.ListAttribute{
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

		"shortcut":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (shortcut) */
		schema.StringAttribute{
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

		"export_list":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (export-list) */
		schema.NumberAttribute{
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

		"import_list":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (import-list) */
		schema.NumberAttribute{
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

		// TagNodes

		// Nodes

		"area_type": schema.SingleNestedAttribute{
			Attributes: ProtocolsOspfAreaAreaType{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Area type

`,
			Description: `Area type

`,
		},
	}
}
