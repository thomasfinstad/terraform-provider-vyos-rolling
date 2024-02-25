// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// FirewallIPvsixOutputFilterRuleDestinationGeoIP describes the resource data model.
type FirewallIPvsixOutputFilterRuleDestinationGeoIP struct {
	// LeafNodes
	LeafFirewallIPvsixOutputFilterRuleDestinationGeoIPCountryCode  types.List `tfsdk:"country_code" vyos:"country-code,omitempty"`
	LeafFirewallIPvsixOutputFilterRuleDestinationGeoIPInverseMatch types.Bool `tfsdk:"inverse_match" vyos:"inverse-match,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixOutputFilterRuleDestinationGeoIP) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"country_code": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `GeoIP country code

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  <country>  &emsp; |  Country code (2 characters)  |

`,
		},

		"inverse_match": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Inverse match of country-codes

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
