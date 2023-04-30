// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ServiceDNSForwardingAuthoritativeDomainRecords describes the resource data model.
type ServiceDNSForwardingAuthoritativeDomainRecords struct {
	// LeafNodes

	// TagNodes
	ServiceDNSForwardingAuthoritativeDomainRecordsA     types.Map `tfsdk:"a" json:"a,omitempty"`
	ServiceDNSForwardingAuthoritativeDomainRecordsAaaa  types.Map `tfsdk:"aaaa" json:"aaaa,omitempty"`
	ServiceDNSForwardingAuthoritativeDomainRecordsCname types.Map `tfsdk:"cname" json:"cname,omitempty"`
	ServiceDNSForwardingAuthoritativeDomainRecordsMx    types.Map `tfsdk:"mx" json:"mx,omitempty"`
	ServiceDNSForwardingAuthoritativeDomainRecordsPtr   types.Map `tfsdk:"ptr" json:"ptr,omitempty"`
	ServiceDNSForwardingAuthoritativeDomainRecordsTxt   types.Map `tfsdk:"txt" json:"txt,omitempty"`
	ServiceDNSForwardingAuthoritativeDomainRecordsSpf   types.Map `tfsdk:"spf" json:"spf,omitempty"`
	ServiceDNSForwardingAuthoritativeDomainRecordsSrv   types.Map `tfsdk:"srv" json:"srv,omitempty"`
	ServiceDNSForwardingAuthoritativeDomainRecordsNaptr types.Map `tfsdk:"naptr" json:"naptr,omitempty"`

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ServiceDNSForwardingAuthoritativeDomainRecords) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		"a": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ServiceDNSForwardingAuthoritativeDomainRecordsA{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `"A" record

|  Format  |  Description  |
|----------|---------------|
|  text  |  A DNS name relative to the root record  |
|  @  |  Root record  |
|  any  |  Wildcard record (any subdomain)  |

`,
		},

		"aaaa": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ServiceDNSForwardingAuthoritativeDomainRecordsAaaa{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `"AAAA" record

|  Format  |  Description  |
|----------|---------------|
|  text  |  A DNS name relative to the root record  |
|  @  |  Root record  |
|  any  |  Wildcard record (any subdomain)  |

`,
		},

		"cname": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ServiceDNSForwardingAuthoritativeDomainRecordsCname{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `"CNAME" record

|  Format  |  Description  |
|----------|---------------|
|  text  |  A DNS name relative to the root record  |
|  @  |  Root record  |

`,
		},

		"mx": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ServiceDNSForwardingAuthoritativeDomainRecordsMx{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `"MX" record

|  Format  |  Description  |
|----------|---------------|
|  text  |  A DNS name relative to the root record  |
|  @  |  Root record  |

`,
		},

		"ptr": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ServiceDNSForwardingAuthoritativeDomainRecordsPtr{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `"PTR" record

|  Format  |  Description  |
|----------|---------------|
|  text  |  A DNS name relative to the root record  |
|  @  |  Root record  |

`,
		},

		"txt": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ServiceDNSForwardingAuthoritativeDomainRecordsTxt{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `"TXT" record

|  Format  |  Description  |
|----------|---------------|
|  text  |  A DNS name relative to the root record  |
|  @  |  Root record  |

`,
		},

		"spf": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ServiceDNSForwardingAuthoritativeDomainRecordsSpf{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `"SPF" record (type=SPF)

|  Format  |  Description  |
|----------|---------------|
|  text  |  A DNS name relative to the root record  |
|  @  |  Root record  |

`,
		},

		"srv": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ServiceDNSForwardingAuthoritativeDomainRecordsSrv{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `"SRV" record

|  Format  |  Description  |
|----------|---------------|
|  text  |  A DNS name relative to the root record  |
|  @  |  Root record  |

`,
		},

		"naptr": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ServiceDNSForwardingAuthoritativeDomainRecordsNaptr{}.ResourceAttributes(),
			},
			Optional: true,
			MarkdownDescription: `"NAPTR" record

|  Format  |  Description  |
|----------|---------------|
|  text  |  A DNS name relative to the root record  |
|  @  |  Root record  |

`,
		},

		// Nodes

	}
}
