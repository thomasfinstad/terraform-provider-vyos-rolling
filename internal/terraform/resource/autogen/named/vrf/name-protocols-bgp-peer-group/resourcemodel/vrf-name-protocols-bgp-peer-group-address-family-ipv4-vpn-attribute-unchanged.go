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

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourVpnAttributeUnchanged{}

// VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourVpnAttributeUnchanged describes the resource data model.
type VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourVpnAttributeUnchanged struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourVpnAttributeUnchangedAsPath  types.Bool `tfsdk:"as_path" vyos:"as-path,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourVpnAttributeUnchangedMed     types.Bool `tfsdk:"med" vyos:"med,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourVpnAttributeUnchangedNextHop types.Bool `tfsdk:"next_hop" vyos:"next-hop,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpPeerGroupAddressFamilyIPvfourVpnAttributeUnchanged) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"as_path":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Send nexthop unchanged

`,
			Description: `Send nexthop unchanged

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
