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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (secondary) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ServiceConfigSyncSecondary{}

// ServiceConfigSyncSecondary describes the resource data model.
// This is a basenode!
// Top level basenode type: `Node`
type ServiceConfigSyncSecondary struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceConfigSyncSecondaryAddress types.String `tfsdk:"address" vyos:"address,omitempty"`
	LeafServiceConfigSyncSecondaryPort    types.Number `tfsdk:"port" vyos:"port,omitempty"`
	LeafServiceConfigSyncSecondaryTimeout types.Number `tfsdk:"timeout" vyos:"timeout,omitempty"`
	LeafServiceConfigSyncSecondaryKey     types.String `tfsdk:"key" vyos:"key,omitempty"`

	// TagNodes

	// Nodes
}

// SetID configures the resource ID
func (o *ServiceConfigSyncSecondary) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceConfigSyncSecondary) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceConfigSyncSecondary) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceConfigSyncSecondary) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"secondary",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceConfigSyncSecondary) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (config-sync) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (service) */
		"service", // Node

		"config-sync", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ServiceConfigSyncSecondary) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (config-sync) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (service) */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceConfigSyncSecondary) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"address":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (address) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address

    |  Format    |  Description            |
    |------------|-------------------------|
    |  ipv4      |  IPv4 address to match  |
    |  ipv6      |  IPv6 address to match  |
    |  hostname  |  FQDN address to match  |
`,
			Description: `IP address

    |  Format    |  Description            |
    |------------|-------------------------|
    |  ipv4      |  IPv4 address to match  |
    |  ipv6      |  IPv6 address to match  |
    |  hostname  |  FQDN address to match  |
`,
		},

		"port":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (port) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Port number used by connection

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Numeric IP port  |
`,
			Description: `Port number used by connection

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Numeric IP port  |
`,

			// Default:          stringdefault.StaticString(`443`),
			Computed: true,
		},

		"timeout":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (timeout) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Connection API timeout

    |  Format  |  Description             |
    |----------|--------------------------|
    |  1-3600  |  Connection API timeout  |
`,
			Description: `Connection API timeout

    |  Format  |  Description             |
    |----------|--------------------------|
    |  1-3600  |  Connection API timeout  |
`,

			// Default:          stringdefault.StaticString(`60`),
			Computed: true,
		},

		"key":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (key) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `HTTP API key

`,
			Description: `HTTP API key

`,
		},

		// TagNodes

		// Nodes

	}
}
