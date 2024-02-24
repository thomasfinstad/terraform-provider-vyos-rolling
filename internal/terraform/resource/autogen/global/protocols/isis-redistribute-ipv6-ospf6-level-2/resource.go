// Package globalprotocolsisisredistributeipvsixospfsixleveltwo code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisisredistributeipvsixospfsixleveltwo

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/global/protocols/isis-redistribute-ipv6-ospf6-level-2/resourcemodel"
)

// NewProtocolsIsisRedistributeIPvsixOspfsixLevelTwo method to return the example resource reference
func NewProtocolsIsisRedistributeIPvsixOspfsixLevelTwo() resource.Resource {
	return &protocolsIsisRedistributeIPvsixOspfsixLevelTwo{
		model: &resourcemodel.ProtocolsIsisRedistributeIPvsixOspfsixLevelTwo{},
	}
}

// protocolsIsisRedistributeIPvsixOspfsixLevelTwo defines the resource implementation.
type protocolsIsisRedistributeIPvsixOspfsixLevelTwo struct {
	client *client.Client
	model  *resourcemodel.ProtocolsIsisRedistributeIPvsixOspfsixLevelTwo
}

// GetClient returns the vyos api client
func (r *protocolsIsisRedistributeIPvsixOspfsixLevelTwo) GetClient() *client.Client {
	return r.client
}

func (r *protocolsIsisRedistributeIPvsixOspfsixLevelTwo) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
