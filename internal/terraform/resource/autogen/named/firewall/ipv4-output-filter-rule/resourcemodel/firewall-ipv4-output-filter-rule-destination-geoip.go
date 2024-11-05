/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvfourOutputFilterRuleDestinationGeoIP{}

// FirewallIPvfourOutputFilterRuleDestinationGeoIP describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type FirewallIPvfourOutputFilterRuleDestinationGeoIP struct {
	// LeafNodes
	LeafFirewallIPvfourOutputFilterRuleDestinationGeoIPCountryCode  types.List `tfsdk:"country_code" vyos:"country-code,omitempty"`
	LeafFirewallIPvfourOutputFilterRuleDestinationGeoIPInverseMatch types.Bool `tfsdk:"inverse_match" vyos:"inverse-match,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvfourOutputFilterRuleDestinationGeoIP) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"country_code":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `GeoIP country code

    |  Format     |  Description                  |
    |-------------|-------------------------------|
    |  <country>  |  Country code (2 characters)  |
`,
			Description: `GeoIP country code

    |  Format     |  Description                  |
    |-------------|-------------------------------|
    |  <country>  |  Country code (2 characters)  |
`,
		},

		"inverse_match":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Inverse match of country-codes

`,
			Description: `Inverse match of country-codes

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
