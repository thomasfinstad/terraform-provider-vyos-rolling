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

var _ helpers.VyosResourceDataModel = &ServiceDNSDynamicNameAddress{}

// ServiceDNSDynamicNameAddress describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type ServiceDNSDynamicNameAddress struct {
	// LeafNodes
	LeafServiceDNSDynamicNameAddressInterface types.String `tfsdk:"interface" vyos:"interface,omitempty"`

	// TagNodes

	// Nodes

	NodeServiceDNSDynamicNameAddressWeb *ServiceDNSDynamicNameAddressWeb `tfsdk:"web" vyos:"web,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDNSDynamicNameAddress) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"interface":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interface

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Interface name  |
`,
			Description: `Interface

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Interface name  |
`,
		},

		// TagNodes

		// Nodes

		"web": schema.SingleNestedAttribute{
			Attributes: ServiceDNSDynamicNameAddressWeb{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `HTTP(S) web request to use

`,
			Description: `HTTP(S) web request to use

`,
		},
	}
}
