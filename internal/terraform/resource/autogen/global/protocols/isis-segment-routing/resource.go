// Package globalprotocolsisissegmentrouting code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisissegmentrouting

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/global/protocols/isis-segment-routing/resourcemodel"
)

// NewProtocolsIsisSegmentRouting method to return the example resource reference
func NewProtocolsIsisSegmentRouting() resource.Resource {
	return &protocolsIsisSegmentRouting{
		model: &resourcemodel.ProtocolsIsisSegmentRouting{},
	}
}

// protocolsIsisSegmentRouting defines the resource implementation.
type protocolsIsisSegmentRouting struct {
	providerData data.ProviderData
	model        *resourcemodel.ProtocolsIsisSegmentRouting
}

// GetClient returns the vyos api client
func (r *protocolsIsisSegmentRouting) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *protocolsIsisSegmentRouting) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *protocolsIsisSegmentRouting) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *protocolsIsisSegmentRouting) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *protocolsIsisSegmentRouting) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
