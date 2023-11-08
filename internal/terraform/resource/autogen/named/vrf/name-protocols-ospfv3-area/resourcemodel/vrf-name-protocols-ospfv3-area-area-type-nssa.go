// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsOspfvthreeAreaAreaTypeNssa describes the resource data model.
type VrfNameProtocolsOspfvthreeAreaAreaTypeNssa struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfvthreeAreaAreaTypeNssaDefaultInformationOriginate types.Bool `tfsdk:"default_information_originate" vyos:"default-information-originate,omitempty"`
	LeafVrfNameProtocolsOspfvthreeAreaAreaTypeNssaNoSummary                   types.Bool `tfsdk:"no_summary" vyos:"no-summary,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfvthreeAreaAreaTypeNssa) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"default_information_originate": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Originate Type 7 default into NSSA area

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"no_summary": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not inject inter-area routes into the stub

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
