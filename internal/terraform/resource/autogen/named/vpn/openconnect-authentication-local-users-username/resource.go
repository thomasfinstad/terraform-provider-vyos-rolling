// Package namedvpnopenconnectauthenticationlocalusersusername code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnopenconnectauthenticationlocalusersusername

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/named/vpn/openconnect-authentication-local-users-username/resourcemodel"
)

// NewVpnOpenconnectAuthenticationLocalUsersUsername method to return the example resource reference
func NewVpnOpenconnectAuthenticationLocalUsersUsername() resource.Resource {
	return &vpnOpenconnectAuthenticationLocalUsersUsername{
		model: &resourcemodel.VpnOpenconnectAuthenticationLocalUsersUsername{},
	}
}

// vpnOpenconnectAuthenticationLocalUsersUsername defines the resource implementation.
type vpnOpenconnectAuthenticationLocalUsersUsername struct {
	providerData data.ProviderData
	model        *resourcemodel.VpnOpenconnectAuthenticationLocalUsersUsername
}

// GetClient returns the vyos api client
func (r *vpnOpenconnectAuthenticationLocalUsersUsername) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *vpnOpenconnectAuthenticationLocalUsersUsername) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *vpnOpenconnectAuthenticationLocalUsersUsername) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *vpnOpenconnectAuthenticationLocalUsersUsername) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
