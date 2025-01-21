/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalprotocolsripngredistributeconnected code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsripngredistributeconnected

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/global/protocols/ripng-redistribute-connected/resourcemodel"
)

/* tools/generate-terraform-resource-full/templates/resources/global/resource-node-based-full.gotmpl #resource-node-based-full (connected) */
// NewProtocolsRIPngRedistributeConnected method to return the example resource reference
func NewProtocolsRIPngRedistributeConnected() resource.Resource {
	return &protocolsRIPngRedistributeConnected{
		model: &resourcemodel.ProtocolsRIPngRedistributeConnected{},
	}
}

// protocolsRIPngRedistributeConnected defines the resource implementation.
type protocolsRIPngRedistributeConnected struct {
	providerData data.ProviderData
	model        *resourcemodel.ProtocolsRIPngRedistributeConnected
}

// GetClient returns the vyos api client
func (r *protocolsRIPngRedistributeConnected) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *protocolsRIPngRedistributeConnected) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *protocolsRIPngRedistributeConnected) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *protocolsRIPngRedistributeConnected) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *protocolsRIPngRedistributeConnected) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
