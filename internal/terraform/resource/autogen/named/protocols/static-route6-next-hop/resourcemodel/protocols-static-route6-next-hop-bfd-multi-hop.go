/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (multi-hop) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &ProtocolsStaticRoutesixNextHopBfdMultiHop{}

// ProtocolsStaticRoutesixNextHopBfdMultiHop describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type ProtocolsStaticRoutesixNextHopBfdMultiHop struct {
	// LeafNodes
	LeafProtocolsStaticRoutesixNextHopBfdMultiHopSourceAddress types.String `tfsdk:"source_address" vyos:"source-address,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsStaticRoutesixNextHopBfdMultiHop) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"source_address":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (source-address) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv6 address used to initiate connection

    |  Format  |  Description          |
    |----------|-----------------------|
    |  ipv6    |  IPv6 source address  |
`,
			Description: `IPv6 address used to initiate connection

    |  Format  |  Description          |
    |----------|-----------------------|
    |  ipv6    |  IPv6 source address  |
`,
		},

		// TagNodes

		// Nodes

	}
}
