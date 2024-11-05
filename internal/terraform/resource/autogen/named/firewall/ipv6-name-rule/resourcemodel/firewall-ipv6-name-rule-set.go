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

var _ helpers.VyosResourceDataModel = &FirewallIPvsixNameRuleSet{}

// FirewallIPvsixNameRuleSet describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type FirewallIPvsixNameRuleSet struct {
	// LeafNodes
	LeafFirewallIPvsixNameRuleSetDscp           types.Number `tfsdk:"dscp" vyos:"dscp,omitempty"`
	LeafFirewallIPvsixNameRuleSetConnectionMark types.Number `tfsdk:"connection_mark" vyos:"connection-mark,omitempty"`
	LeafFirewallIPvsixNameRuleSetMark           types.Number `tfsdk:"mark" vyos:"mark,omitempty"`
	LeafFirewallIPvsixNameRuleSetTCPMss         types.Number `tfsdk:"tcp_mss" vyos:"tcp-mss,omitempty"`
	LeafFirewallIPvsixNameRuleSetHopLimit       types.Number `tfsdk:"hop_limit" vyos:"hop-limit,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixNameRuleSet) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"dscp":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set DSCP (Packet Differentiated Services Codepoint) bits

    |  Format  |  Description  |
    |----------|---------------|
    |  0-63    |  DSCP number  |
`,
			Description: `Set DSCP (Packet Differentiated Services Codepoint) bits

    |  Format  |  Description  |
    |----------|---------------|
    |  0-63    |  DSCP number  |
`,
		},

		"connection_mark":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set connection mark

    |  Format        |  Description      |
    |----------------|-------------------|
    |  0-2147483647  |  Connection mark  |
`,
			Description: `Set connection mark

    |  Format        |  Description      |
    |----------------|-------------------|
    |  0-2147483647  |  Connection mark  |
`,
		},

		"mark":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set packet mark

    |  Format        |  Description  |
    |----------------|---------------|
    |  1-2147483647  |  Packet mark  |
`,
			Description: `Set packet mark

    |  Format        |  Description  |
    |----------------|---------------|
    |  1-2147483647  |  Packet mark  |
`,
		},

		"tcp_mss":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set TCP Maximum Segment Size

    |  Format    |  Description                   |
    |------------|--------------------------------|
    |  500-1460  |  Explicitly set TCP MSS value  |
`,
			Description: `Set TCP Maximum Segment Size

    |  Format    |  Description                   |
    |------------|--------------------------------|
    |  500-1460  |  Explicitly set TCP MSS value  |
`,
		},

		"hop_limit":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set hop limit

    |  Format  |  Description       |
    |----------|--------------------|
    |  0-255   |  Hop limit number  |
`,
			Description: `Set hop limit

    |  Format  |  Description       |
    |----------|--------------------|
    |  0-255   |  Hop limit number  |
`,
		},

		// TagNodes

		// Nodes

	}
}
