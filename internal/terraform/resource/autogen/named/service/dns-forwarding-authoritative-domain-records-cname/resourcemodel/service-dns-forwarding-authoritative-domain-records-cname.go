// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ServiceDNSForwardingAuthoritativeDomainRecordsCname describes the resource data model.
type ServiceDNSForwardingAuthoritativeDomainRecordsCname struct {
	SelfIdentifier types.String `tfsdk:"cname_id" vyos:",self-id"`

	ParentIDServiceDNSForwardingAuthoritativeDomain types.String `tfsdk:"authoritative_domain" vyos:"authoritative-domain,parent-id"`

	// LeafNodes
	LeafServiceDNSForwardingAuthoritativeDomainRecordsCnameTarget  types.String `tfsdk:"target" vyos:"target,omitempty"`
	LeafServiceDNSForwardingAuthoritativeDomainRecordsCnameTTL     types.Number `tfsdk:"ttl" vyos:"ttl,omitempty"`
	LeafServiceDNSForwardingAuthoritativeDomainRecordsCnameDisable types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsCname) GetVyosPath() []string {
	return []string{
		"service",

		"dns",

		"forwarding",

		"authoritative-domain",
		o.ParentIDServiceDNSForwardingAuthoritativeDomain.ValueString(),

		"records",

		"cname",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDNSForwardingAuthoritativeDomainRecordsCname) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"cname_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `"CNAME" record

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  text  &emsp; |  A DNS name relative to the root record  |
    |  @  &emsp; |  Root record  |

`,
		},

		"authoritative_domain_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Domain to host authoritative records for

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  text  &emsp; |  An absolute DNS name  |

`,
		},

		// LeafNodes

		"target": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Target DNS name

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  name.example.com  &emsp; |  An absolute DNS name  |

`,
		},

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

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsCname) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsCname) UnmarshalJSON(_ []byte) error {
	return nil
}
