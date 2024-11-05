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

var _ helpers.VyosResourceDataModel = &ServiceDNSDynamicNameAddressWeb{}

// ServiceDNSDynamicNameAddressWeb describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type ServiceDNSDynamicNameAddressWeb struct {
	// LeafNodes
	LeafServiceDNSDynamicNameAddressWebURL  types.String `tfsdk:"url" vyos:"url,omitempty"`
	LeafServiceDNSDynamicNameAddressWebSkIP types.String `tfsdk:"skip" vyos:"skip,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDNSDynamicNameAddressWeb) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"url":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Remote URL

    |  Format  |  Description         |
    |----------|----------------------|
    |  url     |  Remote HTTP(S) URL  |
`,
			Description: `Remote URL

    |  Format  |  Description         |
    |----------|----------------------|
    |  url     |  Remote HTTP(S) URL  |
`,
		},

		"skip":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Pattern to skip from the HTTP(S) respose

    |  Format  |  Description                                                                  |
    |----------|-------------------------------------------------------------------------------|
    |  txt     |  Pattern to skip from the HTTP(S) respose to extract the external IP address  |
`,
			Description: `Pattern to skip from the HTTP(S) respose

    |  Format  |  Description                                                                  |
    |----------|-------------------------------------------------------------------------------|
    |  txt     |  Pattern to skip from the HTTP(S) respose to extract the external IP address  |
`,
		},

		// TagNodes

		// Nodes

	}
}
