// Package namedprotocolsnhrptunnelmap code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsnhrptunnelmap

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/named/protocols/nhrp-tunnel-map/resourcemodel"
)

// NewProtocolsNhrpTunnelMap method to return the example resource reference
func NewProtocolsNhrpTunnelMap() resource.Resource {
	return &protocolsNhrpTunnelMap{
		model: &resourcemodel.ProtocolsNhrpTunnelMap{},
	}
}

// protocolsNhrpTunnelMap defines the resource implementation.
type protocolsNhrpTunnelMap struct {
	providerData data.ProviderData
	model        *resourcemodel.ProtocolsNhrpTunnelMap
}

// GetClient returns the vyos api client
func (r *protocolsNhrpTunnelMap) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *protocolsNhrpTunnelMap) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *protocolsNhrpTunnelMap) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *protocolsNhrpTunnelMap) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *protocolsNhrpTunnelMap) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
