// Package globalsystemconntrackmodules code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemconntrackmodules

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/global/system/conntrack-modules/resourcemodel"
)

// NewSystemConntrackModules method to return the example resource reference
func NewSystemConntrackModules() resource.Resource {
	return &systemConntrackModules{
		model: &resourcemodel.SystemConntrackModules{},
	}
}

// systemConntrackModules defines the resource implementation.
type systemConntrackModules struct {
	providerData data.ProviderData
	model        *resourcemodel.SystemConntrackModules
}

// GetClient returns the vyos api client
func (r *systemConntrackModules) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *systemConntrackModules) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *systemConntrackModules) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *systemConntrackModules) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
