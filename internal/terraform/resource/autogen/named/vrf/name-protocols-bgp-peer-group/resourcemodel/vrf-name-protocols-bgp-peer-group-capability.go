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

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpPeerGroupCapability{}

// VrfNameProtocolsBgpPeerGroupCapability describes the resource data model.
type VrfNameProtocolsBgpPeerGroupCapability struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpPeerGroupCapabilityDynamic         types.Bool `tfsdk:"dynamic" vyos:"dynamic,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupCapabilityExtendedNexthop types.Bool `tfsdk:"extended_nexthop" vyos:"extended-nexthop,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupCapabilitySoftwareVersion types.Bool `tfsdk:"software_version" vyos:"software-version,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpPeerGroupCapability) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"dynamic":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Advertise dynamic capability to this neighbor

`,
			Description: `Advertise dynamic capability to this neighbor

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"extended_nexthop":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Advertise extended-nexthop capability to this neighbor

`,
			Description: `Advertise extended-nexthop capability to this neighbor

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"software_version":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Advertise Software Version capability to the peer

`,
			Description: `Advertise Software Version capability to the peer

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
