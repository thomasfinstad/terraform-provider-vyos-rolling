/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchanged{}

// VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchanged describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchanged struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchangedAsPath  types.Bool `tfsdk:"as_path" vyos:"as-path,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchangedMed     types.Bool `tfsdk:"med" vyos:"med,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchangedNextHop types.Bool `tfsdk:"next_hop" vyos:"next-hop,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourUnicastAttributeUnchanged) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"as_path":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Send AS path unchanged

`,
			Description: `Send AS path unchanged

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"med":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Send multi-exit discriminator unchanged

`,
			Description: `Send multi-exit discriminator unchanged

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"next_hop":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Send nexthop unchanged

`,
			Description: `Send nexthop unchanged

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
