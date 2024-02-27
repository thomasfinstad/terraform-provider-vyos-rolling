// Package globalfirewallglobaloptionsstatepolicyrelated code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalfirewallglobaloptionsstatepolicyrelated

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/global/firewall/global-options-state-policy-related/resourcemodel"
)

// NewFirewallGlobalOptionsStatePolicyRelated method to return the example resource reference
func NewFirewallGlobalOptionsStatePolicyRelated() resource.Resource {
	return &firewallGlobalOptionsStatePolicyRelated{
		model: &resourcemodel.FirewallGlobalOptionsStatePolicyRelated{},
	}
}

// firewallGlobalOptionsStatePolicyRelated defines the resource implementation.
type firewallGlobalOptionsStatePolicyRelated struct {
	client *client.Client
	model  *resourcemodel.FirewallGlobalOptionsStatePolicyRelated
}

// GetClient returns the vyos api client
func (r *firewallGlobalOptionsStatePolicyRelated) GetClient() *client.Client {
	return r.client
}

// GetModel returns the resource model
func (r *firewallGlobalOptionsStatePolicyRelated) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

func (r *firewallGlobalOptionsStatePolicyRelated) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = &client
}
