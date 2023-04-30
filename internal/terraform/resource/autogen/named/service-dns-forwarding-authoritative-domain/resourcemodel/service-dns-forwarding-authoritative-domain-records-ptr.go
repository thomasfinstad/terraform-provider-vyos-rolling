// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ServiceDNSForwardingAuthoritativeDomainRecordsPtr describes the resource data model.
type ServiceDNSForwardingAuthoritativeDomainRecordsPtr struct {
	// LeafNodes
	ServiceDNSForwardingAuthoritativeDomainRecordsPtrTarget  customtypes.CustomStringValue `tfsdk:"target" json:"target,omitempty"`
	ServiceDNSForwardingAuthoritativeDomainRecordsPtrTTL     customtypes.CustomStringValue `tfsdk:"ttl" json:"ttl,omitempty"`
	ServiceDNSForwardingAuthoritativeDomainRecordsPtrDisable customtypes.CustomStringValue `tfsdk:"disable" json:"disable,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ServiceDNSForwardingAuthoritativeDomainRecordsPtr) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"target": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Target DNS name

|  Format  |  Description  |
|----------|---------------|
|  name.example.com  |  An absolute DNS name  |

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
