// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// VrfNameProtocolsIsisFastReroute describes the resource data model.
type VrfNameProtocolsIsisFastReroute struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsIsisFastRerouteLfa *VrfNameProtocolsIsisFastRerouteLfa `tfsdk:"lfa" vyos:"lfa,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisFastReroute) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"lfa": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisFastRerouteLfa{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Loop free alternate functionality

`,
		},
	}
}
