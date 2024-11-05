/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpPeerGroupTTLSecURIty{}

// VrfNameProtocolsBgpPeerGroupTTLSecURIty describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VrfNameProtocolsBgpPeerGroupTTLSecURIty struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpPeerGroupTTLSecURItyHops types.Number `tfsdk:"hops" vyos:"hops,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpPeerGroupTTLSecURIty) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"hops":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Number of the maximum number of hops to the BGP peer

    |  Format  |  Description     |
    |----------|------------------|
    |  1-254   |  Number of hops  |
`,
			Description: `Number of the maximum number of hops to the BGP peer

    |  Format  |  Description     |
    |----------|------------------|
    |  1-254   |  Number of hops  |
`,
		},

		// TagNodes

		// Nodes

	}
}
