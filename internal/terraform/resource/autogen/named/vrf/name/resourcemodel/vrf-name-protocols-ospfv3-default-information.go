// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// VrfNameProtocolsOspfvthreeDefaultInformation describes the resource data model.
type VrfNameProtocolsOspfvthreeDefaultInformation struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsOspfvthreeDefaultInformationOriginate *VrfNameProtocolsOspfvthreeDefaultInformationOriginate `tfsdk:"originate" vyos:"originate,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfvthreeDefaultInformation) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"originate": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeDefaultInformationOriginate{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Distribute a default route

`,
		},
	}
}
