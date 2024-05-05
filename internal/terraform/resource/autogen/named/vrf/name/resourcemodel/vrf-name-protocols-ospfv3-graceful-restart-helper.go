// Package resourcemodel code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsOspfvthreeGracefulRestartHelper{}

// VrfNameProtocolsOspfvthreeGracefulRestartHelper describes the resource data model.
type VrfNameProtocolsOspfvthreeGracefulRestartHelper struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfvthreeGracefulRestartHelperPlannedOnly        types.Bool   `tfsdk:"planned_only" vyos:"planned-only,omitempty"`
	LeafVrfNameProtocolsOspfvthreeGracefulRestartHelperSupportedGraceTime types.Number `tfsdk:"supported_grace_time" vyos:"supported-grace-time,omitempty"`
	LeafVrfNameProtocolsOspfvthreeGracefulRestartHelperLsaCheckDisable    types.Bool   `tfsdk:"lsa_check_disable" vyos:"lsa-check-disable,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsOspfvthreeGracefulRestartHelperEnable *VrfNameProtocolsOspfvthreeGracefulRestartHelperEnable `tfsdk:"enable" vyos:"enable,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfvthreeGracefulRestartHelper) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"planned_only": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Supported only planned restart

`,
			Description: `Supported only planned restart

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"supported_grace_time": schema.NumberAttribute{
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

		"lsa_check_disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable strict LSA check

`,
			Description: `Disable strict LSA check

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

		"enable": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeGracefulRestartHelperEnable{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Enable helper support

`,
			Description: `Enable helper support

`,
		},
	}
}
