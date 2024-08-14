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

var _ helpers.VyosResourceDataModel = &ProtocolsBgpPeerGroupLocalAsNoPrepend{}

// ProtocolsBgpPeerGroupLocalAsNoPrepend describes the resource data model.
type ProtocolsBgpPeerGroupLocalAsNoPrepend struct {
	// LeafNodes
	LeafProtocolsBgpPeerGroupLocalAsNoPrependReplaceAs types.Bool `tfsdk:"replace_as" vyos:"replace-as,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpPeerGroupLocalAsNoPrepend) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"replace_as": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Prepend only local-as from/to updates for eBGP peers

`,
			Description: `Prepend only local-as from/to updates for eBGP peers

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}