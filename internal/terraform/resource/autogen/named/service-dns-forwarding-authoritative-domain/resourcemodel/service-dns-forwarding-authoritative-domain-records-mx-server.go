// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ServiceDNSForwardingAuthoritativeDomainRecordsMxServer describes the resource data model.
type ServiceDNSForwardingAuthoritativeDomainRecordsMxServer struct {
	// LeafNodes
	ServiceDNSForwardingAuthoritativeDomainRecordsMxServerPriority customtypes.CustomStringValue `tfsdk:"priority" json:"priority,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ServiceDNSForwardingAuthoritativeDomainRecordsMxServer) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"priority": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Server priority

|  Format  |  Description  |
|----------|---------------|
|  u32:1-999  |  Server priority (lower numbers are higher priority)  |

`,

			// Default:          stringdefault.StaticString(`10`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
