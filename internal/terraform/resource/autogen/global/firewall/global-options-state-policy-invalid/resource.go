/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalfirewallglobaloptionsstatepolicyinvalid code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalfirewallglobaloptionsstatepolicyinvalid

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
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/global/firewall/global-options-state-policy-invalid/resourcemodel"
)

/* tools/generate-terraform-resource-full/templates/resources/global/resource-node-based-full.gotmpl */
// NewFirewallGlobalOptionsStatePolicyInvalID method to return the example resource reference
func NewFirewallGlobalOptionsStatePolicyInvalID() resource.Resource {
	return &firewallGlobalOptionsStatePolicyInvalID{
		model: &resourcemodel.FirewallGlobalOptionsStatePolicyInvalID{},
	}
}

// firewallGlobalOptionsStatePolicyInvalID defines the resource implementation.
type firewallGlobalOptionsStatePolicyInvalID struct {
	providerData data.ProviderData
	model        *resourcemodel.FirewallGlobalOptionsStatePolicyInvalID
}

// GetClient returns the vyos api client
func (r *firewallGlobalOptionsStatePolicyInvalID) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *firewallGlobalOptionsStatePolicyInvalID) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *firewallGlobalOptionsStatePolicyInvalID) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *firewallGlobalOptionsStatePolicyInvalID) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *firewallGlobalOptionsStatePolicyInvalID) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
