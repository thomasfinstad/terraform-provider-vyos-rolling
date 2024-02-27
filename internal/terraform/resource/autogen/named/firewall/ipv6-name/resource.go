// Package namedfirewallipvsixname code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallipvsixname

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/firewall/ipv6-name/resourcemodel"
)

// NewFirewallIPvsixName method to return the example resource reference
func NewFirewallIPvsixName() resource.Resource {
	return &firewallIPvsixName{
		model: &resourcemodel.FirewallIPvsixName{},
	}
}

// firewallIPvsixName defines the resource implementation.
type firewallIPvsixName struct {
	client *client.Client
	model  *resourcemodel.FirewallIPvsixName
}

// GetClient returns the vyos api client
func (r *firewallIPvsixName) GetClient() *client.Client {
	return r.client
}

// GetModel returns the resource model
func (r *firewallIPvsixName) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

func (r *firewallIPvsixName) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
