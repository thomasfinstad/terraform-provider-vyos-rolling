/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalprotocolsripdistributelistprefixlist code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsripdistributelistprefixlist

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/global/protocols/rip-distribute-list-prefix-list/resourcemodel"
)

/* tools/generate-terraform-resource-full/templates/resources/global/resource-node-based-full.gotmpl */
// NewProtocolsRIPDistributeListPrefixList method to return the example resource reference
func NewProtocolsRIPDistributeListPrefixList() resource.Resource {
	return &protocolsRIPDistributeListPrefixList{
		model: &resourcemodel.ProtocolsRIPDistributeListPrefixList{},
	}
}

// protocolsRIPDistributeListPrefixList defines the resource implementation.
type protocolsRIPDistributeListPrefixList struct {
	providerData data.ProviderData
	model        *resourcemodel.ProtocolsRIPDistributeListPrefixList
}

// GetClient returns the vyos api client
func (r *protocolsRIPDistributeListPrefixList) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *protocolsRIPDistributeListPrefixList) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *protocolsRIPDistributeListPrefixList) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *protocolsRIPDistributeListPrefixList) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *protocolsRIPDistributeListPrefixList) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
