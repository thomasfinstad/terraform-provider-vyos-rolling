/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedservicednsforwardingauthoritativedomainrecordssrv code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicednsforwardingauthoritativedomainrecordssrv

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

/* tools/generate-terraform-resource-full/templates/resources/common/metadata.gotmpl #metadata (srv) */
// Metadata method to define the resource type name.
func (r serviceDNSForwardingAuthoritativeDomainRecordsSrv) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_dns_forwarding_authoritative_domain_records_srv"
}
