// Package namedprotocolsisisfastreroutelfalocaltiebreakernodeprotectingindex code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsisisfastreroutelfalocaltiebreakernodeprotectingindex

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/named/protocols/isis-fast-reroute-lfa-local-tiebreaker-node-protecting-index/resourcemodel"
)

// NewProtocolsIsisFastRerouteLfaLocalTiebreakerNodeProtectingIndex method to return the example resource reference
func NewProtocolsIsisFastRerouteLfaLocalTiebreakerNodeProtectingIndex() resource.Resource {
	return &protocolsIsisFastRerouteLfaLocalTiebreakerNodeProtectingIndex{
		model: &resourcemodel.ProtocolsIsisFastRerouteLfaLocalTiebreakerNodeProtectingIndex{},
	}
}

// protocolsIsisFastRerouteLfaLocalTiebreakerNodeProtectingIndex defines the resource implementation.
type protocolsIsisFastRerouteLfaLocalTiebreakerNodeProtectingIndex struct {
	providerData data.ProviderData
	model        *resourcemodel.ProtocolsIsisFastRerouteLfaLocalTiebreakerNodeProtectingIndex
}

// GetClient returns the vyos api client
func (r *protocolsIsisFastRerouteLfaLocalTiebreakerNodeProtectingIndex) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *protocolsIsisFastRerouteLfaLocalTiebreakerNodeProtectingIndex) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *protocolsIsisFastRerouteLfaLocalTiebreakerNodeProtectingIndex) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *protocolsIsisFastRerouteLfaLocalTiebreakerNodeProtectingIndex) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *protocolsIsisFastRerouteLfaLocalTiebreakerNodeProtectingIndex) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
