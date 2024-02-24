// Package globalvpnsstpauthenticationradiusratelimit code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnsstpauthenticationradiusratelimit

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/global/vpn/sstp-authentication-radius-rate-limit/resourcemodel"
)

// NewVpnSstpAuthenticationRadiusRateLimit method to return the example resource reference
func NewVpnSstpAuthenticationRadiusRateLimit() resource.Resource {
	return &vpnSstpAuthenticationRadiusRateLimit{
		model: &resourcemodel.VpnSstpAuthenticationRadiusRateLimit{},
	}
}

// vpnSstpAuthenticationRadiusRateLimit defines the resource implementation.
type vpnSstpAuthenticationRadiusRateLimit struct {
	client *client.Client
	model  *resourcemodel.VpnSstpAuthenticationRadiusRateLimit
}

// GetClient returns the vyos api client
func (r *vpnSstpAuthenticationRadiusRateLimit) GetClient() *client.Client {
	return r.client
}

func (r *vpnSstpAuthenticationRadiusRateLimit) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}
