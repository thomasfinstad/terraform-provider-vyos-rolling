// Package globalsystemoptiontcpclient code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemoptiontcpclient

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/global/system/option-ssh-client/resourcemodel"
)

// NewSystemOptionTCPClient method to return the example resource reference
func NewSystemOptionTCPClient() resource.Resource {
	return &systemOptionTCPClient{
		model: &resourcemodel.SystemOptionTCPClient{},
	}
}

// systemOptionTCPClient defines the resource implementation.
type systemOptionTCPClient struct {
	providerData data.ProviderData
	model        *resourcemodel.SystemOptionTCPClient
}

// GetClient returns the vyos api client
func (r *systemOptionTCPClient) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *systemOptionTCPClient) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *systemOptionTCPClient) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *systemOptionTCPClient) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *systemOptionTCPClient) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
