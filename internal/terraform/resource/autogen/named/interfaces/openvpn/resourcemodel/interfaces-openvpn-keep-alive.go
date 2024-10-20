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

var _ helpers.VyosResourceDataModel = &InterfacesOpenvpnKeepAlive{}

// InterfacesOpenvpnKeepAlive describes the resource data model.
type InterfacesOpenvpnKeepAlive struct {
	// LeafNodes
	LeafInterfacesOpenvpnKeepAliveFailureCount types.Number `tfsdk:"failure_count" vyos:"failure-count,omitempty"`
	LeafInterfacesOpenvpnKeepAliveInterval     types.Number `tfsdk:"interval" vyos:"interval,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesOpenvpnKeepAlive) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"failure_count":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum number of keepalive packet failures

    |  Format  |  Description                                  |
    |----------|-----------------------------------------------|
    |  0-1000  |  Maximum number of keepalive packet failures  |
`,
			Description: `Maximum number of keepalive packet failures

    |  Format  |  Description                                  |
    |----------|-----------------------------------------------|
    |  0-1000  |  Maximum number of keepalive packet failures  |
`,

			// Default:          stringdefault.StaticString(`60`),
			Computed: true,
		},

		"interval":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Keepalive packet interval in seconds

    |  Format  |  Description                          |
    |----------|---------------------------------------|
    |  0-600   |  Keepalive packet interval (seconds)  |
`,
			Description: `Keepalive packet interval in seconds

    |  Format  |  Description                          |
    |----------|---------------------------------------|
    |  0-600   |  Keepalive packet interval (seconds)  |
`,

			// Default:          stringdefault.StaticString(`10`),
			Computed: true,
		},

		// Nodes

	}
}
