// Package namedprotocolsospfinterfaceauthenticationmdfivekeyid code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsospfinterfaceauthenticationmdfivekeyid

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/protocols/ospf-interface-authentication-md5-key-id/resourcemodel"
)

// NewProtocolsOspfInterfaceAuthenticationMdfiveKeyID method to return the example resource reference
func NewProtocolsOspfInterfaceAuthenticationMdfiveKeyID() resource.Resource {
	return &protocolsOspfInterfaceAuthenticationMdfiveKeyID{
		model: resourcemodel.ProtocolsOspfInterfaceAuthenticationMdfiveKeyID{},
	}
}

// protocolsOspfInterfaceAuthenticationMdfiveKeyID defines the resource implementation.
type protocolsOspfInterfaceAuthenticationMdfiveKeyID struct {
	ResourceName string
	client       *client.Client
	model        resourcemodel.ProtocolsOspfInterfaceAuthenticationMdfiveKeyID
}

// GetName returns resource name
func (r *protocolsOspfInterfaceAuthenticationMdfiveKeyID) GetName() string {
	return r.ResourceName
}

func (r *protocolsOspfInterfaceAuthenticationMdfiveKeyID) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
