/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalfirewallglobaloptions code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalfirewallglobaloptions

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
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/global/firewall/global-options/resourcemodel"
)

/* tools/generate-terraform-resource-full/templates/resources/global/resource-node-based-full.gotmpl */
// NewFirewallGlobalOptions method to return the example resource reference
func NewFirewallGlobalOptions() resource.Resource {
	return &firewallGlobalOptions{
		model: &resourcemodel.FirewallGlobalOptions{},
	}
}

// firewallGlobalOptions defines the resource implementation.
type firewallGlobalOptions struct {
	providerData data.ProviderData
	model        *resourcemodel.FirewallGlobalOptions
}

// GetClient returns the vyos api client
func (r *firewallGlobalOptions) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *firewallGlobalOptions) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *firewallGlobalOptions) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *firewallGlobalOptions) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *firewallGlobalOptions) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
