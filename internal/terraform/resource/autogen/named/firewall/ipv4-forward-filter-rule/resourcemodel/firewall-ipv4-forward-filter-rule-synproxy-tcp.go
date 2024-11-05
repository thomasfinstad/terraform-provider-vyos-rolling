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

var _ helpers.VyosResourceDataModel = &FirewallIPvfourForwardFilterRuleSynproxyTCP{}

// FirewallIPvfourForwardFilterRuleSynproxyTCP describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type FirewallIPvfourForwardFilterRuleSynproxyTCP struct {
	// LeafNodes
	LeafFirewallIPvfourForwardFilterRuleSynproxyTCPMss         types.Number `tfsdk:"mss" vyos:"mss,omitempty"`
	LeafFirewallIPvfourForwardFilterRuleSynproxyTCPWindowScale types.Number `tfsdk:"window_scale" vyos:"window-scale,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourForwardFilterRuleSynproxyTCP) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"mss":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TCP Maximum segment size

    |  Format     |  Description                                    |
    |-------------|-------------------------------------------------|
    |  501-65535  |  Maximum segment size for synproxy connections  |
`,
			Description: `TCP Maximum segment size

    |  Format     |  Description                                    |
    |-------------|-------------------------------------------------|
    |  501-65535  |  Maximum segment size for synproxy connections  |
`,
		},

		"window_scale":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TCP window scale for synproxy connections

    |  Format  |  Description       |
    |----------|--------------------|
    |  1-14    |  TCP window scale  |
`,
			Description: `TCP window scale for synproxy connections

    |  Format  |  Description       |
    |----------|--------------------|
    |  1-14    |  TCP window scale  |
`,
		},

		// TagNodes

		// Nodes

	}
}
