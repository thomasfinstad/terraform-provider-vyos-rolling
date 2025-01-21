/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (kernel) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsIsisRedistributeIPvfourKernel{}

// VrfNameProtocolsIsisRedistributeIPvfourKernel describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VrfNameProtocolsIsisRedistributeIPvfourKernel struct {
	// LeafNodes

	// TagNodes

	// Nodes

	NodeVrfNameProtocolsIsisRedistributeIPvfourKernelLevelOne *VrfNameProtocolsIsisRedistributeIPvfourKernelLevelOne `tfsdk:"level_1" vyos:"level-1,omitempty"`

	NodeVrfNameProtocolsIsisRedistributeIPvfourKernelLevelTwo *VrfNameProtocolsIsisRedistributeIPvfourKernelLevelTwo `tfsdk:"level_2" vyos:"level-2,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisRedistributeIPvfourKernel) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"level_1": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvfourKernelLevelOne{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Redistribute into level-1

`,
			Description: `Redistribute into level-1

`,
		},

		"level_2": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistributeIPvfourKernelLevelTwo{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Redistribute into level-2

`,
			Description: `Redistribute into level-2

`,
		},
	}
}
