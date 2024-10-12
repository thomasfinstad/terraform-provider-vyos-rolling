// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvfourInputFilterRuleConnectionStatus{}

// FirewallIPvfourInputFilterRuleConnectionStatus describes the resource data model.
type FirewallIPvfourInputFilterRuleConnectionStatus struct {
	// LeafNodes
	LeafFirewallIPvfourInputFilterRuleConnectionStatusNat types.String `tfsdk:"nat" vyos:"nat,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourInputFilterRuleConnectionStatus) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"nat": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `NAT connection status

    |  Format       |  Description                                            |
    |---------------|---------------------------------------------------------|
    |  destination  |  Match connections that are subject to destination NAT  |
    |  source       |  Match connections that are subject to source NAT       |
`,
			Description: `NAT connection status

    |  Format       |  Description                                            |
    |---------------|---------------------------------------------------------|
    |  destination  |  Match connections that are subject to destination NAT  |
    |  source       |  Match connections that are subject to source NAT       |
`,
		},

		// Nodes

	}
}
