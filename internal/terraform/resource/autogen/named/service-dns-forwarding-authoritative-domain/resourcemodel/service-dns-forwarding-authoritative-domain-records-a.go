// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ServiceDNSForwardingAuthoritativeDomainRecordsA describes the resource data model.
type ServiceDNSForwardingAuthoritativeDomainRecordsA struct {
	// LeafNodes
	ServiceDNSForwardingAuthoritativeDomainRecordsAAddress customtypes.CustomStringValue `tfsdk:"address" json:"address,omitempty"`
	ServiceDNSForwardingAuthoritativeDomainRecordsATTL     customtypes.CustomStringValue `tfsdk:"ttl" json:"ttl,omitempty"`
	ServiceDNSForwardingAuthoritativeDomainRecordsADisable customtypes.CustomStringValue `tfsdk:"disable" json:"disable,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ServiceDNSForwardingAuthoritativeDomainRecordsA) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `IPv4 address

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  IPv4 address  |

`,
		},

		"ttl": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Time-to-live (TTL)

|  Format  |  Description  |
|----------|---------------|
|  u32:0-2147483647  |  TTL in seconds  |

`,

			// Default:          stringdefault.StaticString(`300`),
			Computed: true,
		},

		"disable": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Disable instance

`,
		},

		// TagNodes

		// Nodes

	}
}
