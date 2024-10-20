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

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsOspfvthree{}

// VrfNameProtocolsOspfvthree describes the resource data model.
type VrfNameProtocolsOspfvthree struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	ExistsTagVrfNameProtocolsOspfvthreeArea bool `tfsdk:"-" vyos:"area,child"`

	ExistsTagVrfNameProtocolsOspfvthreeInterface bool `tfsdk:"-" vyos:"interface,child"`

	// Nodes
	NodeVrfNameProtocolsOspfvthreeAutoCost            *VrfNameProtocolsOspfvthreeAutoCost            `tfsdk:"auto_cost" vyos:"auto-cost,omitempty"`
	NodeVrfNameProtocolsOspfvthreeDefaultInformation  *VrfNameProtocolsOspfvthreeDefaultInformation  `tfsdk:"default_information" vyos:"default-information,omitempty"`
	NodeVrfNameProtocolsOspfvthreeDistance            *VrfNameProtocolsOspfvthreeDistance            `tfsdk:"distance" vyos:"distance,omitempty"`
	NodeVrfNameProtocolsOspfvthreeGracefulRestart     *VrfNameProtocolsOspfvthreeGracefulRestart     `tfsdk:"graceful_restart" vyos:"graceful-restart,omitempty"`
	NodeVrfNameProtocolsOspfvthreeLogAdjacencyChanges *VrfNameProtocolsOspfvthreeLogAdjacencyChanges `tfsdk:"log_adjacency_changes" vyos:"log-adjacency-changes,omitempty"`
	NodeVrfNameProtocolsOspfvthreeParameters          *VrfNameProtocolsOspfvthreeParameters          `tfsdk:"parameters" vyos:"parameters,omitempty"`
	NodeVrfNameProtocolsOspfvthreeRedistribute        *VrfNameProtocolsOspfvthreeRedistribute        `tfsdk:"redistribute" vyos:"redistribute,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfvthree) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"auto_cost": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeAutoCost{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Calculate interface cost according to bandwidth

`,
			Description: `Calculate interface cost according to bandwidth

`,
		},

		"default_information": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeDefaultInformation{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Default route advertisment settings

`,
			Description: `Default route advertisment settings

`,
		},

		"distance": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeDistance{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Administrative distance

`,
			Description: `Administrative distance

`,
		},

		"graceful_restart": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeGracefulRestart{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Graceful Restart

`,
			Description: `Graceful Restart

`,
		},

		"log_adjacency_changes": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeLogAdjacencyChanges{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Log adjacency state changes

`,
			Description: `Log adjacency state changes

`,
		},

		"parameters": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeParameters{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `OSPFv3 specific parameters

`,
			Description: `OSPFv3 specific parameters

`,
		},

		"redistribute": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthreeRedistribute{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Redistribute information from another routing protocol

`,
			Description: `Redistribute information from another routing protocol

`,
		},
	}
}
