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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ProtocolsOspfTimersThroTTLeSpf{}

// ProtocolsOspfTimersThroTTLeSpf describes the resource data model.
// This is a basenode!
// Top level basenode type: `Node`
type ProtocolsOspfTimersThroTTLeSpf struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafProtocolsOspfTimersThroTTLeSpfDelay           types.Number `tfsdk:"delay" vyos:"delay,omitempty"`
	LeafProtocolsOspfTimersThroTTLeSpfInitialHoldtime types.Number `tfsdk:"initial_holdtime" vyos:"initial-holdtime,omitempty"`
	LeafProtocolsOspfTimersThroTTLeSpfMaxHoldtime     types.Number `tfsdk:"max_holdtime" vyos:"max-holdtime,omitempty"`

	// TagNodes

	// Nodes
}

// SetID configures the resource ID
func (o *ProtocolsOspfTimersThroTTLeSpf) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsOspfTimersThroTTLeSpf) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsOspfTimersThroTTLeSpf) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsOspfTimersThroTTLeSpf) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"spf",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsOspfTimersThroTTLeSpf) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */
		"protocols", // Node

		"ospf", // Node

		"timers", // Node

		"throttle", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ProtocolsOspfTimersThroTTLeSpf) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsOspfTimersThroTTLeSpf) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"delay":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Delay from the first change received to SPF calculation

    |  Format    |  Description            |
    |------------|-------------------------|
    |  0-600000  |  Delay in milliseconds  |
`,
			Description: `Delay from the first change received to SPF calculation

    |  Format    |  Description            |
    |------------|-------------------------|
    |  0-600000  |  Delay in milliseconds  |
`,

			// Default:          stringdefault.StaticString(`200`),
			Computed: true,
		},

		"initial_holdtime":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Initial hold time between consecutive SPF calculations

    |  Format    |  Description                        |
    |------------|-------------------------------------|
    |  0-600000  |  Initial hold time in milliseconds  |
`,
			Description: `Initial hold time between consecutive SPF calculations

    |  Format    |  Description                        |
    |------------|-------------------------------------|
    |  0-600000  |  Initial hold time in milliseconds  |
`,

			// Default:          stringdefault.StaticString(`1000`),
			Computed: true,
		},

		"max_holdtime":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum hold time

    |  Format    |  Description                    |
    |------------|---------------------------------|
    |  0-600000  |  Max hold time in milliseconds  |
`,
			Description: `Maximum hold time

    |  Format    |  Description                    |
    |------------|---------------------------------|
    |  0-600000  |  Max hold time in milliseconds  |
`,

			// Default:          stringdefault.StaticString(`10000`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
