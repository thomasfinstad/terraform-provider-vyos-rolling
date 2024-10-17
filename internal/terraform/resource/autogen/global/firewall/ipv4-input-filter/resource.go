// Package globalfirewallipvfourinputfilter code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalfirewallipvfourinputfilter

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/global/firewall/ipv4-input-filter/resourcemodel"
)

// NewFirewallIPvfourInputFilter method to return the example resource reference
func NewFirewallIPvfourInputFilter() resource.Resource {
	return &firewallIPvfourInputFilter{
		model: &resourcemodel.FirewallIPvfourInputFilter{},
	}
}

// firewallIPvfourInputFilter defines the resource implementation.
type firewallIPvfourInputFilter struct {
	providerData data.ProviderData
	model        *resourcemodel.FirewallIPvfourInputFilter
}

// GetClient returns the vyos api client
func (r *firewallIPvfourInputFilter) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *firewallIPvfourInputFilter) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *firewallIPvfourInputFilter) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *firewallIPvfourInputFilter) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *firewallIPvfourInputFilter) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
