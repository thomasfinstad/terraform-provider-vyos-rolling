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

var _ helpers.VyosResourceDataModel = &LoadBalancingHaproxyBackendTimeout{}

// LoadBalancingHaproxyBackendTimeout describes the resource data model.
type LoadBalancingHaproxyBackendTimeout struct {
	// LeafNodes
	LeafLoadBalancingHaproxyBackendTimeoutCheck   types.Number `tfsdk:"check" vyos:"check,omitempty"`
	LeafLoadBalancingHaproxyBackendTimeoutConnect types.Number `tfsdk:"connect" vyos:"connect,omitempty"`
	LeafLoadBalancingHaproxyBackendTimeoutServer  types.Number `tfsdk:"server" vyos:"server,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o LoadBalancingHaproxyBackendTimeout) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"check":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Timeout in seconds for established connections

    |  Format  |  Description               |
    |----------|----------------------------|
    |  1-3600  |  Check timeout in seconds  |
`,
			Description: `Timeout in seconds for established connections

    |  Format  |  Description               |
    |----------|----------------------------|
    |  1-3600  |  Check timeout in seconds  |
`,
		},

		"connect":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set the maximum time to wait for a connection attempt to a server to succeed

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  1-3600  |  Connect timeout in seconds  |
`,
			Description: `Set the maximum time to wait for a connection attempt to a server to succeed

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  1-3600  |  Connect timeout in seconds  |
`,
		},

		"server":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set the maximum inactivity time on the server side

    |  Format  |  Description                |
    |----------|-----------------------------|
    |  1-3600  |  Server timeout in seconds  |
`,
			Description: `Set the maximum inactivity time on the server side

    |  Format  |  Description                |
    |----------|-----------------------------|
    |  1-3600  |  Server timeout in seconds  |
`,
		},

		// Nodes

	}
}
