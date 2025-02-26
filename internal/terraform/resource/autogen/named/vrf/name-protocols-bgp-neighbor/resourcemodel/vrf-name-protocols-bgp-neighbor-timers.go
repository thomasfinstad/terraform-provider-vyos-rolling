/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (timers) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpNeighborTimers{}

// VrfNameProtocolsBgpNeighborTimers describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VrfNameProtocolsBgpNeighborTimers struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborTimersConnect   types.String `tfsdk:"connect" vyos:"connect,omitempty"`
	LeafVrfNameProtocolsBgpNeighborTimersHoldtime  types.String `tfsdk:"holdtime" vyos:"holdtime,omitempty"`
	LeafVrfNameProtocolsBgpNeighborTimersKeepalive types.Number `tfsdk:"keepalive" vyos:"keepalive,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborTimers) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"connect":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (connect) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `BGP connect timer for this neighbor

    |  Format   |  Description               |
    |-----------|----------------------------|
    |  1-65535  |  Connect timer in seconds  |
    |  0        |  Disable connect timer     |
`,
			Description: `BGP connect timer for this neighbor

    |  Format   |  Description               |
    |-----------|----------------------------|
    |  1-65535  |  Connect timer in seconds  |
    |  0        |  Disable connect timer     |
`,
		},

		"holdtime":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (holdtime) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Hold timer

    |  Format   |  Description            |
    |-----------|-------------------------|
    |  1-65535  |  Hold timer in seconds  |
    |  0        |  Disable hold timer     |
`,
			Description: `Hold timer

    |  Format   |  Description            |
    |-----------|-------------------------|
    |  1-65535  |  Hold timer in seconds  |
    |  0        |  Disable hold timer     |
`,
		},

		"keepalive":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (keepalive) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `BGP keepalive interval for this neighbor

    |  Format   |  Description                    |
    |-----------|---------------------------------|
    |  1-65535  |  Keepalive interval in seconds  |
`,
			Description: `BGP keepalive interval for this neighbor

    |  Format   |  Description                    |
    |-----------|---------------------------------|
    |  1-65535  |  Keepalive interval in seconds  |
`,
		},

		// TagNodes

		// Nodes

	}
}
