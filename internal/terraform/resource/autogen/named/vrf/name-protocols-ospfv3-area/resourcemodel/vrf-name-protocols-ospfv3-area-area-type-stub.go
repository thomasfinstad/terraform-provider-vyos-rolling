// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsOspfvthreeAreaAreaTypeStub{}

// VrfNameProtocolsOspfvthreeAreaAreaTypeStub describes the resource data model.
type VrfNameProtocolsOspfvthreeAreaAreaTypeStub struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfvthreeAreaAreaTypeStubNoSummary types.Bool `tfsdk:"no_summary" vyos:"no-summary,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfvthreeAreaAreaTypeStub) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"no_summary": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not inject inter-area routes into the stub

`,
			Description: `Do not inject inter-area routes into the stub

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
