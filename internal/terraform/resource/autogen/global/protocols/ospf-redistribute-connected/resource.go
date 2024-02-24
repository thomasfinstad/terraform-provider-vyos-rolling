// Package globalprotocolsospfredistributeconnected code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsospfredistributeconnected

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/global/protocols/ospf-redistribute-connected/resourcemodel"
)

// NewProtocolsOspfRedistributeConnected method to return the example resource reference
func NewProtocolsOspfRedistributeConnected() resource.Resource {
	return &protocolsOspfRedistributeConnected{
		model: &resourcemodel.ProtocolsOspfRedistributeConnected{},
	}
}

// protocolsOspfRedistributeConnected defines the resource implementation.
type protocolsOspfRedistributeConnected struct {
	client *client.Client
	model  *resourcemodel.ProtocolsOspfRedistributeConnected
}

// GetClient returns the vyos api client
func (r *protocolsOspfRedistributeConnected) GetClient() *client.Client {
	return r.client
}

func (r *protocolsOspfRedistributeConnected) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
