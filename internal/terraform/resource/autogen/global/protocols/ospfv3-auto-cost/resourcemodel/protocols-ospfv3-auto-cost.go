/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (auto-cost) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ProtocolsOspfvthreeAutoCost{}

// ProtocolsOspfvthreeAutoCost describes the resource data model.
// This is a basenode!
// Top level basenode type: `Node`
type ProtocolsOspfvthreeAutoCost struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafProtocolsOspfvthreeAutoCostReferenceBandwIDth types.Number `tfsdk:"reference_bandwidth" vyos:"reference-bandwidth,omitempty"`

	// TagNodes

	// Nodes
}

// SetID configures the resource ID
func (o *ProtocolsOspfvthreeAutoCost) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsOspfvthreeAutoCost) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsOspfvthreeAutoCost) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsOspfvthreeAutoCost) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"auto-cost",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsOspfvthreeAutoCost) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (ospfv3) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (protocols) */
		"protocols", // Node

		"ospfv3", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ProtocolsOspfvthreeAutoCost) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (ospfv3) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (protocols) */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsOspfvthreeAutoCost) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"reference_bandwidth":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (reference-bandwidth) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Reference bandwidth method to assign cost

    |  Format     |  Description                            |
    |-------------|-----------------------------------------|
    |  1-4294967  |  Reference bandwidth cost in Mbits/sec  |
`,
			Description: `Reference bandwidth method to assign cost

    |  Format     |  Description                            |
    |-------------|-----------------------------------------|
    |  1-4294967  |  Reference bandwidth cost in Mbits/sec  |
`,

			// Default:          stringdefault.StaticString(`100`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
