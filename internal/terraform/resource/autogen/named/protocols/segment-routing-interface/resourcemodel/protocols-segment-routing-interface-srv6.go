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

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &ProtocolsSegmentRoutingInterfaceSrvsix{}

// ProtocolsSegmentRoutingInterfaceSrvsix describes the resource data model.
type ProtocolsSegmentRoutingInterfaceSrvsix struct {
	// LeafNodes
	LeafProtocolsSegmentRoutingInterfaceSrvsixHmac types.String `tfsdk:"hmac" vyos:"hmac,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsSegmentRoutingInterfaceSrvsix) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"hmac":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Define HMAC policy for ingress SR-enabled packets on this interface

    |  Format  |  Description                                              |
    |----------|-----------------------------------------------------------|
    |  accept  |  Accept packets without HMAC, validate packets with HMAC  |
    |  drop    |  Drop packets without HMAC, validate packets with HMAC    |
    |  ignore  |  Ignore HMAC field.                                       |
`,
			Description: `Define HMAC policy for ingress SR-enabled packets on this interface

    |  Format  |  Description                                              |
    |----------|-----------------------------------------------------------|
    |  accept  |  Accept packets without HMAC, validate packets with HMAC  |
    |  drop    |  Drop packets without HMAC, validate packets with HMAC    |
    |  ignore  |  Ignore HMAC field.                                       |
`,

			// Default:          stringdefault.StaticString(`accept`),
			Computed: true,
		},

		// Nodes

	}
}
