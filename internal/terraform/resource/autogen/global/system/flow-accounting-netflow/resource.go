/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalsystemflowaccountingnetflow code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemflowaccountingnetflow

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
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/global/system/flow-accounting-netflow/resourcemodel"
)

/* tools/generate-terraform-resource-full/templates/resources/global/resource-node-based-full.gotmpl #resource-node-based-full (netflow) */
// NewSystemFlowAccountingNetflow method to return the example resource reference
func NewSystemFlowAccountingNetflow() resource.Resource {
	return &systemFlowAccountingNetflow{
		model: &resourcemodel.SystemFlowAccountingNetflow{},
	}
}

// systemFlowAccountingNetflow defines the resource implementation.
type systemFlowAccountingNetflow struct {
	providerData data.ProviderData
	model        *resourcemodel.SystemFlowAccountingNetflow
}

// GetClient returns the vyos api client
func (r *systemFlowAccountingNetflow) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *systemFlowAccountingNetflow) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *systemFlowAccountingNetflow) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *systemFlowAccountingNetflow) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *systemFlowAccountingNetflow) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
