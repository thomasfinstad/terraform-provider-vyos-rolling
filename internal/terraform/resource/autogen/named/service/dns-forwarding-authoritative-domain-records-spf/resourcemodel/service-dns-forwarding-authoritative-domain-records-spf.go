// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ServiceDNSForwardingAuthoritativeDomainRecordsSpf describes the resource data model.
type ServiceDNSForwardingAuthoritativeDomainRecordsSpf struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

	ParentIDServiceDNSForwardingAuthoritativeDomain types.String `tfsdk:"authoritative_domain" vyos:"authoritative-domain_identifier,parent-id"`

	// LeafNodes
	LeafServiceDNSForwardingAuthoritativeDomainRecordsSpfValue   types.String `tfsdk:"value" vyos:"value,omitempty"`
	LeafServiceDNSForwardingAuthoritativeDomainRecordsSpfTTL     types.Number `tfsdk:"ttl" vyos:"ttl,omitempty"`
	LeafServiceDNSForwardingAuthoritativeDomainRecordsSpfDisable types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsSpf) GetVyosPath() []string {
	return []string{
		"service",

		"dns",

		"forwarding",

		"authoritative-domain",
		o.ParentIDServiceDNSForwardingAuthoritativeDomain.ValueString(),

		"records",

		"spf",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDNSForwardingAuthoritativeDomainRecordsSpf) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `"SPF" record (type=SPF)

    |  Format  |  Description  |
    |----------|---------------|
    |  text  |  A DNS name relative to the root record  |
    |  @  |  Root record  |

`,
		},

		"authoritative_domain_identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Domain to host authoritative records for

    |  Format  |  Description  |
    |----------|---------------|
    |  text  |  An absolute DNS name  |

`,
		},

		// LeafNodes

		"value": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Record contents

    |  Format  |  Description  |
    |----------|---------------|
    |  text  |  Record contents  |

`,
		},

		"ttl": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Time-to-live (TTL)

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:0-2147483647  |  TTL in seconds  |

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

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsSpf) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsSpf) UnmarshalJSON(_ []byte) error {
	return nil
}
