// Package namedvrfnameprotocolsstaticroutesixinterface code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsstaticroutesixinterface

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/named/vrf/name-protocols-static-route6-interface/resourcemodel"
)

// NewVrfNameProtocolsStaticRoutesixInterface method to return the example resource reference
func NewVrfNameProtocolsStaticRoutesixInterface() resource.Resource {
	return &vrfNameProtocolsStaticRoutesixInterface{
		model: &resourcemodel.VrfNameProtocolsStaticRoutesixInterface{},
	}
}

// vrfNameProtocolsStaticRoutesixInterface defines the resource implementation.
type vrfNameProtocolsStaticRoutesixInterface struct {
	providerData data.ProviderData
	model        *resourcemodel.VrfNameProtocolsStaticRoutesixInterface
}

// GetClient returns the vyos api client
func (r *vrfNameProtocolsStaticRoutesixInterface) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *vrfNameProtocolsStaticRoutesixInterface) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *vrfNameProtocolsStaticRoutesixInterface) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *vrfNameProtocolsStaticRoutesixInterface) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
