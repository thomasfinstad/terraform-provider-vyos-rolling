// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsRIPngInterfaceSplitHorizon describes the resource data model.
type ProtocolsRIPngInterfaceSplitHorizon struct {
	// LeafNodes
	LeafProtocolsRIPngInterfaceSplitHorizonDisable       types.Bool `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafProtocolsRIPngInterfaceSplitHorizonPoisonReverse types.Bool `tfsdk:"poison_reverse" vyos:"poison-reverse,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsRIPngInterfaceSplitHorizon) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable split horizon on specified interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"poison_reverse": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable split horizon on specified interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
