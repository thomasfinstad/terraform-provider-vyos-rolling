// Package globalprotocolsbgpparameterstcpkeepalive code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgpparameterstcpkeepalive

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/global/protocols/bgp-parameters-tcp-keepalive/resourcemodel"
)

// NewProtocolsBgpParametersTCPKeepalive method to return the example resource reference
func NewProtocolsBgpParametersTCPKeepalive() resource.Resource {
	return &protocolsBgpParametersTCPKeepalive{
		model: &resourcemodel.ProtocolsBgpParametersTCPKeepalive{},
	}
}

// protocolsBgpParametersTCPKeepalive defines the resource implementation.
type protocolsBgpParametersTCPKeepalive struct {
	providerData data.ProviderData
	model        *resourcemodel.ProtocolsBgpParametersTCPKeepalive
}

// GetClient returns the vyos api client
func (r *protocolsBgpParametersTCPKeepalive) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *protocolsBgpParametersTCPKeepalive) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *protocolsBgpParametersTCPKeepalive) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *protocolsBgpParametersTCPKeepalive) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *protocolsBgpParametersTCPKeepalive) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
