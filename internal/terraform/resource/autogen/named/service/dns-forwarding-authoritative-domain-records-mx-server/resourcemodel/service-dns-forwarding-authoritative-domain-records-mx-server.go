// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ServiceDNSForwardingAuthoritativeDomainRecordsMxServer describes the resource data model.
type ServiceDNSForwardingAuthoritativeDomainRecordsMxServer struct {
	SelfIdentifier types.String `tfsdk:"server_id" vyos:",self-id"`

	ParentIDServiceDNSForwardingAuthoritativeDomain types.String `tfsdk:"authoritative_domain" vyos:"authoritative-domain,parent-id"`

	ParentIDServiceDNSForwardingAuthoritativeDomainRecordsMx types.String `tfsdk:"mx" vyos:"mx,parent-id"`

	// LeafNodes
	LeafServiceDNSForwardingAuthoritativeDomainRecordsMxServerPriority types.Number `tfsdk:"priority" vyos:"priority,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsMxServer) GetVyosPath() []string {
	return []string{
		"service",

		"dns",

		"forwarding",

		"authoritative-domain",
		o.ParentIDServiceDNSForwardingAuthoritativeDomain.ValueString(),

		"records",

		"mx",
		o.ParentIDServiceDNSForwardingAuthoritativeDomainRecordsMx.ValueString(),

		"server",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDNSForwardingAuthoritativeDomainRecordsMxServer) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, an amalgamation of the `server_id` and the parents `*_id` fields seperated by dunder `__` starting with top level ancestor.",
		},
		"server_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Mail server

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  name.example.com  &emsp; |  An absolute DNS name  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		"authoritative_domain_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Domain to host authoritative records for

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  text  &emsp; |  An absolute DNS name  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		"mx_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `"MX" record

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  text  &emsp; |  A DNS name relative to the root record  |
    |  @  &emsp; |  Root record  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"priority": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Server priority

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-999  &emsp; |  Server priority (lower numbers are higher priority)  |

`,

			// Default:          stringdefault.StaticString(`10`),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsMxServer) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsMxServer) UnmarshalJSON(_ []byte) error {
	return nil
}
