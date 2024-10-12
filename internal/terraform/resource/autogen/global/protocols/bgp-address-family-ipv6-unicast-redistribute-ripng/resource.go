// Package globalprotocolsbgpaddressfamilyipvsixunicastredistributeripng code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgpaddressfamilyipvsixunicastredistributeripng

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/global/protocols/bgp-address-family-ipv6-unicast-redistribute-ripng/resourcemodel"
)

// NewProtocolsBgpAddressFamilyIPvsixUnicastRedistributeRIPng method to return the example resource reference
func NewProtocolsBgpAddressFamilyIPvsixUnicastRedistributeRIPng() resource.Resource {
	return &protocolsBgpAddressFamilyIPvsixUnicastRedistributeRIPng{
		model: &resourcemodel.ProtocolsBgpAddressFamilyIPvsixUnicastRedistributeRIPng{},
	}
}

// protocolsBgpAddressFamilyIPvsixUnicastRedistributeRIPng defines the resource implementation.
type protocolsBgpAddressFamilyIPvsixUnicastRedistributeRIPng struct {
	providerData data.ProviderData
	model        *resourcemodel.ProtocolsBgpAddressFamilyIPvsixUnicastRedistributeRIPng
}

// GetClient returns the vyos api client
func (r *protocolsBgpAddressFamilyIPvsixUnicastRedistributeRIPng) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *protocolsBgpAddressFamilyIPvsixUnicastRedistributeRIPng) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *protocolsBgpAddressFamilyIPvsixUnicastRedistributeRIPng) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *protocolsBgpAddressFamilyIPvsixUnicastRedistributeRIPng) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
