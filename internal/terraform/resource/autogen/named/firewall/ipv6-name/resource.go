// Package namedfirewallipvsixname code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallipvsixname

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/firewall/ipv6-name/resourcemodel"
)

// NewFirewallIPvsixName method to return the example resource reference
func NewFirewallIPvsixName() resource.Resource {
	return &firewallIPvsixName{
		model: resourcemodel.FirewallIPvsixName{},
	}
}

// firewallIPvsixName defines the resource implementation.
type firewallIPvsixName struct {
	ResourceName string
	client       *client.Client
	model        resourcemodel.FirewallIPvsixName
}

// GetName returns resource name
func (r *firewallIPvsixName) GetName() string {
	return r.ResourceName
}

func (r *firewallIPvsixName) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}
