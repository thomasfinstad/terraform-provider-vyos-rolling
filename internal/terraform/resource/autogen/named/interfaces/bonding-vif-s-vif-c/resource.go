// Package namedinterfacesbondingvifsvifc code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesbondingvifsvifc

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/named/interfaces/bonding-vif-s-vif-c/resourcemodel"
)

// NewInterfacesBondingVifSVifC method to return the example resource reference
func NewInterfacesBondingVifSVifC() resource.Resource {
	return &interfacesBondingVifSVifC{
		model: &resourcemodel.InterfacesBondingVifSVifC{},
	}
}

// interfacesBondingVifSVifC defines the resource implementation.
type interfacesBondingVifSVifC struct {
	providerData data.ProviderData
	model        *resourcemodel.InterfacesBondingVifSVifC
}

// GetClient returns the vyos api client
func (r *interfacesBondingVifSVifC) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *interfacesBondingVifSVifC) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *interfacesBondingVifSVifC) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *interfacesBondingVifSVifC) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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