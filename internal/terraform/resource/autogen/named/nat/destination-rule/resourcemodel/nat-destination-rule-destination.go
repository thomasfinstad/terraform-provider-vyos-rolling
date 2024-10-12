// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &NatDestinationRuleDestination{}

// NatDestinationRuleDestination describes the resource data model.
type NatDestinationRuleDestination struct {
	// LeafNodes
	LeafNatDestinationRuleDestinationFqdn    types.String `tfsdk:"fqdn" vyos:"fqdn,omitempty"`
	LeafNatDestinationRuleDestinationAddress types.String `tfsdk:"address" vyos:"address,omitempty"`
	LeafNatDestinationRuleDestinationPort    types.String `tfsdk:"port" vyos:"port,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeNatDestinationRuleDestinationGroup *NatDestinationRuleDestinationGroup `tfsdk:"group" vyos:"group,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o NatDestinationRuleDestination) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"fqdn": schema.StringAttribute{
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

		"address": schema.StringAttribute{
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

		"port": schema.StringAttribute{
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

		// Nodes

		"group": schema.SingleNestedAttribute{
			Attributes: NatDestinationRuleDestinationGroup{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Group

`,
			Description: `Group

`,
		},
	}
}
