// Package namedvrfnameprotocolsbgpaddressfamilyltwovpnevpnvni code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsbgpaddressfamilyltwovpnevpnvni

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/vrf/name-protocols-bgp-address-family-l2vpn-evpn-vni/resourcemodel"
)

// NewVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVni method to return the example resource reference
func NewVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVni() resource.Resource {
	return &vrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVni{
		model: &resourcemodel.VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVni{},
	}
}

// vrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVni defines the resource implementation.
type vrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVni struct {
	client *client.Client
	model  *resourcemodel.VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVni
}

// GetClient returns the vyos api client
func (r *vrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVni) GetClient() *client.Client {
	return r.client
}

// GetModel returns the resource model
func (r *vrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVni) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

func (r *vrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVni) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = &client
}
