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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (dynamic) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ServiceDNSDynamic{}

// ServiceDNSDynamic describes the resource data model.
// This is a basenode!
// Top level basenode type: `Node`
type ServiceDNSDynamic struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceDNSDynamicInterval types.Number `tfsdk:"interval" vyos:"interval,omitempty"`
	LeafServiceDNSDynamicVrf      types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`

	// TagNodes

	ExistsTagServiceDNSDynamicName bool `tfsdk:"-" vyos:"name,child"`

	// Nodes
}

// SetID configures the resource ID
func (o *ServiceDNSDynamic) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceDNSDynamic) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceDNSDynamic) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceDNSDynamic) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"dynamic",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceDNSDynamic) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (dns) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (service) */
		"service", // Node

		"dns", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ServiceDNSDynamic) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (dns) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (service) */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDNSDynamic) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"interval":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (interval) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interval in seconds to wait between Dynamic DNS updates

    |  Format   |  Description      |
    |-----------|-------------------|
    |  60-3600  |  Time in seconds  |
`,
			Description: `Interval in seconds to wait between Dynamic DNS updates

    |  Format   |  Description      |
    |-----------|-------------------|
    |  60-3600  |  Time in seconds  |
`,

			// Default:          stringdefault.StaticString(`300`),
			Computed: true,
		},

		"vrf":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (vrf) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VRF instance name

    |  Format  |  Description        |
    |----------|---------------------|
    |  txt     |  VRF instance name  |
`,
			Description: `VRF instance name

    |  Format  |  Description        |
    |----------|---------------------|
    |  txt     |  VRF instance name  |
`,
		},

		// TagNodes

		// Nodes

	}
}
