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

var _ helpers.VyosResourceDataModel = &InterfacesVxlanParametersIPvsix{}

// InterfacesVxlanParametersIPvsix describes the resource data model.
type InterfacesVxlanParametersIPvsix struct {
	// LeafNodes
	LeafInterfacesVxlanParametersIPvsixFlowlabel types.String `tfsdk:"flowlabel" vyos:"flowlabel,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesVxlanParametersIPvsix) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"flowlabel":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specifies the flow label to use in outgoing packets

    |  Format        |  Description                      |
    |----------------|-----------------------------------|
    |  inherit       |  Copy field from original header  |
    |  0x0-0x0fffff  |  Tunnel key, or hex value         |
`,
			Description: `Specifies the flow label to use in outgoing packets

    |  Format        |  Description                      |
    |----------------|-----------------------------------|
    |  inherit       |  Copy field from original header  |
    |  0x0-0x0fffff  |  Tunnel key, or hex value         |
`,
		},

		// Nodes

	}
}
