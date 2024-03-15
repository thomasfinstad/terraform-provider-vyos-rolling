// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// VrfNameProtocolsOspfvthreeRedistribute describes the resource data model.
type VrfNameProtocolsOspfvthreeRedistribute struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsOspfvthreeRedistributeBabel     *VrfNameProtocolsOspfvthreeRedistributeBabel     `tfsdk:"babel" vyos:"babel,omitempty"`
	NodeVrfNameProtocolsOspfvthreeRedistributeBgp       *VrfNameProtocolsOspfvthreeRedistributeBgp       `tfsdk:"bgp" vyos:"bgp,omitempty"`
	NodeVrfNameProtocolsOspfvthreeRedistributeConnected *VrfNameProtocolsOspfvthreeRedistributeConnected `tfsdk:"connected" vyos:"connected,omitempty"`
	NodeVrfNameProtocolsOspfvthreeRedistributeIsis      *VrfNameProtocolsOspfvthreeRedistributeIsis      `tfsdk:"isis" vyos:"isis,omitempty"`
	NodeVrfNameProtocolsOspfvthreeRedistributeKernel    *VrfNameProtocolsOspfvthreeRedistributeKernel    `tfsdk:"kernel" vyos:"kernel,omitempty"`
	NodeVrfNameProtocolsOspfvthreeRedistributeRIPng     *VrfNameProtocolsOspfvthreeRedistributeRIPng     `tfsdk:"ripng" vyos:"ripng,omitempty"`
	NodeVrfNameProtocolsOspfvthreeRedistributeStatic    *VrfNameProtocolsOspfvthreeRedistributeStatic    `tfsdk:"static" vyos:"static,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfvthreeRedistribute) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"babel": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeRedistributeBabel{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute Babel routes

`,
		},

		"bgp": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeRedistributeBgp{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute BGP routes

`,
		},

		"connected": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeRedistributeConnected{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute connected routes

`,
		},

		"isis": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeRedistributeIsis{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute IS-IS routes

`,
		},

		"kernel": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeRedistributeKernel{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute kernel routes

`,
		},

		"ripng": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeRedistributeRIPng{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute RIPNG routes

`,
		},

		"static": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeRedistributeStatic{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Redistribute static routes

`,
		},
	}
}
