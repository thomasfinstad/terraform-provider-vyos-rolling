/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalservicepppoeserverextendedscripts code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicepppoeserverextendedscripts

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
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/global/service/pppoe-server-extended-scripts/resourcemodel"
)

/* tools/generate-terraform-resource-full/templates/resources/global/resource-node-based-full.gotmpl #resource-node-based-full (extended-scripts) */
// NewServicePppoeServerExtendedScrIPts method to return the example resource reference
func NewServicePppoeServerExtendedScrIPts() resource.Resource {
	return &servicePppoeServerExtendedScrIPts{
		model: &resourcemodel.ServicePppoeServerExtendedScrIPts{},
	}
}

// servicePppoeServerExtendedScrIPts defines the resource implementation.
type servicePppoeServerExtendedScrIPts struct {
	providerData data.ProviderData
	model        *resourcemodel.ServicePppoeServerExtendedScrIPts
}

// GetClient returns the vyos api client
func (r *servicePppoeServerExtendedScrIPts) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *servicePppoeServerExtendedScrIPts) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *servicePppoeServerExtendedScrIPts) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *servicePppoeServerExtendedScrIPts) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *servicePppoeServerExtendedScrIPts) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
