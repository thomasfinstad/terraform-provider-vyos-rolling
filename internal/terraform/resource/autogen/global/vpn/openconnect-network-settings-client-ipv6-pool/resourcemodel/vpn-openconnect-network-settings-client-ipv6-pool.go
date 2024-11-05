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

var _ helpers.VyosTopResourceDataModel = &VpnOpenconnectNetworkSettingsClientIPvsixPool{}

// VpnOpenconnectNetworkSettingsClientIPvsixPool describes the resource data model.
// This is a basenode!
// Top level basenode type: `Node`
type VpnOpenconnectNetworkSettingsClientIPvsixPool struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafVpnOpenconnectNetworkSettingsClientIPvsixPoolPrefix types.String `tfsdk:"prefix" vyos:"prefix,omitempty"`
	LeafVpnOpenconnectNetworkSettingsClientIPvsixPoolMask   types.Number `tfsdk:"mask" vyos:"mask,omitempty"`

	// TagNodes

	// Nodes
}

// SetID configures the resource ID
func (o *VpnOpenconnectNetworkSettingsClientIPvsixPool) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *VpnOpenconnectNetworkSettingsClientIPvsixPool) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *VpnOpenconnectNetworkSettingsClientIPvsixPool) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnOpenconnectNetworkSettingsClientIPvsixPool) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"client-ipv6-pool",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *VpnOpenconnectNetworkSettingsClientIPvsixPool) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */
		"vpn", // Node

		"openconnect", // Node

		"network-settings", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *VpnOpenconnectNetworkSettingsClientIPvsixPool) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnOpenconnectNetworkSettingsClientIPvsixPool) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"prefix":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Pool of addresses used to assign to clients

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  ipv6net  |  IPv6 address and prefix length  |
`,
			Description: `Pool of addresses used to assign to clients

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  ipv6net  |  IPv6 address and prefix length  |
`,
		},

		"mask":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Prefix length used for individual client

    |  Format  |  Description           |
    |----------|------------------------|
    |  48-128  |  Client prefix length  |
`,
			Description: `Prefix length used for individual client

    |  Format  |  Description           |
    |----------|------------------------|
    |  48-128  |  Client prefix length  |
`,

			// Default:          stringdefault.StaticString(`64`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
