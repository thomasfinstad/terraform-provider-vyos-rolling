// Package namedsystemipvsixprotocol code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedsystemipvsixprotocol

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/named/system/ipv6-protocol/resourcemodel"
)

// NewSystemIPvsixProtocol method to return the example resource reference
func NewSystemIPvsixProtocol() resource.Resource {
	return &systemIPvsixProtocol{
		model: &resourcemodel.SystemIPvsixProtocol{},
	}
}

// systemIPvsixProtocol defines the resource implementation.
type systemIPvsixProtocol struct {
	providerData data.ProviderData
	model        *resourcemodel.SystemIPvsixProtocol
}

// GetClient returns the vyos api client
func (r *systemIPvsixProtocol) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *systemIPvsixProtocol) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *systemIPvsixProtocol) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *systemIPvsixProtocol) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *systemIPvsixProtocol) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
