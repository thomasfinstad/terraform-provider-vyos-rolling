// Package namedvpnipsecremoteaccessconnectionauthenticationlocalusersusername code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnipsecremoteaccessconnectionauthenticationlocalusersusername

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/vpn/ipsec-remote-access-connection-authentication-local-users-username/resourcemodel"
)

// NewVpnIPsecRemoteAccessConnectionAuthenticationLocalUsersUsername method to return the example resource reference
func NewVpnIPsecRemoteAccessConnectionAuthenticationLocalUsersUsername() resource.Resource {
	return &vpnIPsecRemoteAccessConnectionAuthenticationLocalUsersUsername{
		model: resourcemodel.VpnIPsecRemoteAccessConnectionAuthenticationLocalUsersUsername{},
	}
}

// vpnIPsecRemoteAccessConnectionAuthenticationLocalUsersUsername defines the resource implementation.
type vpnIPsecRemoteAccessConnectionAuthenticationLocalUsersUsername struct {
	ResourceName string
	client       *client.Client
	model        resourcemodel.VpnIPsecRemoteAccessConnectionAuthenticationLocalUsersUsername
}

// GetName returns resource name
func (r *vpnIPsecRemoteAccessConnectionAuthenticationLocalUsersUsername) GetName() string {
	return r.ResourceName
}

func (r *vpnIPsecRemoteAccessConnectionAuthenticationLocalUsersUsername) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
