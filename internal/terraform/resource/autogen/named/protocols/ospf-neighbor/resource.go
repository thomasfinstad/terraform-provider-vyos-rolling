// Package namedprotocolsospfneighbor code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsospfneighbor

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/named/protocols/ospf-neighbor/resourcemodel"
)

// NewProtocolsOspfNeighbor method to return the example resource reference
func NewProtocolsOspfNeighbor() resource.Resource {
	return &protocolsOspfNeighbor{
		model: &resourcemodel.ProtocolsOspfNeighbor{},
	}
}

// protocolsOspfNeighbor defines the resource implementation.
type protocolsOspfNeighbor struct {
	providerData data.ProviderData
	model        *resourcemodel.ProtocolsOspfNeighbor
}

// GetClient returns the vyos api client
func (r *protocolsOspfNeighbor) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *protocolsOspfNeighbor) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *protocolsOspfNeighbor) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *protocolsOspfNeighbor) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
