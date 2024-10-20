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
var _ helpers.VyosTopResourceDataModel = &SystemFlowAccountingNetflowTimeout{}

// SystemFlowAccountingNetflowTimeout describes the resource data model.
type SystemFlowAccountingNetflowTimeout struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafSystemFlowAccountingNetflowTimeoutExpiryInterval types.Number `tfsdk:"expiry_interval" vyos:"expiry-interval,omitempty"`
	LeafSystemFlowAccountingNetflowTimeoutFlowGeneric    types.Number `tfsdk:"flow_generic" vyos:"flow-generic,omitempty"`
	LeafSystemFlowAccountingNetflowTimeoutIcmp           types.Number `tfsdk:"icmp" vyos:"icmp,omitempty"`
	LeafSystemFlowAccountingNetflowTimeoutMaxActiveLife  types.Number `tfsdk:"max_active_life" vyos:"max-active-life,omitempty"`
	LeafSystemFlowAccountingNetflowTimeoutTCPFin         types.Number `tfsdk:"tcp_fin" vyos:"tcp-fin,omitempty"`
	LeafSystemFlowAccountingNetflowTimeoutTCPGeneric     types.Number `tfsdk:"tcp_generic" vyos:"tcp-generic,omitempty"`
	LeafSystemFlowAccountingNetflowTimeoutTCPRst         types.Number `tfsdk:"tcp_rst" vyos:"tcp-rst,omitempty"`
	LeafSystemFlowAccountingNetflowTimeoutUDP            types.Number `tfsdk:"udp" vyos:"udp,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *SystemFlowAccountingNetflowTimeout) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *SystemFlowAccountingNetflowTimeout) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *SystemFlowAccountingNetflowTimeout) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *SystemFlowAccountingNetflowTimeout) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"timeout",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *SystemFlowAccountingNetflowTimeout) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/global/resource-model-parent-vyos-path-hack.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/global/resource-model-parent-vyos-path-hack.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/global/resource-model-parent-vyos-path-hack.gotmpl */
		"system",

		"flow-accounting",

		"netflow",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *SystemFlowAccountingNetflowTimeout) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack-for-non-global.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack-for-non-global.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack-for-non-global.gotmpl */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemFlowAccountingNetflowTimeout) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"expiry_interval":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Expiry scan interval

    |  Format        |  Description           |
    |----------------|------------------------|
    |  0-2147483647  |  Expiry scan interval  |
`,
			Description: `Expiry scan interval

    |  Format        |  Description           |
    |----------------|------------------------|
    |  0-2147483647  |  Expiry scan interval  |
`,

			// Default:          stringdefault.StaticString(`60`),
			Computed: true,
		},

		"flow_generic":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Generic flow timeout value

    |  Format        |  Description                      |
    |----------------|-----------------------------------|
    |  0-2147483647  |  Generic flow timeout in seconds  |
`,
			Description: `Generic flow timeout value

    |  Format        |  Description                      |
    |----------------|-----------------------------------|
    |  0-2147483647  |  Generic flow timeout in seconds  |
`,

			// Default:          stringdefault.StaticString(`3600`),
			Computed: true,
		},

		"icmp":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `ICMP timeout value

    |  Format        |  Description              |
    |----------------|---------------------------|
    |  0-2147483647  |  ICMP timeout in seconds  |
`,
			Description: `ICMP timeout value

    |  Format        |  Description              |
    |----------------|---------------------------|
    |  0-2147483647  |  ICMP timeout in seconds  |
`,

			// Default:          stringdefault.StaticString(`300`),
			Computed: true,
		},

		"max_active_life":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Max active timeout value

    |  Format        |  Description                    |
    |----------------|---------------------------------|
    |  0-2147483647  |  Max active timeout in seconds  |
`,
			Description: `Max active timeout value

    |  Format        |  Description                    |
    |----------------|---------------------------------|
    |  0-2147483647  |  Max active timeout in seconds  |
`,

			// Default:          stringdefault.StaticString(`604800`),
			Computed: true,
		},

		"tcp_fin":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TCP finish timeout value

    |  Format        |  Description                 |
    |----------------|------------------------------|
    |  0-2147483647  |  TCP FIN timeout in seconds  |
`,
			Description: `TCP finish timeout value

    |  Format        |  Description                 |
    |----------------|------------------------------|
    |  0-2147483647  |  TCP FIN timeout in seconds  |
`,

			// Default:          stringdefault.StaticString(`300`),
			Computed: true,
		},

		"tcp_generic":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TCP generic timeout value

    |  Format        |  Description                     |
    |----------------|----------------------------------|
    |  0-2147483647  |  TCP generic timeout in seconds  |
`,
			Description: `TCP generic timeout value

    |  Format        |  Description                     |
    |----------------|----------------------------------|
    |  0-2147483647  |  TCP generic timeout in seconds  |
`,

			// Default:          stringdefault.StaticString(`3600`),
			Computed: true,
		},

		"tcp_rst":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TCP reset timeout value

    |  Format        |  Description                 |
    |----------------|------------------------------|
    |  0-2147483647  |  TCP RST timeout in seconds  |
`,
			Description: `TCP reset timeout value

    |  Format        |  Description                 |
    |----------------|------------------------------|
    |  0-2147483647  |  TCP RST timeout in seconds  |
`,

			// Default:          stringdefault.StaticString(`120`),
			Computed: true,
		},

		"udp":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `UDP timeout value

    |  Format        |  Description             |
    |----------------|--------------------------|
    |  0-2147483647  |  UDP timeout in seconds  |
`,
			Description: `UDP timeout value

    |  Format        |  Description             |
    |----------------|--------------------------|
    |  0-2147483647  |  UDP timeout in seconds  |
`,

			// Default:          stringdefault.StaticString(`300`),
			Computed: true,
		},
	}
}
