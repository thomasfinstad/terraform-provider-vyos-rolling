// Package namedvrfnameprotocolsbgppeergrouplocalas code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsbgppeergrouplocalas

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/vrf/name-protocols-bgp-peer-group-local-as/resourcemodel"
)

// NewVrfNameProtocolsBgpPeerGroupLocalAs method to return the example resource reference
func NewVrfNameProtocolsBgpPeerGroupLocalAs() resource.Resource {
	return &vrfNameProtocolsBgpPeerGroupLocalAs{
		model: &resourcemodel.VrfNameProtocolsBgpPeerGroupLocalAs{},
	}
}

// vrfNameProtocolsBgpPeerGroupLocalAs defines the resource implementation.
type vrfNameProtocolsBgpPeerGroupLocalAs struct {
	providerData data.ProviderData
	model        *resourcemodel.VrfNameProtocolsBgpPeerGroupLocalAs
}

// GetClient returns the vyos api client
func (r *vrfNameProtocolsBgpPeerGroupLocalAs) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *vrfNameProtocolsBgpPeerGroupLocalAs) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

func (r *vrfNameProtocolsBgpPeerGroupLocalAs) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
