/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &PolicyRouteMapRuleSetEvpn{}

// PolicyRouteMapRuleSetEvpn describes the resource data model.
type PolicyRouteMapRuleSetEvpn struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodePolicyRouteMapRuleSetEvpnGateway *PolicyRouteMapRuleSetEvpnGateway `tfsdk:"gateway" vyos:"gateway,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRuleSetEvpn) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"gateway": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleSetEvpnGateway{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Set gateway IP for prefix advertisement route

`,
			Description: `Set gateway IP for prefix advertisement route

`,
		},
	}
}
