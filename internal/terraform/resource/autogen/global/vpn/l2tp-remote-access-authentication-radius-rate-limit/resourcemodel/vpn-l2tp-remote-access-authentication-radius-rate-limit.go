/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (rate-limit) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &VpnLtwotpRemoteAccessAuthenticationRadiusRateLimit{}

// VpnLtwotpRemoteAccessAuthenticationRadiusRateLimit describes the resource data model.
// This is a basenode!
// Top level basenode type: `Node`
type VpnLtwotpRemoteAccessAuthenticationRadiusRateLimit struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafVpnLtwotpRemoteAccessAuthenticationRadiusRateLimitAttribute  types.String `tfsdk:"attribute" vyos:"attribute,omitempty"`
	LeafVpnLtwotpRemoteAccessAuthenticationRadiusRateLimitVendor     types.String `tfsdk:"vendor" vyos:"vendor,omitempty"`
	LeafVpnLtwotpRemoteAccessAuthenticationRadiusRateLimitEnable     types.Bool   `tfsdk:"enable" vyos:"enable,omitempty"`
	LeafVpnLtwotpRemoteAccessAuthenticationRadiusRateLimitMultIPlier types.String `tfsdk:"multiplier" vyos:"multiplier,omitempty"`

	// TagNodes

	// Nodes
}

// SetID configures the resource ID
func (o *VpnLtwotpRemoteAccessAuthenticationRadiusRateLimit) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *VpnLtwotpRemoteAccessAuthenticationRadiusRateLimit) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *VpnLtwotpRemoteAccessAuthenticationRadiusRateLimit) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnLtwotpRemoteAccessAuthenticationRadiusRateLimit) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"rate-limit",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *VpnLtwotpRemoteAccessAuthenticationRadiusRateLimit) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (radius) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (authentication) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (remote-access) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (l2tp) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (vpn) */
		"vpn", // Node

		"l2tp", // Node

		"remote-access", // Node

		"authentication", // Node

		"radius", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *VpnLtwotpRemoteAccessAuthenticationRadiusRateLimit) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (radius) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (authentication) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (remote-access) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (l2tp) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (vpn) */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnLtwotpRemoteAccessAuthenticationRadiusRateLimit) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"attribute":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (attribute) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `RADIUS attribute that contains rate information

`,
			Description: `RADIUS attribute that contains rate information

`,

			// Default:          stringdefault.StaticString(`Filter-Id`),
			Computed: true,
		},

		"vendor":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (vendor) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Vendor dictionary

`,
			Description: `Vendor dictionary

`,
		},

		"enable":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (enable) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable bandwidth shaping via RADIUS

`,
			Description: `Enable bandwidth shaping via RADIUS

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"multiplier":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (multiplier) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Shaper multiplier

    |  Format        |  Description        |
    |----------------|---------------------|
    |  <0.001-1000>  |  Shaper multiplier  |
`,
			Description: `Shaper multiplier

    |  Format        |  Description        |
    |----------------|---------------------|
    |  <0.001-1000>  |  Shaper multiplier  |
`,

			// Default:          stringdefault.StaticString(`1`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
