/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalprotocolsbgpparametersdampening code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgpparametersdampening

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
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/global/protocols/bgp-parameters-dampening/resourcemodel"
)

/* tools/generate-terraform-resource-full/templates/resources/global/resource-node-based-full.gotmpl #resource-node-based-full (dampening) */
// NewProtocolsBgpParametersDampening method to return the example resource reference
func NewProtocolsBgpParametersDampening() resource.Resource {
	return &protocolsBgpParametersDampening{
		model: &resourcemodel.ProtocolsBgpParametersDampening{},
	}
}

// protocolsBgpParametersDampening defines the resource implementation.
type protocolsBgpParametersDampening struct {
	providerData data.ProviderData
	model        *resourcemodel.ProtocolsBgpParametersDampening
}

// GetClient returns the vyos api client
func (r *protocolsBgpParametersDampening) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *protocolsBgpParametersDampening) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *protocolsBgpParametersDampening) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *protocolsBgpParametersDampening) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *protocolsBgpParametersDampening) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
