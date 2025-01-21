/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (set) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvfourOutputFilterRuleSet{}

// FirewallIPvfourOutputFilterRuleSet describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type FirewallIPvfourOutputFilterRuleSet struct {
	// LeafNodes
	LeafFirewallIPvfourOutputFilterRuleSetDscp           types.Number `tfsdk:"dscp" vyos:"dscp,omitempty"`
	LeafFirewallIPvfourOutputFilterRuleSetConnectionMark types.Number `tfsdk:"connection_mark" vyos:"connection-mark,omitempty"`
	LeafFirewallIPvfourOutputFilterRuleSetMark           types.Number `tfsdk:"mark" vyos:"mark,omitempty"`
	LeafFirewallIPvfourOutputFilterRuleSetTCPMss         types.Number `tfsdk:"tcp_mss" vyos:"tcp-mss,omitempty"`
	LeafFirewallIPvfourOutputFilterRuleSetTTL            types.Number `tfsdk:"ttl" vyos:"ttl,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourOutputFilterRuleSet) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"dscp":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (dscp) */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (connection-mark) */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (mark) */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (tcp-mss) */
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

		"ttl":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (ttl) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set TTL (time to live)

    |  Format  |  Description  |
    |----------|---------------|
    |  0-255   |  TTL number   |
`,
			Description: `Set TTL (time to live)

    |  Format  |  Description  |
    |----------|---------------|
    |  0-255   |  TTL number   |
`,
		},

		// TagNodes

		// Nodes

	}
}
