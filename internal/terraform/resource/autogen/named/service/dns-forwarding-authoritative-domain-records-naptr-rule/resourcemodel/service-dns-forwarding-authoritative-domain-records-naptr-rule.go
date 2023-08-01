// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ServiceDNSForwardingAuthoritativeDomainRecordsNaptrRule describes the resource data model.
type ServiceDNSForwardingAuthoritativeDomainRecordsNaptrRule struct {
	SelfIdentifier types.Number `tfsdk:"rule_id" vyos:",self-id"`

	ParentIDServiceDNSForwardingAuthoritativeDomain types.String `tfsdk:"authoritative_domain" vyos:"authoritative-domain,parent-id"`

	ParentIDServiceDNSForwardingAuthoritativeDomainRecordsNaptr types.String `tfsdk:"naptr" vyos:"naptr,parent-id"`

	// LeafNodes
	LeafServiceDNSForwardingAuthoritativeDomainRecordsNaptrRuleOrder            types.Number `tfsdk:"order" vyos:"order,omitempty"`
	LeafServiceDNSForwardingAuthoritativeDomainRecordsNaptrRulePreference       types.Number `tfsdk:"preference" vyos:"preference,omitempty"`
	LeafServiceDNSForwardingAuthoritativeDomainRecordsNaptrRuleLookupSrv        types.Bool   `tfsdk:"lookup_srv" vyos:"lookup-srv,omitempty"`
	LeafServiceDNSForwardingAuthoritativeDomainRecordsNaptrRuleLookupA          types.Bool   `tfsdk:"lookup_a" vyos:"lookup-a,omitempty"`
	LeafServiceDNSForwardingAuthoritativeDomainRecordsNaptrRuleResolveURI       types.Bool   `tfsdk:"resolve_uri" vyos:"resolve-uri,omitempty"`
	LeafServiceDNSForwardingAuthoritativeDomainRecordsNaptrRuleProtocolSpecific types.Bool   `tfsdk:"protocol_specific" vyos:"protocol-specific,omitempty"`
	LeafServiceDNSForwardingAuthoritativeDomainRecordsNaptrRuleService          types.String `tfsdk:"service" vyos:"service,omitempty"`
	LeafServiceDNSForwardingAuthoritativeDomainRecordsNaptrRuleRegexp           types.String `tfsdk:"regexp" vyos:"regexp,omitempty"`
	LeafServiceDNSForwardingAuthoritativeDomainRecordsNaptrRuleReplacement      types.String `tfsdk:"replacement" vyos:"replacement,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsNaptrRule) GetVyosPath() []string {
	return []string{
		"service",

		"dns",

		"forwarding",

		"authoritative-domain",
		o.ParentIDServiceDNSForwardingAuthoritativeDomain.ValueString(),

		"records",

		"naptr",
		o.ParentIDServiceDNSForwardingAuthoritativeDomainRecordsNaptr.ValueString(),

		"rule",
		o.SelfIdentifier.ValueBigFloat().String(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDNSForwardingAuthoritativeDomainRecordsNaptrRule) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"rule_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `NAPTR rule

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-65535  &emsp; |  Rule number  |

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

		"naptr_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `"NAPTR" record

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  text  &emsp; |  A DNS name relative to the root record  |
    |  @  &emsp; |  Root record  |

`,
		},

		// LeafNodes

		"order": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Rule order

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-65535  &emsp; |  Rule order (lower order is evaluated first)  |

`,
		},

		"preference": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Rule preference

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-65535  &emsp; |  Rule preference  |

`,

			// Default:          stringdefault.StaticString(`0`),
			Computed: true,
		},

		"lookup_srv": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `"S" flag

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"lookup_a": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `"A" flag

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"resolve_uri": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `"U" flag

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"protocol_specific": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `"P" flag

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"service": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Service type

`,
		},

		"regexp": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Regular expression

`,
		},

		"replacement": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Replacement DNS name

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  name.example.com  &emsp; |  An absolute DNS name  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsNaptrRule) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsNaptrRule) UnmarshalJSON(_ []byte) error {
	return nil
}