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

var _ helpers.VyosResourceDataModel = &LoadBalancingHaproxyBackendHTTPCheckExpect{}

// LoadBalancingHaproxyBackendHTTPCheckExpect describes the resource data model.
type LoadBalancingHaproxyBackendHTTPCheckExpect struct {
	// LeafNodes
	LeafLoadBalancingHaproxyBackendHTTPCheckExpectStatus types.Number `tfsdk:"status" vyos:"status,omitempty"`
	LeafLoadBalancingHaproxyBackendHTTPCheckExpectString types.String `tfsdk:"string" vyos:"string,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o LoadBalancingHaproxyBackendHTTPCheckExpect) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"status":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Expected response status code for the health check to pass

    |  Format   |  Description             |
    |-----------|--------------------------|
    |  200-399  |  Expected response code  |
`,
			Description: `Expected response status code for the health check to pass

    |  Format   |  Description             |
    |-----------|--------------------------|
    |  200-399  |  Expected response code  |
`,
		},

		"string":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Expected to be in response body for the health check to pass

    |  Format  |  Description                              |
    |----------|-------------------------------------------|
    |  txt     |  A string expected to be in the response  |
`,
			Description: `Expected to be in response body for the health check to pass

    |  Format  |  Description                              |
    |----------|-------------------------------------------|
    |  txt     |  A string expected to be in the response  |
`,
		},

		// Nodes

	}
}
