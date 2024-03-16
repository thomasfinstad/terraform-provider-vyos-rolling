// Package namedfirewallipvsixoutputfilterrule code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallipvsixoutputfilterrule

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/firewall/ipv6-output-filter-rule/resourcemodel"
)

// NewFirewallIPvsixOutputFilterRule method to return the example resource reference
func NewFirewallIPvsixOutputFilterRule() resource.Resource {
	return &firewallIPvsixOutputFilterRule{
		model: &resourcemodel.FirewallIPvsixOutputFilterRule{},
	}
}

// firewallIPvsixOutputFilterRule defines the resource implementation.
type firewallIPvsixOutputFilterRule struct {
	providerData data.ProviderData
	model        *resourcemodel.FirewallIPvsixOutputFilterRule
}

// GetClient returns the vyos api client
func (r *firewallIPvsixOutputFilterRule) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *firewallIPvsixOutputFilterRule) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

func (r *firewallIPvsixOutputFilterRule) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
