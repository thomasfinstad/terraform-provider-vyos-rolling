/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ProtocolsRIPInterfaceAuthenticationMdfive{}

// ProtocolsRIPInterfaceAuthenticationMdfive describes the resource data model.
type ProtocolsRIPInterfaceAuthenticationMdfive struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl */
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafProtocolsRIPInterfaceAuthenticationMdfivePassword types.String `tfsdk:"password" vyos:"password,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// SetID configures the resource ID
func (o *ProtocolsRIPInterfaceAuthenticationMdfive) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsRIPInterfaceAuthenticationMdfive) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsRIPInterfaceAuthenticationMdfive) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsRIPInterfaceAuthenticationMdfive) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"md5",
		o.SelfIdentifier.Attributes()["md5"].(types.Number).ValueBigFloat().String(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsRIPInterfaceAuthenticationMdfive) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-vyos-path-hack.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-vyos-path-hack.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-vyos-path-hack.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-vyos-path-hack.gotmpl */
		"protocols",

		"rip",

		"interface",

		o.SelfIdentifier.Attributes()["interface"].(types.String).ValueString(),

		"authentication",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ProtocolsRIPInterfaceAuthenticationMdfive) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack-for-non-global.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack-for-non-global.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-vyos-path-hack.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-vyos-path-hack.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-vyos-path-hack.gotmpl */
		"protocols",

		"rip",

		"interface",

		o.SelfIdentifier.Attributes()["interface"].(types.String).ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsRIPInterfaceAuthenticationMdfive) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"md5": schema.NumberAttribute{
					Required: true,
					MarkdownDescription: `MD5 key id

    |  Format  |  Description  |
    |----------|---------------|
    |  1-255   |  OSPF key id  |
`,
					Description: `MD5 key id

    |  Format  |  Description  |
    |----------|---------------|
    |  1-255   |  OSPF key id  |
`,
					PlanModifiers: []planmodifier.Number{
						numberplanmodifier.RequiresReplace(),
					},
				},

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl */

				"interface": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Interface name

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Interface name  |
`,
					Description: `Interface name

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Interface name  |
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
								regexp.MustCompile(`^[.:a-zA-Z0-9-_]+$`),
								"illegal character in  interface, value must match: ^[.:a-zA-Z0-9-_]+$",
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

		"password":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Authentication password

    |  Format  |  Description                      |
    |----------|-----------------------------------|
    |  txt     |  MD5 Key (16 characters or less)  |
`,
			Description: `Authentication password

    |  Format  |  Description                      |
    |----------|-----------------------------------|
    |  txt     |  MD5 Key (16 characters or less)  |
`,
		},

		// Nodes

	}
}
