// Package namedcontainernetwork code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedcontainernetwork

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/named/container/network/resourcemodel"
)

// NewContainerNetwork method to return the example resource reference
func NewContainerNetwork() resource.Resource {
	return &containerNetwork{
		model: &resourcemodel.ContainerNetwork{},
	}
}

// containerNetwork defines the resource implementation.
type containerNetwork struct {
	providerData data.ProviderData
	model        *resourcemodel.ContainerNetwork
}

// GetClient returns the vyos api client
func (r *containerNetwork) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *containerNetwork) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *containerNetwork) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *containerNetwork) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *containerNetwork) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
