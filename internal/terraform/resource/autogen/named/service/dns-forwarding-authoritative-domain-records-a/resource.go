/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedservicednsforwardingauthoritativedomainrecordsa code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicednsforwardingauthoritativedomainrecordsa

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/named/service/dns-forwarding-authoritative-domain-records-a/resourcemodel"
)

/* tools/generate-terraform-resource-full/templates/resources/named/resource-tagnode-based-full.gotmpl */
// NewServiceDNSForwardingAuthoritativeDomainRecordsA method to return the example resource reference
func NewServiceDNSForwardingAuthoritativeDomainRecordsA() resource.Resource {
	return &serviceDNSForwardingAuthoritativeDomainRecordsA{
		model: &resourcemodel.ServiceDNSForwardingAuthoritativeDomainRecordsA{},
	}
}

// serviceDNSForwardingAuthoritativeDomainRecordsA defines the resource implementation.
type serviceDNSForwardingAuthoritativeDomainRecordsA struct {
	providerData data.ProviderData
	model        *resourcemodel.ServiceDNSForwardingAuthoritativeDomainRecordsA
}

// GetClient returns the vyos api client
func (r *serviceDNSForwardingAuthoritativeDomainRecordsA) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *serviceDNSForwardingAuthoritativeDomainRecordsA) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *serviceDNSForwardingAuthoritativeDomainRecordsA) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *serviceDNSForwardingAuthoritativeDomainRecordsA) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	data, ok := req.ProviderData.(data.ProviderData)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.providerData = data
}

func (r *serviceDNSForwardingAuthoritativeDomainRecordsA) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
