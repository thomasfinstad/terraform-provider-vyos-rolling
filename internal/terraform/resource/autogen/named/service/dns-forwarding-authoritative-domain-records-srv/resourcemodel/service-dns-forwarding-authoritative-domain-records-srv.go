// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ServiceDNSForwardingAuthoritativeDomainRecordsSrv describes the resource data model.
type ServiceDNSForwardingAuthoritativeDomainRecordsSrv struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"srv_id" vyos:",self-id"`

	ParentIDServiceDNSForwardingAuthoritativeDomain types.String `tfsdk:"authoritative_domain" vyos:"authoritative-domain,parent-id"`

	// LeafNodes
	LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvTTL     types.Number `tfsdk:"ttl" vyos:"ttl,omitempty"`
	LeafServiceDNSForwardingAuthoritativeDomainRecordsSrvDisable types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagServiceDNSForwardingAuthoritativeDomainRecordsSrvEntry bool `tfsdk:"entry" vyos:"entry,child"`

	// Nodes
}

// GetID returns the resource ID
func (o ServiceDNSForwardingAuthoritativeDomainRecordsSrv) GetID() *types.String {
	return &o.ID
}

// SetID configures the resource ID
func (o ServiceDNSForwardingAuthoritativeDomainRecordsSrv) SetID(id types.String) {
	o.ID = id
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsSrv) GetVyosPath() []string {
	return []string{
		"service",

		"dns",

		"forwarding",

		"authoritative-domain",
		o.ParentIDServiceDNSForwardingAuthoritativeDomain.ValueString(),

		"records",

		"srv",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDNSForwardingAuthoritativeDomainRecordsSrv) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"srv_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `"SRV" record

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  text  &emsp; |  A DNS name relative to the root record  |
    |  @  &emsp; |  Root record  |

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

		// LeafNodes

		"ttl": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Time-to-live (TTL)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-2147483647  &emsp; |  TTL in seconds  |

`,

			// Default:          stringdefault.StaticString(`300`),
			Computed: true,
		},

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
