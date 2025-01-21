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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (send) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &ProtocolsRIPInterfaceSend{}

// ProtocolsRIPInterfaceSend describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type ProtocolsRIPInterfaceSend struct {
	// LeafNodes
	LeafProtocolsRIPInterfaceSendVersion types.String `tfsdk:"version" vyos:"version,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsRIPInterfaceSend) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"version":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (version) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Limit RIP protocol version

    |  Format  |  Description       |
    |----------|--------------------|
    |  1       |  Allow RIPv1 only  |
    |  2       |  Allow RIPv2 only  |
`,
			Description: `Limit RIP protocol version

    |  Format  |  Description       |
    |----------|--------------------|
    |  1       |  Allow RIPv1 only  |
    |  2       |  Allow RIPv2 only  |
`,
		},

		// TagNodes

		// Nodes

	}
}
