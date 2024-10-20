/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/global/resource-model.gotmpl */
// Validate compliance
var _ helpers.VyosTopResourceDataModel = &ProtocolsOspfDistanceOspf{}

// ProtocolsOspfDistanceOspf describes the resource data model.
type ProtocolsOspfDistanceOspf struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafProtocolsOspfDistanceOspfExternal  types.Number `tfsdk:"external" vyos:"external,omitempty"`
	LeafProtocolsOspfDistanceOspfInterArea types.Number `tfsdk:"inter_area" vyos:"inter-area,omitempty"`
	LeafProtocolsOspfDistanceOspfIntraArea types.Number `tfsdk:"intra_area" vyos:"intra-area,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *ProtocolsOspfDistanceOspf) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsOspfDistanceOspf) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsOspfDistanceOspf) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsOspfDistanceOspf) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"ospf",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsOspfDistanceOspf) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/global/resource-model-parent-vyos-path-hack.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/global/resource-model-parent-vyos-path-hack.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/global/resource-model-parent-vyos-path-hack.gotmpl */
		"protocols",

		"ospf",

		"distance",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *ProtocolsOspfDistanceOspf) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack-for-non-global.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack-for-non-global.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack-for-non-global.gotmpl */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsOspfDistanceOspf) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"external":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Distance for external routes

    |  Format  |  Description                   |
    |----------|--------------------------------|
    |  1-255   |  Distance for external routes  |
`,
			Description: `Distance for external routes

    |  Format  |  Description                   |
    |----------|--------------------------------|
    |  1-255   |  Distance for external routes  |
`,
		},

		"inter_area":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Distance for inter-area routes

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  1-255   |  Distance for inter-area routes  |
`,
			Description: `Distance for inter-area routes

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  1-255   |  Distance for inter-area routes  |
`,
		},

		"intra_area":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Distance for intra-area routes

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  1-255   |  Distance for intra-area routes  |
`,
			Description: `Distance for intra-area routes

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  1-255   |  Distance for intra-area routes  |
`,
		},
	}
}
