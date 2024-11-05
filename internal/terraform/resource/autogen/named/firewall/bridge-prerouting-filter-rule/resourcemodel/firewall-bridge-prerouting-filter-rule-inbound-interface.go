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

var _ helpers.VyosResourceDataModel = &FirewallBrIDgePreroutingFilterRuleInboundInterface{}

// FirewallBrIDgePreroutingFilterRuleInboundInterface describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type FirewallBrIDgePreroutingFilterRuleInboundInterface struct {
	// LeafNodes
	LeafFirewallBrIDgePreroutingFilterRuleInboundInterfaceName  types.String `tfsdk:"name" vyos:"name,omitempty"`
	LeafFirewallBrIDgePreroutingFilterRuleInboundInterfaceGroup types.String `tfsdk:"group" vyos:"group,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallBrIDgePreroutingFilterRuleInboundInterface) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"name":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match interface

    |  Format  |  Description                       |
    |----------|------------------------------------|
    |  txt     |  Interface name                    |
    |  txt&    |  Interface name with wildcard      |
    |  !txt    |  Inverted interface name to match  |
`,
			Description: `Match interface

    |  Format  |  Description                       |
    |----------|------------------------------------|
    |  txt     |  Interface name                    |
    |  txt&    |  Interface name with wildcard      |
    |  !txt    |  Inverted interface name to match  |
`,
		},

		"group":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match interface-group

    |  Format  |  Description                             |
    |----------|------------------------------------------|
    |  txt     |  Interface-group name to match           |
    |  !txt    |  Inverted interface-group name to match  |
`,
			Description: `Match interface-group

    |  Format  |  Description                             |
    |----------|------------------------------------------|
    |  txt     |  Interface-group name to match           |
    |  !txt    |  Inverted interface-group name to match  |
`,
		},

		// TagNodes

		// Nodes

	}
}
