/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalserviceidsddosprotectionthresholdudp code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalserviceidsddosprotectionthresholdudp

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
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/global/service/ids-ddos-protection-threshold-udp/resourcemodel"
)

/* tools/generate-terraform-resource-full/templates/resources/global/resource-node-based-full.gotmpl #resource-node-based-full (udp) */
// NewServiceIDsDdosProtectionThresholdUDP method to return the example resource reference
func NewServiceIDsDdosProtectionThresholdUDP() resource.Resource {
	return &serviceIDsDdosProtectionThresholdUDP{
		model: &resourcemodel.ServiceIDsDdosProtectionThresholdUDP{},
	}
}

// serviceIDsDdosProtectionThresholdUDP defines the resource implementation.
type serviceIDsDdosProtectionThresholdUDP struct {
	providerData data.ProviderData
	model        *resourcemodel.ServiceIDsDdosProtectionThresholdUDP
}

// GetClient returns the vyos api client
func (r *serviceIDsDdosProtectionThresholdUDP) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *serviceIDsDdosProtectionThresholdUDP) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *serviceIDsDdosProtectionThresholdUDP) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *serviceIDsDdosProtectionThresholdUDP) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *serviceIDsDdosProtectionThresholdUDP) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
