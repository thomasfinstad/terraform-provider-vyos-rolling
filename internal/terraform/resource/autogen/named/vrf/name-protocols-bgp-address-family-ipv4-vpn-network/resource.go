// Package namedvrfnameprotocolsbgpaddressfamilyipvfourvpnnetwork code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsbgpaddressfamilyipvfourvpnnetwork

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/vrf/name-protocols-bgp-address-family-ipv4-vpn-network/resourcemodel"
)

// NewVrfNameProtocolsBgpAddressFamilyIPvfourVpnNetwork method to return the example resource reference
func NewVrfNameProtocolsBgpAddressFamilyIPvfourVpnNetwork() resource.Resource {
	return &vrfNameProtocolsBgpAddressFamilyIPvfourVpnNetwork{
		model: &resourcemodel.VrfNameProtocolsBgpAddressFamilyIPvfourVpnNetwork{},
	}
}

// vrfNameProtocolsBgpAddressFamilyIPvfourVpnNetwork defines the resource implementation.
type vrfNameProtocolsBgpAddressFamilyIPvfourVpnNetwork struct {
	providerData data.ProviderData
	model        *resourcemodel.VrfNameProtocolsBgpAddressFamilyIPvfourVpnNetwork
}

// GetClient returns the vyos api client
func (r *vrfNameProtocolsBgpAddressFamilyIPvfourVpnNetwork) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *vrfNameProtocolsBgpAddressFamilyIPvfourVpnNetwork) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

func (r *vrfNameProtocolsBgpAddressFamilyIPvfourVpnNetwork) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
