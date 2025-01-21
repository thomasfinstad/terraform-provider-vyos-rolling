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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (source) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &NatDestinationRuleSource{}

// NatDestinationRuleSource describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type NatDestinationRuleSource struct {
	// LeafNodes
	LeafNatDestinationRuleSourceFqdn    types.String `tfsdk:"fqdn" vyos:"fqdn,omitempty"`
	LeafNatDestinationRuleSourceAddress types.String `tfsdk:"address" vyos:"address,omitempty"`
	LeafNatDestinationRuleSourcePort    types.String `tfsdk:"port" vyos:"port,omitempty"`

	// TagNodes

	// Nodes

	NodeNatDestinationRuleSourceGroup *NatDestinationRuleSourceGroup `tfsdk:"group" vyos:"group,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o NatDestinationRuleSource) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"fqdn":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (fqdn) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Fully qualified domain name

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  <fqdn>  |  Fully qualified domain name  |
`,
			Description: `Fully qualified domain name

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  <fqdn>  |  Fully qualified domain name  |
`,
		},

		"address":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (address) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address, subnet, or range

    |  Format      |  Description                                    |
    |--------------|-------------------------------------------------|
    |  ipv4        |  IPv4 address to match                          |
    |  ipv4net     |  IPv4 prefix to match                           |
    |  ipv4range   |  IPv4 address range to match                    |
    |  !ipv4       |  Match everything except the specified address  |
    |  !ipv4net    |  Match everything except the specified prefix   |
    |  !ipv4range  |  Match everything except the specified range    |
`,
			Description: `IP address, subnet, or range

    |  Format      |  Description                                    |
    |--------------|-------------------------------------------------|
    |  ipv4        |  IPv4 address to match                          |
    |  ipv4net     |  IPv4 prefix to match                           |
    |  ipv4range   |  IPv4 address range to match                    |
    |  !ipv4       |  Match everything except the specified address  |
    |  !ipv4net    |  Match everything except the specified prefix   |
    |  !ipv4range  |  Match everything except the specified range    |
`,
		},

		"port":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (port) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Port number

    |  Format     |  Description                                                                                                                                                                          |
    |-------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
    |  txt        |  Named port (any name in /etc/services, e.g., http)                                                                                                                                   |
    |  1-65535    |  Numeric IP port                                                                                                                                                                      |
    |  start-end  |  Numbered port range (e.g. 1001-1005)                                                                                                                                                 |
    |             |  </br></br>Multiple destination ports can be specified as a comma-separated list.</br>The whole list can also be negated using '!'.</br>For example: '!22,telnet,http,123,1001-1005'  |
`,
			Description: `Port number

    |  Format     |  Description                                                                                                                                                                          |
    |-------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
    |  txt        |  Named port (any name in /etc/services, e.g., http)                                                                                                                                   |
    |  1-65535    |  Numeric IP port                                                                                                                                                                      |
    |  start-end  |  Numbered port range (e.g. 1001-1005)                                                                                                                                                 |
    |             |  </br></br>Multiple destination ports can be specified as a comma-separated list.</br>The whole list can also be negated using '!'.</br>For example: '!22,telnet,http,123,1001-1005'  |
`,
		},

		// TagNodes

		// Nodes

		"group": schema.SingleNestedAttribute{
			Attributes: NatDestinationRuleSourceGroup{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Group

`,
			Description: `Group

`,
		},
	}
}
