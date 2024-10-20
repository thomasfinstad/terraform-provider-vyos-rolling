/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsOspfvthreeRedistribute{}

// VrfNameProtocolsOspfvthreeRedistribute describes the resource data model.
type VrfNameProtocolsOspfvthreeRedistribute struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

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
func (o VrfNameProtocolsOspfvthreeRedistribute) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"babel": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeRedistributeBabel{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Redistribute Babel routes

`,
			Description: `Redistribute Babel routes

`,
		},

		"bgp": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeRedistributeBgp{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Redistribute BGP routes

`,
			Description: `Redistribute BGP routes

`,
		},

		"connected": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeRedistributeConnected{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Redistribute connected routes

`,
			Description: `Redistribute connected routes

`,
		},

		"isis": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeRedistributeIsis{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Redistribute IS-IS routes

`,
			Description: `Redistribute IS-IS routes

`,
		},

		"kernel": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeRedistributeKernel{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Redistribute kernel routes

`,
			Description: `Redistribute kernel routes

`,
		},

		"ripng": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeRedistributeRIPng{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Redistribute RIPNG routes

`,
			Description: `Redistribute RIPNG routes

`,
		},

		"static": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeRedistributeStatic{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Redistribute static routes

`,
			Description: `Redistribute static routes

`,
		},
	}
}
