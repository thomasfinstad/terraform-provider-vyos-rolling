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

var _ helpers.VyosResourceDataModel = &InterfacesEthernetVifMirror{}

// InterfacesEthernetVifMirror describes the resource data model.
type InterfacesEthernetVifMirror struct {
	// LeafNodes
	LeafInterfacesEthernetVifMirrorIngress types.String `tfsdk:"ingress" vyos:"ingress,omitempty"`
	LeafInterfacesEthernetVifMirrorEgress  types.String `tfsdk:"egress" vyos:"egress,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesEthernetVifMirror) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"ingress":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Mirror ingress traffic to destination interface

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  txt     |  Destination interface name  |
`,
			Description: `Mirror ingress traffic to destination interface

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  txt     |  Destination interface name  |
`,
		},

		"egress":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Mirror egress traffic to destination interface

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  txt     |  Destination interface name  |
`,
			Description: `Mirror egress traffic to destination interface

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  txt     |  Destination interface name  |
`,
		},

		// Nodes

	}
}
