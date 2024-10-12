// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsOspfMaxMetricRouterLsa{}

// VrfNameProtocolsOspfMaxMetricRouterLsa describes the resource data model.
type VrfNameProtocolsOspfMaxMetricRouterLsa struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfMaxMetricRouterLsaAdministrative types.Bool   `tfsdk:"administrative" vyos:"administrative,omitempty"`
	LeafVrfNameProtocolsOspfMaxMetricRouterLsaOnShutdown     types.Number `tfsdk:"on_shutdown" vyos:"on-shutdown,omitempty"`
	LeafVrfNameProtocolsOspfMaxMetricRouterLsaOnStartup      types.Number `tfsdk:"on_startup" vyos:"on-startup,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfMaxMetricRouterLsa) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"administrative": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Administratively apply, for an indefinite period

`,
			Description: `Administratively apply, for an indefinite period

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"on_shutdown": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Advertise stub-router prior to full shutdown of OSPF

    |  Format  |  Description                                      |
    |----------|---------------------------------------------------|
    |  5-100   |  Time (seconds) to advertise self as stub-router  |
`,
			Description: `Advertise stub-router prior to full shutdown of OSPF

    |  Format  |  Description                                      |
    |----------|---------------------------------------------------|
    |  5-100   |  Time (seconds) to advertise self as stub-router  |
`,
		},

		"on_startup": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Automatically advertise stub Router-LSA on startup of OSPF

    |  Format   |  Description                                      |
    |-----------|---------------------------------------------------|
    |  5-86400  |  Time (seconds) to advertise self as stub-router  |
`,
			Description: `Automatically advertise stub Router-LSA on startup of OSPF

    |  Format   |  Description                                      |
    |-----------|---------------------------------------------------|
    |  5-86400  |  Time (seconds) to advertise self as stub-router  |
`,
		},

		// Nodes

	}
}
