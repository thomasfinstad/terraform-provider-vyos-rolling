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

var _ helpers.VyosResourceDataModel = &QosPolicyLimiterClassMatchIPDestination{}

// QosPolicyLimiterClassMatchIPDestination describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type QosPolicyLimiterClassMatchIPDestination struct {
	// LeafNodes
	LeafQosPolicyLimiterClassMatchIPDestinationAddress types.String `tfsdk:"address" vyos:"address,omitempty"`
	LeafQosPolicyLimiterClassMatchIPDestinationPort    types.Number `tfsdk:"port" vyos:"port,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyLimiterClassMatchIPDestination) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv4 destination address for this match

    |  Format   |  Description   |
    |-----------|----------------|
    |  ipv4     |  IPv4 address  |
    |  ipv4net  |  IPv4 prefix   |
`,
			Description: `IPv4 destination address for this match

    |  Format   |  Description   |
    |-----------|----------------|
    |  ipv4     |  IPv4 address  |
    |  ipv4net  |  IPv4 prefix   |
`,
		},

		"port":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Port number used by connection

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Numeric IP port  |
`,
			Description: `Port number used by connection

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Numeric IP port  |
`,
		},

		// TagNodes

		// Nodes

	}
}
