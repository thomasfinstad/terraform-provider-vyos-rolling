// Package namedvrfnameprotocolsospfareavirtuallinkauthenticationmdfivekeyid code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsospfareavirtuallinkauthenticationmdfivekeyid

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/vrf/name-protocols-ospf-area-virtual-link-authentication-md5-key-id/resourcemodel"
)

// NewVrfNameProtocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyID method to return the example resource reference
func NewVrfNameProtocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyID() resource.Resource {
	return &vrfNameProtocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyID{
		model: &resourcemodel.VrfNameProtocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyID{},
	}
}

// vrfNameProtocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyID defines the resource implementation.
type vrfNameProtocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyID struct {
	providerData data.ProviderData
	model        *resourcemodel.VrfNameProtocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyID
}

// GetClient returns the vyos api client
func (r *vrfNameProtocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyID) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *vrfNameProtocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyID) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

func (r *vrfNameProtocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyID) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
