/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalprotocolsisisdefaultinformationoriginateipvfourleveltwo code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisisdefaultinformationoriginateipvfourleveltwo

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
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/global/protocols/isis-default-information-originate-ipv4-level-2/resourcemodel"
)

/* tools/generate-terraform-resource-full/templates/resources/global/resource-node-based-full.gotmpl */
// NewProtocolsIsisDefaultInformationOriginateIPvfourLevelTwo method to return the example resource reference
func NewProtocolsIsisDefaultInformationOriginateIPvfourLevelTwo() resource.Resource {
	return &protocolsIsisDefaultInformationOriginateIPvfourLevelTwo{
		model: &resourcemodel.ProtocolsIsisDefaultInformationOriginateIPvfourLevelTwo{},
	}
}

// protocolsIsisDefaultInformationOriginateIPvfourLevelTwo defines the resource implementation.
type protocolsIsisDefaultInformationOriginateIPvfourLevelTwo struct {
	providerData data.ProviderData
	model        *resourcemodel.ProtocolsIsisDefaultInformationOriginateIPvfourLevelTwo
}

// GetClient returns the vyos api client
func (r *protocolsIsisDefaultInformationOriginateIPvfourLevelTwo) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *protocolsIsisDefaultInformationOriginateIPvfourLevelTwo) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *protocolsIsisDefaultInformationOriginateIPvfourLevelTwo) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *protocolsIsisDefaultInformationOriginateIPvfourLevelTwo) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *protocolsIsisDefaultInformationOriginateIPvfourLevelTwo) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
