// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// ServiceDNSForwardingAuthoritativeDomainRecords describes the resource data model.
type ServiceDNSForwardingAuthoritativeDomainRecords struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagServiceDNSForwardingAuthoritativeDomainRecordsA     bool `tfsdk:"a" vyos:"a,child"`
	ExistsTagServiceDNSForwardingAuthoritativeDomainRecordsAaaa  bool `tfsdk:"aaaa" vyos:"aaaa,child"`
	ExistsTagServiceDNSForwardingAuthoritativeDomainRecordsCname bool `tfsdk:"cname" vyos:"cname,child"`
	ExistsTagServiceDNSForwardingAuthoritativeDomainRecordsMx    bool `tfsdk:"mx" vyos:"mx,child"`
	ExistsTagServiceDNSForwardingAuthoritativeDomainRecordsPtr   bool `tfsdk:"ptr" vyos:"ptr,child"`
	ExistsTagServiceDNSForwardingAuthoritativeDomainRecordsTxt   bool `tfsdk:"txt" vyos:"txt,child"`
	ExistsTagServiceDNSForwardingAuthoritativeDomainRecordsSpf   bool `tfsdk:"spf" vyos:"spf,child"`
	ExistsTagServiceDNSForwardingAuthoritativeDomainRecordsSrv   bool `tfsdk:"srv" vyos:"srv,child"`
	ExistsTagServiceDNSForwardingAuthoritativeDomainRecordsNaptr bool `tfsdk:"naptr" vyos:"naptr,child"`

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDNSForwardingAuthoritativeDomainRecords) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{

		// LeafNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServiceDNSForwardingAuthoritativeDomainRecords) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ServiceDNSForwardingAuthoritativeDomainRecords) UnmarshalJSON(_ []byte) error {
	return nil
}