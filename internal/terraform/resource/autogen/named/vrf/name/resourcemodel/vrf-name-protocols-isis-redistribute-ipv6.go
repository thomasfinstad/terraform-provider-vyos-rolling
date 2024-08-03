// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsIsisRedistributeIPvsix{}

// VrfNameProtocolsIsisRedistributeIPvsix describes the resource data model.
type VrfNameProtocolsIsisRedistributeIPvsix struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeVrfNameProtocolsIsisRedistributeIPvsixBgp       *VrfNameProtocolsIsisRedistributeIPvsixBgp       `tfsdk:"bgp" vyos:"bgp,omitempty"`
	NodeVrfNameProtocolsIsisRedistributeIPvsixConnected *VrfNameProtocolsIsisRedistributeIPvsixConnected `tfsdk:"connected" vyos:"connected,omitempty"`
	NodeVrfNameProtocolsIsisRedistributeIPvsixKernel    *VrfNameProtocolsIsisRedistributeIPvsixKernel    `tfsdk:"kernel" vyos:"kernel,omitempty"`
	NodeVrfNameProtocolsIsisRedistributeIPvsixOspfsix   *VrfNameProtocolsIsisRedistributeIPvsixOspfsix   `tfsdk:"ospf6" vyos:"ospf6,omitempty"`
	NodeVrfNameProtocolsIsisRedistributeIPvsixRIPng     *VrfNameProtocolsIsisRedistributeIPvsixRIPng     `tfsdk:"ripng" vyos:"ripng,omitempty"`
	NodeVrfNameProtocolsIsisRedistributeIPvsixBabel     *VrfNameProtocolsIsisRedistributeIPvsixBabel     `tfsdk:"babel" vyos:"babel,omitempty"`
	NodeVrfNameProtocolsIsisRedistributeIPvsixStatic    *VrfNameProtocolsIsisRedistributeIPvsixStatic    `tfsdk:"static" vyos:"static,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisRedistributeIPvsix) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"bgp": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvsixBgp{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Redistribute BGP routes into IS-IS

`,
			Description: `Redistribute BGP routes into IS-IS

`,
		},

		"connected": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvsixConnected{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Redistribute connected routes into IS-IS

`,
			Description: `Redistribute connected routes into IS-IS

`,
		},

		"kernel": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvsixKernel{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Redistribute kernel routes into IS-IS

`,
			Description: `Redistribute kernel routes into IS-IS

`,
		},

		"ospf6": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvsixOspfsix{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Redistribute OSPFv3 routes into IS-IS

`,
			Description: `Redistribute OSPFv3 routes into IS-IS

`,
		},

		"ripng": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvsixRIPng{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Redistribute RIPng routes into IS-IS

`,
			Description: `Redistribute RIPng routes into IS-IS

`,
		},

		"babel": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvsixBabel{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Redistribute Babel routes into IS-IS

`,
			Description: `Redistribute Babel routes into IS-IS

`,
		},

		"static": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvsixStatic{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Redistribute static routes into IS-IS

`,
			Description: `Redistribute static routes into IS-IS

`,
		},
	}
}
