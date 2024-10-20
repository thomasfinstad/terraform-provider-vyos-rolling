/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &InterfacesGeneveParameters{}

// InterfacesGeneveParameters describes the resource data model.
type InterfacesGeneveParameters struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeInterfacesGeneveParametersIP     *InterfacesGeneveParametersIP     `tfsdk:"ip" vyos:"ip,omitempty"`
	NodeInterfacesGeneveParametersIPvsix *InterfacesGeneveParametersIPvsix `tfsdk:"ipv6" vyos:"ipv6,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesGeneveParameters) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesGeneveParametersIP{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv4 specific tunnel parameters

`,
			Description: `IPv4 specific tunnel parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesGeneveParametersIPvsix{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv6 specific tunnel parameters

`,
			Description: `IPv6 specific tunnel parameters

`,
		},
	}
}
