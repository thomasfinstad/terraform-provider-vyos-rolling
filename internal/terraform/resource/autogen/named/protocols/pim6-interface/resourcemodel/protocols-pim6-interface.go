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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (interface) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ProtocolsPimsixInterface{}

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-tag-node-identifier (interface) */
// ProtocolsPimsixInterfaceSelfIdentifier used by TagNodes to keep the id info
type ProtocolsPimsixInterfaceSelfIdentifier struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (interface) */

	ProtocolsPimsixInterface types.String `tfsdk:"interface" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (pim6) */

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (protocols) */
}

// ProtocolsPimsixInterface describes the resource data model.
// This is a basenode!
// Top level basenode type: `TagNode`
type ProtocolsPimsixInterface struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	SelfIdentifier *ProtocolsPimsixInterfaceSelfIdentifier `tfsdk:"identifier" vyos:"-,self-id"`

	// LeafNodes
	LeafProtocolsPimsixInterfaceNoBsm        types.Bool   `tfsdk:"no_bsm" vyos:"no-bsm,omitempty"`
	LeafProtocolsPimsixInterfaceNoUnicastBsm types.Bool   `tfsdk:"no_unicast_bsm" vyos:"no-unicast-bsm,omitempty"`
	LeafProtocolsPimsixInterfaceDrPriority   types.Number `tfsdk:"dr_priority" vyos:"dr-priority,omitempty"`
	LeafProtocolsPimsixInterfaceHello        types.Number `tfsdk:"hello" vyos:"hello,omitempty"`
	LeafProtocolsPimsixInterfacePassive      types.Bool   `tfsdk:"passive" vyos:"passive,omitempty"`

	// TagNodes

	// Nodes

	NodeProtocolsPimsixInterfaceMld *ProtocolsPimsixInterfaceMld `tfsdk:"mld" vyos:"mld,omitempty"`
}

// SetID configures the resource ID
func (o *ProtocolsPimsixInterface) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsPimsixInterface) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsPimsixInterface) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsPimsixInterface) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"interface",
		o.SelfIdentifier.ProtocolsPimsixInterface.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsPimsixInterface) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (pim6) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (protocols) */
		"protocols", // Node

		"pim6", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ProtocolsPimsixInterface) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (pim6) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (protocols) */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsPimsixInterface) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"interface": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `PIMv6 interface

`,
					Description: `PIMv6 interface

`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in interface, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  interface, value must match: ^[.:a-zA-Z0-9-_/]+$",
							),
						),
					},
				},

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (pim6) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (protocols) */

			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"no_bsm":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (no-bsm) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not process bootstrap messages

`,
			Description: `Do not process bootstrap messages

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"no_unicast_bsm":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (no-unicast-bsm) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not process unicast bootstrap messages

`,
			Description: `Do not process unicast bootstrap messages

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"dr_priority":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (dr-priority) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Designated router election priority

    |  Format        |  Description  |
    |----------------|---------------|
    |  1-4294967295  |  DR Priority  |
`,
			Description: `Designated router election priority

    |  Format        |  Description  |
    |----------------|---------------|
    |  1-4294967295  |  DR Priority  |
`,
		},

		"hello":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (hello) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Hello Interval

    |  Format  |  Description                |
    |----------|-----------------------------|
    |  1-180   |  Hello Interval in seconds  |
`,
			Description: `Hello Interval

    |  Format  |  Description                |
    |----------|-----------------------------|
    |  1-180   |  Hello Interval in seconds  |
`,
		},

		"passive":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (passive) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable sending and receiving PIM control packets on the interface

`,
			Description: `Disable sending and receiving PIM control packets on the interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// TagNodes

		// Nodes

		"mld": schema.SingleNestedAttribute{
			Attributes: ProtocolsPimsixInterfaceMld{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Multicast Listener Discovery (MLD)

`,
			Description: `Multicast Listener Discovery (MLD)

`,
		},
	}
}
