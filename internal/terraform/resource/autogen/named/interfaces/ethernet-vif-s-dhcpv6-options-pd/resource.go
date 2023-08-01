// Package namedinterfacesethernetvifsdhcpvsixoptionspd code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesethernetvifsdhcpvsixoptionspd

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/interfaces/ethernet-vif-s-dhcpv6-options-pd/resourcemodel"
)

// NewInterfacesEthernetVifSDhcpvsixOptionsPd method to return the example resource reference
func NewInterfacesEthernetVifSDhcpvsixOptionsPd() resource.Resource {
	return &interfacesEthernetVifSDhcpvsixOptionsPd{
		model: resourcemodel.InterfacesEthernetVifSDhcpvsixOptionsPd{},
	}
}

// interfacesEthernetVifSDhcpvsixOptionsPd defines the resource implementation.
type interfacesEthernetVifSDhcpvsixOptionsPd struct {
	ResourceName string
	client       *client.Client
	model        resourcemodel.InterfacesEthernetVifSDhcpvsixOptionsPd
}

func (r *interfacesEthernetVifSDhcpvsixOptionsPd) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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