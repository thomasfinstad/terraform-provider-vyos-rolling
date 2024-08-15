// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &SystemConntrackTimeoutCustomIPvsixRuleSource{}

// SystemConntrackTimeoutCustomIPvsixRuleSource describes the resource data model.
type SystemConntrackTimeoutCustomIPvsixRuleSource struct {
	// LeafNodes
	LeafSystemConntrackTimeoutCustomIPvsixRuleSourceAddress types.String `tfsdk:"address" vyos:"address,omitempty"`
	LeafSystemConntrackTimeoutCustomIPvsixRuleSourcePort    types.String `tfsdk:"port" vyos:"port,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemConntrackTimeoutCustomIPvsixRuleSource) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address, subnet, or range

    |  Format      |  Description                                    |
    |--------------|-------------------------------------------------|
    |  ipv6        |  IP address to match                            |
    |  ipv6net     |  Subnet to match                                |
    |  ipv6range   |  IP range to match                              |
    |  !ipv6       |  Match everything except the specified address  |
    |  !ipv6net    |  Match everything except the specified prefix   |
    |  !ipv6range  |  Match everything except the specified range    |
`,
			Description: `IP address, subnet, or range

    |  Format      |  Description                                    |
    |--------------|-------------------------------------------------|
    |  ipv6        |  IP address to match                            |
    |  ipv6net     |  Subnet to match                                |
    |  ipv6range   |  IP range to match                              |
    |  !ipv6       |  Match everything except the specified address  |
    |  !ipv6net    |  Match everything except the specified prefix   |
    |  !ipv6range  |  Match everything except the specified range    |
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

	}
}
