// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// ServiceDNSForwardingAuthoritativeDomainRecords describes the resource data model.
type ServiceDNSForwardingAuthoritativeDomainRecords struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagServiceDNSForwardingAuthoritativeDomainRecordsA     bool `tfsdk:"a" vyos:"a,ignore,child"`
	ExistsTagServiceDNSForwardingAuthoritativeDomainRecordsAaaa  bool `tfsdk:"aaaa" vyos:"aaaa,ignore,child"`
	ExistsTagServiceDNSForwardingAuthoritativeDomainRecordsCname bool `tfsdk:"cname" vyos:"cname,ignore,child"`
	ExistsTagServiceDNSForwardingAuthoritativeDomainRecordsMx    bool `tfsdk:"mx" vyos:"mx,ignore,child"`
	ExistsTagServiceDNSForwardingAuthoritativeDomainRecordsPtr   bool `tfsdk:"ptr" vyos:"ptr,ignore,child"`
	ExistsTagServiceDNSForwardingAuthoritativeDomainRecordsTxt   bool `tfsdk:"txt" vyos:"txt,ignore,child"`
	ExistsTagServiceDNSForwardingAuthoritativeDomainRecordsSpf   bool `tfsdk:"spf" vyos:"spf,ignore,child"`
	ExistsTagServiceDNSForwardingAuthoritativeDomainRecordsSrv   bool `tfsdk:"srv" vyos:"srv,ignore,child"`
	ExistsTagServiceDNSForwardingAuthoritativeDomainRecordsNaptr bool `tfsdk:"naptr" vyos:"naptr,ignore,child"`

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDNSForwardingAuthoritativeDomainRecords) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{

		// LeafNodes

		// Nodes

	}
}
