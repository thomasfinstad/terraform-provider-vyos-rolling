/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalservicetcpaccesscontrolallow code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicetcpaccesscontrolallow

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
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/global/service/ssh-access-control-allow/resourcemodel"
)

/* tools/generate-terraform-resource-full/templates/resources/global/resource-node-based-full.gotmpl #resource-node-based-full (allow) */
// NewServiceTCPAccessControlAllow method to return the example resource reference
func NewServiceTCPAccessControlAllow() resource.Resource {
	return &serviceTCPAccessControlAllow{
		model: &resourcemodel.ServiceTCPAccessControlAllow{},
	}
}

// serviceTCPAccessControlAllow defines the resource implementation.
type serviceTCPAccessControlAllow struct {
	providerData data.ProviderData
	model        *resourcemodel.ServiceTCPAccessControlAllow
}

// GetClient returns the vyos api client
func (r *serviceTCPAccessControlAllow) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *serviceTCPAccessControlAllow) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *serviceTCPAccessControlAllow) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *serviceTCPAccessControlAllow) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *serviceTCPAccessControlAllow) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
