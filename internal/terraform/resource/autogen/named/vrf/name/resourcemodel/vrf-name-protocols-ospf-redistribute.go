// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsOspfRedistribute{}

// VrfNameProtocolsOspfRedistribute describes the resource data model.
type VrfNameProtocolsOspfRedistribute struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	ExistsTagVrfNameProtocolsOspfRedistributeTable bool `tfsdk:"table" vyos:"table,child"`

	// Nodes
	NodeVrfNameProtocolsOspfRedistributeBgp       *VrfNameProtocolsOspfRedistributeBgp       `tfsdk:"bgp" vyos:"bgp,omitempty"`
	NodeVrfNameProtocolsOspfRedistributeConnected *VrfNameProtocolsOspfRedistributeConnected `tfsdk:"connected" vyos:"connected,omitempty"`
	NodeVrfNameProtocolsOspfRedistributeIsis      *VrfNameProtocolsOspfRedistributeIsis      `tfsdk:"isis" vyos:"isis,omitempty"`
	NodeVrfNameProtocolsOspfRedistributeKernel    *VrfNameProtocolsOspfRedistributeKernel    `tfsdk:"kernel" vyos:"kernel,omitempty"`
	NodeVrfNameProtocolsOspfRedistributeRIP       *VrfNameProtocolsOspfRedistributeRIP       `tfsdk:"rip" vyos:"rip,omitempty"`
	NodeVrfNameProtocolsOspfRedistributeBabel     *VrfNameProtocolsOspfRedistributeBabel     `tfsdk:"babel" vyos:"babel,omitempty"`
	NodeVrfNameProtocolsOspfRedistributeStatic    *VrfNameProtocolsOspfRedistributeStatic    `tfsdk:"static" vyos:"static,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfRedistribute) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"bgp": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfRedistributeBgp{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Redistribute BGP routes

`,
			Description: `Redistribute BGP routes

`,
		},

		"connected": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfRedistributeConnected{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Redistribute connected routes

`,
			Description: `Redistribute connected routes

`,
		},

		"isis": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfRedistributeIsis{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Redistribute IS-IS routes

`,
			Description: `Redistribute IS-IS routes

`,
		},

		"kernel": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfRedistributeKernel{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Redistribute Kernel routes

`,
			Description: `Redistribute Kernel routes

`,
		},

		"rip": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfRedistributeRIP{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Redistribute RIP routes

`,
			Description: `Redistribute RIP routes

`,
		},

		"babel": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfRedistributeBabel{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Redistribute Babel routes

`,
			Description: `Redistribute Babel routes

`,
		},

		"static": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfRedistributeStatic{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Redistribute statically configured routes

`,
			Description: `Redistribute statically configured routes

`,
		},
	}
}
