/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &ServiceDNSForwardingZoneCacheOptionsRefresh{}

// ServiceDNSForwardingZoneCacheOptionsRefresh describes the resource data model.
type ServiceDNSForwardingZoneCacheOptionsRefresh struct {
	// LeafNodes
	LeafServiceDNSForwardingZoneCacheOptionsRefreshOnReload types.Bool   `tfsdk:"on_reload" vyos:"on-reload,omitempty"`
	LeafServiceDNSForwardingZoneCacheOptionsRefreshInterval types.Number `tfsdk:"interval" vyos:"interval,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDNSForwardingZoneCacheOptionsRefresh) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"on_reload":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Retrieval zone only at startup and on reload

`,
			Description: `Retrieval zone only at startup and on reload

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"interval":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Periodic zone retrieval interval

    |  Format      |  Description                    |
    |--------------|---------------------------------|
    |  0-31536000  |  Retrieval interval in seconds  |
`,
			Description: `Periodic zone retrieval interval

    |  Format      |  Description                    |
    |--------------|---------------------------------|
    |  0-31536000  |  Retrieval interval in seconds  |
`,

			// Default:          stringdefault.StaticString(`86400`),
			Computed: true,
		},

		// Nodes

	}
}
