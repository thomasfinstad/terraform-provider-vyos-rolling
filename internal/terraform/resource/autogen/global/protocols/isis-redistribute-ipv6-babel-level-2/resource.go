// Package globalprotocolsisisredistributeipvsixbabelleveltwo code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisisredistributeipvsixbabelleveltwo

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/global/protocols/isis-redistribute-ipv6-babel-level-2/resourcemodel"
)

// NewProtocolsIsisRedistributeIPvsixBabelLevelTwo method to return the example resource reference
func NewProtocolsIsisRedistributeIPvsixBabelLevelTwo() resource.Resource {
	return &protocolsIsisRedistributeIPvsixBabelLevelTwo{
		model: &resourcemodel.ProtocolsIsisRedistributeIPvsixBabelLevelTwo{},
	}
}

// protocolsIsisRedistributeIPvsixBabelLevelTwo defines the resource implementation.
type protocolsIsisRedistributeIPvsixBabelLevelTwo struct {
	providerData data.ProviderData
	model        *resourcemodel.ProtocolsIsisRedistributeIPvsixBabelLevelTwo
}

// GetClient returns the vyos api client
func (r *protocolsIsisRedistributeIPvsixBabelLevelTwo) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *protocolsIsisRedistributeIPvsixBabelLevelTwo) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *protocolsIsisRedistributeIPvsixBabelLevelTwo) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *protocolsIsisRedistributeIPvsixBabelLevelTwo) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
