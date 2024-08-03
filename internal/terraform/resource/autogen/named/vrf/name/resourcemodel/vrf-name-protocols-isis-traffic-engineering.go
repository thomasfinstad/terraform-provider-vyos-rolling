// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsIsisTrafficEngineering{}

// VrfNameProtocolsIsisTrafficEngineering describes the resource data model.
type VrfNameProtocolsIsisTrafficEngineering struct {
	// LeafNodes
	LeafVrfNameProtocolsIsisTrafficEngineeringEnable  types.Bool   `tfsdk:"enable" vyos:"enable,omitempty"`
	LeafVrfNameProtocolsIsisTrafficEngineeringAddress types.String `tfsdk:"address" vyos:"address,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisTrafficEngineering) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"enable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable MPLS traffic engineering extensions

`,
			Description: `Enable MPLS traffic engineering extensions

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `MPLS traffic engineering router ID

    |  Format  |  Description   |
    |----------|----------------|
    |  ipv4    |  IPv4 address  |
`,
			Description: `MPLS traffic engineering router ID

    |  Format  |  Description   |
    |----------|----------------|
    |  ipv4    |  IPv4 address  |
`,
		},

		// Nodes

	}
}
