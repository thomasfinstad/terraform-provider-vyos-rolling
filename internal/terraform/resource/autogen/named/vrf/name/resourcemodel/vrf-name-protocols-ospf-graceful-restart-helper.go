/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (helper) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsOspfGracefulRestartHelper{}

// VrfNameProtocolsOspfGracefulRestartHelper describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VrfNameProtocolsOspfGracefulRestartHelper struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfGracefulRestartHelperPlannedOnly         types.Bool   `tfsdk:"planned_only" vyos:"planned-only,omitempty"`
	LeafVrfNameProtocolsOspfGracefulRestartHelperSupportedGraceTime  types.Number `tfsdk:"supported_grace_time" vyos:"supported-grace-time,omitempty"`
	LeafVrfNameProtocolsOspfGracefulRestartHelperNoStrictLsaChecking types.Bool   `tfsdk:"no_strict_lsa_checking" vyos:"no-strict-lsa-checking,omitempty"`

	// TagNodes

	// Nodes

	NodeVrfNameProtocolsOspfGracefulRestartHelperEnable *VrfNameProtocolsOspfGracefulRestartHelperEnable `tfsdk:"enable" vyos:"enable,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfGracefulRestartHelper) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"planned_only":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (planned-only) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Supported only planned restart

`,
			Description: `Supported only planned restart

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"supported_grace_time":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (supported-grace-time) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Supported grace timer

    |  Format   |  Description                |
    |-----------|-----------------------------|
    |  10-1800  |  Grace interval in seconds  |
`,
			Description: `Supported grace timer

    |  Format   |  Description                |
    |-----------|-----------------------------|
    |  10-1800  |  Grace interval in seconds  |
`,
		},

		"no_strict_lsa_checking":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (no-strict-lsa-checking) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable strict LSA check

`,
			Description: `Disable strict LSA check

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// TagNodes

		// Nodes

		"enable": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfGracefulRestartHelperEnable{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Enable helper support

`,
			Description: `Enable helper support

`,
		},
	}
}
